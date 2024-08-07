// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: BackgroundStyle.proto

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

type BackgroundStyleUnknown int32

const (
	BackgroundStyle_c BackgroundStyleUnknown = 0
	BackgroundStyle_d BackgroundStyleUnknown = 1
	BackgroundStyle_e BackgroundStyleUnknown = 2
)

// Enum value maps for BackgroundStyleUnknown.
var (
	BackgroundStyleUnknown_name = map[int32]string{
		0: "c",
		1: "d",
		2: "e",
	}
	BackgroundStyleUnknown_value = map[string]int32{
		"c": 0,
		"d": 1,
		"e": 2,
	}
)

func (x BackgroundStyleUnknown) Enum() *BackgroundStyleUnknown {
	p := new(BackgroundStyleUnknown)
	*p = x
	return p
}

func (x BackgroundStyleUnknown) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BackgroundStyleUnknown) Descriptor() protoreflect.EnumDescriptor {
	return file_BackgroundStyle_proto_enumTypes[0].Descriptor()
}

func (BackgroundStyleUnknown) Type() protoreflect.EnumType {
	return &file_BackgroundStyle_proto_enumTypes[0]
}

func (x BackgroundStyleUnknown) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BackgroundStyleUnknown.Descriptor instead.
func (BackgroundStyleUnknown) EnumDescriptor() ([]byte, []int) {
	return file_BackgroundStyle_proto_rawDescGZIP(), []int{0, 0}
}

type BackgroundStyle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string          `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B []*ImageCdnNode `protobuf:"bytes,2,rep,name=b,proto3" json:"b,omitempty"`
}

func (x *BackgroundStyle) Reset() {
	*x = BackgroundStyle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BackgroundStyle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackgroundStyle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackgroundStyle) ProtoMessage() {}

func (x *BackgroundStyle) ProtoReflect() protoreflect.Message {
	mi := &file_BackgroundStyle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackgroundStyle.ProtoReflect.Descriptor instead.
func (*BackgroundStyle) Descriptor() ([]byte, []int) {
	return file_BackgroundStyle_proto_rawDescGZIP(), []int{0}
}

func (x *BackgroundStyle) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *BackgroundStyle) GetB() []*ImageCdnNode {
	if x != nil {
		return x.B
	}
	return nil
}

var File_BackgroundStyle_proto protoreflect.FileDescriptor

var file_BackgroundStyle_proto_rawDesc = []byte{
	0x0a, 0x15, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x79, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61,
	0x6e, 0x6d, 0x75, 0x1a, 0x12, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x64, 0x6e, 0x4e, 0x6f, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0f, 0x42, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x26, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x64, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x01, 0x62,
	0x22, 0x1e, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x05, 0x0a, 0x01, 0x63,
	0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x64, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x65, 0x10, 0x02,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75,
	0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BackgroundStyle_proto_rawDescOnce sync.Once
	file_BackgroundStyle_proto_rawDescData = file_BackgroundStyle_proto_rawDesc
)

func file_BackgroundStyle_proto_rawDescGZIP() []byte {
	file_BackgroundStyle_proto_rawDescOnce.Do(func() {
		file_BackgroundStyle_proto_rawDescData = protoimpl.X.CompressGZIP(file_BackgroundStyle_proto_rawDescData)
	})
	return file_BackgroundStyle_proto_rawDescData
}

var file_BackgroundStyle_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BackgroundStyle_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BackgroundStyle_proto_goTypes = []any{
	(BackgroundStyleUnknown)(0), // 0: AcFunDanmu.BackgroundStyle.unknown
	(*BackgroundStyle)(nil),     // 1: AcFunDanmu.BackgroundStyle
	(*ImageCdnNode)(nil),        // 2: AcFunDanmu.ImageCdnNode
}
var file_BackgroundStyle_proto_depIdxs = []int32{
	2, // 0: AcFunDanmu.BackgroundStyle.b:type_name -> AcFunDanmu.ImageCdnNode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BackgroundStyle_proto_init() }
func file_BackgroundStyle_proto_init() {
	if File_BackgroundStyle_proto != nil {
		return
	}
	file_ImageCdnNode_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BackgroundStyle_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BackgroundStyle); i {
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
			RawDescriptor: file_BackgroundStyle_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BackgroundStyle_proto_goTypes,
		DependencyIndexes: file_BackgroundStyle_proto_depIdxs,
		EnumInfos:         file_BackgroundStyle_proto_enumTypes,
		MessageInfos:      file_BackgroundStyle_proto_msgTypes,
	}.Build()
	File_BackgroundStyle_proto = out.File
	file_BackgroundStyle_proto_rawDesc = nil
	file_BackgroundStyle_proto_goTypes = nil
	file_BackgroundStyle_proto_depIdxs = nil
}
