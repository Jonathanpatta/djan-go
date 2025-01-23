package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestProduct(t *testing.T) {
	jsonData, err := json.Marshal(Product{
		Id:    "test5",
		Code:  "asdfzzx",
		Price: 234234,
	})
	if err != nil {
		fmt.Printf("Error marshalling struct: %v\n", err)
		return
	}

	// Define the URL for the POST request
	url := "http://localhost:8000/api/product"

	// Create a new POST request with the JSON as the body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request using http.DefaultClient
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	bodyString := string(body)

	// Print the response status
	fmt.Printf("Response status: %s\n", resp.Status)

	fmt.Printf("Response body: %s\n", bodyString)
}
