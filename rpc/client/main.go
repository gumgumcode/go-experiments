package main

import (
	"context"
	"fmt"
	"log"

	pb "rpc/calculator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)

	req := &pb.AddRequest{
		A: 10,
		B: 5,
	}

	resp, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("RPC failed: %v", err)
	}

	fmt.Printf("Result: %d\n", resp.Result)
}
