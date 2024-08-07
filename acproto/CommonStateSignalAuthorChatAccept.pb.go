// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: CommonStateSignalAuthorChatAccept.proto

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

type CommonStateSignalAuthorChatAccept struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorChatId   string `protobuf:"bytes,1,opt,name=authorChatId,proto3" json:"authorChatId,omitempty"`
	AryaSignalInfo string `protobuf:"bytes,2,opt,name=aryaSignalInfo,proto3" json:"aryaSignalInfo,omitempty"`
}

func (x *CommonStateSignalAuthorChatAccept) Reset() {
	*x = CommonStateSignalAuthorChatAccept{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonStateSignalAuthorChatAccept_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonStateSignalAuthorChatAccept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonStateSignalAuthorChatAccept) ProtoMessage() {}

func (x *CommonStateSignalAuthorChatAccept) ProtoReflect() protoreflect.Message {
	mi := &file_CommonStateSignalAuthorChatAccept_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonStateSignalAuthorChatAccept.ProtoReflect.Descriptor instead.
func (*CommonStateSignalAuthorChatAccept) Descriptor() ([]byte, []int) {
	return file_CommonStateSignalAuthorChatAccept_proto_rawDescGZIP(), []int{0}
}

func (x *CommonStateSignalAuthorChatAccept) GetAuthorChatId() string {
	if x != nil {
		return x.AuthorChatId
	}
	return ""
}

func (x *CommonStateSignalAuthorChatAccept) GetAryaSignalInfo() string {
	if x != nil {
		return x.AryaSignalInfo
	}
	return ""
}

var File_CommonStateSignalAuthorChatAccept_proto protoreflect.FileDescriptor

var file_CommonStateSignalAuthorChatAccept_proto_rawDesc = []byte{
	0x0a, 0x27, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e,
	0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0x6f, 0x0a, 0x21, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x43, 0x68, 0x61, 0x74, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x61, 0x72, 0x79, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x72, 0x79, 0x61, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75,
	0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonStateSignalAuthorChatAccept_proto_rawDescOnce sync.Once
	file_CommonStateSignalAuthorChatAccept_proto_rawDescData = file_CommonStateSignalAuthorChatAccept_proto_rawDesc
)

func file_CommonStateSignalAuthorChatAccept_proto_rawDescGZIP() []byte {
	file_CommonStateSignalAuthorChatAccept_proto_rawDescOnce.Do(func() {
		file_CommonStateSignalAuthorChatAccept_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonStateSignalAuthorChatAccept_proto_rawDescData)
	})
	return file_CommonStateSignalAuthorChatAccept_proto_rawDescData
}

var file_CommonStateSignalAuthorChatAccept_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonStateSignalAuthorChatAccept_proto_goTypes = []any{
	(*CommonStateSignalAuthorChatAccept)(nil), // 0: AcFunDanmu.CommonStateSignalAuthorChatAccept
}
var file_CommonStateSignalAuthorChatAccept_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CommonStateSignalAuthorChatAccept_proto_init() }
func file_CommonStateSignalAuthorChatAccept_proto_init() {
	if File_CommonStateSignalAuthorChatAccept_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CommonStateSignalAuthorChatAccept_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommonStateSignalAuthorChatAccept); i {
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
			RawDescriptor: file_CommonStateSignalAuthorChatAccept_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonStateSignalAuthorChatAccept_proto_goTypes,
		DependencyIndexes: file_CommonStateSignalAuthorChatAccept_proto_depIdxs,
		MessageInfos:      file_CommonStateSignalAuthorChatAccept_proto_msgTypes,
	}.Build()
	File_CommonStateSignalAuthorChatAccept_proto = out.File
	file_CommonStateSignalAuthorChatAccept_proto_rawDesc = nil
	file_CommonStateSignalAuthorChatAccept_proto_goTypes = nil
	file_CommonStateSignalAuthorChatAccept_proto_depIdxs = nil
}
