// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: comment.proto

package users

import (
	general "github.com/anton-uvarenko/cinema/broker-service/protobufs/general"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FilmId      int32                `protobuf:"varint,2,opt,name=FilmId,proto3" json:"FilmId,omitempty"`
	Text        string               `protobuf:"bytes,3,opt,name=Text,proto3" json:"Text,omitempty"`
	UserId      int32                `protobuf:"varint,4,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ReplyTo     int32                `protobuf:"varint,5,opt,name=ReplyTo,proto3" json:"ReplyTo,omitempty"`
	CommentType string               `protobuf:"bytes,6,opt,name=CommentType,proto3" json:"CommentType,omitempty"`
	WasEdited   bool                 `protobuf:"varint,7,opt,name=WasEdited,proto3" json:"WasEdited,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,9,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	AvatarLink  string               `protobuf:"bytes,10,opt,name=AvatarLink,proto3" json:"AvatarLink,omitempty"`
	Username    string               `protobuf:"bytes,11,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetFilmId() int32 {
	if x != nil {
		return x.FilmId
	}
	return 0
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Comment) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Comment) GetReplyTo() int32 {
	if x != nil {
		return x.ReplyTo
	}
	return 0
}

func (x *Comment) GetCommentType() string {
	if x != nil {
		return x.CommentType
	}
	return ""
}

func (x *Comment) GetWasEdited() bool {
	if x != nil {
		return x.WasEdited
	}
	return false
}

func (x *Comment) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Comment) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Comment) GetAvatarLink() string {
	if x != nil {
		return x.AvatarLink
	}
	return ""
}

func (x *Comment) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *CommentsResponse) Reset() {
	*x = CommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentsResponse) ProtoMessage() {}

func (x *CommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentsResponse.ProtoReflect.Descriptor instead.
func (*CommentsResponse) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentsResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type CommentPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilmId      int32  `protobuf:"varint,1,opt,name=FilmId,proto3" json:"FilmId,omitempty"`
	Text        string `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
	UserId      int32  `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ReplyTo     int32  `protobuf:"varint,4,opt,name=ReplyTo,proto3" json:"ReplyTo,omitempty"`
	CommentType string `protobuf:"bytes,6,opt,name=CommentType,proto3" json:"CommentType,omitempty"`
}

func (x *CommentPayload) Reset() {
	*x = CommentPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentPayload) ProtoMessage() {}

func (x *CommentPayload) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentPayload.ProtoReflect.Descriptor instead.
func (*CommentPayload) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CommentPayload) GetFilmId() int32 {
	if x != nil {
		return x.FilmId
	}
	return 0
}

func (x *CommentPayload) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CommentPayload) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CommentPayload) GetReplyTo() int32 {
	if x != nil {
		return x.ReplyTo
	}
	return 0
}

func (x *CommentPayload) GetCommentType() string {
	if x != nil {
		return x.CommentType
	}
	return ""
}

type LikeCommentPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilmId int32 `protobuf:"varint,1,opt,name=FilmId,proto3" json:"FilmId,omitempty"`
	UserId int32 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *LikeCommentPayload) Reset() {
	*x = LikeCommentPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeCommentPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeCommentPayload) ProtoMessage() {}

func (x *LikeCommentPayload) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeCommentPayload.ProtoReflect.Descriptor instead.
func (*LikeCommentPayload) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{3}
}

func (x *LikeCommentPayload) GetFilmId() int32 {
	if x != nil {
		return x.FilmId
	}
	return 0
}

func (x *LikeCommentPayload) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetPrivateCommentsPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	FilmId int32 `protobuf:"varint,2,opt,name=FilmId,proto3" json:"FilmId,omitempty"`
	Page   int32 `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty"`
	Amount int32 `protobuf:"varint,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *GetPrivateCommentsPayload) Reset() {
	*x = GetPrivateCommentsPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivateCommentsPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateCommentsPayload) ProtoMessage() {}

func (x *GetPrivateCommentsPayload) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateCommentsPayload.ProtoReflect.Descriptor instead.
func (*GetPrivateCommentsPayload) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{4}
}

func (x *GetPrivateCommentsPayload) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetPrivateCommentsPayload) GetFilmId() int32 {
	if x != nil {
		return x.FilmId
	}
	return 0
}

func (x *GetPrivateCommentsPayload) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetPrivateCommentsPayload) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type GetPublicCommentsPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilmId          int32 `protobuf:"varint,1,opt,name=FilmId,proto3" json:"FilmId,omitempty"`
	Page            int32 `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	Amount          int32 `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	ResponsesAmount int32 `protobuf:"varint,4,opt,name=ResponsesAmount,proto3" json:"ResponsesAmount,omitempty"`
}

func (x *GetPublicCommentsPayload) Reset() {
	*x = GetPublicCommentsPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublicCommentsPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublicCommentsPayload) ProtoMessage() {}

func (x *GetPublicCommentsPayload) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublicCommentsPayload.ProtoReflect.Descriptor instead.
func (*GetPublicCommentsPayload) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{5}
}

func (x *GetPublicCommentsPayload) GetFilmId() int32 {
	if x != nil {
		return x.FilmId
	}
	return 0
}

func (x *GetPublicCommentsPayload) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetPublicCommentsPayload) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *GetPublicCommentsPayload) GetResponsesAmount() int32 {
	if x != nil {
		return x.ResponsesAmount
	}
	return 0
}

type GetResponsesToCommentPaylaod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId int32 `protobuf:"varint,1,opt,name=CommentId,proto3" json:"CommentId,omitempty"`
	Page      int32 `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	Amount    int32 `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *GetResponsesToCommentPaylaod) Reset() {
	*x = GetResponsesToCommentPaylaod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponsesToCommentPaylaod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponsesToCommentPaylaod) ProtoMessage() {}

func (x *GetResponsesToCommentPaylaod) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponsesToCommentPaylaod.ProtoReflect.Descriptor instead.
func (*GetResponsesToCommentPaylaod) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{6}
}

func (x *GetResponsesToCommentPaylaod) GetCommentId() int32 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *GetResponsesToCommentPaylaod) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetResponsesToCommentPaylaod) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_comment_proto protoreflect.FileDescriptor

var file_comment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x57, 0x61, 0x73, 0x45, 0x64, 0x69, 0x74, 0x65, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x57, 0x61, 0x73, 0x45, 0x64, 0x69, 0x74, 0x65, 0x64,
	0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69,
	0x6e, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x3e, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x90, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54,
	0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x44, 0x0a, 0x12, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c,
	0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x6d, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x77, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x46, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x46, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x46, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x68, 0x0a,
	0x1c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x54, 0x6f, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x61, 0x6f, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x50,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xf0, 0x02, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a,
	0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x54, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x54, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x79, 0x6c, 0x61, 0x6f, 0x64, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0e, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2d, 0x75,
	0x76, 0x61, 0x72, 0x65, 0x6e, 0x6b, 0x6f, 0x2f, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_comment_proto_rawDescOnce sync.Once
	file_comment_proto_rawDescData = file_comment_proto_rawDesc
)

func file_comment_proto_rawDescGZIP() []byte {
	file_comment_proto_rawDescOnce.Do(func() {
		file_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_proto_rawDescData)
	})
	return file_comment_proto_rawDescData
}

var file_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_comment_proto_goTypes = []interface{}{
	(*Comment)(nil),                      // 0: users.Comment
	(*CommentsResponse)(nil),             // 1: users.CommentsResponse
	(*CommentPayload)(nil),               // 2: users.CommentPayload
	(*LikeCommentPayload)(nil),           // 3: users.LikeCommentPayload
	(*GetPrivateCommentsPayload)(nil),    // 4: users.GetPrivateCommentsPayload
	(*GetPublicCommentsPayload)(nil),     // 5: users.GetPublicCommentsPayload
	(*GetResponsesToCommentPaylaod)(nil), // 6: users.GetResponsesToCommentPaylaod
	(*timestamp.Timestamp)(nil),          // 7: google.protobuf.Timestamp
	(*general.Empty)(nil),                // 8: general.Empty
}
var file_comment_proto_depIdxs = []int32{
	7, // 0: users.Comment.CreatedAt:type_name -> google.protobuf.Timestamp
	7, // 1: users.Comment.UpdatedAt:type_name -> google.protobuf.Timestamp
	0, // 2: users.CommentsResponse.comments:type_name -> users.Comment
	2, // 3: users.Comments.AddComment:input_type -> users.CommentPayload
	5, // 4: users.Comments.GetPublicComments:input_type -> users.GetPublicCommentsPayload
	4, // 5: users.Comments.GetPrivateComments:input_type -> users.GetPrivateCommentsPayload
	6, // 6: users.Comments.GetResponsesToComment:input_type -> users.GetResponsesToCommentPaylaod
	3, // 7: users.Comments.LikeComment:input_type -> users.LikeCommentPayload
	0, // 8: users.Comments.AddComment:output_type -> users.Comment
	1, // 9: users.Comments.GetPublicComments:output_type -> users.CommentsResponse
	1, // 10: users.Comments.GetPrivateComments:output_type -> users.CommentsResponse
	1, // 11: users.Comments.GetResponsesToComment:output_type -> users.CommentsResponse
	8, // 12: users.Comments.LikeComment:output_type -> general.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_comment_proto_init() }
func file_comment_proto_init() {
	if File_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentsResponse); i {
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
		file_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentPayload); i {
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
		file_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeCommentPayload); i {
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
		file_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrivateCommentsPayload); i {
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
		file_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPublicCommentsPayload); i {
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
		file_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponsesToCommentPaylaod); i {
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
			RawDescriptor: file_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comment_proto_goTypes,
		DependencyIndexes: file_comment_proto_depIdxs,
		MessageInfos:      file_comment_proto_msgTypes,
	}.Build()
	File_comment_proto = out.File
	file_comment_proto_rawDesc = nil
	file_comment_proto_goTypes = nil
	file_comment_proto_depIdxs = nil
}
