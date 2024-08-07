// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: CommonStateSignalPKSoundConfigChanged.proto

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

type CommonStateSignalPKSoundConfigChangedUnknown int32

const (
	CommonStateSignalPKSoundConfigChanged_c CommonStateSignalPKSoundConfigChangedUnknown = 0
	CommonStateSignalPKSoundConfigChanged_d CommonStateSignalPKSoundConfigChangedUnknown = 1
	CommonStateSignalPKSoundConfigChanged_e CommonStateSignalPKSoundConfigChangedUnknown = 2
)

// Enum value maps for CommonStateSignalPKSoundConfigChangedUnknown.
var (
	CommonStateSignalPKSoundConfigChangedUnknown_name = map[int32]string{
		0: "c",
		1: "d",
		2: "e",
	}
	CommonStateSignalPKSoundConfigChangedUnknown_value = map[string]int32{
		"c": 0,
		"d": 1,
		"e": 2,
	}
)

func (x CommonStateSignalPKSoundConfigChangedUnknown) Enum() *CommonStateSignalPKSoundConfigChangedUnknown {
	p := new(CommonStateSignalPKSoundConfigChangedUnknown)
	*p = x
	return p
}

func (x CommonStateSignalPKSoundConfigChangedUnknown) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonStateSignalPKSoundConfigChangedUnknown) Descriptor() protoreflect.EnumDescriptor {
	return file_CommonStateSignalPKSoundConfigChanged_proto_enumTypes[0].Descriptor()
}

func (CommonStateSignalPKSoundConfigChangedUnknown) Type() protoreflect.EnumType {
	return &file_CommonStateSignalPKSoundConfigChanged_proto_enumTypes[0]
}

func (x CommonStateSignalPKSoundConfigChangedUnknown) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonStateSignalPKSoundConfigChangedUnknown.Descriptor instead.
func (CommonStateSignalPKSoundConfigChangedUnknown) EnumDescriptor() ([]byte, []int) {
	return file_CommonStateSignalPKSoundConfigChanged_proto_rawDescGZIP(), []int{0, 0}
}

type CommonStateSignalPKSoundConfigChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string                                       `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B CommonStateSignalPKSoundConfigChangedUnknown `protobuf:"varint,2,opt,name=b,proto3,enum=AcFunDanmu.CommonStateSignalPKSoundConfigChangedUnknown" json:"b,omitempty"`
}

func (x *CommonStateSignalPKSoundConfigChanged) Reset() {
	*x = CommonStateSignalPKSoundConfigChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonStateSignalPKSoundConfigChanged_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonStateSignalPKSoundConfigChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonStateSignalPKSoundConfigChanged) ProtoMessage() {}

func (x *CommonStateSignalPKSoundConfigChanged) ProtoReflect() protoreflect.Message {
	mi := &file_CommonStateSignalPKSoundConfigChanged_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonStateSignalPKSoundConfigChanged.ProtoReflect.Descriptor instead.
func (*CommonStateSignalPKSoundConfigChanged) Descriptor() ([]byte, []int) {
	return file_CommonStateSignalPKSoundConfigChanged_proto_rawDescGZIP(), []int{0}
}

func (x *CommonStateSignalPKSoundConfigChanged) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *CommonStateSignalPKSoundConfigChanged) GetB() CommonStateSignalPKSoundConfigChangedUnknown {
	if x != nil {
		return x.B
	}
	return CommonStateSignalPKSoundConfigChanged_c
}

var File_CommonStateSignalPKSoundConfigChanged_proto protoreflect.FileDescriptor

var file_CommonStateSignalPKSoundConfigChanged_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x50, 0x4b, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41,
	0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0x9e, 0x01, 0x0a, 0x25, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x50,
	0x4b, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x61, 0x12, 0x47, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x41,
	0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x50, 0x4b, 0x53, 0x6f, 0x75,
	0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x2e,
	0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x01, 0x62, 0x22, 0x1e, 0x0a, 0x07, 0x75, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x05, 0x0a, 0x01, 0x63, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01,
	0x64, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x65, 0x10, 0x02, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f,
	0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonStateSignalPKSoundConfigChanged_proto_rawDescOnce sync.Once
	file_CommonStateSignalPKSoundConfigChanged_proto_rawDescData = file_CommonStateSignalPKSoundConfigChanged_proto_rawDesc
)

func file_CommonStateSignalPKSoundConfigChanged_proto_rawDescGZIP() []byte {
	file_CommonStateSignalPKSoundConfigChanged_proto_rawDescOnce.Do(func() {
		file_CommonStateSignalPKSoundConfigChanged_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonStateSignalPKSoundConfigChanged_proto_rawDescData)
	})
	return file_CommonStateSignalPKSoundConfigChanged_proto_rawDescData
}

var file_CommonStateSignalPKSoundConfigChanged_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CommonStateSignalPKSoundConfigChanged_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonStateSignalPKSoundConfigChanged_proto_goTypes = []any{
	(CommonStateSignalPKSoundConfigChangedUnknown)(0), // 0: AcFunDanmu.CommonStateSignalPKSoundConfigChanged.unknown
	(*CommonStateSignalPKSoundConfigChanged)(nil),     // 1: AcFunDanmu.CommonStateSignalPKSoundConfigChanged
}
var file_CommonStateSignalPKSoundConfigChanged_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.CommonStateSignalPKSoundConfigChanged.b:type_name -> AcFunDanmu.CommonStateSignalPKSoundConfigChanged.unknown
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CommonStateSignalPKSoundConfigChanged_proto_init() }
func file_CommonStateSignalPKSoundConfigChanged_proto_init() {
	if File_CommonStateSignalPKSoundConfigChanged_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CommonStateSignalPKSoundConfigChanged_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommonStateSignalPKSoundConfigChanged); i {
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
			RawDescriptor: file_CommonStateSignalPKSoundConfigChanged_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonStateSignalPKSoundConfigChanged_proto_goTypes,
		DependencyIndexes: file_CommonStateSignalPKSoundConfigChanged_proto_depIdxs,
		EnumInfos:         file_CommonStateSignalPKSoundConfigChanged_proto_enumTypes,
		MessageInfos:      file_CommonStateSignalPKSoundConfigChanged_proto_msgTypes,
	}.Build()
	File_CommonStateSignalPKSoundConfigChanged_proto = out.File
	file_CommonStateSignalPKSoundConfigChanged_proto_rawDesc = nil
	file_CommonStateSignalPKSoundConfigChanged_proto_goTypes = nil
	file_CommonStateSignalPKSoundConfigChanged_proto_depIdxs = nil
}
