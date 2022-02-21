package main

import (
	"fmt"
	"log"
	"net"
	"time"

	demo "github.com/AnsonShie/grpc_practice/proto/demo"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	ips []string
}

func (s *Server) Ingest(req *demo.IngestRequest, resp demo.IngestService_IngestServer) error {
	fmt.Printf("Sum function is invoked with %v \n", req)

	//a := req.GetUuid()

	for _, ip := range s.ips {
		if ip == "3.3.3.3" {
			panic(ip)
		}
		res := &demo.IngestResponse{
			Result: ip,
		}
		if err := resp.Send(res); err != nil {
			fmt.Println(err)
			return err
		}
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

var (
	customFunc grpc_recovery.RecoveryHandlerFunc
)

func main() {
	fmt.Println("starting gRPC server...")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	customFunc = func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(customFunc),
	}

	grpcServer := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_recovery.StreamServerInterceptor(opts...),
	)))
	server := Server{
		ips: []string{"1.1.1.1", "2.2.2.2", "3.3.3.3"},
	}
	demo.RegisterIngestServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
