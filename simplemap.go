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
func Version() string {
	return "0.0.3"
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
func (m *simpleMap) ExistKey(key string) bool {
	if _, ok := m.data.(map[string]interface{})[key]; ok {
		return true
	}

	return false
}

// 新增或者更改map key对应的value
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

// 删除key所对应的那一条item，如果key不存在，不做任何操作
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