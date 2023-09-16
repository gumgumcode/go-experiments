package main

import "sync"

type Message struct {
	User    string
	Content string
}

func main() {
	sendCh := make(chan Message)
	receiveCh := make(chan Message)

	var wg = sync.WaitGroup

	go chatRoom(sendCh, receiveCh)

	users := []string{"user1", "user2", "user3"}

}
