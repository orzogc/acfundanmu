// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ZtLivePkProto.proto

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

type ZtLivePkProto int32

const (
	ZtLivePkProto_ZtLivePkProtoA ZtLivePkProto = 0
	ZtLivePkProto_ZtLivePkProtoB ZtLivePkProto = 1
	ZtLivePkProto_ZtLivePkProtoC ZtLivePkProto = 2
	ZtLivePkProto_ZtLivePkProtoD ZtLivePkProto = 3
)

// Enum value maps for ZtLivePkProto.
var (
	ZtLivePkProto_name = map[int32]string{
		0: "ZtLivePkProtoA",
		1: "ZtLivePkProtoB",
		2: "ZtLivePkProtoC",
		3: "ZtLivePkProtoD",
	}
	ZtLivePkProto_value = map[string]int32{
		"ZtLivePkProtoA": 0,
		"ZtLivePkProtoB": 1,
		"ZtLivePkProtoC": 2,
		"ZtLivePkProtoD": 3,
	}
)

func (x ZtLivePkProto) Enum() *ZtLivePkProto {
	p := new(ZtLivePkProto)
	*p = x
	return p
}

func (x ZtLivePkProto) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZtLivePkProto) Descriptor() protoreflect.EnumDescriptor {
	return file_ZtLivePkProto_proto_enumTypes[0].Descriptor()
}

func (ZtLivePkProto) Type() protoreflect.EnumType {
	return &file_ZtLivePkProto_proto_enumTypes[0]
}

func (x ZtLivePkProto) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZtLivePkProto.Descriptor instead.
func (ZtLivePkProto) EnumDescriptor() ([]byte, []int) {
	return file_ZtLivePkProto_proto_rawDescGZIP(), []int{0}
}

var File_ZtLivePkProto_proto protoreflect.FileDescriptor

var file_ZtLivePkProto_proto_rawDesc = []byte{
	0x0a, 0x13, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x50, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d,
	0x75, 0x2a, 0x5f, 0x0a, 0x0d, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x50, 0x6b, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x0a, 0x0e, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x50, 0x6b, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x41, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65,
	0x50, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x5a, 0x74,
	0x4c, 0x69, 0x76, 0x65, 0x50, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x50, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x44,
	0x10, 0x03, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e,
	0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ZtLivePkProto_proto_rawDescOnce sync.Once
	file_ZtLivePkProto_proto_rawDescData = file_ZtLivePkProto_proto_rawDesc
)

func file_ZtLivePkProto_proto_rawDescGZIP() []byte {
	file_ZtLivePkProto_proto_rawDescOnce.Do(func() {
		file_ZtLivePkProto_proto_rawDescData = protoimpl.X.CompressGZIP(file_ZtLivePkProto_proto_rawDescData)
	})
	return file_ZtLivePkProto_proto_rawDescData
}

var file_ZtLivePkProto_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ZtLivePkProto_proto_goTypes = []interface{}{
	(ZtLivePkProto)(0), // 0: AcFunDanmu.ZtLivePkProto
}
var file_ZtLivePkProto_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ZtLivePkProto_proto_init() }
func file_ZtLivePkProto_proto_init() {
	if File_ZtLivePkProto_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ZtLivePkProto_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ZtLivePkProto_proto_goTypes,
		DependencyIndexes: file_ZtLivePkProto_proto_depIdxs,
		EnumInfos:         file_ZtLivePkProto_proto_enumTypes,
	}.Build()
	File_ZtLivePkProto_proto = out.File
	file_ZtLivePkProto_proto_rawDesc = nil
	file_ZtLivePkProto_proto_goTypes = nil
	file_ZtLivePkProto_proto_depIdxs = nil
}
