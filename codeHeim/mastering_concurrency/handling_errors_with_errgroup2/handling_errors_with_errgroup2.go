// What is errgroup?
// /provides a way to group goroutines
// run Goroutines concurrently while handling erros.
// On error, other goroutines can be cancelled.
// Before running you need to import the package "golang.org/x/sync/errgroup"
// go get golang.org/x/sync/errgroup
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"

	"golang.org/x/sync/errgroup"
)

var urls = []string{
	"https://invalid-url",
	"https://golang.org",
	"https://www.codeheim.io",
	"https://pkg.go.dev/golang.org/x/sync/errgroup",
}

func fetchPage(ctx context.Context, url string, mu *sync.Mutex, responses *map[string]string) error {
	//we need to check if the context is cancelled before proceeding
	select {
	//This is useful on going operations when another go routine encounters an error
	case <-ctx.Done():
		fmt.Println("Context is cancelled", ctx.Err())
		return nil
	default:
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("failed to fetch %s: %s\n", url, err)
			//at the moment we are not returning the error the context will be cancelled of all the go routines
			return fmt.Errorf("failed to fetch %s: %w", url, err)
		}
		defer resp.Body.Close()

		fmt.Printf("Successfully fetched %s\n", url)
		//Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body from: %s: %w", url, err)
		}

		//Store the response body in the map
		//But we need a Mutex cause we are writing to the map from multiple goroutines
		//We need to lock that no race condition occurs
		mu.Lock()
		(*responses)[url] = string(body)
		mu.Unlock()

		fmt.Printf("Successfully read response body of %s\n", url)
	}

	return nil // return nil if no error
}

// Execute go run ./codeHeim/mastering_concurrency/handling_errors_with_errgroup2/handling_errors_with_errgroup2.go
// we want to featch the content of the urls concurrently and we don't want to fail the entire process if one of the urls fails.
// it will skip the senario  if one of the go rotines fails
func main() {
	g, ctx := errgroup.WithContext(context.Background()) // create a new errgroup with a context
	g.SetLimit(2)                                        // set the limit of the number of goroutines to run concurrently
	// create a map to store the responses
	responses := make(map[string]string)
	var mu sync.Mutex

	for _, url := range urls {
		//Start a new goroutine for each url
		g.Go(func() error {
			return fetchPage(ctx, url, &mu, &responses)
		})
	}

	// Wait for all the goroutines to finish
	if err := g.Wait(); err != nil {
		fmt.Println("Error occurred: ", err)
	} else {
		fmt.Println("All urls fetched successfully")
		//Print the responses
		for url, content := range responses {
			fmt.Printf("Respone from %s: %s\n", url, content[:100]) //Print only the first 100 characters
		}
	}

}
