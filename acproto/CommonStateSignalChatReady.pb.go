// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: CommonStateSignalChatReady.proto

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

type CommonStateSignalChatReady struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId        string          `protobuf:"bytes,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
	GuestUserInfo *ZtLiveUserInfo `protobuf:"bytes,2,opt,name=guestUserInfo,proto3" json:"guestUserInfo,omitempty"`
	MediaType     ChatMediaType   `protobuf:"varint,3,opt,name=mediaType,proto3,enum=AcFunDanmu.ChatMediaType" json:"mediaType,omitempty"`
}

func (x *CommonStateSignalChatReady) Reset() {
	*x = CommonStateSignalChatReady{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonStateSignalChatReady_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonStateSignalChatReady) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonStateSignalChatReady) ProtoMessage() {}

func (x *CommonStateSignalChatReady) ProtoReflect() protoreflect.Message {
	mi := &file_CommonStateSignalChatReady_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonStateSignalChatReady.ProtoReflect.Descriptor instead.
func (*CommonStateSignalChatReady) Descriptor() ([]byte, []int) {
	return file_CommonStateSignalChatReady_proto_rawDescGZIP(), []int{0}
}

func (x *CommonStateSignalChatReady) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *CommonStateSignalChatReady) GetGuestUserInfo() *ZtLiveUserInfo {
	if x != nil {
		return x.GuestUserInfo
	}
	return nil
}

func (x *CommonStateSignalChatReady) GetMediaType() ChatMediaType {
	if x != nil {
		return x.MediaType
	}
	return ChatMediaType_UNKNOWN
}

var File_CommonStateSignalChatReady_proto protoreflect.FileDescriptor

var file_CommonStateSignalChatReady_proto_rawDesc = []byte{
	0x0a, 0x20, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x1a, 0x13,
	0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x1a, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64,
	0x12, 0x40, 0x0a, 0x0d, 0x67, 0x75, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44,
	0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x67, 0x75, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x37, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e,
	0x6d, 0x75, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63,
	0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonStateSignalChatReady_proto_rawDescOnce sync.Once
	file_CommonStateSignalChatReady_proto_rawDescData = file_CommonStateSignalChatReady_proto_rawDesc
)

func file_CommonStateSignalChatReady_proto_rawDescGZIP() []byte {
	file_CommonStateSignalChatReady_proto_rawDescOnce.Do(func() {
		file_CommonStateSignalChatReady_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonStateSignalChatReady_proto_rawDescData)
	})
	return file_CommonStateSignalChatReady_proto_rawDescData
}

var file_CommonStateSignalChatReady_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonStateSignalChatReady_proto_goTypes = []interface{}{
	(*CommonStateSignalChatReady)(nil), // 0: AcFunDanmu.CommonStateSignalChatReady
	(*ZtLiveUserInfo)(nil),             // 1: AcFunDanmu.ZtLiveUserInfo
	(ChatMediaType)(0),                 // 2: AcFunDanmu.ChatMediaType
}
var file_CommonStateSignalChatReady_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.CommonStateSignalChatReady.guestUserInfo:type_name -> AcFunDanmu.ZtLiveUserInfo
	2, // 1: AcFunDanmu.CommonStateSignalChatReady.mediaType:type_name -> AcFunDanmu.ChatMediaType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_CommonStateSignalChatReady_proto_init() }
func file_CommonStateSignalChatReady_proto_init() {
	if File_CommonStateSignalChatReady_proto != nil {
		return
	}
	file_ChatMediaType_proto_init()
	file_ZtLiveUserInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CommonStateSignalChatReady_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonStateSignalChatReady); i {
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
			RawDescriptor: file_CommonStateSignalChatReady_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonStateSignalChatReady_proto_goTypes,
		DependencyIndexes: file_CommonStateSignalChatReady_proto_depIdxs,
		MessageInfos:      file_CommonStateSignalChatReady_proto_msgTypes,
	}.Build()
	File_CommonStateSignalChatReady_proto = out.File
	file_CommonStateSignalChatReady_proto_rawDesc = nil
	file_CommonStateSignalChatReady_proto_goTypes = nil
	file_CommonStateSignalChatReady_proto_depIdxs = nil
}
