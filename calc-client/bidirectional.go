package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/henry-jackson/grpc-calculator-server/calculatorpb"
)

func doBiDirectionalStreaming(c calculatorpb.CalculatorServiceClient) {
	// place to store our max return val
	var max int32

	rand.Seed(time.Now().UnixNano()) // set our random seed

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Failed to create stream: %v\n", err)
	}

	waitChan := make(chan struct{})
	// start sending requests
	go func() {
		for i := 0; i < 100; i++ {
			stream.Send(&calculatorpb.FindMaximumRequest{
				Input: rand.Int31n(10000),
			})
			time.Sleep(time.Millisecond * 25)
		}
		// send a closing message to server (EOF)
		stream.CloseSend()
	}()

	// start receiving requests
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Unexpected error when streaming responses: %v\n", err)
				break
			}
			log.Printf("New maximum number: %v\n", res.GetCurrentMax())
			max = res.GetCurrentMax()
		}
		close(waitChan)
	}()

	// wait for receives to close
	<-waitChan
	log.Printf("Maximum number of all random inputs was: %v\n", max)
}
