// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: password-verify.proto

package auth

import (
	context "context"
	general "github.com/anton-uvarenko/cinema/broker-service/protobufs/general"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PassVerifyClient is the client API for PassVerify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PassVerifyClient interface {
	SendRecoveryCode(ctx context.Context, in *EmailPayload, opts ...grpc.CallOption) (*general.Empty, error)
	VerifyRecoveryCode(ctx context.Context, in *CodePayload, opts ...grpc.CallOption) (*general.JwtResponse, error)
	UpdatePassword(ctx context.Context, in *PasswordPayload, opts ...grpc.CallOption) (*general.Empty, error)
}

type passVerifyClient struct {
	cc grpc.ClientConnInterface
}

func NewPassVerifyClient(cc grpc.ClientConnInterface) PassVerifyClient {
	return &passVerifyClient{cc}
}

func (c *passVerifyClient) SendRecoveryCode(ctx context.Context, in *EmailPayload, opts ...grpc.CallOption) (*general.Empty, error) {
	out := new(general.Empty)
	err := c.cc.Invoke(ctx, "/auth.PassVerify/SendRecoveryCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passVerifyClient) VerifyRecoveryCode(ctx context.Context, in *CodePayload, opts ...grpc.CallOption) (*general.JwtResponse, error) {
	out := new(general.JwtResponse)
	err := c.cc.Invoke(ctx, "/auth.PassVerify/VerifyRecoveryCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passVerifyClient) UpdatePassword(ctx context.Context, in *PasswordPayload, opts ...grpc.CallOption) (*general.Empty, error) {
	out := new(general.Empty)
	err := c.cc.Invoke(ctx, "/auth.PassVerify/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PassVerifyServer is the server API for PassVerify service.
// All implementations must embed UnimplementedPassVerifyServer
// for forward compatibility
type PassVerifyServer interface {
	SendRecoveryCode(context.Context, *EmailPayload) (*general.Empty, error)
	VerifyRecoveryCode(context.Context, *CodePayload) (*general.JwtResponse, error)
	UpdatePassword(context.Context, *PasswordPayload) (*general.Empty, error)
	mustEmbedUnimplementedPassVerifyServer()
}

// UnimplementedPassVerifyServer must be embedded to have forward compatible implementations.
type UnimplementedPassVerifyServer struct {
}

func (UnimplementedPassVerifyServer) SendRecoveryCode(context.Context, *EmailPayload) (*general.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRecoveryCode not implemented")
}
func (UnimplementedPassVerifyServer) VerifyRecoveryCode(context.Context, *CodePayload) (*general.JwtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyRecoveryCode not implemented")
}
func (UnimplementedPassVerifyServer) UpdatePassword(context.Context, *PasswordPayload) (*general.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedPassVerifyServer) mustEmbedUnimplementedPassVerifyServer() {}

// UnsafePassVerifyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PassVerifyServer will
// result in compilation errors.
type UnsafePassVerifyServer interface {
	mustEmbedUnimplementedPassVerifyServer()
}

func RegisterPassVerifyServer(s grpc.ServiceRegistrar, srv PassVerifyServer) {
	s.RegisterService(&PassVerify_ServiceDesc, srv)
}

func _PassVerify_SendRecoveryCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassVerifyServer).SendRecoveryCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.PassVerify/SendRecoveryCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassVerifyServer).SendRecoveryCode(ctx, req.(*EmailPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassVerify_VerifyRecoveryCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassVerifyServer).VerifyRecoveryCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.PassVerify/VerifyRecoveryCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassVerifyServer).VerifyRecoveryCode(ctx, req.(*CodePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassVerify_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassVerifyServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.PassVerify/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassVerifyServer).UpdatePassword(ctx, req.(*PasswordPayload))
	}
	return interceptor(ctx, in, info, handler)
}

// PassVerify_ServiceDesc is the grpc.ServiceDesc for PassVerify service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PassVerify_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.PassVerify",
	HandlerType: (*PassVerifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRecoveryCode",
			Handler:    _PassVerify_SendRecoveryCode_Handler,
		},
		{
			MethodName: "VerifyRecoveryCode",
			Handler:    _PassVerify_VerifyRecoveryCode_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _PassVerify_UpdatePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "password-verify.proto",
}
