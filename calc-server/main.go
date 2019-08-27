package main

import (
	"log"
	"net"

	"github.com/henry-jackson/grpc-calculator-server/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %s\n", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	log.Println("Server running at http://localhost:50051 ...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s\n", err)
	}
}
