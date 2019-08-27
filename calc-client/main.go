package main

import (
	"log"

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
	// doClientStreaming(c)

	// uncomment to test bidirectional request/response streaming behavior
	doBiDirectionalStreaming(c)
}
