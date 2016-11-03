package models

import (
	"encoding/json"
	"fmt"
)

// Parse parses string JSON to a model
func Parse(stringModel string) interface{} {
	var model interface{}
	err := json.Unmarshal([]byte(stringModel), model)
	if err != nil {
		fmt.Println("model Parse error:", err)
		fmt.Println("model Parse error json:", stringModel)
	}
	return model
}

// Stringnify convets model to JSON string
func Stringnify(model interface{}) string {
	byteModel, err := json.Marshal(model)
	if err != nil {
		fmt.Println("model Stringnify error:", err)
	}
	return string(byteModel)
}
