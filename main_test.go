package main

import (
	"reflect"
	"testing"
)

func TestParseTextToJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]interface{}
	}{
		{
			name:  "Test simple parsing",
			input: `name:"Fikar" age:10 address:{city:"Jakarta Selatan" state:"DKI Jakarta"} gender:"male" is_married:false`,
			expected: map[string]interface{}{
				"name": "Fikar",
				"age":  10,
				"address": map[string]interface{}{
					"city":  "Jakarta Selatan",
					"state": "DKI Jakarta",
				},
				"gender":     "male",
				"is_married": false,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := parseTextToJSON(test.input)
			if err != nil {
				t.Errorf("Error parsing text: %v", err)
				return
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected: %+v\nGot: %+v", test.expected, result)
			}
		})
	}
}
