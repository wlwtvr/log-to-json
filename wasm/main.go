//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"syscall/js"

	"wlwtvr/log-to-json/internal/parser"
)

func parse(this js.Value, p []js.Value) interface{} {
	text := p[0].String()
	result, err := parser.ParseTextToJSON(text)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	return string(jsonResult)
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("parseTextToJSON", js.FuncOf(parse))
	<-c
}
