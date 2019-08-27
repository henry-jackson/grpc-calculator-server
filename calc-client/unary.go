package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/henry-jackson/grpc-calculator-server/calculatorpb"
)

func doUnary(c calculatorpb.CalculatorServiceClient) {
	rand.Seed(time.Now().UnixNano()) // set our random seed

	req := &calculatorpb.SumRequest{
		A: rand.Int31n(1000), // generate random int32 between 0-1000
		B: rand.Int31n(1000), // generate random int32 between 0-1000
	}

	fmt.Printf("Sending sum request: %v\n", req)

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to send calculation request: %s\n", err)
	}

	fmt.Printf("Got response: %v\n", res.GetResult())
}
