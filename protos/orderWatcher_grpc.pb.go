// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.17.3
// source: orderWatcher.proto

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

// OrderWatcherClient is the client API for OrderWatcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderWatcherClient interface {
	Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (OrderWatcher_SubscribeClient, error)
	Unsubscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type orderWatcherClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderWatcherClient(cc grpc.ClientConnInterface) OrderWatcherClient {
	return &orderWatcherClient{cc}
}

func (c *orderWatcherClient) Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (OrderWatcher_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderWatcher_ServiceDesc.Streams[0], "/protos.OrderWatcher/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderWatcherSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderWatcher_SubscribeClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type orderWatcherSubscribeClient struct {
	grpc.ClientStream
}

func (x *orderWatcherSubscribeClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderWatcherClient) Unsubscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.OrderWatcher/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderWatcherServer is the server API for OrderWatcher service.
// All implementations must embed UnimplementedOrderWatcherServer
// for forward compatibility
type OrderWatcherServer interface {
	Subscribe(*Request, OrderWatcher_SubscribeServer) error
	Unsubscribe(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedOrderWatcherServer()
}

// UnimplementedOrderWatcherServer must be embedded to have forward compatible implementations.
type UnimplementedOrderWatcherServer struct {
}

func (UnimplementedOrderWatcherServer) Subscribe(*Request, OrderWatcher_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedOrderWatcherServer) Unsubscribe(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedOrderWatcherServer) mustEmbedUnimplementedOrderWatcherServer() {}

// UnsafeOrderWatcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderWatcherServer will
// result in compilation errors.
type UnsafeOrderWatcherServer interface {
	mustEmbedUnimplementedOrderWatcherServer()
}

func RegisterOrderWatcherServer(s grpc.ServiceRegistrar, srv OrderWatcherServer) {
	s.RegisterService(&OrderWatcher_ServiceDesc, srv)
}

func _OrderWatcher_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderWatcherServer).Subscribe(m, &orderWatcherSubscribeServer{stream})
}

type OrderWatcher_SubscribeServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type orderWatcherSubscribeServer struct {
	grpc.ServerStream
}

func (x *orderWatcherSubscribeServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderWatcher_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderWatcherServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.OrderWatcher/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderWatcherServer).Unsubscribe(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderWatcher_ServiceDesc is the grpc.ServiceDesc for OrderWatcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderWatcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.OrderWatcher",
	HandlerType: (*OrderWatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unsubscribe",
			Handler:    _OrderWatcher_Unsubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _OrderWatcher_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "orderWatcher.proto",
}
