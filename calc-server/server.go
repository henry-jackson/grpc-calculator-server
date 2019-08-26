package main

import (
	"context"
	"fmt"
	"io"
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

func (s *server) PrimeNumberDecomposition(in *calculatorpb.PrimeNumberDecompositionRequest, srv calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Received request: %v\n", in)
	n := in.GetInput()

	var divisor int32 = 2
	for n > 1 {
		if n%divisor == 0 {
			// this is a factor, send a response
			srv.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				Result: divisor,
			})
			n = n / divisor // divide so that we have the rest of the number left.
		} else {
			divisor++
			fmt.Printf("Divisor is: %v\n", divisor)
		}
	}
	return nil
}

func (s *server) CalculateAverage(stream calculatorpb.CalculatorService_CalculateAverageServer) error {
	// slice to store all streamed inputs
	nums := []float64{}
	for {
		req, err := stream.Recv()
		fmt.Printf("Received request: %v\n", req.GetInput())
		// if client says no more data is coming, calculate our response and
		// send it
		if err == io.EOF {
			var sum float64
			for _, num := range nums {
				sum += num
			}
			avg := sum / float64(len(nums))
			return stream.SendAndClose(&calculatorpb.CalculateAverageResponse{
				Result: avg,
			})
		}
		// check that err was nil
		if err != nil {
			log.Fatalf("Unexpected error from client stream: %v\n", err)
		}
		// grab input
		nums = append(nums, float64(req.GetInput()))
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %s\n", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s\n", err)
	}
}
