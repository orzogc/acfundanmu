// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: acfun.live.proto

package acproto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AcFunUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AcFunUserInfo) Reset() {
	*x = AcFunUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acfun_live_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcFunUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcFunUserInfo) ProtoMessage() {}

func (x *AcFunUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_acfun_live_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcFunUserInfo.ProtoReflect.Descriptor instead.
func (*AcFunUserInfo) Descriptor() ([]byte, []int) {
	return file_acfun_live_proto_rawDescGZIP(), []int{0}
}

func (x *AcFunUserInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AcFunUserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AcfunActionSignalThrowBanana struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Visitor    *AcFunUserInfo `protobuf:"bytes,1,opt,name=visitor,proto3" json:"visitor,omitempty"`
	Count      int32          `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	SendTimeMs int64          `protobuf:"varint,3,opt,name=sendTimeMs,proto3" json:"sendTimeMs,omitempty"`
}

func (x *AcfunActionSignalThrowBanana) Reset() {
	*x = AcfunActionSignalThrowBanana{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acfun_live_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcfunActionSignalThrowBanana) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcfunActionSignalThrowBanana) ProtoMessage() {}

func (x *AcfunActionSignalThrowBanana) ProtoReflect() protoreflect.Message {
	mi := &file_acfun_live_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcfunActionSignalThrowBanana.ProtoReflect.Descriptor instead.
func (*AcfunActionSignalThrowBanana) Descriptor() ([]byte, []int) {
	return file_acfun_live_proto_rawDescGZIP(), []int{1}
}

func (x *AcfunActionSignalThrowBanana) GetVisitor() *AcFunUserInfo {
	if x != nil {
		return x.Visitor
	}
	return nil
}

func (x *AcfunActionSignalThrowBanana) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *AcfunActionSignalThrowBanana) GetSendTimeMs() int64 {
	if x != nil {
		return x.SendTimeMs
	}
	return 0
}

type AcfunStateSignalDisplayInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BananaCount string `protobuf:"bytes,1,opt,name=bananaCount,proto3" json:"bananaCount,omitempty"`
}

func (x *AcfunStateSignalDisplayInfo) Reset() {
	*x = AcfunStateSignalDisplayInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acfun_live_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcfunStateSignalDisplayInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcfunStateSignalDisplayInfo) ProtoMessage() {}

func (x *AcfunStateSignalDisplayInfo) ProtoReflect() protoreflect.Message {
	mi := &file_acfun_live_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcfunStateSignalDisplayInfo.ProtoReflect.Descriptor instead.
func (*AcfunStateSignalDisplayInfo) Descriptor() ([]byte, []int) {
	return file_acfun_live_proto_rawDescGZIP(), []int{2}
}

func (x *AcfunStateSignalDisplayInfo) GetBananaCount() string {
	if x != nil {
		return x.BananaCount
	}
	return ""
}

var File_acfun_live_proto protoreflect.FileDescriptor

var file_acfun_live_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0x3b,
	0x0a, 0x0d, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x1c,
	0x41, 0x63, 0x66, 0x75, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x54, 0x68, 0x72, 0x6f, 0x77, 0x42, 0x61, 0x6e, 0x61, 0x6e, 0x61, 0x12, 0x33, 0x0a, 0x07,
	0x76, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x4d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x22, 0x3f, 0x0a, 0x1b, 0x41, 0x63, 0x66, 0x75, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x61, 0x6e, 0x61, 0x6e, 0x61,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x6e,
	0x61, 0x6e, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x61, 0x63,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acfun_live_proto_rawDescOnce sync.Once
	file_acfun_live_proto_rawDescData = file_acfun_live_proto_rawDesc
)

func file_acfun_live_proto_rawDescGZIP() []byte {
	file_acfun_live_proto_rawDescOnce.Do(func() {
		file_acfun_live_proto_rawDescData = protoimpl.X.CompressGZIP(file_acfun_live_proto_rawDescData)
	})
	return file_acfun_live_proto_rawDescData
}

var file_acfun_live_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_acfun_live_proto_goTypes = []interface{}{
	(*AcFunUserInfo)(nil),                // 0: AcFunDanmu.AcFunUserInfo
	(*AcfunActionSignalThrowBanana)(nil), // 1: AcFunDanmu.AcfunActionSignalThrowBanana
	(*AcfunStateSignalDisplayInfo)(nil),  // 2: AcFunDanmu.AcfunStateSignalDisplayInfo
}
var file_acfun_live_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.AcfunActionSignalThrowBanana.visitor:type_name -> AcFunDanmu.AcFunUserInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_acfun_live_proto_init() }
func file_acfun_live_proto_init() {
	if File_acfun_live_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_acfun_live_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcFunUserInfo); i {
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
		file_acfun_live_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcfunActionSignalThrowBanana); i {
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
		file_acfun_live_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcfunStateSignalDisplayInfo); i {
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
			RawDescriptor: file_acfun_live_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_acfun_live_proto_goTypes,
		DependencyIndexes: file_acfun_live_proto_depIdxs,
		MessageInfos:      file_acfun_live_proto_msgTypes,
	}.Build()
	File_acfun_live_proto = out.File
	file_acfun_live_proto_rawDesc = nil
	file_acfun_live_proto_goTypes = nil
	file_acfun_live_proto_depIdxs = nil
}
