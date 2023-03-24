# gokvpairs

KeyValuePairs is a type that can hold the order of keys of json object in Golang.

The keys of map in Golang do not guarantee the order, so if we parse json in Golang with map[string]interface{} ,the keys will sorted by alphabet when it is unmarshalled.

This is a problem in some projects, so I defined a type KeyValuePairs[interface{}] to replace map[string]interface{} when I decode those json data whose key order need be preserved.

## example
main.go
```
package main

import (
	"encoding/json"
	"fmt"

	"github.com/kkqy/gokvpairs"
)

type OldTestStruct struct {
	Name          string            `json:"name"`
	OrderedObject map[string]string `json:"ordered_object"`
}
type NewTestStruct struct {
	Name          string                          `json:"name"`
	OrderedObject gokvpairs.KeyValuePairs[string] `json:"ordered_object"`
}

func main() {
	data := []byte(`
	{
		"name": "Test",
		"ordered_object": {
			"c": "value_c",
			"b": "value_b"
		}
	}`)
	fmt.Println("OriginalData:", string(data))
	fmt.Println("====================")

	// Old
	oldTestStruct := OldTestStruct{}
	json.Unmarshal([]byte(data), &oldTestStruct)
	result, _ := json.Marshal(oldTestStruct)
	fmt.Println("Old：", string(result))
	oldTestStruct.OrderedObject["a"] = "value_a"
	result, _ = json.Marshal(oldTestStruct)
	fmt.Println("Old：", string(result))

	fmt.Println("====================")

	// New
	newTestStruct := NewTestStruct{}
	json.Unmarshal([]byte(data), &newTestStruct)
	result, _ = json.Marshal(newTestStruct)
	fmt.Println("New：", string(result))
	newTestStruct.OrderedObject = append(newTestStruct.OrderedObject, gokvpairs.KeyValuePair[string]{Key: "a", Value: "value_a"})
	result, _ = json.Marshal(newTestStruct)
	fmt.Println("New：", string(result))
}
```
output:
```
OriginalData: 
        {
                "name": "Test",
                "ordered_object": {
                        "c": "value_c",
                        "b": "value_b"
                }
        }
====================
Old： {"name":"Test","ordered_object":{"b":"value_b","c":"value_c"}}
Old： {"name":"Test","ordered_object":{"a":"value_a","b":"value_b","c":"value_c"}}
====================
New： {"name":"Test","ordered_object":{"c":"value_c","b":"value_b"}}
New： {"name":"Test","ordered_object":{"c":"value_c","b":"value_b","a":"value_a"}}
```
