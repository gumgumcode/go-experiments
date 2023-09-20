package main

import (
	"fmt"
	"time"
)

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	go foo()
	go bar()

	// Wait for some time to allow goroutines to execute
	time.Sleep(time.Second * 2)
}
