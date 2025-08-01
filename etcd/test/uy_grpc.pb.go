// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: uy.proto

package proto

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

// GreetClient is the client API for Greet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetClient interface {
	Morning(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	Night(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}

type greetClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetClient(cc grpc.ClientConnInterface) GreetClient {
	return &greetClient{cc}
}

func (c *greetClient) Morning(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/Greet/Morning", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetClient) Night(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/Greet/Night", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServer is the server API for Greet service.
// All implementations should embed UnimplementedGreetServer
// for forward compatibility
type GreetServer interface {
	Morning(context.Context, *GreetRequest) (*GreetResponse, error)
	Night(context.Context, *GreetRequest) (*GreetResponse, error)
}

// UnimplementedGreetServer should be embedded to have forward compatible implementations.
type UnimplementedGreetServer struct {
}

func (UnimplementedGreetServer) Morning(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Morning not implemented")
}
func (UnimplementedGreetServer) Night(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Night not implemented")
}

// UnsafeGreetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServer will
// result in compilation errors.
type UnsafeGreetServer interface {
	mustEmbedUnimplementedGreetServer()
}

func RegisterGreetServer(s grpc.ServiceRegistrar, srv GreetServer) {
	s.RegisterService(&Greet_ServiceDesc, srv)
}

func _Greet_Morning_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServer).Morning(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greet/Morning",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServer).Morning(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greet_Night_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServer).Night(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greet/Night",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServer).Night(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greet_ServiceDesc is the grpc.ServiceDesc for Greet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Greet",
	HandlerType: (*GreetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Morning",
			Handler:    _Greet_Morning_Handler,
		},
		{
			MethodName: "Night",
			Handler:    _Greet_Night_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uy.proto",
}
