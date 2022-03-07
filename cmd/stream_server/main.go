package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"sync"
	"time"

	"github.com/AnsonShie/grpc_practice/proto/demo"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	sum        int32
	quotient   int32
	mod        int32
	enableMod  bool
	modEnabled bool
	sync.RWMutex
}

func (s *server) EnableMod(ctx context.Context, req *demo.ModRequest) (*demo.ModResponse, error) {
	s.enableMod = req.GetEnable()
	return &demo.ModResponse{Response: fmt.Sprintf("Mod mode: %v", s.enableMod)}, nil
}

func (s *server) doMod(ctx context.Context) {
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			s.RLock()
			currentSum := s.sum
			s.RUnlock()
			if currentSum > s.mod {
				s.Lock()
				s.quotient = s.sum / s.mod
				s.sum = s.sum % s.mod
				s.Unlock()
			}
		}
	}
	s.modEnabled = false
	fmt.Println("mod goroutine exit")
}

func (s *server) Count(stream demo.CounterService_CountServer) error {
	ctx, _ := context.WithTimeout(stream.Context(), time.Second*20)
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			req, err := stream.Recv()
			if err != nil {
				fmt.Printf("Receive error: %v \n", err)
				return err
			}
			if s.enableMod && !s.modEnabled {
				go s.doMod(ctx)
				s.modEnabled = true
			}
			s.Lock()
			s.sum += req.GetInput()
			resp := &demo.CounterResponse{
				Sum:      s.sum,
				Quotient: s.quotient,
			}
			s.Unlock()
			if err := stream.Send(resp); err != nil {
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
		mod: int32(math.Pow10(4)),
	}
	demo.RegisterCounterServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
