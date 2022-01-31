// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: hello/hellopb/hello.proto

package hellopb

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

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	// Unary
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// Server Streaming
	HelloManyTimes(ctx context.Context, in *HelloManyTimesRequest, opts ...grpc.CallOption) (HelloService_HelloManyTimesClient, error)
	// Client Streaming
	LongHello(ctx context.Context, opts ...grpc.CallOption) (HelloService_LongHelloClient, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/greet.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) HelloManyTimes(ctx context.Context, in *HelloManyTimesRequest, opts ...grpc.CallOption) (HelloService_HelloManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], "/greet.HelloService/HelloManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceHelloManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_HelloManyTimesClient interface {
	Recv() (*HelloManyTimesResponse, error)
	grpc.ClientStream
}

type helloServiceHelloManyTimesClient struct {
	grpc.ClientStream
}

func (x *helloServiceHelloManyTimesClient) Recv() (*HelloManyTimesResponse, error) {
	m := new(HelloManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) LongHello(ctx context.Context, opts ...grpc.CallOption) (HelloService_LongHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[1], "/greet.HelloService/LongHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceLongHelloClient{stream}
	return x, nil
}

type HelloService_LongHelloClient interface {
	Send(*LongHelloRequest) error
	CloseAndRecv() (*LongHelloResponse, error)
	grpc.ClientStream
}

type helloServiceLongHelloClient struct {
	grpc.ClientStream
}

func (x *helloServiceLongHelloClient) Send(m *LongHelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceLongHelloClient) CloseAndRecv() (*LongHelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LongHelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	// Unary
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	// Server Streaming
	HelloManyTimes(*HelloManyTimesRequest, HelloService_HelloManyTimesServer) error
	// Client Streaming
	LongHello(HelloService_LongHelloServer) error
	//mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedHelloServiceServer) HelloManyTimes(*HelloManyTimesRequest, HelloService_HelloManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloManyTimes not implemented")
}
func (UnimplementedHelloServiceServer) LongHello(HelloService_LongHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method LongHello not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_HelloManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).HelloManyTimes(m, &helloServiceHelloManyTimesServer{stream})
}

type HelloService_HelloManyTimesServer interface {
	Send(*HelloManyTimesResponse) error
	grpc.ServerStream
}

type helloServiceHelloManyTimesServer struct {
	grpc.ServerStream
}

func (x *helloServiceHelloManyTimesServer) Send(m *HelloManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloService_LongHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).LongHello(&helloServiceLongHelloServer{stream})
}

type HelloService_LongHelloServer interface {
	SendAndClose(*LongHelloResponse) error
	Recv() (*LongHelloRequest, error)
	grpc.ServerStream
}

type helloServiceLongHelloServer struct {
	grpc.ServerStream
}

func (x *helloServiceLongHelloServer) SendAndClose(m *LongHelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceLongHelloServer) Recv() (*LongHelloRequest, error) {
	m := new(LongHelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloManyTimes",
			Handler:       _HelloService_HelloManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LongHello",
			Handler:       _HelloService_LongHello_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "hello/hellopb/hello.proto",
}
