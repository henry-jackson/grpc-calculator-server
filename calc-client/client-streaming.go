package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/henry-jackson/grpc-calculator-server/calculatorpb"
)

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	rand.Seed(time.Now().UnixNano()) // set our random seed

	stream, err := c.CalculateAverage(context.Background())
	if err != nil {
		log.Fatalf("Failed to stream request: %v\n", err)
	}

	for i := 0; i < 5; i++ {
		err := stream.Send(&calculatorpb.CalculateAverageRequest{
			Input: rand.Int31n(100),
		})
		if err != nil {
			log.Fatalf("Failed to send request: %s\n", err)
		}
		// simulate delay between inputs
		time.Sleep(time.Millisecond * 200)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to close and receive response: %s\n", err)
	}

	fmt.Printf("Got response from CalculateAverage rpc: %v\n", response)
}
