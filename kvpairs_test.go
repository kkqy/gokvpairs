package gokvpairs_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kkqy/gokvpairs"
)

type TestStruct struct {
	Name          string                          `json:"name"`
	OrderedObject gokvpairs.KeyValuePairs[string] `json:"ordered_object"`
}

func TestKeyValuePairs(t *testing.T) {
	testStruct := TestStruct{}
	json.Unmarshal([]byte(`{
		"name": "Test",
		"ordered_object": {
			"c": "value_c",
			"b": "value_b"
		}
	}`), &testStruct)
	result, _ := json.Marshal(testStruct)
	fmt.Println(string(result))
	testStruct.OrderedObject = append(testStruct.OrderedObject, gokvpairs.KeyValuePair[string]{Key: "a", Value: "value_a"})
	result, _ = json.Marshal(testStruct)
	fmt.Println(string(result))
}
