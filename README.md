# grpc-calculator-server
> implements a basic gRPC API server that has unary, server streaming, client
> streaming, and bidirectional streaming methods along with sample client side
> code.

## How to Run
> Assumes you have already installed protoc, grpc-go, the Go language, etc. See [this
> repo](https://github.com/protocolbuffers/protobuf) for any help.
#### Server
In one session:
```sh
go run calc-server/server.go
```

#### Client
> Client behavior can be altered by uncommenting certain lines in main() of
> client.go
In another session:
```sh
go run calc-client/client.go
```

## Example
> See screenshots below for example with server running on left and client
> executing on right.

#### Unary - Sum two numbers
![unary-screenshot](images/unary-screenshot.png)


#### Server streaming - Find prime factors of number
![sever-streaming-screenshot](images/server-streaming.gif)

#### Client streaming - Calculate average of streamed inputs
![client-streaming-screenshot](images/client-streaming.gif)
