//What is context?
// A built-in package that provides tools to manage concurrent operations.
// Context carries deadlines, cancelation signals and request-scoped values between goroutines.

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Execute go run ./codeHeim/mastering_context/enhance_concurrency_control_timeout/main.go
// This example demonstrates how to use context to cancel a slow search operation after a timeout
func main() {
	// Create a context with a timeout of 3 seconds
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc() //cancel the context when the main function is done, the time start in the invoke of the function context.WithTimeout
	// returns a ctx context and a cancelFunc function
	//context should signal when operation should be stopped
	//the cancel function trigger this signal used in line 54

	// Perform a search with the context
	res, err := Search(ctx, "random string")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("response: %s\n", res)
}

// Search simulates a search operation that can be cancelled via context
func Search(ctx context.Context, query string) (string, error) {
	//Make a channel to receive the result from the slow function
	resp := make(chan string)
	go func() {
		resp <- RandomSleepAndReturnAPI(query)
		close(resp)
	}()

	//Wait for either the response of the function that finished or the context to be done cause ot a time out
	for {
		select {
		case dst := <-resp:
			return dst, nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

// RandomSleepAndReturnAPI simulates a slow API call by sleeping for a random duration
func RandomSleepAndReturnAPI(query string) string {
	//create a new random number generator with a custom seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	//generate a random duration between 0 and 5 seconds
	randomDuration := time.Duration(rng.Int63n(int64(5 * time.Second)))
	//sleep for the random duration
	time.Sleep(randomDuration)
	return fmt.Sprintf("It took us %v... Hope it was worth the wait!", randomDuration)
}
