package simplemap

import (
	"encoding/json"
	"errors"
	"reflect"
)

type simpleMap struct {
	data interface{}
}

// return the current version
func Version() string {
	return "0.0.1"
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