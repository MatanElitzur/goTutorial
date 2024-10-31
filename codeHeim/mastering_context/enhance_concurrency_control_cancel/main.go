//What is context?
// A built-in package that provides tools to manage concurrent operations.
// Context carries deadlines, cancelation signals and request-scoped values between goroutines.

package main

import (
	"context"
	"fmt"
	"time"
)

// Execute go run ./codeHeim/mastering_context/enhance_concurrency_control_cancel/main.go
// This example demonstrates how to use context to cancel all running goroutines
func main() {
	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	// returns a ctx context and a cancel function
	//context should signal when operation should be stopped
	//the cancel function trigger this signal used in line 54

	//Launce a goroutine that will listen to the context cancellation signal
	go func() {
		for {
			select {
			case <-ctx.Done(): //Check if the context is cancelled
				fmt.Println("Goroutine 1 cancelled", ctx.Err())
				return
			default:
				fmt.Println("Goroutine 1 is working....")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Launce another goroutine that simulates some work
	go func() {
		for {
			select {
			case <-ctx.Done(): //Check if the context is cancelled
				fmt.Println("Goroutine 2 cancelled", ctx.Err())
				return
			default:
				fmt.Println("Goroutine 2 is working....")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Simulate some work in the main function
	fmt.Println("Main function is working...")
	time.Sleep(2 * time.Second) // in this time the goroutines are working

	// Cancel the context, which will signal the goroutines to stop
	fmt.Println("Cancelling the context...")
	cancel()

	//Give goroutines time to finish
	time.Sleep(1 * time.Second)
	fmt.Println("Main function done.")
}
