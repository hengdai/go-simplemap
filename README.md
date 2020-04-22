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
	jsonStr := `
{
  "status": 0,
  "message": "",
  "data": {
    "search_data": [
      {
        "elements": [
          {
            "rating": 0,
            "name": "奈良市",
            "url": "/scenic/3/10052/",
            "wish_to_go_count": 328,
            "name_orig": "奈良市",
            "visited_count": 1958,
            "comments_count": 0,
            "location": {
              "lat": 34.685087,
              "lng": 135.805
            },
            "has_experience": false,
            "rating_users": 0,
            "name_zh": "奈良市",
            "name_en": "Nara",
            "type": 3,
            "id": 10052,
            "has_route_maps": false,
            "icon": "http://media.breadtrip.com/images/icons/2/city.png"
          },
          {
            "rating": 0,
            "name": "小樽市",
            "url": "/scenic/3/26772/",
            "wish_to_go_count": 266,
            "name_orig": "小樽市",
            "visited_count": 954,
            "comments_count": 0,
            "location": {
              "lat": 43.190717,
              "lng": 140.994662
            },
            "has_experience": false,
            "rating_users": 0,
            "name_zh": "小樽市",
            "name_en": "Otaru",
            "type": 3,
            "id": 26772,
            "has_route_maps": false,
            "icon": "http://media.breadtrip.com/images/icons/2/city.png"
          }
        ]
      }
    ]
  }
}
`
	smap, err := simplemap.NewMap(jsonStr)
	if err != nil {
	    panic(err.Error())
	}

	s, err := smap.GetItem("data.search_data.0.elements.1.location")
	fmt.Println(s, err)
}
}
```
输出

```
{"lat":43.190717,"lng":140.994662} <nil>
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
// 根据key获取对应的value，支持多层获取，例如a.b.c.d，支持数组索引，例如：a.0.b
func (m *simpleMap) GetItem(keyStr string) (string, error)
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

```
// 返回第一层级的数据keys的长度，如果
func (m *simpleMap) Length() int
```

```
// 获取key对应的value的长度
// 如果value是map则返回map的keys的长度，
// 如果是list，则直接返回长度
// 如果是string，则返回string的长度
func (m *simpleMap) ValueLength(keyStr string) int
```

```
// 判断key对应的value是不是map
func (m *simpleMap) IsValueMap(keyStr string) bool
```

```
// 判断key对应的value是不是array
func (m *simpleMap) IsValueArr(keyStr string) bool
```
