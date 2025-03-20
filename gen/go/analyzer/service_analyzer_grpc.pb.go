// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: service_analyzer.proto

package analyzer

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
	Analyzer_Start_FullMethodName = "/analyzer.Analyzer/Start"
	Analyzer_Stop_FullMethodName  = "/analyzer.Analyzer/Stop"
)

// AnalyzerClient is the client API for Analyzer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyzerClient interface {
	Start(ctx context.Context, in *StartAnalyzerRequest, opts ...grpc.CallOption) (*StartAnalyzerResponse, error)
	Stop(ctx context.Context, in *StopAnalyzerRequest, opts ...grpc.CallOption) (*StopAnalyzerResponse, error)
}

type analyzerClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyzerClient(cc grpc.ClientConnInterface) AnalyzerClient {
	return &analyzerClient{cc}
}

func (c *analyzerClient) Start(ctx context.Context, in *StartAnalyzerRequest, opts ...grpc.CallOption) (*StartAnalyzerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartAnalyzerResponse)
	err := c.cc.Invoke(ctx, Analyzer_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) Stop(ctx context.Context, in *StopAnalyzerRequest, opts ...grpc.CallOption) (*StopAnalyzerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopAnalyzerResponse)
	err := c.cc.Invoke(ctx, Analyzer_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyzerServer is the server API for Analyzer service.
// All implementations must embed UnimplementedAnalyzerServer
// for forward compatibility.
type AnalyzerServer interface {
	Start(context.Context, *StartAnalyzerRequest) (*StartAnalyzerResponse, error)
	Stop(context.Context, *StopAnalyzerRequest) (*StopAnalyzerResponse, error)
	mustEmbedUnimplementedAnalyzerServer()
}

// UnimplementedAnalyzerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAnalyzerServer struct{}

func (UnimplementedAnalyzerServer) Start(context.Context, *StartAnalyzerRequest) (*StartAnalyzerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedAnalyzerServer) Stop(context.Context, *StopAnalyzerRequest) (*StopAnalyzerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedAnalyzerServer) mustEmbedUnimplementedAnalyzerServer() {}
func (UnimplementedAnalyzerServer) testEmbeddedByValue()                  {}

// UnsafeAnalyzerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyzerServer will
// result in compilation errors.
type UnsafeAnalyzerServer interface {
	mustEmbedUnimplementedAnalyzerServer()
}

func RegisterAnalyzerServer(s grpc.ServiceRegistrar, srv AnalyzerServer) {
	// If the following call pancis, it indicates UnimplementedAnalyzerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Analyzer_ServiceDesc, srv)
}

func _Analyzer_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartAnalyzerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Analyzer_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).Start(ctx, req.(*StartAnalyzerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopAnalyzerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Analyzer_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).Stop(ctx, req.(*StopAnalyzerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Analyzer_ServiceDesc is the grpc.ServiceDesc for Analyzer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Analyzer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "analyzer.Analyzer",
	HandlerType: (*AnalyzerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Analyzer_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Analyzer_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_analyzer.proto",
}
