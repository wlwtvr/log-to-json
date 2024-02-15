package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Function to parse text to JSON
func parseTextToJSON(text string) (map[string]interface{}, error) {
	// Regular expression pattern to match field names and values
	pattern := `(\w+):(".*?"|\{.*?\}|\w+)`
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
			nestedObject, err := parseTextToJSON(fieldValue[1 : len(fieldValue)-1])
			if err != nil {
				return nil, err
			}
			parsedData[fieldName] = nestedObject
		} else {
			// Parse field value based on its type
			value, err := parseValue(fieldValue)
			if err != nil {
				return nil, err
			}
			parsedData[fieldName] = value
		}
	}

	return parsedData, nil
}

// Function to parse field value based on its type
func parseValue(value string) (interface{}, error) {
	// Try to parse as integer
	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue, nil
	}

	// Try to parse as float
	if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
		return floatValue, nil
	}

	// Try to parse as boolean
	if boolValue, err := strconv.ParseBool(value); err == nil {
		return boolValue, nil
	}

	// If not an integer, float, or boolean, treat as string
	return value, nil
}

func main() {
	var inputText string

	// Check if input is being piped
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Data is being piped, read from stdin
		bytes, err := io.ReadAll(os.Stdin)
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
	parsedJSON, err := parseTextToJSON(inputText)
	if err != nil {
		fmt.Println("Error parsing text:", err)
		return
	}

	// Convert parsed data to JSON string
	jsonStr, err := json.MarshalIndent(parsedJSON, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print the parsed JSON
	fmt.Println(string(jsonStr))
}
