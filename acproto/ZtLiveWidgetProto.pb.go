// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: ZtLiveWidgetProto.proto

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

type ZtLiveWidgetProtoA int32

const (
	ZtLiveWidgetProtoA_ZtLiveWidgetProtoAa ZtLiveWidgetProtoA = 0
	ZtLiveWidgetProtoA_ZtLiveWidgetProtoAb ZtLiveWidgetProtoA = 1
	ZtLiveWidgetProtoA_ZtLiveWidgetProtoAc ZtLiveWidgetProtoA = 2
	ZtLiveWidgetProtoA_ZtLiveWidgetProtoAd ZtLiveWidgetProtoA = 3
	ZtLiveWidgetProtoA_ZtLiveWidgetProtoAe ZtLiveWidgetProtoA = 4
)

// Enum value maps for ZtLiveWidgetProtoA.
var (
	ZtLiveWidgetProtoA_name = map[int32]string{
		0: "ZtLiveWidgetProtoAa",
		1: "ZtLiveWidgetProtoAb",
		2: "ZtLiveWidgetProtoAc",
		3: "ZtLiveWidgetProtoAd",
		4: "ZtLiveWidgetProtoAe",
	}
	ZtLiveWidgetProtoA_value = map[string]int32{
		"ZtLiveWidgetProtoAa": 0,
		"ZtLiveWidgetProtoAb": 1,
		"ZtLiveWidgetProtoAc": 2,
		"ZtLiveWidgetProtoAd": 3,
		"ZtLiveWidgetProtoAe": 4,
	}
)

func (x ZtLiveWidgetProtoA) Enum() *ZtLiveWidgetProtoA {
	p := new(ZtLiveWidgetProtoA)
	*p = x
	return p
}

func (x ZtLiveWidgetProtoA) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZtLiveWidgetProtoA) Descriptor() protoreflect.EnumDescriptor {
	return file_ZtLiveWidgetProto_proto_enumTypes[0].Descriptor()
}

func (ZtLiveWidgetProtoA) Type() protoreflect.EnumType {
	return &file_ZtLiveWidgetProto_proto_enumTypes[0]
}

func (x ZtLiveWidgetProtoA) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZtLiveWidgetProtoA.Descriptor instead.
func (ZtLiveWidgetProtoA) EnumDescriptor() ([]byte, []int) {
	return file_ZtLiveWidgetProto_proto_rawDescGZIP(), []int{0}
}

type ZtLiveWidgetProtoB int32

const (
	ZtLiveWidgetProtoB_ZtLiveWidgetProtoBa ZtLiveWidgetProtoB = 0
	ZtLiveWidgetProtoB_ZtLiveWidgetProtoBb ZtLiveWidgetProtoB = 1
	ZtLiveWidgetProtoB_ZtLiveWidgetProtoBc ZtLiveWidgetProtoB = 2
	ZtLiveWidgetProtoB_ZtLiveWidgetProtoBd ZtLiveWidgetProtoB = 3
	ZtLiveWidgetProtoB_ZtLiveWidgetProtoBe ZtLiveWidgetProtoB = 4
	ZtLiveWidgetProtoB_ZtLiveWidgetProtoBf ZtLiveWidgetProtoB = 5
)

// Enum value maps for ZtLiveWidgetProtoB.
var (
	ZtLiveWidgetProtoB_name = map[int32]string{
		0: "ZtLiveWidgetProtoBa",
		1: "ZtLiveWidgetProtoBb",
		2: "ZtLiveWidgetProtoBc",
		3: "ZtLiveWidgetProtoBd",
		4: "ZtLiveWidgetProtoBe",
		5: "ZtLiveWidgetProtoBf",
	}
	ZtLiveWidgetProtoB_value = map[string]int32{
		"ZtLiveWidgetProtoBa": 0,
		"ZtLiveWidgetProtoBb": 1,
		"ZtLiveWidgetProtoBc": 2,
		"ZtLiveWidgetProtoBd": 3,
		"ZtLiveWidgetProtoBe": 4,
		"ZtLiveWidgetProtoBf": 5,
	}
)

func (x ZtLiveWidgetProtoB) Enum() *ZtLiveWidgetProtoB {
	p := new(ZtLiveWidgetProtoB)
	*p = x
	return p
}

func (x ZtLiveWidgetProtoB) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZtLiveWidgetProtoB) Descriptor() protoreflect.EnumDescriptor {
	return file_ZtLiveWidgetProto_proto_enumTypes[1].Descriptor()
}

func (ZtLiveWidgetProtoB) Type() protoreflect.EnumType {
	return &file_ZtLiveWidgetProto_proto_enumTypes[1]
}

func (x ZtLiveWidgetProtoB) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZtLiveWidgetProtoB.Descriptor instead.
func (ZtLiveWidgetProtoB) EnumDescriptor() ([]byte, []int) {
	return file_ZtLiveWidgetProto_proto_rawDescGZIP(), []int{1}
}

type ZtLiveWidgetProtoC int32

const (
	ZtLiveWidgetProtoC_ZtLiveWidgetProtoCa ZtLiveWidgetProtoC = 0
	ZtLiveWidgetProtoC_ZtLiveWidgetProtoCb ZtLiveWidgetProtoC = 1
	ZtLiveWidgetProtoC_ZtLiveWidgetProtoCc ZtLiveWidgetProtoC = 2
	ZtLiveWidgetProtoC_ZtLiveWidgetProtoCd ZtLiveWidgetProtoC = 3
)

// Enum value maps for ZtLiveWidgetProtoC.
var (
	ZtLiveWidgetProtoC_name = map[int32]string{
		0: "ZtLiveWidgetProtoCa",
		1: "ZtLiveWidgetProtoCb",
		2: "ZtLiveWidgetProtoCc",
		3: "ZtLiveWidgetProtoCd",
	}
	ZtLiveWidgetProtoC_value = map[string]int32{
		"ZtLiveWidgetProtoCa": 0,
		"ZtLiveWidgetProtoCb": 1,
		"ZtLiveWidgetProtoCc": 2,
		"ZtLiveWidgetProtoCd": 3,
	}
)

func (x ZtLiveWidgetProtoC) Enum() *ZtLiveWidgetProtoC {
	p := new(ZtLiveWidgetProtoC)
	*p = x
	return p
}

func (x ZtLiveWidgetProtoC) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZtLiveWidgetProtoC) Descriptor() protoreflect.EnumDescriptor {
	return file_ZtLiveWidgetProto_proto_enumTypes[2].Descriptor()
}

func (ZtLiveWidgetProtoC) Type() protoreflect.EnumType {
	return &file_ZtLiveWidgetProto_proto_enumTypes[2]
}

func (x ZtLiveWidgetProtoC) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZtLiveWidgetProtoC.Descriptor instead.
func (ZtLiveWidgetProtoC) EnumDescriptor() ([]byte, []int) {
	return file_ZtLiveWidgetProto_proto_rawDescGZIP(), []int{2}
}

var File_ZtLiveWidgetProto_proto protoreflect.FileDescriptor

var file_ZtLiveWidgetProto_proto_rawDesc = []byte{
	0x0a, 0x17, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e,
	0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2a, 0x91, 0x01, 0x0a, 0x12, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x12, 0x17, 0x0a, 0x13,
	0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x41, 0x61, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x62, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x13, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x41, 0x63, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x5a, 0x74, 0x4c, 0x69, 0x76,
	0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x64, 0x10, 0x03,
	0x12, 0x17, 0x0a, 0x13, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x65, 0x10, 0x04, 0x2a, 0xaa, 0x01, 0x0a, 0x12, 0x5a, 0x74,
	0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42,
	0x12, 0x17, 0x0a, 0x13, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x61, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x5a, 0x74, 0x4c,
	0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x62,
	0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x63, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x5a,
	0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x42, 0x64, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x65, 0x10, 0x04, 0x12, 0x17, 0x0a,
	0x13, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x66, 0x10, 0x05, 0x2a, 0x78, 0x0a, 0x12, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x12, 0x17, 0x0a, 0x13,
	0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x43, 0x61, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x62, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x13, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x43, 0x63, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x5a, 0x74, 0x4c, 0x69, 0x76,
	0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x64, 0x10, 0x03,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75,
	0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ZtLiveWidgetProto_proto_rawDescOnce sync.Once
	file_ZtLiveWidgetProto_proto_rawDescData = file_ZtLiveWidgetProto_proto_rawDesc
)

func file_ZtLiveWidgetProto_proto_rawDescGZIP() []byte {
	file_ZtLiveWidgetProto_proto_rawDescOnce.Do(func() {
		file_ZtLiveWidgetProto_proto_rawDescData = protoimpl.X.CompressGZIP(file_ZtLiveWidgetProto_proto_rawDescData)
	})
	return file_ZtLiveWidgetProto_proto_rawDescData
}

var file_ZtLiveWidgetProto_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_ZtLiveWidgetProto_proto_goTypes = []interface{}{
	(ZtLiveWidgetProtoA)(0), // 0: AcFunDanmu.ZtLiveWidgetProtoA
	(ZtLiveWidgetProtoB)(0), // 1: AcFunDanmu.ZtLiveWidgetProtoB
	(ZtLiveWidgetProtoC)(0), // 2: AcFunDanmu.ZtLiveWidgetProtoC
}
var file_ZtLiveWidgetProto_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ZtLiveWidgetProto_proto_init() }
func file_ZtLiveWidgetProto_proto_init() {
	if File_ZtLiveWidgetProto_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ZtLiveWidgetProto_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ZtLiveWidgetProto_proto_goTypes,
		DependencyIndexes: file_ZtLiveWidgetProto_proto_depIdxs,
		EnumInfos:         file_ZtLiveWidgetProto_proto_enumTypes,
	}.Build()
	File_ZtLiveWidgetProto_proto = out.File
	file_ZtLiveWidgetProto_proto_rawDesc = nil
	file_ZtLiveWidgetProto_proto_goTypes = nil
	file_ZtLiveWidgetProto_proto_depIdxs = nil
}
