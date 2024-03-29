// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EulerServiceClient is the client API for EulerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EulerServiceClient interface {
	Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error)
}

type eulerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEulerServiceClient(cc grpc.ClientConnInterface) EulerServiceClient {
	return &eulerServiceClient{cc}
}

func (c *eulerServiceClient) Solve(ctx context.Context, in *SolveRequest, opts ...grpc.CallOption) (*SolveResponse, error) {
	out := new(SolveResponse)
	err := c.cc.Invoke(ctx, "/euler.EulerService/Solve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EulerServiceServer is the server API for EulerService service.
// All implementations should embed UnimplementedEulerServiceServer
// for forward compatibility
type EulerServiceServer interface {
	Solve(context.Context, *SolveRequest) (*SolveResponse, error)
}

// UnimplementedEulerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEulerServiceServer struct {
}

func (UnimplementedEulerServiceServer) Solve(context.Context, *SolveRequest) (*SolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Solve not implemented")
}

// UnsafeEulerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EulerServiceServer will
// result in compilation errors.
type UnsafeEulerServiceServer interface {
	mustEmbedUnimplementedEulerServiceServer()
}

func RegisterEulerServiceServer(s grpc.ServiceRegistrar, srv EulerServiceServer) {
	s.RegisterService(&_EulerService_serviceDesc, srv)
}

func _EulerService_Solve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EulerServiceServer).Solve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/euler.EulerService/Solve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EulerServiceServer).Solve(ctx, req.(*SolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EulerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "euler.EulerService",
	HandlerType: (*EulerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Solve",
			Handler:    _EulerService_Solve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "euler.proto",
}
