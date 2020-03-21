# go-simplemap

#### 将golang的常用map操作进行封装，使用更加顺手

示例：
```
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
    simpleMap, _ := testMap.GetMap()
    fmt.Println(simpleMap)
}
```
输出

```
map[age:25 name:daiheng]
```

所有方法如下：
```
// 返回map类型的数据
func (m *simpleMap) GetMap() (map[string]interface{}, error)
```

```
// 判断map是否存在key
func (m *simpleMap) ExistKey(key string) bool
```

```
// 新增或者更改map key对应的value
func (m *simpleMap) SetItem(key string, value interface{}) error
```

```
// 删除key所对应的那一条item，如果key不存在，不做任何操作
func (m *simpleMap) DelItem(key string) error
```

```
// 返回item对应的所有key
func (m *simpleMap) Keys() ([]interface{}, error)
```

```
// 返回item对应的所有value
func (m *simpleMap) Values() ([]interface{}, error)
```

```
// map转换为json字符串
func (m *simpleMap) JsonStr() (string, error)
```

