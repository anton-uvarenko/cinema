// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: user_data.proto

package users

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

// UserDataUploaderClient is the client API for UserDataUploader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserDataUploaderClient interface {
	UploadData(ctx context.Context, opts ...grpc.CallOption) (UserDataUploader_UploadDataClient, error)
	GetData(ctx context.Context, in *GetDataPayload, opts ...grpc.CallOption) (*GetDataResponse, error)
	UpdateData(ctx context.Context, opts ...grpc.CallOption) (UserDataUploader_UpdateDataClient, error)
	DeleteData(ctx context.Context, in *DeleteDataPayload, opts ...grpc.CallOption) (*general.Empty, error)
}

type userDataUploaderClient struct {
	cc grpc.ClientConnInterface
}

func NewUserDataUploaderClient(cc grpc.ClientConnInterface) UserDataUploaderClient {
	return &userDataUploaderClient{cc}
}

func (c *userDataUploaderClient) UploadData(ctx context.Context, opts ...grpc.CallOption) (UserDataUploader_UploadDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserDataUploader_ServiceDesc.Streams[0], "/users.UserDataUploader/UploadData", opts...)
	if err != nil {
		return nil, err
	}
	x := &userDataUploaderUploadDataClient{stream}
	return x, nil
}

type UserDataUploader_UploadDataClient interface {
	Send(*UploadDataPayload) error
	CloseAndRecv() (*general.Empty, error)
	grpc.ClientStream
}

type userDataUploaderUploadDataClient struct {
	grpc.ClientStream
}

func (x *userDataUploaderUploadDataClient) Send(m *UploadDataPayload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userDataUploaderUploadDataClient) CloseAndRecv() (*general.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(general.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userDataUploaderClient) GetData(ctx context.Context, in *GetDataPayload, opts ...grpc.CallOption) (*GetDataResponse, error) {
	out := new(GetDataResponse)
	err := c.cc.Invoke(ctx, "/users.UserDataUploader/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataUploaderClient) UpdateData(ctx context.Context, opts ...grpc.CallOption) (UserDataUploader_UpdateDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserDataUploader_ServiceDesc.Streams[1], "/users.UserDataUploader/UpdateData", opts...)
	if err != nil {
		return nil, err
	}
	x := &userDataUploaderUpdateDataClient{stream}
	return x, nil
}

type UserDataUploader_UpdateDataClient interface {
	Send(*UploadDataPayload) error
	CloseAndRecv() (*general.Empty, error)
	grpc.ClientStream
}

type userDataUploaderUpdateDataClient struct {
	grpc.ClientStream
}

func (x *userDataUploaderUpdateDataClient) Send(m *UploadDataPayload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userDataUploaderUpdateDataClient) CloseAndRecv() (*general.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(general.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userDataUploaderClient) DeleteData(ctx context.Context, in *DeleteDataPayload, opts ...grpc.CallOption) (*general.Empty, error) {
	out := new(general.Empty)
	err := c.cc.Invoke(ctx, "/users.UserDataUploader/DeleteData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDataUploaderServer is the server API for UserDataUploader service.
// All implementations must embed UnimplementedUserDataUploaderServer
// for forward compatibility
type UserDataUploaderServer interface {
	UploadData(UserDataUploader_UploadDataServer) error
	GetData(context.Context, *GetDataPayload) (*GetDataResponse, error)
	UpdateData(UserDataUploader_UpdateDataServer) error
	DeleteData(context.Context, *DeleteDataPayload) (*general.Empty, error)
	mustEmbedUnimplementedUserDataUploaderServer()
}

// UnimplementedUserDataUploaderServer must be embedded to have forward compatible implementations.
type UnimplementedUserDataUploaderServer struct {
}

func (UnimplementedUserDataUploaderServer) UploadData(UserDataUploader_UploadDataServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadData not implemented")
}
func (UnimplementedUserDataUploaderServer) GetData(context.Context, *GetDataPayload) (*GetDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedUserDataUploaderServer) UpdateData(UserDataUploader_UpdateDataServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateData not implemented")
}
func (UnimplementedUserDataUploaderServer) DeleteData(context.Context, *DeleteDataPayload) (*general.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteData not implemented")
}
func (UnimplementedUserDataUploaderServer) mustEmbedUnimplementedUserDataUploaderServer() {}

// UnsafeUserDataUploaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserDataUploaderServer will
// result in compilation errors.
type UnsafeUserDataUploaderServer interface {
	mustEmbedUnimplementedUserDataUploaderServer()
}

func RegisterUserDataUploaderServer(s grpc.ServiceRegistrar, srv UserDataUploaderServer) {
	s.RegisterService(&UserDataUploader_ServiceDesc, srv)
}

func _UserDataUploader_UploadData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserDataUploaderServer).UploadData(&userDataUploaderUploadDataServer{stream})
}

type UserDataUploader_UploadDataServer interface {
	SendAndClose(*general.Empty) error
	Recv() (*UploadDataPayload, error)
	grpc.ServerStream
}

type userDataUploaderUploadDataServer struct {
	grpc.ServerStream
}

func (x *userDataUploaderUploadDataServer) SendAndClose(m *general.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userDataUploaderUploadDataServer) Recv() (*UploadDataPayload, error) {
	m := new(UploadDataPayload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserDataUploader_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataUploaderServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserDataUploader/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataUploaderServer).GetData(ctx, req.(*GetDataPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataUploader_UpdateData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserDataUploaderServer).UpdateData(&userDataUploaderUpdateDataServer{stream})
}

type UserDataUploader_UpdateDataServer interface {
	SendAndClose(*general.Empty) error
	Recv() (*UploadDataPayload, error)
	grpc.ServerStream
}

type userDataUploaderUpdateDataServer struct {
	grpc.ServerStream
}

func (x *userDataUploaderUpdateDataServer) SendAndClose(m *general.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userDataUploaderUpdateDataServer) Recv() (*UploadDataPayload, error) {
	m := new(UploadDataPayload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserDataUploader_DeleteData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDataPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataUploaderServer).DeleteData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserDataUploader/DeleteData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataUploaderServer).DeleteData(ctx, req.(*DeleteDataPayload))
	}
	return interceptor(ctx, in, info, handler)
}

// UserDataUploader_ServiceDesc is the grpc.ServiceDesc for UserDataUploader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserDataUploader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.UserDataUploader",
	HandlerType: (*UserDataUploaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _UserDataUploader_GetData_Handler,
		},
		{
			MethodName: "DeleteData",
			Handler:    _UserDataUploader_DeleteData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadData",
			Handler:       _UserDataUploader_UploadData_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateData",
			Handler:       _UserDataUploader_UpdateData_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "user_data.proto",
}
