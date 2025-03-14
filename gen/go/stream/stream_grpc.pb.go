// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: stream/stream.proto

package stream

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	StreamService_NewStream_FullMethodName = "/stream.StreamService/NewStream"
)

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Stream 서비스를 정의합니다.
type StreamServiceClient interface {
	// 클라이언트가 심볼(symbol)을 요청하면, 서버가 250ms 간격으로 해당 심볼 데이터를 스트리밍으로 전송합니다.
	NewStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamResponse], error)
}

type streamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServiceClient(cc grpc.ClientConnInterface) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) NewStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[0], StreamService_NewStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamRequest, StreamResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StreamService_NewStreamClient = grpc.ServerStreamingClient[StreamResponse]

// StreamServiceServer is the server API for StreamService service.
// All implementations must embed UnimplementedStreamServiceServer
// for forward compatibility.
//
// Stream 서비스를 정의합니다.
type StreamServiceServer interface {
	// 클라이언트가 심볼(symbol)을 요청하면, 서버가 250ms 간격으로 해당 심볼 데이터를 스트리밍으로 전송합니다.
	NewStream(*StreamRequest, grpc.ServerStreamingServer[StreamResponse]) error
	mustEmbedUnimplementedStreamServiceServer()
}

// UnimplementedStreamServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStreamServiceServer struct{}

func (UnimplementedStreamServiceServer) NewStream(*StreamRequest, grpc.ServerStreamingServer[StreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method NewStream not implemented")
}
func (UnimplementedStreamServiceServer) mustEmbedUnimplementedStreamServiceServer() {}
func (UnimplementedStreamServiceServer) testEmbeddedByValue()                       {}

// UnsafeStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServiceServer will
// result in compilation errors.
type UnsafeStreamServiceServer interface {
	mustEmbedUnimplementedStreamServiceServer()
}

func RegisterStreamServiceServer(s grpc.ServiceRegistrar, srv StreamServiceServer) {
	// If the following call pancis, it indicates UnimplementedStreamServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StreamService_ServiceDesc, srv)
}

func _StreamService_NewStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).NewStream(m, &grpc.GenericServerStream[StreamRequest, StreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StreamService_NewStreamServer = grpc.ServerStreamingServer[StreamResponse]

// StreamService_ServiceDesc is the grpc.ServiceDesc for StreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NewStream",
			Handler:       _StreamService_NewStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stream/stream.proto",
}
