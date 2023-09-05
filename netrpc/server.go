package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type GreetingService struct{}

func (s *GreetingService) SayHello(request *string, response *string) error {
	*response = "Hello, " + *request + "!"
	return nil
}

func main() {
	gs := new(GreetingService)
	rpc.Register(gs)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
