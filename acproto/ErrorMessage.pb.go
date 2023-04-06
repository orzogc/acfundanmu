// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: ErrorMessage.proto

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

type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocaleMessages []*LocaleMessage `protobuf:"bytes,1,rep,name=localeMessages,proto3" json:"localeMessages,omitempty"`
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ErrorMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ErrorMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_ErrorMessage_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorMessage) GetLocaleMessages() []*LocaleMessage {
	if x != nil {
		return x.LocaleMessages
	}
	return nil
}

var File_ErrorMessage_proto protoreflect.FileDescriptor

var file_ErrorMessage_proto_rawDesc = []byte{
	0x0a, 0x12, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x2e, 0x49, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x1a, 0x13, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a,
	0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4a,
	0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61,
	0x6e, 0x6d, 0x75, 0x2e, 0x49, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f,
	0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ErrorMessage_proto_rawDescOnce sync.Once
	file_ErrorMessage_proto_rawDescData = file_ErrorMessage_proto_rawDesc
)

func file_ErrorMessage_proto_rawDescGZIP() []byte {
	file_ErrorMessage_proto_rawDescOnce.Do(func() {
		file_ErrorMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_ErrorMessage_proto_rawDescData)
	})
	return file_ErrorMessage_proto_rawDescData
}

var file_ErrorMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ErrorMessage_proto_goTypes = []interface{}{
	(*ErrorMessage)(nil),  // 0: AcFunDanmu.Im.Basic.ErrorMessage
	(*LocaleMessage)(nil), // 1: AcFunDanmu.Im.Basic.LocaleMessage
}
var file_ErrorMessage_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.Im.Basic.ErrorMessage.localeMessages:type_name -> AcFunDanmu.Im.Basic.LocaleMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ErrorMessage_proto_init() }
func file_ErrorMessage_proto_init() {
	if File_ErrorMessage_proto != nil {
		return
	}
	file_LocaleMessage_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ErrorMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMessage); i {
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
			RawDescriptor: file_ErrorMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ErrorMessage_proto_goTypes,
		DependencyIndexes: file_ErrorMessage_proto_depIdxs,
		MessageInfos:      file_ErrorMessage_proto_msgTypes,
	}.Build()
	File_ErrorMessage_proto = out.File
	file_ErrorMessage_proto_rawDesc = nil
	file_ErrorMessage_proto_goTypes = nil
	file_ErrorMessage_proto_depIdxs = nil
}
