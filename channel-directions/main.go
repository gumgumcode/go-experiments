package main

import "fmt"

// process1 func accepts "send-only" channel
// represented by array going INTO the keyword "chan"
func process1(p1 chan<- string, msg string) {
	p1 <- msg
}

// process2 func accepts both "receive-only" and "send-only" channels
// "receive-only" is represented by array going OUT of the keyword "chan"
// "send-only" is represented by array going INTO the keyword "chan"
func process2(p1 <-chan string, p2 chan<- string) {
	transfer := <-p1
	p2 <- transfer
}

func main() {
	pipe1 := make(chan string, 1)
	pipe2 := make(chan string, 1)

	// send payload to pipe
	process1(pipe1, "payload")

	// transfer from pipe1 to pipe2
	process2(pipe1, pipe2)

	// receive from pipe2
	fmt.Println(<-pipe2)
}
