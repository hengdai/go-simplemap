package simplemap

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type simpleJson struct {
	data interface{}
}

func NewJson(v interface{}) (*simpleJson, error) {
	if reflect.TypeOf(v).Name() == "string" {
		var dat interface{}
		v, _ := v.(string)
		err := json.Unmarshal([]byte(v), &dat)
		if nil != err {
			return nil, err
		}

		return &simpleJson{data: dat}, nil
	}

	if _, ok := v.(map[string]interface{}); ok {
		return &simpleJson{data: v}, nil
	}

	if _, ok := v.([]interface{}); ok {
		return &simpleJson{data: v}, nil
	}

	return nil, errors.New("json string is not illegal")
}

// ExistKey 判断key是否在json中存在
func (sm *simpleJson) ExistKey(keyStr string) (bool, error) {
	keyList := strings.Split(keyStr, ".")
	data := sm.data
	length := len(keyList)

	for i, key := range keyList {
		datMap, isMap := data.(map[string]interface{})
		datArr, isArray := data.([]interface{})
		if !isMap && !isArray{
			return false, errors.New("invalid map")
		}

		if isMap {
			if value, ok := datMap[key]; ok {
				if length-1 == i {
					return true, nil
				}
				data = value
			}
		}

		if isArray {
			interKey, err := strconv.Atoi(key)
			if err != nil {
				return false, errors.New(fmt.Sprintf("key [%v] is not number", keyStr))
			}

			if len(datArr) < interKey + 1 {
				return false, errors.New("index out of range [" + key +"] with length " + strconv.Itoa(len(datArr)))
			}

			value := datArr[interKey]
			if length-1 == i {
				return true, nil
			}

			data = value
		}
	}

	return false, errors.New("key '" + keyStr + "' not exist")
}

// GetValue 根据key获取对应的json value，支持多层获取，例如a.b.c.d，支持数组索引，例如：a.0.b
func (sm *simpleJson) GetValue(keyStr string) (string, error) {
	keyList := strings.Split(keyStr, ".")
	data := sm.data
	length := len(keyList)

	for i, key := range keyList {
		datMap, isMap := data.(map[string]interface{})
		datArr, isArray := data.([]interface{})
		if !isMap && !isArray{
			return "", errors.New("invalid map")
		}

		if isMap {
			if value, ok := datMap[key]; ok {
				if length-1 == i {
					b, err := json.Marshal(value)
					if err != nil {
						return "", err
					}
					return string(b), nil
				}
				data = value
			}
		}

		if isArray {
			interKey, err := strconv.Atoi(key)
			if err != nil {
				return "", errors.New(fmt.Sprintf("key [%v] is not number", keyStr))
			}

			if len(datArr) < interKey + 1 {
				return "", errors.New("index out of range [" + key +"] with length " + strconv.Itoa(len(datArr)))
			}

			value := datArr[interKey]
			if length-1 == i {
				b, err := json.Marshal(value)
				if err != nil {
					return "", err
				}
				return string(b), nil
			}
			data = value
		}
	}

	return "", errors.New("key '" + keyStr + "' not exist")
}

// 返回key对应的item中所有key
func (sm *simpleJson) Keys(keyStr string) (string, error) {
	value, err := sm.GetValue(keyStr)
	if err != nil {
		return "", err
	}

	var vInterface interface{}
	err = json.Unmarshal([]byte(value), &vInterface)
	if err != nil {
		return "", err
	}

	// map
	if item, ok := vInterface.(map[string]interface{}); ok {
		var keys []string
		for key := range item {
			keys = append(keys, key)
		}
		b, err := json.Marshal(keys)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	// []
	if items, ok := vInterface.([]interface{}); ok {
		var keys []string
		for i := range items {
			keys = append(keys, strconv.Itoa(i))
		}
		b, err := json.Marshal(keys)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	// string
	if items, ok := vInterface.(string); ok {
		return "", errors.New(fmt.Sprintf("key [%v] get string [%v], string has no key", keyStr, items))
	}

	return "", errors.New("invalid map")
}

// 返回key对应的item中所有Values
func (sm *simpleJson) Values(keyStr string) (string, error) {
	value, err := sm.GetValue(keyStr)
	if err != nil {
		return "", err
	}

	var vInterface interface{}
	err = json.Unmarshal([]byte(value), &vInterface)
	if err != nil {
		return "", err
	}

	// map
	if item, ok := vInterface.(map[string]interface{}); ok {
		var values []interface{}
		for _, value := range item {
			values = append(values, value)
		}

		b, err := json.Marshal(values)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	// []
	if items, ok := vInterface.([]interface{}); ok {
		var values []interface{}
		for _, value := range items {
			values = append(values, value)
		}

		b, err := json.Marshal(values)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	// string
	if item, ok := vInterface.(string); ok {
		return item, nil
	}

	return "", errors.New("invalid map")
}

// 判断key对应的value是不是map
func (sm *simpleJson) IsValueMap(keyStr string) bool {
	values, err := sm.GetValue(keyStr)
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
func (sm *simpleJson) IsValueArr(keyStr string) bool {
	values, err := sm.GetValue(keyStr)
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