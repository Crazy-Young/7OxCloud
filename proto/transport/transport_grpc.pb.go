// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.0
// source: transport.proto

package transportProto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TransportServiceClient is the client API for TransportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransportServiceClient interface {
	SendCSVFile(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRecommendedFeed(ctx context.Context, in *GetRecommendedFeedRequest, opts ...grpc.CallOption) (*GetRecommendedFeedResponse, error)
}

type transportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportServiceClient(cc grpc.ClientConnInterface) TransportServiceClient {
	return &transportServiceClient{cc}
}

func (c *transportServiceClient) SendCSVFile(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/transportProto.TransportService/SendCSVFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportServiceClient) GetRecommendedFeed(ctx context.Context, in *GetRecommendedFeedRequest, opts ...grpc.CallOption) (*GetRecommendedFeedResponse, error) {
	out := new(GetRecommendedFeedResponse)
	err := c.cc.Invoke(ctx, "/transportProto.TransportService/GetRecommendedFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportServiceServer is the server API for TransportService service.
// All implementations must embed UnimplementedTransportServiceServer
// for forward compatibility
type TransportServiceServer interface {
	SendCSVFile(context.Context, *SendRequest) (*emptypb.Empty, error)
	GetRecommendedFeed(context.Context, *GetRecommendedFeedRequest) (*GetRecommendedFeedResponse, error)
	mustEmbedUnimplementedTransportServiceServer()
}

// UnimplementedTransportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransportServiceServer struct {
}

func (UnimplementedTransportServiceServer) SendCSVFile(context.Context, *SendRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCSVFile not implemented")
}
func (UnimplementedTransportServiceServer) GetRecommendedFeed(context.Context, *GetRecommendedFeedRequest) (*GetRecommendedFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendedFeed not implemented")
}
func (UnimplementedTransportServiceServer) mustEmbedUnimplementedTransportServiceServer() {}

// UnsafeTransportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransportServiceServer will
// result in compilation errors.
type UnsafeTransportServiceServer interface {
	mustEmbedUnimplementedTransportServiceServer()
}

func RegisterTransportServiceServer(s grpc.ServiceRegistrar, srv TransportServiceServer) {
	s.RegisterService(&TransportService_ServiceDesc, srv)
}

func _TransportService_SendCSVFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).SendCSVFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transportProto.TransportService/SendCSVFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).SendCSVFile(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportService_GetRecommendedFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendedFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).GetRecommendedFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transportProto.TransportService/GetRecommendedFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).GetRecommendedFeed(ctx, req.(*GetRecommendedFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransportService_ServiceDesc is the grpc.ServiceDesc for TransportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transportProto.TransportService",
	HandlerType: (*TransportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCSVFile",
			Handler:    _TransportService_SendCSVFile_Handler,
		},
		{
			MethodName: "GetRecommendedFeed",
			Handler:    _TransportService_GetRecommendedFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport.proto",
}