// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.0
// source: interaction.proto

package interactionProto

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

// InteractionServiceClient is the client API for InteractionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractionServiceClient interface {
	LikeVideo(ctx context.Context, in *LikeVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CancelLikeVideo(ctx context.Context, in *CancelLikeVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetLikeList(ctx context.Context, in *GetLikeListRequest, opts ...grpc.CallOption) (*FeedResponse, error)
	CollectVideo(ctx context.Context, in *CollectVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CancelCollectVideo(ctx context.Context, in *CancelCollectVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCollectList(ctx context.Context, in *GetCollectListRequest, opts ...grpc.CallOption) (*FeedResponse, error)
	Comment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCommentList(ctx context.Context, in *GetCommentListRequest, opts ...grpc.CallOption) (*GetCommentListResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type interactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractionServiceClient(cc grpc.ClientConnInterface) InteractionServiceClient {
	return &interactionServiceClient{cc}
}

func (c *interactionServiceClient) LikeVideo(ctx context.Context, in *LikeVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/interactionProto.InteractionService/LikeVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) CancelLikeVideo(ctx context.Context, in *CancelLikeVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/interactionProto.InteractionService/CancelLikeVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) GetLikeList(ctx context.Context, in *GetLikeListRequest, opts ...grpc.CallOption) (*FeedResponse, error) {
	out := new(FeedResponse)
	err := c.cc.Invoke(ctx, "/interactionProto.InteractionService/GetLikeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) CollectVideo(ctx context.Context, in *CollectVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/interactionProto.InteractionService/CollectVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) CancelCollectVideo(ctx context.Context, in *CancelCollectVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/interactionProto.InteractionService/CancelCollectVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) GetCollectList(ctx context.Context, in *GetCollectListRequest, opts ...grpc.CallOption) (*FeedResponse, error) {
	out := new(FeedResponse)
	err := c.cc.Invoke(ctx, "/interactionProto.InteractionService/GetCollectList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) Comment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/interactionProto.InteractionService/Comment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) GetCommentList(ctx context.Context, in *GetCommentListRequest, opts ...grpc.CallOption) (*GetCommentListResponse, error) {
	out := new(GetCommentListResponse)
	err := c.cc.Invoke(ctx, "/interactionProto.InteractionService/GetCommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/interactionProto.InteractionService/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractionServiceServer is the server API for InteractionService service.
// All implementations must embed UnimplementedInteractionServiceServer
// for forward compatibility
type InteractionServiceServer interface {
	LikeVideo(context.Context, *LikeVideoRequest) (*emptypb.Empty, error)
	CancelLikeVideo(context.Context, *CancelLikeVideoRequest) (*emptypb.Empty, error)
	GetLikeList(context.Context, *GetLikeListRequest) (*FeedResponse, error)
	CollectVideo(context.Context, *CollectVideoRequest) (*emptypb.Empty, error)
	CancelCollectVideo(context.Context, *CancelCollectVideoRequest) (*emptypb.Empty, error)
	GetCollectList(context.Context, *GetCollectListRequest) (*FeedResponse, error)
	Comment(context.Context, *CommentRequest) (*emptypb.Empty, error)
	GetCommentList(context.Context, *GetCommentListRequest) (*GetCommentListResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedInteractionServiceServer()
}

// UnimplementedInteractionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInteractionServiceServer struct {
}

func (UnimplementedInteractionServiceServer) LikeVideo(context.Context, *LikeVideoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeVideo not implemented")
}
func (UnimplementedInteractionServiceServer) CancelLikeVideo(context.Context, *CancelLikeVideoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelLikeVideo not implemented")
}
func (UnimplementedInteractionServiceServer) GetLikeList(context.Context, *GetLikeListRequest) (*FeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikeList not implemented")
}
func (UnimplementedInteractionServiceServer) CollectVideo(context.Context, *CollectVideoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectVideo not implemented")
}
func (UnimplementedInteractionServiceServer) CancelCollectVideo(context.Context, *CancelCollectVideoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelCollectVideo not implemented")
}
func (UnimplementedInteractionServiceServer) GetCollectList(context.Context, *GetCollectListRequest) (*FeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollectList not implemented")
}
func (UnimplementedInteractionServiceServer) Comment(context.Context, *CommentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comment not implemented")
}
func (UnimplementedInteractionServiceServer) GetCommentList(context.Context, *GetCommentListRequest) (*GetCommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentList not implemented")
}
func (UnimplementedInteractionServiceServer) DeleteComment(context.Context, *DeleteCommentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedInteractionServiceServer) mustEmbedUnimplementedInteractionServiceServer() {}

// UnsafeInteractionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractionServiceServer will
// result in compilation errors.
type UnsafeInteractionServiceServer interface {
	mustEmbedUnimplementedInteractionServiceServer()
}

func RegisterInteractionServiceServer(s grpc.ServiceRegistrar, srv InteractionServiceServer) {
	s.RegisterService(&InteractionService_ServiceDesc, srv)
}

func _InteractionService_LikeVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).LikeVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactionProto.InteractionService/LikeVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).LikeVideo(ctx, req.(*LikeVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_CancelLikeVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelLikeVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).CancelLikeVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactionProto.InteractionService/CancelLikeVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).CancelLikeVideo(ctx, req.(*CancelLikeVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_GetLikeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLikeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).GetLikeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactionProto.InteractionService/GetLikeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).GetLikeList(ctx, req.(*GetLikeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_CollectVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).CollectVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactionProto.InteractionService/CollectVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).CollectVideo(ctx, req.(*CollectVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_CancelCollectVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelCollectVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).CancelCollectVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactionProto.InteractionService/CancelCollectVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).CancelCollectVideo(ctx, req.(*CancelCollectVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_GetCollectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).GetCollectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactionProto.InteractionService/GetCollectList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).GetCollectList(ctx, req.(*GetCollectListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_Comment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).Comment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactionProto.InteractionService/Comment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).Comment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_GetCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).GetCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactionProto.InteractionService/GetCommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).GetCommentList(ctx, req.(*GetCommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactionProto.InteractionService/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServiceServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InteractionService_ServiceDesc is the grpc.ServiceDesc for InteractionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InteractionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interactionProto.InteractionService",
	HandlerType: (*InteractionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LikeVideo",
			Handler:    _InteractionService_LikeVideo_Handler,
		},
		{
			MethodName: "CancelLikeVideo",
			Handler:    _InteractionService_CancelLikeVideo_Handler,
		},
		{
			MethodName: "GetLikeList",
			Handler:    _InteractionService_GetLikeList_Handler,
		},
		{
			MethodName: "CollectVideo",
			Handler:    _InteractionService_CollectVideo_Handler,
		},
		{
			MethodName: "CancelCollectVideo",
			Handler:    _InteractionService_CancelCollectVideo_Handler,
		},
		{
			MethodName: "GetCollectList",
			Handler:    _InteractionService_GetCollectList_Handler,
		},
		{
			MethodName: "Comment",
			Handler:    _InteractionService_Comment_Handler,
		},
		{
			MethodName: "GetCommentList",
			Handler:    _InteractionService_GetCommentList_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _InteractionService_DeleteComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interaction.proto",
}
