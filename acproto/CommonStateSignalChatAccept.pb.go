// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: CommonStateSignalChatAccept.proto

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

type CommonStateSignalChatAccept struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId          string        `protobuf:"bytes,1,opt,name=chatId,proto3" json:"chatId,omitempty"`
	MediaType       ChatMediaType `protobuf:"varint,2,opt,name=mediaType,proto3,enum=AcFunDanmu.ChatMediaType" json:"mediaType,omitempty"`
	ArraySignalInfo string        `protobuf:"bytes,3,opt,name=arraySignalInfo,proto3" json:"arraySignalInfo,omitempty"`
}

func (x *CommonStateSignalChatAccept) Reset() {
	*x = CommonStateSignalChatAccept{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonStateSignalChatAccept_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonStateSignalChatAccept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonStateSignalChatAccept) ProtoMessage() {}

func (x *CommonStateSignalChatAccept) ProtoReflect() protoreflect.Message {
	mi := &file_CommonStateSignalChatAccept_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonStateSignalChatAccept.ProtoReflect.Descriptor instead.
func (*CommonStateSignalChatAccept) Descriptor() ([]byte, []int) {
	return file_CommonStateSignalChatAccept_proto_rawDescGZIP(), []int{0}
}

func (x *CommonStateSignalChatAccept) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *CommonStateSignalChatAccept) GetMediaType() ChatMediaType {
	if x != nil {
		return x.MediaType
	}
	return ChatMediaType_UNKNOWN
}

func (x *CommonStateSignalChatAccept) GetArraySignalInfo() string {
	if x != nil {
		return x.ArraySignalInfo
	}
	return ""
}

var File_CommonStateSignalChatAccept_proto protoreflect.FileDescriptor

var file_CommonStateSignalChatAccept_proto_rawDesc = []byte{
	0x0a, 0x21, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x1a,
	0x13, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x1b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x09,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x61, 0x72, 0x72, 0x61, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonStateSignalChatAccept_proto_rawDescOnce sync.Once
	file_CommonStateSignalChatAccept_proto_rawDescData = file_CommonStateSignalChatAccept_proto_rawDesc
)

func file_CommonStateSignalChatAccept_proto_rawDescGZIP() []byte {
	file_CommonStateSignalChatAccept_proto_rawDescOnce.Do(func() {
		file_CommonStateSignalChatAccept_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonStateSignalChatAccept_proto_rawDescData)
	})
	return file_CommonStateSignalChatAccept_proto_rawDescData
}

var file_CommonStateSignalChatAccept_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonStateSignalChatAccept_proto_goTypes = []interface{}{
	(*CommonStateSignalChatAccept)(nil), // 0: AcFunDanmu.CommonStateSignalChatAccept
	(ChatMediaType)(0),                  // 1: AcFunDanmu.ChatMediaType
}
var file_CommonStateSignalChatAccept_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.CommonStateSignalChatAccept.mediaType:type_name -> AcFunDanmu.ChatMediaType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CommonStateSignalChatAccept_proto_init() }
func file_CommonStateSignalChatAccept_proto_init() {
	if File_CommonStateSignalChatAccept_proto != nil {
		return
	}
	file_ChatMediaType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CommonStateSignalChatAccept_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonStateSignalChatAccept); i {
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
			RawDescriptor: file_CommonStateSignalChatAccept_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonStateSignalChatAccept_proto_goTypes,
		DependencyIndexes: file_CommonStateSignalChatAccept_proto_depIdxs,
		MessageInfos:      file_CommonStateSignalChatAccept_proto_msgTypes,
	}.Build()
	File_CommonStateSignalChatAccept_proto = out.File
	file_CommonStateSignalChatAccept_proto_rawDesc = nil
	file_CommonStateSignalChatAccept_proto_goTypes = nil
	file_CommonStateSignalChatAccept_proto_depIdxs = nil
}
