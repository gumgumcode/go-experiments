package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "hi from my goroutine!" }()

	inbox := <-messages
	fmt.Println(inbox)
}

// $ go run channels-intro.go
// hi from my goroutine!
