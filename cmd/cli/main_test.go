package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		stdin          string
		useStdin       bool
		expectedOutput map[string]interface{}
		expectError    bool
	}{
		{
			name:     "Command line argument simple parsing",
			args:     []string{"program", `name:"Fikar" age:10 gender:"male" is_married:false`},
			useStdin: false,
			expectedOutput: map[string]interface{}{
				"name":       "Fikar",
				"age":        float64(10),
				"gender":     "male",
				"is_married": false,
			},
			expectError: false,
		},
		{
			name:     "Stdin input with nested object",
			stdin:    `name:"John" age:30 address:{city:"New York" state:"NY"}`,
			useStdin: true,
			expectedOutput: map[string]interface{}{
				"name":    "John",
				"age":     float64(30),
				"address": map[string]interface{}{"city": "New York", "state": "NY"},
			},
			expectError: false,
		},
		{
			name:        "No arguments provided",
			args:        []string{"program"},
			useStdin:    false,
			expectError: true,
		},
		{
			name:     "Input with array-like repeated fields",
			stdin:    `item:"apple" item:"orange" item:"banana"`,
			useStdin: true,
			expectedOutput: map[string]interface{}{
				"item": []interface{}{"apple", "orange", "banana"},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original stdin, stdout, and args
			oldStdin := os.Stdin
			oldStdout := os.Stdout
			oldArgs := os.Args

			// Restore after test
			defer func() {
				os.Stdin = oldStdin
				os.Stdout = oldStdout
				os.Args = oldArgs
			}()

			// Set up test args
			if tt.args != nil {
				os.Args = tt.args
			}

			// Set up stdin if needed
			if tt.useStdin {
				r, w, _ := os.Pipe()
				os.Stdin = r
				w.Write([]byte(tt.stdin))
				w.Close()
			}

			// Capture stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Run the main function in a separate goroutine
			exit := make(chan struct{})
			go func() {
				main()
				close(exit)
			}()

			// Close the write end of the pipe to avoid deadlock
			w.Close()
			<-exit

			// Read captured output
			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			// Check for expected error message
			if tt.expectError {
				if !strings.Contains(output, "Usage:") && !strings.Contains(output, "Error") {
					t.Errorf("Expected error message, got: %s", output)
				}
				return
			}

			// Parse the JSON output
			var result map[string]interface{}
			err := json.Unmarshal([]byte(output), &result)
			if err != nil {
				t.Fatalf("Failed to parse JSON output: %v\nOutput was: %s", err, output)
			}

			// Compare with expected output
			if !compareJSON(result, tt.expectedOutput) {
				expectedJSON, _ := json.Marshal(tt.expectedOutput)
				resultJSON, _ := json.Marshal(result)
				t.Errorf("Expected JSON %s, got %s", expectedJSON, resultJSON)
			}
		})
	}
}

// Helper function to compare JSON objects
func compareJSON(actual, expected map[string]interface{}) bool {
	if len(actual) != len(expected) {
		return false
	}

	for k, expectedVal := range expected {
		actualVal, exists := actual[k]
		if !exists {
			return false
		}

		// Handle nested maps
		expectedMap, expectedIsMap := expectedVal.(map[string]interface{})
		actualMap, actualIsMap := actualVal.(map[string]interface{})
		if expectedIsMap && actualIsMap {
			if !compareJSON(actualMap, expectedMap) {
				return false
			}
			continue
		}

		// Handle arrays
		expectedArr, expectedIsArr := expectedVal.([]interface{})
		actualArr, actualIsArr := actualVal.([]interface{})
		if expectedIsArr && actualIsArr {
			if len(expectedArr) != len(actualArr) {
				return false
			}
			for i := range expectedArr {
				if expectedArr[i] != actualArr[i] {
					return false
				}
			}
			continue
		}

		// Simple value comparison
		if expectedVal != actualVal {
			return false
		}
	}
	return true
}
