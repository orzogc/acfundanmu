// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: KwaiStateSignalEcommerceCart.proto

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

type KwaiStateSignalEcommerceCartUnknown int32

const (
	KwaiStateSignalEcommerceCart_b KwaiStateSignalEcommerceCartUnknown = 0
	KwaiStateSignalEcommerceCart_c KwaiStateSignalEcommerceCartUnknown = 1
)

// Enum value maps for KwaiStateSignalEcommerceCartUnknown.
var (
	KwaiStateSignalEcommerceCartUnknown_name = map[int32]string{
		0: "b",
		1: "c",
	}
	KwaiStateSignalEcommerceCartUnknown_value = map[string]int32{
		"b": 0,
		"c": 1,
	}
)

func (x KwaiStateSignalEcommerceCartUnknown) Enum() *KwaiStateSignalEcommerceCartUnknown {
	p := new(KwaiStateSignalEcommerceCartUnknown)
	*p = x
	return p
}

func (x KwaiStateSignalEcommerceCartUnknown) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KwaiStateSignalEcommerceCartUnknown) Descriptor() protoreflect.EnumDescriptor {
	return file_KwaiStateSignalEcommerceCart_proto_enumTypes[0].Descriptor()
}

func (KwaiStateSignalEcommerceCartUnknown) Type() protoreflect.EnumType {
	return &file_KwaiStateSignalEcommerceCart_proto_enumTypes[0]
}

func (x KwaiStateSignalEcommerceCartUnknown) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KwaiStateSignalEcommerceCartUnknown.Descriptor instead.
func (KwaiStateSignalEcommerceCartUnknown) EnumDescriptor() ([]byte, []int) {
	return file_KwaiStateSignalEcommerceCart_proto_rawDescGZIP(), []int{0, 0}
}

type KwaiStateSignalEcommerceCart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A KwaiStateSignalEcommerceCartUnknown `protobuf:"varint,1,opt,name=a,proto3,enum=AcFunDanmu.KwaiStateSignalEcommerceCartUnknown" json:"a,omitempty"`
}

func (x *KwaiStateSignalEcommerceCart) Reset() {
	*x = KwaiStateSignalEcommerceCart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KwaiStateSignalEcommerceCart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KwaiStateSignalEcommerceCart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KwaiStateSignalEcommerceCart) ProtoMessage() {}

func (x *KwaiStateSignalEcommerceCart) ProtoReflect() protoreflect.Message {
	mi := &file_KwaiStateSignalEcommerceCart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KwaiStateSignalEcommerceCart.ProtoReflect.Descriptor instead.
func (*KwaiStateSignalEcommerceCart) Descriptor() ([]byte, []int) {
	return file_KwaiStateSignalEcommerceCart_proto_rawDescGZIP(), []int{0}
}

func (x *KwaiStateSignalEcommerceCart) GetA() KwaiStateSignalEcommerceCartUnknown {
	if x != nil {
		return x.A
	}
	return KwaiStateSignalEcommerceCart_b
}

var File_KwaiStateSignalEcommerceCart_proto protoreflect.FileDescriptor

var file_KwaiStateSignalEcommerceCart_proto_rawDesc = []byte{
	0x0a, 0x22, 0x4b, 0x77, 0x61, 0x69, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x45, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x43, 0x61, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x22, 0x77, 0x0a, 0x1c, 0x4b, 0x77, 0x61, 0x69, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x12, 0x3e, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x41, 0x63,
	0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x4b, 0x77, 0x61, 0x69, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x2e, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x01, 0x61,
	0x22, 0x17, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x05, 0x0a, 0x01, 0x62,
	0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x63, 0x10, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61,
	0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_KwaiStateSignalEcommerceCart_proto_rawDescOnce sync.Once
	file_KwaiStateSignalEcommerceCart_proto_rawDescData = file_KwaiStateSignalEcommerceCart_proto_rawDesc
)

func file_KwaiStateSignalEcommerceCart_proto_rawDescGZIP() []byte {
	file_KwaiStateSignalEcommerceCart_proto_rawDescOnce.Do(func() {
		file_KwaiStateSignalEcommerceCart_proto_rawDescData = protoimpl.X.CompressGZIP(file_KwaiStateSignalEcommerceCart_proto_rawDescData)
	})
	return file_KwaiStateSignalEcommerceCart_proto_rawDescData
}

var file_KwaiStateSignalEcommerceCart_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_KwaiStateSignalEcommerceCart_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_KwaiStateSignalEcommerceCart_proto_goTypes = []any{
	(KwaiStateSignalEcommerceCartUnknown)(0), // 0: AcFunDanmu.KwaiStateSignalEcommerceCart.unknown
	(*KwaiStateSignalEcommerceCart)(nil),     // 1: AcFunDanmu.KwaiStateSignalEcommerceCart
}
var file_KwaiStateSignalEcommerceCart_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.KwaiStateSignalEcommerceCart.a:type_name -> AcFunDanmu.KwaiStateSignalEcommerceCart.unknown
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_KwaiStateSignalEcommerceCart_proto_init() }
func file_KwaiStateSignalEcommerceCart_proto_init() {
	if File_KwaiStateSignalEcommerceCart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_KwaiStateSignalEcommerceCart_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*KwaiStateSignalEcommerceCart); i {
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
			RawDescriptor: file_KwaiStateSignalEcommerceCart_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KwaiStateSignalEcommerceCart_proto_goTypes,
		DependencyIndexes: file_KwaiStateSignalEcommerceCart_proto_depIdxs,
		EnumInfos:         file_KwaiStateSignalEcommerceCart_proto_enumTypes,
		MessageInfos:      file_KwaiStateSignalEcommerceCart_proto_msgTypes,
	}.Build()
	File_KwaiStateSignalEcommerceCart_proto = out.File
	file_KwaiStateSignalEcommerceCart_proto_rawDesc = nil
	file_KwaiStateSignalEcommerceCart_proto_goTypes = nil
	file_KwaiStateSignalEcommerceCart_proto_depIdxs = nil
}
