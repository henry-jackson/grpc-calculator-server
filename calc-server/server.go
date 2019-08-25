package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/henry-jackson/grpc-calculator-server/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, in *calculatorpb.SumRequest) (out *calculatorpb.SumResponse, err error) {
	fmt.Printf("Received request: %v\n", in)
	a := in.GetA()
	b := in.GetB()
	sum := a + b
	out = &calculatorpb.SumResponse{
		Result: sum,
	}
	fmt.Printf("Sending response: %v\n", out)
	return out, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
