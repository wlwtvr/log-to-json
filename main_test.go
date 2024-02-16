package main

import (
	"reflect"
	"testing"
)

func TestParseTextToJSON(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput interface{}
		expectedError  error
	}{
		{
			name:  "Test simple parsing",
			input: `name:"Fikar" age:10 address:{city:"Jakarta Selatan" state:"DKI Jakarta"} gender:"male" is_married:false`,
			expectedOutput: map[string]interface{}{
				"name":       "Fikar",
				"age":        10,
				"address":    map[string]interface{}{"city": "Jakarta Selatan", "state": "DKI Jakarta"},
				"gender":     "male",
				"is_married": false,
			},
			expectedError: nil,
		},
		{
			name:  "Test scientific positive notation number",
			input: `name:"John" age:30 salary:2e+06 married:true`,
			expectedOutput: map[string]interface{}{
				"name":    "John",
				"age":     30,
				"salary":  "2e+06",
				"married": true,
			},
			expectedError: nil,
		},
		{
			name:  "Test scientific negative notation number",
			input: `atom_size:"1e-23"`,
			expectedOutput: map[string]interface{}{
				"atom_size": "1e-23",
			},
			expectedError: nil,
		},
		{
			name:  "Test identical field name, with simple value",
			input: `name:"Alice" age:25 name:"Bob"`,
			expectedOutput: map[string]interface{}{
				"name": []interface{}{
					"Alice",
					"Bob",
				},
				"age": 25,
			},
			expectedError: nil,
		},
		{
			name:  "Test identical field name, with object value",
			input: `display_data:{header:"header_1" title:"title_1"} display_data:{header:"header_2"  title:"title_2"}`,
			expectedOutput: map[string]interface{}{
				"display_data": []interface{}{
					map[string]interface{}{
						"header": "header_1",
						"title":  "title_1",
					},
					map[string]interface{}{
						"header": "header_2",
						"title":  "title_2",
					},
				},
			},
			expectedError: nil,
		},
		{
			name:  "Test floating number (dot)",
			input: `number:35.30`,
			expectedOutput: map[string]interface{}{
				"number": 35.30,
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := parseTextToJSON(test.input)
			if err != test.expectedError {
				t.Errorf("Expected error: %v, got: %v", test.expectedError, err)
				return
			}
			if !reflect.DeepEqual(output, test.expectedOutput) {
				t.Errorf("Expected output: %v, got: %v", test.expectedOutput, output)
			}
		})
	}
}
