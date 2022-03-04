package main

import (
	"fmt"
	"log"
	"net"

	"github.com/AnsonShie/grpc_practice/proto/demo"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	uuidMapSum map[string]int32
}

func (s *server) Count(stream demo.CounterService_CountServer) error {
	for {
		select {
		case <-stream.Context().Done():
			fmt.Println("cancel by ctx")
			return nil
		default:
			req, err := stream.Recv()
			if err != nil {
				fmt.Printf("Receive error: %v \n", err)
				return err
			}
			s.uuidMapSum[req.GetUuid()] += req.GetInput()
			if err := stream.Send(&demo.CounterResponse{Sum: s.uuidMapSum[req.GetUuid()]}); err != nil {
				fmt.Printf("Send error: %v \n", err)
				return err
			}
		}
	}
	return nil
}

var (
	customFunc grpc_recovery.RecoveryHandlerFunc
)

func main() {
	fmt.Println("starting gRPC server...")

	lis, err := net.Listen("tcp", "localhost:50049")
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
	server := server{
		uuidMapSum: make(map[string]int32),
	}
	demo.RegisterCounterServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
