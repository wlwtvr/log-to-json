package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"wlwtvr/log-to-json/internal/parser"
)

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
	parsedJSON, err := parser.ParseTextToJSON(inputText)
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
