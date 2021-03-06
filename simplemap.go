package simplemap

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type simpleMap struct {
	data interface{}
}

// return the current version
func (m *simpleMap) Version() string {
	return "0.0.6"
}

// 初始化map，入参支持字符串json和map，且key只能是string类型
func NewMap(v interface{}) (*simpleMap, error) {
	var data simpleMap
	if reflect.TypeOf(v).Name() == "string" {
		v, _ := v.(string)
		err := json.Unmarshal([]byte(v), &data.data)
		if nil != err {
			return nil, err
		}

		return &data, nil
	}

	if _, ok := v.(map[string]interface{}); ok {
		data.data = v
		return &data, nil
	}

	panic("NewMap need json string or map[string]interface")
}

// 返回map类型的数据
func (m *simpleMap) GetMap() (map[string]interface{}, error) {
	if item, ok := m.data.(map[string]interface{}); ok {
		return item, nil
	}

	return nil, errors.New("invalid map")
}

// 判断map是否存在key
func (m *simpleMap) ExistKey(keyStr string) (bool, error) {
	keyList := strings.Split(keyStr, ".")
	data := m.data
	length := len(keyList)

	for i, key := range keyList {
		_, isMap := data.(map[string]interface{})
		_, isArray := data.([]interface{})
		if !isMap && !isArray{
			return false, errors.New("invalid map")
		}

		if isMap {
			if value, ok := data.(map[string]interface{})[key]; ok {
				if length-1 == i {
					return true, nil
				}
				data = value
			}
		}

		if isArray {
			interKey, err := strconv.Atoi(key)
			if arr, ok := data.([]interface{}); ok && err == nil {
				if len(arr) < interKey + 1 {
					return false, errors.New("index out of range [" + key +"] with length " + strconv.Itoa(len(arr)))
				}
				value := arr[interKey]
				if length-1 == i {
					return true, nil
				}
				data = value
			}
		}
	}

	return false, errors.New("key '" + keyStr + "' not exist")
}

// 新增或者更改map key对应的value，只支持一层嵌套
func (m *simpleMap) SetItem(key string, value interface{}) error {
	item, ok := m.data.(map[string]interface{})
	if !ok {
		err :=  errors.New("invalid map")
		return err
	}
	item[key] = value
	m.data = item

	return nil
}

// 根据key获取对应的value，支持多层获取，例如a.b.c.d，支持数组索引，例如：a.0.b
func (m *simpleMap) GetItem(keyStr string) (string, error) {
	keyList := strings.Split(keyStr, ".")
	data := m.data
	length := len(keyList)

	for i, key := range keyList {
		_, isMap := data.(map[string]interface{})
		_, isArray := data.([]interface{})
		if !isMap && !isArray{
			return "", errors.New("invalid map")
		}

		if isMap {
			if value, ok := data.(map[string]interface{})[key]; ok {
				if length-1 == i {
					str, err := json.Marshal(value)
					if err != nil {
						return "", err
					}
					return string(str), nil
				}
				data = value
			}
		}

		if isArray {
			interKey, err := strconv.Atoi(key)
			if arr, ok := data.([]interface{}); ok && err == nil {
				if len(arr) < interKey + 1 {
					return "", errors.New("index out of range [" + key +"] with length " + strconv.Itoa(len(arr)))
				}
				value := arr[interKey]
				if length-1 == i {
					str, err := json.Marshal(value)
					if err != nil {
						return "", err
					}
					return string(str), nil
				}
				data = value
			}
		}
	}

	return "", errors.New("key '" + keyStr + "' not exist")
}

// 删除key所对应的那一条item，如果key不存在，不做任何操作。只支持一层嵌套
func (m *simpleMap) DelItem(key string) error {
	item, ok := m.data.(map[string]interface{})
	if !ok {
		err :=  errors.New("invalid map")
		return err
	}
	delete(item, key)
	m.data = item

	return nil
}

// 返回item对应的所有key
func (m *simpleMap) Keys() ([]interface{}, error) {
	item, ok := m.data.(map[string]interface{})
	if !ok {
		err :=  errors.New("invalid map")
		return nil, err
	}

	var keys []interface{}
	for key := range item {
		keys = append(keys, key)
	}

	return keys, nil
}

// 返回item对应的所有value
func (m *simpleMap) Values() ([]interface{}, error) {
	item, ok := m.data.(map[string]interface{})
	if !ok {
		err :=  errors.New("invalid map")
		return nil, err
	}

	var values []interface{}
	for _, value := range item {
		values = append(values, value)
	}

	return values, nil
}

// map转换为json字符串
func (m *simpleMap) JsonStr() (string, error) {
	jsonRes, err := json.Marshal(m.data)
	if nil != err {
		return "", err
	}

	return string(jsonRes), nil
}

// 返回第一层级的数据keys的长度，如果
func (m *simpleMap) Length() int {
	keys, err := m.Keys()
	if nil != err {
		return 0
	}

	return len(keys)
}

// 获取key对应的value的长度
// 如果value是map则返回map的keys的长度，
// 如果是list，则直接返回长度
// 如果是string，则返回string的长度
func (m *simpleMap) ValueLength(keyStr string) int {
	values, err := m.GetItem(keyStr)
	if nil != err {
		return 0
	}

	var valMap map[string]interface{}
	var valList []interface{}
	isMap := json.Unmarshal([]byte(values), &valMap)
	isArray := json.Unmarshal([]byte(values), &valList)

	if nil == isMap {
		length := 0
		for range valMap {
			length++
		}
		return length
	}

	if nil == isArray {
		return len(valList)
	}

	return len(values)
}

// 判断key对应的value是不是map
func (m *simpleMap) IsValueMap(keyStr string) bool {
	values, err := m.GetItem(keyStr)
	if nil != err {
		return false
	}

	var valMap map[string]interface{}
	isMap := json.Unmarshal([]byte(values), &valMap)

	if nil != isMap {
		return false
	}

	return true
}

// 判断key对应的value是不是array
func (m *simpleMap) IsValueArr(keyStr string) bool {
	values, err := m.GetItem(keyStr)
	if nil != err {
		return false
	}

	var valList []interface{}
	isArray := json.Unmarshal([]byte(values), &valList)

	if nil != isArray {
		return false
	}

	return true
}