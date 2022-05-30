// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.0
// source: health.proto

package protos

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

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthServiceClient interface {
	GetSystemHealth(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	GetCpuData(ctx context.Context, opts ...grpc.CallOption) (HealthService_GetCpuDataClient, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) GetSystemHealth(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/main.HealthService/GetSystemHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) GetCpuData(ctx context.Context, opts ...grpc.CallOption) (HealthService_GetCpuDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &HealthService_ServiceDesc.Streams[0], "/main.HealthService/GetCpuData", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthServiceGetCpuDataClient{stream}
	return x, nil
}

type HealthService_GetCpuDataClient interface {
	Send(*CpuRequest) error
	Recv() (*CpuResponse, error)
	grpc.ClientStream
}

type healthServiceGetCpuDataClient struct {
	grpc.ClientStream
}

func (x *healthServiceGetCpuDataClient) Send(m *CpuRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *healthServiceGetCpuDataClient) Recv() (*CpuResponse, error) {
	m := new(CpuResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HealthServiceServer is the server API for HealthService service.
// All implementations must embed UnimplementedHealthServiceServer
// for forward compatibility
type HealthServiceServer interface {
	GetSystemHealth(context.Context, *HealthRequest) (*HealthResponse, error)
	GetCpuData(HealthService_GetCpuDataServer) error
	mustEmbedUnimplementedHealthServiceServer()
}

// UnimplementedHealthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHealthServiceServer struct {
}

func (UnimplementedHealthServiceServer) GetSystemHealth(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemHealth not implemented")
}
func (UnimplementedHealthServiceServer) GetCpuData(HealthService_GetCpuDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCpuData not implemented")
}
func (UnimplementedHealthServiceServer) mustEmbedUnimplementedHealthServiceServer() {}

// UnsafeHealthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServiceServer will
// result in compilation errors.
type UnsafeHealthServiceServer interface {
	mustEmbedUnimplementedHealthServiceServer()
}

func RegisterHealthServiceServer(s grpc.ServiceRegistrar, srv HealthServiceServer) {
	s.RegisterService(&HealthService_ServiceDesc, srv)
}

func _HealthService_GetSystemHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).GetSystemHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.HealthService/GetSystemHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).GetSystemHealth(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_GetCpuData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HealthServiceServer).GetCpuData(&healthServiceGetCpuDataServer{stream})
}

type HealthService_GetCpuDataServer interface {
	Send(*CpuResponse) error
	Recv() (*CpuRequest, error)
	grpc.ServerStream
}

type healthServiceGetCpuDataServer struct {
	grpc.ServerStream
}

func (x *healthServiceGetCpuDataServer) Send(m *CpuResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *healthServiceGetCpuDataServer) Recv() (*CpuRequest, error) {
	m := new(CpuRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HealthService_ServiceDesc is the grpc.ServiceDesc for HealthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSystemHealth",
			Handler:    _HealthService_GetSystemHealth_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCpuData",
			Handler:       _HealthService_GetCpuData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "health.proto",
}
