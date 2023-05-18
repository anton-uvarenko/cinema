// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: comment.proto

package users

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

// CommentsClient is the client API for Comments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentsClient interface {
	AddComment(ctx context.Context, in *CommentPayload, opts ...grpc.CallOption) (*Comment, error)
	GetComments(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CommentsResponse, error)
	LikeComment(ctx context.Context, in *LikeCommentPayload, opts ...grpc.CallOption) (*Empty, error)
}

type commentsClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentsClient(cc grpc.ClientConnInterface) CommentsClient {
	return &commentsClient{cc}
}

func (c *commentsClient) AddComment(ctx context.Context, in *CommentPayload, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, "/users.Comments/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsClient) GetComments(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CommentsResponse, error) {
	out := new(CommentsResponse)
	err := c.cc.Invoke(ctx, "/users.Comments/GetComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsClient) LikeComment(ctx context.Context, in *LikeCommentPayload, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/users.Comments/LikeComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentsServer is the server API for Comments service.
// All implementations must embed UnimplementedCommentsServer
// for forward compatibility
type CommentsServer interface {
	AddComment(context.Context, *CommentPayload) (*Comment, error)
	GetComments(context.Context, *Empty) (*CommentsResponse, error)
	LikeComment(context.Context, *LikeCommentPayload) (*Empty, error)
	mustEmbedUnimplementedCommentsServer()
}

// UnimplementedCommentsServer must be embedded to have forward compatible implementations.
type UnimplementedCommentsServer struct {
}

func (UnimplementedCommentsServer) AddComment(context.Context, *CommentPayload) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedCommentsServer) GetComments(context.Context, *Empty) (*CommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedCommentsServer) LikeComment(context.Context, *LikeCommentPayload) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeComment not implemented")
}
func (UnimplementedCommentsServer) mustEmbedUnimplementedCommentsServer() {}

// UnsafeCommentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentsServer will
// result in compilation errors.
type UnsafeCommentsServer interface {
	mustEmbedUnimplementedCommentsServer()
}

func RegisterCommentsServer(s grpc.ServiceRegistrar, srv CommentsServer) {
	s.RegisterService(&Comments_ServiceDesc, srv)
}

func _Comments_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Comments/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServer).AddComment(ctx, req.(*CommentPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comments_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Comments/GetComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServer).GetComments(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comments_LikeComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeCommentPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServer).LikeComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Comments/LikeComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServer).LikeComment(ctx, req.(*LikeCommentPayload))
	}
	return interceptor(ctx, in, info, handler)
}

// Comments_ServiceDesc is the grpc.ServiceDesc for Comments service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comments_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.Comments",
	HandlerType: (*CommentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddComment",
			Handler:    _Comments_AddComment_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _Comments_GetComments_Handler,
		},
		{
			MethodName: "LikeComment",
			Handler:    _Comments_LikeComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
