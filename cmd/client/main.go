package main

import (
	"context"
	"io"
	"log"
	"sync"

	demo "github.com/AnsonShie/grpc_practice/proto/demo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ipInfo struct {
	customerId string
	companyId  string
}

func main() {
	ipMap := map[string]ipInfo{
		"2.2.2.2": {
			companyId:  "123",
			customerId: "456",
		},
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	conn2, err := grpc.Dial("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	client := demo.NewIngestServiceClient(conn)
	client2 := demo.NewRCAServiceClient(conn2)

	result := doIngestUnary(client)
	for ip := range result {
		if _, ok := ipMap[ip]; ok {
			doFirstviewUnary(client2, ipMap[ip])
		}
	}
}

func doFirstviewUnary(client demo.RCAServiceClient, info ipInfo) {
	req := &demo.RCARequest{
		CustomerId: info.customerId,
		CompanyId:  info.companyId,
	}
	res, err := client.Firstview(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to call firstview: %v", err)
	}
	log.Printf("response from firstview: %v", res)
}

func doIngestUnary(client demo.IngestServiceClient) <-chan string {
	ch := make(chan string)
	req := &demo.IngestRequest{
		Uuid: 123,
	}

	stream, err := client.Ingest(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling CalculatorService: %v \n", err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
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
			log.Printf("response from Ingest: %v", resp)

			ch <- resp.Result
		}
	}()
	go func() {
		wg.Wait()
		defer close(ch)
	}()
	return ch
}
