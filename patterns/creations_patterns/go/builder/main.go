package main

import (
	"builder"
	"fmt"
)

func main() {
	// Example 1: Simple GET request
	fmt.Println("=== Example 1: Simple GET request ===")
	getBuilder, err := builder.NewHttpRequestBuilder("https://httpbin.org/get")
	if err != nil {
		fmt.Printf("Error creating builder: %v\n", err)
		return
	}

	getRequest := getBuilder.
		Header("Accept", "application/json").
		QueryParam("name", "builder-pattern").
		QueryParam("version", "1.0").
		Timeout(5000).
		Build()

	fmt.Printf("Request: %s\n", getRequest)

	resp, err := getRequest.Execute()
	if err != nil {
		fmt.Printf("Error executing request: %v\n", err)
	} else {
		fmt.Printf("Response: %s\n", resp)
		fmt.Printf("Body preview: %.200s...\n", resp.Body)
	}

	// Example 2: POST request with JSON body
	fmt.Println("\n=== Example 2: POST request with JSON body ===")
	postBuilder, err := builder.NewHttpRequestBuilder("https://httpbin.org/post")
	if err != nil {
		fmt.Printf("Error creating builder: %v\n", err)
		return
	}

	postRequest := postBuilder.
		Method("POST").
		Header("Content-Type", "application/json").
		Header("Accept", "application/json").
		Body(`{"message": "Hello from Go Builder Pattern!"}`).
		Timeout(5000).
		Build()

	fmt.Printf("Request: %s\n", postRequest)

	resp, err = postRequest.Execute()
	if err != nil {
		fmt.Printf("Error executing request: %v\n", err)
	} else {
		fmt.Printf("Response: %s\n", resp)
		fmt.Printf("Body preview: %.300s...\n", resp.Body)
	}
}
