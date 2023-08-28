package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrease the counter when the function exits
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate some work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	numWorkers := 3

	fmt.Println("Main starting")

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Increment the counter for each worker
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all workers to finish

	fmt.Println("Main done")
}

// The main function creates a number of worker Goroutines,
// increments the WaitGroup counter for each one,
// and then waits for all the workers to finish using wg.Wait().
// This ensures that the main function doesn't proceed until all the workers are done.

// Using the sync.WaitGroup properly is important
// to avoid race conditions and ensure proper synchronization in a Go program.
