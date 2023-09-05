package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer client.Close()

	request := "Omkar"
	var response string
	err = client.Call("GreetingService.SayHello", &request, &response)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Server says: ", response)
}
