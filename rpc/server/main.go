package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "rpc/calculator"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	result := req.A + req.B
	return &pb.AddResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	gs := grpc.NewServer()
	pb.RegisterCalculatorServer(gs, &server{})

	fmt.Println("Server is running on :50051...")
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
