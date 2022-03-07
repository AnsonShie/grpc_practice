package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/AnsonShie/grpc_practice/proto/demo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.Dial("localhost:50049", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	// create a channel
	client := demo.NewCounterServiceClient(conn)
	doCount(client, 10)

}

func doCount(client demo.CounterServiceClient, timeout time.Duration) {
	req := &demo.CounterRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	metaCtx := metadata.AppendToOutgoingContext(ctx, "uuid", "c1111111-1111-1111-1111-11111111111d")
	stream, err := client.Count(metaCtx)
	go func() {
		time.Sleep(time.Second * 2)
		resp, err := client.EnableMod(ctx, &demo.ModRequest{Enable: true})
		if err != nil {
			log.Fatalf("failed to enable mod: %v", err)
		} else {
			fmt.Println(resp.Response)
		}
		time.Sleep(time.Second * 2)
		resp, err = client.EnableMod(ctx, &demo.ModRequest{Enable: false})
		if err != nil {
			log.Fatalf("failed to enable mod: %v", err)
		} else {
			fmt.Println(resp.Response)
		}
		time.Sleep(time.Second * 2)
		resp, err = client.EnableMod(ctx, &demo.ModRequest{Enable: true})
		if err != nil {
			log.Fatalf("failed to enable mod: %v", err)
		} else {
			fmt.Println(resp.Response)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		time.Sleep(time.Second * 8)
		stream.CloseSend()
	}()
	for {
		rand.Seed(time.Now().UnixNano())
		req.Input = rand.Int31n(100000) + 5
		stream.Send(req)
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("reply : %v\n", reply)
		}
		time.Sleep(time.Millisecond * 500)
	}
}
