package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"wlwtvr/log-to-json/internal/parser"
)

func parseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	parsedJSON, err := parser.ParseTextToJSON(string(body))
	if err != nil {
		http.Error(w, "Error parsing text", http.StatusBadRequest)
		return
	}

	jsonResponse, err := json.MarshalIndent(parsedJSON, "", "  ")
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/parse", parseHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
