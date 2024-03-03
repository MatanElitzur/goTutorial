package knowledge

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func Limiter_1() {
	// Create a rate limiter allowing 10 requests per second with a capacity of 2
	limiter := rate.NewLimiter(rate.Every(time.Second), 10)

	// Function simulating resource access (replace with your actual logic)
	accessResource := func() {
		fmt.Println("Accessing resource...")
		// Simulate some work
		time.Sleep(time.Second * 1) // Simulate 1 second of work
	}

	for i := 0; i < 20; i++ {
		// Check if the limiter allows the request
		allowed := limiter.Allow()

		if !allowed {
			fmt.Println("Rate limit exceeded, waiting...")
			// Wait for the next available token (optional, implement backoff strategy)
			time.Sleep(time.Second * 1) // Wait for 1 second before retrying
			continue
		}

		// Access the resource if allowed
		accessResource()
	}
}

func Limiter_2() {
	fmt.Println("**** Limiter_2 ****")
	// Create a rate limiter allowing 5 requests per second with a capacity of 3
	limiter := rate.NewLimiter(rate.Every(time.Second), 5)

	// Function simulating resource access (replace with your actual logic)
	accessResource := func() {
		fmt.Println("Accessing resource...")
		// Simulate some work
		time.Sleep(time.Second * 1) // Simulate 1 second of work
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	for i := 0; i < 10; i++ {
		// Wait for the limiter to allow N (2) events
		err := limiter.WaitN(ctx, 2)
		if err != nil {
			fmt.Println("Error waiting for tokens:", err)
			// Handle context cancellation or exceeding wait time limit
			break
		}

		// Access the resource if allowed
		accessResource()
	}
}
