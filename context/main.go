package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningOperation(ctx context.Context) error {
	// Do some long-running operation.
	time.Sleep(5 * time.Second)
	return nil
}

func main() {
	// Create a context with a deadline of 1 second.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Start the long-running operation in a goroutine.
	go longRunningOperation(ctx)

	// Wait for the operation to complete or for the context to be canceled.
	select {
	case <-ctx.Done():
		// The operation was canceled.
		fmt.Println("Operation was canceled.")
	case <-time.After(10 * time.Second):
		// The operation did not complete within the deadline.
		fmt.Println("Operation timed out.")
	}
}
