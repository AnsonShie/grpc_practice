// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: demo.proto

package demo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IngestServiceClient is the client API for IngestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IngestServiceClient interface {
	Ingest(ctx context.Context, in *IngestRequest, opts ...grpc.CallOption) (IngestService_IngestClient, error)
}

type ingestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIngestServiceClient(cc grpc.ClientConnInterface) IngestServiceClient {
	return &ingestServiceClient{cc}
}

func (c *ingestServiceClient) Ingest(ctx context.Context, in *IngestRequest, opts ...grpc.CallOption) (IngestService_IngestClient, error) {
	stream, err := c.cc.NewStream(ctx, &IngestService_ServiceDesc.Streams[0], "/demo.IngestService/Ingest", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingestServiceIngestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IngestService_IngestClient interface {
	Recv() (*IngestResponse, error)
	grpc.ClientStream
}

type ingestServiceIngestClient struct {
	grpc.ClientStream
}

func (x *ingestServiceIngestClient) Recv() (*IngestResponse, error) {
	m := new(IngestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IngestServiceServer is the server API for IngestService service.
// All implementations should embed UnimplementedIngestServiceServer
// for forward compatibility
type IngestServiceServer interface {
	Ingest(*IngestRequest, IngestService_IngestServer) error
}

// UnimplementedIngestServiceServer should be embedded to have forward compatible implementations.
type UnimplementedIngestServiceServer struct {
}

func (UnimplementedIngestServiceServer) Ingest(*IngestRequest, IngestService_IngestServer) error {
	return status.Errorf(codes.Unimplemented, "method Ingest not implemented")
}

// UnsafeIngestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IngestServiceServer will
// result in compilation errors.
type UnsafeIngestServiceServer interface {
	mustEmbedUnimplementedIngestServiceServer()
}

func RegisterIngestServiceServer(s grpc.ServiceRegistrar, srv IngestServiceServer) {
	s.RegisterService(&IngestService_ServiceDesc, srv)
}

func _IngestService_Ingest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IngestRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IngestServiceServer).Ingest(m, &ingestServiceIngestServer{stream})
}

type IngestService_IngestServer interface {
	Send(*IngestResponse) error
	grpc.ServerStream
}

type ingestServiceIngestServer struct {
	grpc.ServerStream
}

func (x *ingestServiceIngestServer) Send(m *IngestResponse) error {
	return x.ServerStream.SendMsg(m)
}

// IngestService_ServiceDesc is the grpc.ServiceDesc for IngestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IngestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.IngestService",
	HandlerType: (*IngestServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Ingest",
			Handler:       _IngestService_Ingest_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "demo.proto",
}

// RCAServiceClient is the client API for RCAService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RCAServiceClient interface {
	Firstview(ctx context.Context, in *RCARequest, opts ...grpc.CallOption) (*RCAResponse, error)
}

type rCAServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRCAServiceClient(cc grpc.ClientConnInterface) RCAServiceClient {
	return &rCAServiceClient{cc}
}

func (c *rCAServiceClient) Firstview(ctx context.Context, in *RCARequest, opts ...grpc.CallOption) (*RCAResponse, error) {
	out := new(RCAResponse)
	err := c.cc.Invoke(ctx, "/demo.RCAService/Firstview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RCAServiceServer is the server API for RCAService service.
// All implementations should embed UnimplementedRCAServiceServer
// for forward compatibility
type RCAServiceServer interface {
	Firstview(context.Context, *RCARequest) (*RCAResponse, error)
}

// UnimplementedRCAServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRCAServiceServer struct {
}

func (UnimplementedRCAServiceServer) Firstview(context.Context, *RCARequest) (*RCAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Firstview not implemented")
}

// UnsafeRCAServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RCAServiceServer will
// result in compilation errors.
type UnsafeRCAServiceServer interface {
	mustEmbedUnimplementedRCAServiceServer()
}

func RegisterRCAServiceServer(s grpc.ServiceRegistrar, srv RCAServiceServer) {
	s.RegisterService(&RCAService_ServiceDesc, srv)
}

func _RCAService_Firstview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RCARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RCAServiceServer).Firstview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.RCAService/Firstview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RCAServiceServer).Firstview(ctx, req.(*RCARequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RCAService_ServiceDesc is the grpc.ServiceDesc for RCAService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RCAService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.RCAService",
	HandlerType: (*RCAServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Firstview",
			Handler:    _RCAService_Firstview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// CounterServiceClient is the client API for CounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CounterServiceClient interface {
	Count(ctx context.Context, opts ...grpc.CallOption) (CounterService_CountClient, error)
}

type counterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCounterServiceClient(cc grpc.ClientConnInterface) CounterServiceClient {
	return &counterServiceClient{cc}
}

func (c *counterServiceClient) Count(ctx context.Context, opts ...grpc.CallOption) (CounterService_CountClient, error) {
	stream, err := c.cc.NewStream(ctx, &CounterService_ServiceDesc.Streams[0], "/demo.CounterService/Count", opts...)
	if err != nil {
		return nil, err
	}
	x := &counterServiceCountClient{stream}
	return x, nil
}

type CounterService_CountClient interface {
	Send(*CounterRequest) error
	Recv() (*CounterResponse, error)
	grpc.ClientStream
}

type counterServiceCountClient struct {
	grpc.ClientStream
}

func (x *counterServiceCountClient) Send(m *CounterRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *counterServiceCountClient) Recv() (*CounterResponse, error) {
	m := new(CounterResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CounterServiceServer is the server API for CounterService service.
// All implementations should embed UnimplementedCounterServiceServer
// for forward compatibility
type CounterServiceServer interface {
	Count(CounterService_CountServer) error
}

// UnimplementedCounterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCounterServiceServer struct {
}

func (UnimplementedCounterServiceServer) Count(CounterService_CountServer) error {
	return status.Errorf(codes.Unimplemented, "method Count not implemented")
}

// UnsafeCounterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CounterServiceServer will
// result in compilation errors.
type UnsafeCounterServiceServer interface {
	mustEmbedUnimplementedCounterServiceServer()
}

func RegisterCounterServiceServer(s grpc.ServiceRegistrar, srv CounterServiceServer) {
	s.RegisterService(&CounterService_ServiceDesc, srv)
}

func _CounterService_Count_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CounterServiceServer).Count(&counterServiceCountServer{stream})
}

type CounterService_CountServer interface {
	Send(*CounterResponse) error
	Recv() (*CounterRequest, error)
	grpc.ServerStream
}

type counterServiceCountServer struct {
	grpc.ServerStream
}

func (x *counterServiceCountServer) Send(m *CounterResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *counterServiceCountServer) Recv() (*CounterRequest, error) {
	m := new(CounterRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CounterService_ServiceDesc is the grpc.ServiceDesc for CounterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CounterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.CounterService",
	HandlerType: (*CounterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Count",
			Handler:       _CounterService_Count_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "demo.proto",
}
