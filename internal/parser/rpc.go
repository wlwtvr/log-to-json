package parser

import (
	"regexp"
	"strconv"
	"strings"
)

// ParseTextToJSON parses the given text and converts it into a JSON-like map.
// It uses a regular expression to match field names and values.
func ParseTextToJSON(text string) (map[string]interface{}, error) {
	// Regular expression pattern to match field names and values
	pattern := `(\w+):(".*?"|\{.*?\}|[\w+eE.-]+|true|false)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(text, -1)

	// Initialize map to store parsed data
	parsedData := make(map[string]interface{})

	for _, match := range matches {
		fieldName, fieldValue := match[1], match[2]
		value, err := parseFieldValue(fieldValue)
		if err != nil {
			return nil, err
		}
		updateParsedData(parsedData, fieldName, value)
	}

	return parsedData, nil
}

// parseFieldValue parses the given field value and returns the corresponding interface{} value.
// If the field value represents a nested object, it recursively calls parseTextToJSON to parse it.
func parseFieldValue(fieldValue string) (interface{}, error) {
	if strings.HasPrefix(fieldValue, "{") && strings.HasSuffix(fieldValue, "}") {
		return ParseTextToJSON(fieldValue[1 : len(fieldValue)-1])
	}
	return parseValue(fieldValue)
}

// updateParsedData updates the given parsedData map with the provided fieldName and value.
// If the fieldName already exists in parsedData, it converts the existing value to an array
// and appends the new value to it; otherwise, it updates the parsedData with the new value.
func updateParsedData(parsedData map[string]interface{}, fieldName string, value interface{}) {
	if existingValue, ok := parsedData[fieldName]; ok {
		arr, isArray := existingValue.([]interface{})
		if !isArray {
			arr = []interface{}{existingValue}
		}
		arr = append(arr, value)
		parsedData[fieldName] = arr
	} else {
		parsedData[fieldName] = value
	}
}

// parseValue parses the given value and returns the corresponding interface{} value.
func parseValue(value string) (interface{}, error) {
	// Remove quotes if value is a string
	if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
		return value[1 : len(value)-1], nil
	}

	// Try to parse floating/decimal number
	if strings.Contains(value, ".") {
		if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			return floatValue, nil
		}
	}

	// Try to parse as integer
	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue, nil
	}

	// Try to parse as boolean
	if boolValue, err := strconv.ParseBool(value); err == nil {
		return boolValue, nil
	}

	// Handle scientific notation numbers
	if isScientificNotation(value) {
		return value, nil
	}

	// If not an integer, float, or boolean, treat as string
	return value, nil
}

// isScientificNotation checks if the given string represents a number in scientific notation.
func isScientificNotation(s string) bool {
	// Regular expression pattern to match scientific notation
	pattern := `[eE][-+]?\d+`
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}
