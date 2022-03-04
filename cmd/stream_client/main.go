package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/AnsonShie/grpc_practice/proto/demo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("localhost:50049", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	client := demo.NewCounterServiceClient(conn)
	doCount(client)

}

func doCount(client demo.CounterServiceClient) {
	req := &demo.CounterRequest{
		Uuid: "c1111111-1111-1111-1111-11111111111d",
	}
	ctx, cancel := context.WithCancel(context.Background())
	stream, err := client.Count(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= 10; i = i + 1 {
		rand.Seed(time.Now().UnixNano())
		req.Input = rand.Int31n(10) + 1
		stream.Send(req)
		reply, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("reply : %v\n", reply)
		}
	}
	cancel()
	for i := 1; i <= 10; i = i + 1 {
		rand.Seed(time.Now().UnixNano())
		req.Input = rand.Int31n(10) + 1
		stream.Send(req)
		reply, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("reply : %v\n", reply)
		}
	}

}
