package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	demo "github.com/AnsonShie/grpc_practice/proto/demo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientDeadline := time.Now().Add(time.Duration(100) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()
	defer cancel()

	client := demo.NewIngestServiceClient(conn)

	doUnary(client)
}

func doUnary(client demo.IngestServiceClient) {
	fmt.Println("Staring to do a Unary RPC")
	req := &demo.IngestRequest{
		Uuid: 123,
	}

	stream, err := client.Ingest(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling CalculatorService: %v \n", err)
	}

	for {
		resp, err := stream.Recv()

		// means we have finished receiving
		if err == io.EOF {
			break
		}

		// error handling
		if err != nil {
			log.Fatalf("%v.Ingest(_) = _, %v", client, err)
		}

		// print the response
		log.Println(resp)
	}
}
