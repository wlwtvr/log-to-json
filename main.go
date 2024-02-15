package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// Function to parse text to JSON
func parseTextToJSON(text string) map[string]interface{} {
	// Regular expression pattern to match field names and values
	pattern := `(\w+):(".*?"|\{.*?\})`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(text, -1)

	// Initialize map to store parsed data
	parsedData := make(map[string]interface{})

	for _, match := range matches {
		fieldName := match[1]
		fieldValue := match[2]

		// Remove quotes if value is a string
		if strings.HasPrefix(fieldValue, `"`) && strings.HasSuffix(fieldValue, `"`) {
			fieldValue = fieldValue[1 : len(fieldValue)-1]
		}

		// Check if fieldValue represents a nested object
		if strings.HasPrefix(fieldValue, "{") && strings.HasSuffix(fieldValue, "}") {
			// Parse nested object recursively
			parsedData[fieldName] = parseNestedObject(fieldValue[1 : len(fieldValue)-1])
		} else {
			parsedData[fieldName] = fieldValue
		}
	}

	return parsedData
}

// Function to parse nested object
func parseNestedObject(text string) map[string]interface{} {
	nestedData := make(map[string]interface{})

	// Regular expression pattern to match nested field names and values
	pattern := `(\w+):(".*?"|\{.*?\})`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		fieldName := match[1]
		fieldValue := match[2]

		// Remove quotes if value is a string
		if strings.HasPrefix(fieldValue, `"`) && strings.HasSuffix(fieldValue, `"`) {
			fieldValue = fieldValue[1 : len(fieldValue)-1]
		}

		// Check if fieldValue represents a nested object
		if strings.HasPrefix(fieldValue, "{") && strings.HasSuffix(fieldValue, "}") {
			// Parse nested object recursively
			nestedData[fieldName] = parseNestedObject(fieldValue[1 : len(fieldValue)-1])
		} else {
			nestedData[fieldName] = fieldValue
		}
	}

	return nestedData
}

func main() {
	var inputText string

	// Check if input is being piped
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Data is being piped, read from stdin
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			return
		}
		inputText = string(bytes)
	} else {
		// No piped input, read from command line argument
		if len(os.Args) < 2 {
			fmt.Println("Usage: cli_command <text>")
			return
		}
		inputText = os.Args[1]
	}

	// Parse text to JSON
	parsedJSON := parseTextToJSON(inputText)

	// Convert parsed data to JSON string
	jsonStr, err := json.MarshalIndent(parsedJSON, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print the parsed JSON
	fmt.Println(string(jsonStr))
}
