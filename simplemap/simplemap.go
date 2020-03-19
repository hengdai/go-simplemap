package simplemap

import (
	"encoding/json"
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

	panic("NewMap 入参json字符串或者map")
}

// 判断map是否存在key
func (m *simpleMap) ExistKey(key string) bool {
	if _, ok := m.data.(map[string]interface{})[key]; ok {
		return true
	}

	return false
}
