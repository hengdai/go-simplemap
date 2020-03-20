package main

import (
	"fmt"
	"go-simplemap/simplemap"
)

func main() {
	data := `{
		"name": "daiheng",
		"age": "25"
		}`

	testMap, err := simplemap.NewMap(data)
	if nil != err {
		fmt.Println(err.Error())
		return
	}

	keys, _ := testMap.Values()
	fmt.Println(keys)
}
