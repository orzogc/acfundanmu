// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: CommonStateSignalPkEnd.proto

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

type CommonStateSignalPkEndUnknown int32

const (
	CommonStateSignalPkEnd_d CommonStateSignalPkEndUnknown = 0
	CommonStateSignalPkEnd_e CommonStateSignalPkEndUnknown = 1
	CommonStateSignalPkEnd_f CommonStateSignalPkEndUnknown = 2
	CommonStateSignalPkEnd_g CommonStateSignalPkEndUnknown = 3
	CommonStateSignalPkEnd_h CommonStateSignalPkEndUnknown = 4
	CommonStateSignalPkEnd_i CommonStateSignalPkEndUnknown = 5
	CommonStateSignalPkEnd_j CommonStateSignalPkEndUnknown = 6
	CommonStateSignalPkEnd_k CommonStateSignalPkEndUnknown = 7
	CommonStateSignalPkEnd_l CommonStateSignalPkEndUnknown = 8
	CommonStateSignalPkEnd_m CommonStateSignalPkEndUnknown = 9
	CommonStateSignalPkEnd_n CommonStateSignalPkEndUnknown = 10
	CommonStateSignalPkEnd_o CommonStateSignalPkEndUnknown = 11
)

// Enum value maps for CommonStateSignalPkEndUnknown.
var (
	CommonStateSignalPkEndUnknown_name = map[int32]string{
		0:  "d",
		1:  "e",
		2:  "f",
		3:  "g",
		4:  "h",
		5:  "i",
		6:  "j",
		7:  "k",
		8:  "l",
		9:  "m",
		10: "n",
		11: "o",
	}
	CommonStateSignalPkEndUnknown_value = map[string]int32{
		"d": 0,
		"e": 1,
		"f": 2,
		"g": 3,
		"h": 4,
		"i": 5,
		"j": 6,
		"k": 7,
		"l": 8,
		"m": 9,
		"n": 10,
		"o": 11,
	}
)

func (x CommonStateSignalPkEndUnknown) Enum() *CommonStateSignalPkEndUnknown {
	p := new(CommonStateSignalPkEndUnknown)
	*p = x
	return p
}

func (x CommonStateSignalPkEndUnknown) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonStateSignalPkEndUnknown) Descriptor() protoreflect.EnumDescriptor {
	return file_CommonStateSignalPkEnd_proto_enumTypes[0].Descriptor()
}

func (CommonStateSignalPkEndUnknown) Type() protoreflect.EnumType {
	return &file_CommonStateSignalPkEnd_proto_enumTypes[0]
}

func (x CommonStateSignalPkEndUnknown) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonStateSignalPkEndUnknown.Descriptor instead.
func (CommonStateSignalPkEndUnknown) EnumDescriptor() ([]byte, []int) {
	return file_CommonStateSignalPkEnd_proto_rawDescGZIP(), []int{0, 0}
}

type CommonStateSignalPkEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string                        `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B CommonStateSignalPkEndUnknown `protobuf:"varint,2,opt,name=b,proto3,enum=AcFunDanmu.CommonStateSignalPkEndUnknown" json:"b,omitempty"`
	C string                        `protobuf:"bytes,3,opt,name=c,proto3" json:"c,omitempty"`
}

func (x *CommonStateSignalPkEnd) Reset() {
	*x = CommonStateSignalPkEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonStateSignalPkEnd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonStateSignalPkEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonStateSignalPkEnd) ProtoMessage() {}

func (x *CommonStateSignalPkEnd) ProtoReflect() protoreflect.Message {
	mi := &file_CommonStateSignalPkEnd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonStateSignalPkEnd.ProtoReflect.Descriptor instead.
func (*CommonStateSignalPkEnd) Descriptor() ([]byte, []int) {
	return file_CommonStateSignalPkEnd_proto_rawDescGZIP(), []int{0}
}

func (x *CommonStateSignalPkEnd) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *CommonStateSignalPkEnd) GetB() CommonStateSignalPkEndUnknown {
	if x != nil {
		return x.B
	}
	return CommonStateSignalPkEnd_d
}

func (x *CommonStateSignalPkEnd) GetC() string {
	if x != nil {
		return x.C
	}
	return ""
}

var File_CommonStateSignalPkEnd_proto protoreflect.FileDescriptor

var file_CommonStateSignalPkEnd_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x50, 0x6b, 0x45, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0xcd, 0x01, 0x0a, 0x16, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x50, 0x6b, 0x45, 0x6e, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x61, 0x12, 0x38, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a,
	0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x50, 0x6b, 0x45,
	0x6e, 0x64, 0x2e, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a,
	0x01, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x63, 0x22, 0x5d, 0x0a, 0x07, 0x75,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x05, 0x0a, 0x01, 0x64, 0x10, 0x00, 0x12, 0x05, 0x0a,
	0x01, 0x65, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x66, 0x10, 0x02, 0x12, 0x05, 0x0a, 0x01, 0x67,
	0x10, 0x03, 0x12, 0x05, 0x0a, 0x01, 0x68, 0x10, 0x04, 0x12, 0x05, 0x0a, 0x01, 0x69, 0x10, 0x05,
	0x12, 0x05, 0x0a, 0x01, 0x6a, 0x10, 0x06, 0x12, 0x05, 0x0a, 0x01, 0x6b, 0x10, 0x07, 0x12, 0x05,
	0x0a, 0x01, 0x6c, 0x10, 0x08, 0x12, 0x05, 0x0a, 0x01, 0x6d, 0x10, 0x09, 0x12, 0x05, 0x0a, 0x01,
	0x6e, 0x10, 0x0a, 0x12, 0x05, 0x0a, 0x01, 0x6f, 0x10, 0x0b, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f,
	0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonStateSignalPkEnd_proto_rawDescOnce sync.Once
	file_CommonStateSignalPkEnd_proto_rawDescData = file_CommonStateSignalPkEnd_proto_rawDesc
)

func file_CommonStateSignalPkEnd_proto_rawDescGZIP() []byte {
	file_CommonStateSignalPkEnd_proto_rawDescOnce.Do(func() {
		file_CommonStateSignalPkEnd_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonStateSignalPkEnd_proto_rawDescData)
	})
	return file_CommonStateSignalPkEnd_proto_rawDescData
}

var file_CommonStateSignalPkEnd_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CommonStateSignalPkEnd_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonStateSignalPkEnd_proto_goTypes = []any{
	(CommonStateSignalPkEndUnknown)(0), // 0: AcFunDanmu.CommonStateSignalPkEnd.unknown
	(*CommonStateSignalPkEnd)(nil),     // 1: AcFunDanmu.CommonStateSignalPkEnd
}
var file_CommonStateSignalPkEnd_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.CommonStateSignalPkEnd.b:type_name -> AcFunDanmu.CommonStateSignalPkEnd.unknown
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CommonStateSignalPkEnd_proto_init() }
func file_CommonStateSignalPkEnd_proto_init() {
	if File_CommonStateSignalPkEnd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CommonStateSignalPkEnd_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommonStateSignalPkEnd); i {
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
			RawDescriptor: file_CommonStateSignalPkEnd_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonStateSignalPkEnd_proto_goTypes,
		DependencyIndexes: file_CommonStateSignalPkEnd_proto_depIdxs,
		EnumInfos:         file_CommonStateSignalPkEnd_proto_enumTypes,
		MessageInfos:      file_CommonStateSignalPkEnd_proto_msgTypes,
	}.Build()
	File_CommonStateSignalPkEnd_proto = out.File
	file_CommonStateSignalPkEnd_proto_rawDesc = nil
	file_CommonStateSignalPkEnd_proto_goTypes = nil
	file_CommonStateSignalPkEnd_proto_depIdxs = nil
}
