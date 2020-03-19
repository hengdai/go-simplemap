package simplemap

import "encoding/json"

// return the current version
func Version() string {
	return "0.0.1"
}

func NewMap(data string) (interface{}, error) {
	var simpleMap interface{}
	err := json.Unmarshal([]byte(data), &simpleMap)
	if nil != err {
		return nil, err
	}

	return simpleMap, nil
}

