package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/henry-jackson/grpc-calculator-server/calculatorpb"
)

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	rand.Seed(time.Now().UnixNano()) // set our random seed

	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Input: rand.Int31n(1000000), // generate random int32 between 0 and 1 million
	}

	fmt.Printf("Sending sum request: %v\n", req)

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to stream responses: %v\n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("reached end of stream responses")
			break
		}
		if err != nil {
			log.Fatalf("Unexpected error: %s\n", err)
		}
		fmt.Printf("Got prime factor of: %v\n", res.GetResult())
	}
}
