package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/henry-jackson/grpc-calculator-server/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // set our random seed

	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	req := &calculatorpb.SumRequest{
		A: rand.Int31n(1000), // generate random int32 between 0-1000
		B: rand.Int31n(1000), // generate random int32 between 0-1000
	}

	fmt.Printf("Sending sum request: %v\n", req)

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to send calculation request: %s", err)
	}

	fmt.Printf("Get response: %v", res.GetResult())
}
