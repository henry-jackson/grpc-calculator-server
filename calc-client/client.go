package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/henry-jackson/grpc-calculator-server/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// uncomment to test unary gRPC API behavior
	// doUnary(c)

	// uncomment to test response streaming behavior
	doServerStreaming(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	rand.Seed(time.Now().UnixNano()) // set our random seed

	req := &calculatorpb.SumRequest{
		A: rand.Int31n(1000), // generate random int32 between 0-1000
		B: rand.Int31n(1000), // generate random int32 between 0-1000
	}

	fmt.Printf("Sending sum request: %v\n", req)

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to send calculation request: %s", err)
	}

	fmt.Printf("Got response: %v\n", res.GetResult())
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	rand.Seed(time.Now().UnixNano()) // set our random seed

	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Input: rand.Int31n(1000000), // generate random int32 between 0 and 1 million
	}

	fmt.Printf("Sending sum request: %v\n", req)

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to connect stream server response: %s", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("reached end of stream responses")
			break
		}
		if err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}
		fmt.Printf("Got prime factor of: %v\n", res.GetResult())
	}
}
