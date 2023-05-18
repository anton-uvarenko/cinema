// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: password-verify.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EmailPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *EmailPayload) Reset() {
	*x = EmailPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_verify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailPayload) ProtoMessage() {}

func (x *EmailPayload) ProtoReflect() protoreflect.Message {
	mi := &file_password_verify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailPayload.ProtoReflect.Descriptor instead.
func (*EmailPayload) Descriptor() ([]byte, []int) {
	return file_password_verify_proto_rawDescGZIP(), []int{0}
}

func (x *EmailPayload) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CodePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CodePayload) Reset() {
	*x = CodePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_verify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodePayload) ProtoMessage() {}

func (x *CodePayload) ProtoReflect() protoreflect.Message {
	mi := &file_password_verify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodePayload.ProtoReflect.Descriptor instead.
func (*CodePayload) Descriptor() ([]byte, []int) {
	return file_password_verify_proto_rawDescGZIP(), []int{1}
}

func (x *CodePayload) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CodePayload) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type PasswordPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *PasswordPayload) Reset() {
	*x = PasswordPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_verify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordPayload) ProtoMessage() {}

func (x *PasswordPayload) ProtoReflect() protoreflect.Message {
	mi := &file_password_verify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordPayload.ProtoReflect.Descriptor instead.
func (*PasswordPayload) Descriptor() ([]byte, []int) {
	return file_password_verify_proto_rawDescGZIP(), []int{2}
}

func (x *PasswordPayload) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PasswordPayload) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_password_verify_proto protoreflect.FileDescriptor

var file_password_verify_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2d, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x0d, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0c,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x37, 0x0a, 0x0b, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3d, 0x0a, 0x0f, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0xbc, 0x01, 0x0a, 0x0a, 0x50,
	0x61, 0x73, 0x73, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x36, 0x0a, 0x10, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x1a, 0x0e, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x3d, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x14, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4a, 0x77, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0e, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x75, 0x76,
	0x61, 0x72, 0x65, 0x6e, 0x6b, 0x6f, 0x2f, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_password_verify_proto_rawDescOnce sync.Once
	file_password_verify_proto_rawDescData = file_password_verify_proto_rawDesc
)

func file_password_verify_proto_rawDescGZIP() []byte {
	file_password_verify_proto_rawDescOnce.Do(func() {
		file_password_verify_proto_rawDescData = protoimpl.X.CompressGZIP(file_password_verify_proto_rawDescData)
	})
	return file_password_verify_proto_rawDescData
}

var file_password_verify_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_password_verify_proto_goTypes = []interface{}{
	(*EmailPayload)(nil),    // 0: auth.EmailPayload
	(*CodePayload)(nil),     // 1: auth.CodePayload
	(*PasswordPayload)(nil), // 2: auth.PasswordPayload
	(*Empty)(nil),           // 3: general.Empty
	(*JwtResponse)(nil),     // 4: general.JwtResponse
}
var file_password_verify_proto_depIdxs = []int32{
	0, // 0: auth.PassVerify.SendRecoveryCode:input_type -> auth.EmailPayload
	1, // 1: auth.PassVerify.VerifyRecoveryCode:input_type -> auth.CodePayload
	2, // 2: auth.PassVerify.UpdatePassword:input_type -> auth.PasswordPayload
	3, // 3: auth.PassVerify.SendRecoveryCode:output_type -> general.Empty
	4, // 4: auth.PassVerify.VerifyRecoveryCode:output_type -> general.JwtResponse
	3, // 5: auth.PassVerify.UpdatePassword:output_type -> general.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_password_verify_proto_init() }
func file_password_verify_proto_init() {
	if File_password_verify_proto != nil {
		return
	}
	file_general_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_password_verify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_password_verify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodePayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_password_verify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_password_verify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_password_verify_proto_goTypes,
		DependencyIndexes: file_password_verify_proto_depIdxs,
		MessageInfos:      file_password_verify_proto_msgTypes,
	}.Build()
	File_password_verify_proto = out.File
	file_password_verify_proto_rawDesc = nil
	file_password_verify_proto_goTypes = nil
	file_password_verify_proto_depIdxs = nil
}
