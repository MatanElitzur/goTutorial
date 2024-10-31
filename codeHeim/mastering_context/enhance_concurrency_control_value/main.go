package main

import (
	"context"
	"fmt"
)

// Execute go run ./codeHeim/mastering_context/enhance_concurrency_control_value/main.go
// This example demonstrates how to use context to pass request-scoped values in the code and not definding specific parameters in the function signature
func main() {
	// Create a context with a value and store a value in it
	ctx := context.WithValue(context.Background(), "userID", 42)

	// Pass the context to the function that needs access to the value
	ProcessRequest(ctx)
}

// ProcessRequest simulates processing a request with access to the context value
func ProcessRequest(ctx context.Context) {
	// Extract the value from the context

	userID, ok := ctx.Value("userID").(int)
	if !ok {
		println("userID not found in context")
		return
	}

	//Use the extracted value in your code
	fmt.Printf("Processing request for user ID: %d\n", userID)

	//Simulate passing the context further down the call stack
	FurtherProcessing(ctx)
}

func FurtherProcessing(ctx context.Context) {
	// Extract the value from the context
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		println("userID not found in context")
		return
	}

	// Use the extracted value in your code
	fmt.Printf("Further processing for user ID: %d\n", userID)
}
