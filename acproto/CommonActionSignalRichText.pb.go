// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: CommonActionSignalRichText.proto

package acproto

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

type CommonActionSignalRichText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Segments   []*CommonActionSignalRichText_RichTextSegment `protobuf:"bytes,1,rep,name=segments,proto3" json:"segments,omitempty"` //segment
	SendTimeMs int64                                         `protobuf:"varint,2,opt,name=sendTimeMs,proto3" json:"sendTimeMs,omitempty"`
}

func (x *CommonActionSignalRichText) Reset() {
	*x = CommonActionSignalRichText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonActionSignalRichText_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonActionSignalRichText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonActionSignalRichText) ProtoMessage() {}

func (x *CommonActionSignalRichText) ProtoReflect() protoreflect.Message {
	mi := &file_CommonActionSignalRichText_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonActionSignalRichText.ProtoReflect.Descriptor instead.
func (*CommonActionSignalRichText) Descriptor() ([]byte, []int) {
	return file_CommonActionSignalRichText_proto_rawDescGZIP(), []int{0}
}

func (x *CommonActionSignalRichText) GetSegments() []*CommonActionSignalRichText_RichTextSegment {
	if x != nil {
		return x.Segments
	}
	return nil
}

func (x *CommonActionSignalRichText) GetSendTimeMs() int64 {
	if x != nil {
		return x.SendTimeMs
	}
	return 0
}

type CommonActionSignalRichText_ImageSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pictures         []*ImageCdnNode `protobuf:"bytes,1,rep,name=pictures,proto3" json:"pictures,omitempty"` //cdnNode
	AlternativeText  string          `protobuf:"bytes,2,opt,name=alternativeText,proto3" json:"alternativeText,omitempty"`
	AlternativeColor string          `protobuf:"bytes,3,opt,name=alternativeColor,proto3" json:"alternativeColor,omitempty"`
}

func (x *CommonActionSignalRichText_ImageSegment) Reset() {
	*x = CommonActionSignalRichText_ImageSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonActionSignalRichText_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonActionSignalRichText_ImageSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonActionSignalRichText_ImageSegment) ProtoMessage() {}

func (x *CommonActionSignalRichText_ImageSegment) ProtoReflect() protoreflect.Message {
	mi := &file_CommonActionSignalRichText_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonActionSignalRichText_ImageSegment.ProtoReflect.Descriptor instead.
func (*CommonActionSignalRichText_ImageSegment) Descriptor() ([]byte, []int) {
	return file_CommonActionSignalRichText_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CommonActionSignalRichText_ImageSegment) GetPictures() []*ImageCdnNode {
	if x != nil {
		return x.Pictures
	}
	return nil
}

func (x *CommonActionSignalRichText_ImageSegment) GetAlternativeText() string {
	if x != nil {
		return x.AlternativeText
	}
	return ""
}

func (x *CommonActionSignalRichText_ImageSegment) GetAlternativeColor() string {
	if x != nil {
		return x.AlternativeColor
	}
	return ""
}

type CommonActionSignalRichText_PlainSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Color string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *CommonActionSignalRichText_PlainSegment) Reset() {
	*x = CommonActionSignalRichText_PlainSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonActionSignalRichText_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonActionSignalRichText_PlainSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonActionSignalRichText_PlainSegment) ProtoMessage() {}

func (x *CommonActionSignalRichText_PlainSegment) ProtoReflect() protoreflect.Message {
	mi := &file_CommonActionSignalRichText_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonActionSignalRichText_PlainSegment.ProtoReflect.Descriptor instead.
func (*CommonActionSignalRichText_PlainSegment) Descriptor() ([]byte, []int) {
	return file_CommonActionSignalRichText_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CommonActionSignalRichText_PlainSegment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CommonActionSignalRichText_PlainSegment) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type CommonActionSignalRichText_RichTextSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Segment:
	//	*CommonActionSignalRichText_RichTextSegment_UserInfo
	//	*CommonActionSignalRichText_RichTextSegment_Plain
	//	*CommonActionSignalRichText_RichTextSegment_Image
	Segment isCommonActionSignalRichText_RichTextSegment_Segment `protobuf_oneof:"segment"`
}

func (x *CommonActionSignalRichText_RichTextSegment) Reset() {
	*x = CommonActionSignalRichText_RichTextSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonActionSignalRichText_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonActionSignalRichText_RichTextSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonActionSignalRichText_RichTextSegment) ProtoMessage() {}

func (x *CommonActionSignalRichText_RichTextSegment) ProtoReflect() protoreflect.Message {
	mi := &file_CommonActionSignalRichText_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonActionSignalRichText_RichTextSegment.ProtoReflect.Descriptor instead.
func (*CommonActionSignalRichText_RichTextSegment) Descriptor() ([]byte, []int) {
	return file_CommonActionSignalRichText_proto_rawDescGZIP(), []int{0, 2}
}

func (m *CommonActionSignalRichText_RichTextSegment) GetSegment() isCommonActionSignalRichText_RichTextSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (x *CommonActionSignalRichText_RichTextSegment) GetUserInfo() *CommonActionSignalRichText_UserInfoSegment {
	if x, ok := x.GetSegment().(*CommonActionSignalRichText_RichTextSegment_UserInfo); ok {
		return x.UserInfo
	}
	return nil
}

func (x *CommonActionSignalRichText_RichTextSegment) GetPlain() *CommonActionSignalRichText_PlainSegment {
	if x, ok := x.GetSegment().(*CommonActionSignalRichText_RichTextSegment_Plain); ok {
		return x.Plain
	}
	return nil
}

func (x *CommonActionSignalRichText_RichTextSegment) GetImage() *CommonActionSignalRichText_ImageSegment {
	if x, ok := x.GetSegment().(*CommonActionSignalRichText_RichTextSegment_Image); ok {
		return x.Image
	}
	return nil
}

type isCommonActionSignalRichText_RichTextSegment_Segment interface {
	isCommonActionSignalRichText_RichTextSegment_Segment()
}

type CommonActionSignalRichText_RichTextSegment_UserInfo struct {
	UserInfo *CommonActionSignalRichText_UserInfoSegment `protobuf:"bytes,1,opt,name=userInfo,proto3,oneof"`
}

type CommonActionSignalRichText_RichTextSegment_Plain struct {
	Plain *CommonActionSignalRichText_PlainSegment `protobuf:"bytes,2,opt,name=plain,proto3,oneof"`
}

type CommonActionSignalRichText_RichTextSegment_Image struct {
	Image *CommonActionSignalRichText_ImageSegment `protobuf:"bytes,3,opt,name=image,proto3,oneof"`
}

func (*CommonActionSignalRichText_RichTextSegment_UserInfo) isCommonActionSignalRichText_RichTextSegment_Segment() {
}

func (*CommonActionSignalRichText_RichTextSegment_Plain) isCommonActionSignalRichText_RichTextSegment_Segment() {
}

func (*CommonActionSignalRichText_RichTextSegment_Image) isCommonActionSignalRichText_RichTextSegment_Segment() {
}

type CommonActionSignalRichText_UserInfoSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  *ZtLiveUserInfo `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Color string          `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *CommonActionSignalRichText_UserInfoSegment) Reset() {
	*x = CommonActionSignalRichText_UserInfoSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonActionSignalRichText_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonActionSignalRichText_UserInfoSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonActionSignalRichText_UserInfoSegment) ProtoMessage() {}

func (x *CommonActionSignalRichText_UserInfoSegment) ProtoReflect() protoreflect.Message {
	mi := &file_CommonActionSignalRichText_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonActionSignalRichText_UserInfoSegment.ProtoReflect.Descriptor instead.
func (*CommonActionSignalRichText_UserInfoSegment) Descriptor() ([]byte, []int) {
	return file_CommonActionSignalRichText_proto_rawDescGZIP(), []int{0, 3}
}

func (x *CommonActionSignalRichText_UserInfoSegment) GetUser() *ZtLiveUserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CommonActionSignalRichText_UserInfoSegment) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

var File_CommonActionSignalRichText_proto protoreflect.FileDescriptor

var file_CommonActionSignalRichText_proto_rawDesc = []byte{
	0x0a, 0x20, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x52, 0x69, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x1a, 0x14,
	0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x64, 0x6e, 0x4e, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x05, 0x0a, 0x1a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52,
	0x69, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x12, 0x52, 0x0a, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x41, 0x63, 0x46, 0x75,
	0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x69, 0x63, 0x68, 0x54, 0x65, 0x78,
	0x74, 0x2e, 0x52, 0x69, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x1a, 0x9a, 0x01, 0x0a, 0x0c,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x08,
	0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x43, 0x64, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6c, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x2a, 0x0a, 0x10,
	0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x1a, 0x38, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x69,
	0x6e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x1a, 0x8c, 0x02, 0x0a, 0x0f, 0x52, 0x69, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x54, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e,
	0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x69, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4b, 0x0a, 0x05,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x41, 0x63,
	0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x69, 0x63, 0x68, 0x54,
	0x65, 0x78, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x12, 0x4b, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e,
	0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x69, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x1a, 0x57, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e,
	0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f,
	0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonActionSignalRichText_proto_rawDescOnce sync.Once
	file_CommonActionSignalRichText_proto_rawDescData = file_CommonActionSignalRichText_proto_rawDesc
)

func file_CommonActionSignalRichText_proto_rawDescGZIP() []byte {
	file_CommonActionSignalRichText_proto_rawDescOnce.Do(func() {
		file_CommonActionSignalRichText_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonActionSignalRichText_proto_rawDescData)
	})
	return file_CommonActionSignalRichText_proto_rawDescData
}

var file_CommonActionSignalRichText_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_CommonActionSignalRichText_proto_goTypes = []interface{}{
	(*CommonActionSignalRichText)(nil),                 // 0: AcFunDanmu.CommonActionSignalRichText
	(*CommonActionSignalRichText_ImageSegment)(nil),    // 1: AcFunDanmu.CommonActionSignalRichText.ImageSegment
	(*CommonActionSignalRichText_PlainSegment)(nil),    // 2: AcFunDanmu.CommonActionSignalRichText.PlainSegment
	(*CommonActionSignalRichText_RichTextSegment)(nil), // 3: AcFunDanmu.CommonActionSignalRichText.RichTextSegment
	(*CommonActionSignalRichText_UserInfoSegment)(nil), // 4: AcFunDanmu.CommonActionSignalRichText.UserInfoSegment
	(*ImageCdnNode)(nil),                               // 5: AcFunDanmu.ImageCdnNode
	(*ZtLiveUserInfo)(nil),                             // 6: AcFunDanmu.ZtLiveUserInfo
}
var file_CommonActionSignalRichText_proto_depIdxs = []int32{
	3, // 0: AcFunDanmu.CommonActionSignalRichText.segments:type_name -> AcFunDanmu.CommonActionSignalRichText.RichTextSegment
	5, // 1: AcFunDanmu.CommonActionSignalRichText.ImageSegment.pictures:type_name -> AcFunDanmu.ImageCdnNode
	4, // 2: AcFunDanmu.CommonActionSignalRichText.RichTextSegment.userInfo:type_name -> AcFunDanmu.CommonActionSignalRichText.UserInfoSegment
	2, // 3: AcFunDanmu.CommonActionSignalRichText.RichTextSegment.plain:type_name -> AcFunDanmu.CommonActionSignalRichText.PlainSegment
	1, // 4: AcFunDanmu.CommonActionSignalRichText.RichTextSegment.image:type_name -> AcFunDanmu.CommonActionSignalRichText.ImageSegment
	6, // 5: AcFunDanmu.CommonActionSignalRichText.UserInfoSegment.user:type_name -> AcFunDanmu.ZtLiveUserInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_CommonActionSignalRichText_proto_init() }
func file_CommonActionSignalRichText_proto_init() {
	if File_CommonActionSignalRichText_proto != nil {
		return
	}
	file_ZtLiveUserInfo_proto_init()
	file_ImageCdnNode_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CommonActionSignalRichText_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonActionSignalRichText); i {
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
		file_CommonActionSignalRichText_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonActionSignalRichText_ImageSegment); i {
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
		file_CommonActionSignalRichText_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonActionSignalRichText_PlainSegment); i {
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
		file_CommonActionSignalRichText_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonActionSignalRichText_RichTextSegment); i {
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
		file_CommonActionSignalRichText_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonActionSignalRichText_UserInfoSegment); i {
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
	file_CommonActionSignalRichText_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*CommonActionSignalRichText_RichTextSegment_UserInfo)(nil),
		(*CommonActionSignalRichText_RichTextSegment_Plain)(nil),
		(*CommonActionSignalRichText_RichTextSegment_Image)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CommonActionSignalRichText_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonActionSignalRichText_proto_goTypes,
		DependencyIndexes: file_CommonActionSignalRichText_proto_depIdxs,
		MessageInfos:      file_CommonActionSignalRichText_proto_msgTypes,
	}.Build()
	File_CommonActionSignalRichText_proto = out.File
	file_CommonActionSignalRichText_proto_rawDesc = nil
	file_CommonActionSignalRichText_proto_goTypes = nil
	file_CommonActionSignalRichText_proto_depIdxs = nil
}
