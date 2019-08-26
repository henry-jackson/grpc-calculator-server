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
		log.Fatalf("Failed to dial: %s\n", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// uncomment to test unary gRPC API behavior
	// doUnary(c)

	// uncomment to test response streaming behavior
	// doServerStreaming(c)

	// uncomment to test request streaming behavior
	doClientStreaming(c)

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
		log.Fatalf("Failed to send calculation request: %s\n", err)
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
		log.Fatalf("Failed to stream responses: %s\n", err)
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

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	rand.Seed(time.Now().UnixNano()) // set our random seed

	stream, err := c.CalculateAverage(context.Background())
	if err != nil {
		log.Fatalf("Failed to stream requestse %s\n", err)
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
