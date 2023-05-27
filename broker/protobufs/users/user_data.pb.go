// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: user_data.proto

package users

import (
	general "github.com/anton-uvarenko/cinema/broker-service/protobufs/general"
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

type UploadDataPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image    []byte `protobuf:"bytes,1,opt,name=Image,proto3" json:"Image,omitempty"`
	FavActor int32  `protobuf:"varint,2,opt,name=FavActor,proto3" json:"FavActor,omitempty"`
	FavGenre int32  `protobuf:"varint,3,opt,name=FavGenre,proto3" json:"FavGenre,omitempty"`
	FavFilm  int32  `protobuf:"varint,4,opt,name=FavFilm,proto3" json:"FavFilm,omitempty"`
	UserId   int32  `protobuf:"varint,5,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *UploadDataPayload) Reset() {
	*x = UploadDataPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadDataPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadDataPayload) ProtoMessage() {}

func (x *UploadDataPayload) ProtoReflect() protoreflect.Message {
	mi := &file_user_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadDataPayload.ProtoReflect.Descriptor instead.
func (*UploadDataPayload) Descriptor() ([]byte, []int) {
	return file_user_data_proto_rawDescGZIP(), []int{0}
}

func (x *UploadDataPayload) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *UploadDataPayload) GetFavActor() int32 {
	if x != nil {
		return x.FavActor
	}
	return 0
}

func (x *UploadDataPayload) GetFavGenre() int32 {
	if x != nil {
		return x.FavGenre
	}
	return 0
}

func (x *UploadDataPayload) GetFavFilm() int32 {
	if x != nil {
		return x.FavFilm
	}
	return 0
}

func (x *UploadDataPayload) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageLink string `protobuf:"bytes,1,opt,name=ImageLink,proto3" json:"ImageLink,omitempty"`
	FavActor  int32  `protobuf:"varint,2,opt,name=FavActor,proto3" json:"FavActor,omitempty"`
	FavGenre  int32  `protobuf:"varint,3,opt,name=FavGenre,proto3" json:"FavGenre,omitempty"`
	FavFilm   int32  `protobuf:"varint,4,opt,name=FavFilm,proto3" json:"FavFilm,omitempty"`
	UserId    int32  `protobuf:"varint,5,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetDataResponse) Reset() {
	*x = GetDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataResponse) ProtoMessage() {}

func (x *GetDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataResponse.ProtoReflect.Descriptor instead.
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return file_user_data_proto_rawDescGZIP(), []int{1}
}

func (x *GetDataResponse) GetImageLink() string {
	if x != nil {
		return x.ImageLink
	}
	return ""
}

func (x *GetDataResponse) GetFavActor() int32 {
	if x != nil {
		return x.FavActor
	}
	return 0
}

func (x *GetDataResponse) GetFavGenre() int32 {
	if x != nil {
		return x.FavGenre
	}
	return 0
}

func (x *GetDataResponse) GetFavFilm() int32 {
	if x != nil {
		return x.FavFilm
	}
	return 0
}

func (x *GetDataResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeleteDataPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *DeleteDataPayload) Reset() {
	*x = DeleteDataPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDataPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDataPayload) ProtoMessage() {}

func (x *DeleteDataPayload) ProtoReflect() protoreflect.Message {
	mi := &file_user_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDataPayload.ProtoReflect.Descriptor instead.
func (*DeleteDataPayload) Descriptor() ([]byte, []int) {
	return file_user_data_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteDataPayload) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetDataPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetDataPayload) Reset() {
	*x = GetDataPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataPayload) ProtoMessage() {}

func (x *GetDataPayload) ProtoReflect() protoreflect.Message {
	mi := &file_user_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataPayload.ProtoReflect.Descriptor instead.
func (*GetDataPayload) Descriptor() ([]byte, []int) {
	return file_user_data_proto_rawDescGZIP(), []int{3}
}

func (x *GetDataPayload) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_user_data_proto protoreflect.FileDescriptor

var file_user_data_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x76, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x61, 0x76, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x46, 0x61, 0x76, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x46, 0x61, 0x76, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x46,
	0x61, 0x76, 0x46, 0x69, 0x6c, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x46, 0x61,
	0x76, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x99, 0x01,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x1a, 0x0a, 0x08, 0x46, 0x61, 0x76, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x46, 0x61, 0x76, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x46,
	0x61, 0x76, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46,
	0x61, 0x76, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x61, 0x76, 0x46, 0x69,
	0x6c, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x46, 0x61, 0x76, 0x46, 0x69, 0x6c,
	0x6d, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x32, 0xf8, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0e, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01, 0x12,
	0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x1a, 0x0e, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x28, 0x01, 0x12, 0x36, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0e, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3f, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d,
	0x75, 0x76, 0x61, 0x72, 0x65, 0x6e, 0x6b, 0x6f, 0x2f, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_data_proto_rawDescOnce sync.Once
	file_user_data_proto_rawDescData = file_user_data_proto_rawDesc
)

func file_user_data_proto_rawDescGZIP() []byte {
	file_user_data_proto_rawDescOnce.Do(func() {
		file_user_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_data_proto_rawDescData)
	})
	return file_user_data_proto_rawDescData
}

var file_user_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_data_proto_goTypes = []interface{}{
	(*UploadDataPayload)(nil), // 0: users.UploadDataPayload
	(*GetDataResponse)(nil),   // 1: users.GetDataResponse
	(*DeleteDataPayload)(nil), // 2: users.DeleteDataPayload
	(*GetDataPayload)(nil),    // 3: users.GetDataPayload
	(*general.Empty)(nil),     // 4: general.Empty
}
var file_user_data_proto_depIdxs = []int32{
	0, // 0: users.UserDataUploader.UploadData:input_type -> users.UploadDataPayload
	3, // 1: users.UserDataUploader.GetData:input_type -> users.GetDataPayload
	0, // 2: users.UserDataUploader.UpdateData:input_type -> users.UploadDataPayload
	2, // 3: users.UserDataUploader.DeleteData:input_type -> users.DeleteDataPayload
	4, // 4: users.UserDataUploader.UploadData:output_type -> general.Empty
	1, // 5: users.UserDataUploader.GetData:output_type -> users.GetDataResponse
	4, // 6: users.UserDataUploader.UpdateData:output_type -> general.Empty
	4, // 7: users.UserDataUploader.DeleteData:output_type -> general.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_data_proto_init() }
func file_user_data_proto_init() {
	if File_user_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadDataPayload); i {
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
		file_user_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataResponse); i {
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
		file_user_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDataPayload); i {
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
		file_user_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataPayload); i {
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
			RawDescriptor: file_user_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_data_proto_goTypes,
		DependencyIndexes: file_user_data_proto_depIdxs,
		MessageInfos:      file_user_data_proto_msgTypes,
	}.Build()
	File_user_data_proto = out.File
	file_user_data_proto_rawDesc = nil
	file_user_data_proto_goTypes = nil
	file_user_data_proto_depIdxs = nil
}
