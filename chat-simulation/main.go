package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a buffered channel for communication between users
	messageChannel := make(chan string, 10) // Buffered with a capacity of 10

	// Use a WaitGroup to wait for both users to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// User 1
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			message := fmt.Sprintf("User 1: Message %d", i)
			messageChannel <- message
			time.Sleep(time.Millisecond * 100) // Simulate a delay
		}
	}()

	// User 2
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			message := fmt.Sprintf("User 2: Message %d", i)
			messageChannel <- message
			time.Sleep(time.Millisecond * 200) // Simulate a delay
		}
	}()

	// Use a goroutine to receive and display messages
	go func() {
		for message := range messageChannel {
			fmt.Println(message)
		}
	}()

	// Wait for both users to finish
	wg.Wait()

	// Close the message channel when done
	close(messageChannel)

	// Wait for the display goroutine to finish
	time.Sleep(time.Second * 2) // Give it some time to process remaining messages
}
