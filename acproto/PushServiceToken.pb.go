// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: PushServiceToken.proto

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

type PushServiceToken_PushType int32

const (
	PushServiceToken_kPushTypeInvalid PushServiceToken_PushType = 0
	PushServiceToken_kPushTypeAPNS    PushServiceToken_PushType = 1
	PushServiceToken_kPushTypeXmPush  PushServiceToken_PushType = 2
	PushServiceToken_kPushTypeJgPush  PushServiceToken_PushType = 3
	PushServiceToken_kPushTypeGtPush  PushServiceToken_PushType = 4
	PushServiceToken_kPushTypeOpPush  PushServiceToken_PushType = 5
	PushServiceToken_kPushTypeVvPush  PushServiceToken_PushType = 6
	PushServiceToken_kPushTypeHwPush  PushServiceToken_PushType = 7
	PushServiceToken_kPushTypeFcm     PushServiceToken_PushType = 8
)

// Enum value maps for PushServiceToken_PushType.
var (
	PushServiceToken_PushType_name = map[int32]string{
		0: "kPushTypeInvalid",
		1: "kPushTypeAPNS",
		2: "kPushTypeXmPush",
		3: "kPushTypeJgPush",
		4: "kPushTypeGtPush",
		5: "kPushTypeOpPush",
		6: "kPushTypeVvPush",
		7: "kPushTypeHwPush",
		8: "kPushTypeFcm",
	}
	PushServiceToken_PushType_value = map[string]int32{
		"kPushTypeInvalid": 0,
		"kPushTypeAPNS":    1,
		"kPushTypeXmPush":  2,
		"kPushTypeJgPush":  3,
		"kPushTypeGtPush":  4,
		"kPushTypeOpPush":  5,
		"kPushTypeVvPush":  6,
		"kPushTypeHwPush":  7,
		"kPushTypeFcm":     8,
	}
)

func (x PushServiceToken_PushType) Enum() *PushServiceToken_PushType {
	p := new(PushServiceToken_PushType)
	*p = x
	return p
}

func (x PushServiceToken_PushType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushServiceToken_PushType) Descriptor() protoreflect.EnumDescriptor {
	return file_PushServiceToken_proto_enumTypes[0].Descriptor()
}

func (PushServiceToken_PushType) Type() protoreflect.EnumType {
	return &file_PushServiceToken_proto_enumTypes[0]
}

func (x PushServiceToken_PushType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushServiceToken_PushType.Descriptor instead.
func (PushServiceToken_PushType) EnumDescriptor() ([]byte, []int) {
	return file_PushServiceToken_proto_rawDescGZIP(), []int{0, 0}
}

type PushServiceToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushType      PushServiceToken_PushType `protobuf:"varint,1,opt,name=pushType,proto3,enum=AcFunDanmu.Im.Basic.PushServiceToken_PushType" json:"pushType,omitempty"`
	Token         []byte                    `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	IsPassThrough bool                      `protobuf:"varint,3,opt,name=isPassThrough,proto3" json:"isPassThrough,omitempty"`
}

func (x *PushServiceToken) Reset() {
	*x = PushServiceToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PushServiceToken_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushServiceToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushServiceToken) ProtoMessage() {}

func (x *PushServiceToken) ProtoReflect() protoreflect.Message {
	mi := &file_PushServiceToken_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushServiceToken.ProtoReflect.Descriptor instead.
func (*PushServiceToken) Descriptor() ([]byte, []int) {
	return file_PushServiceToken_proto_rawDescGZIP(), []int{0}
}

func (x *PushServiceToken) GetPushType() PushServiceToken_PushType {
	if x != nil {
		return x.PushType
	}
	return PushServiceToken_kPushTypeInvalid
}

func (x *PushServiceToken) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *PushServiceToken) GetIsPassThrough() bool {
	if x != nil {
		return x.IsPassThrough
	}
	return false
}

var File_PushServiceToken_proto protoreflect.FileDescriptor

var file_PushServiceToken_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44,
	0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x49, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x22, 0xe0, 0x02,
	0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x4a, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d,
	0x75, 0x2e, 0x49, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x73, 0x50, 0x61, 0x73, 0x73, 0x54, 0x68,
	0x72, 0x6f, 0x75, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x50,
	0x61, 0x73, 0x73, 0x54, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x22, 0xc3, 0x01, 0x0a, 0x08, 0x50,
	0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x6b, 0x50, 0x75, 0x73, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x6b, 0x50, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x41, 0x50, 0x4e, 0x53, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x6b, 0x50, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x58, 0x6d, 0x50,
	0x75, 0x73, 0x68, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x6b, 0x50, 0x75, 0x73, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x4a, 0x67, 0x50, 0x75, 0x73, 0x68, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x6b, 0x50,
	0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x47, 0x74, 0x50, 0x75, 0x73, 0x68, 0x10, 0x04, 0x12,
	0x13, 0x0a, 0x0f, 0x6b, 0x50, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x4f, 0x70, 0x50, 0x75,
	0x73, 0x68, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x6b, 0x50, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x56, 0x76, 0x50, 0x75, 0x73, 0x68, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x6b, 0x50, 0x75,
	0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x48, 0x77, 0x50, 0x75, 0x73, 0x68, 0x10, 0x07, 0x12, 0x10,
	0x0a, 0x0c, 0x6b, 0x50, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x46, 0x63, 0x6d, 0x10, 0x08,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75,
	0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PushServiceToken_proto_rawDescOnce sync.Once
	file_PushServiceToken_proto_rawDescData = file_PushServiceToken_proto_rawDesc
)

func file_PushServiceToken_proto_rawDescGZIP() []byte {
	file_PushServiceToken_proto_rawDescOnce.Do(func() {
		file_PushServiceToken_proto_rawDescData = protoimpl.X.CompressGZIP(file_PushServiceToken_proto_rawDescData)
	})
	return file_PushServiceToken_proto_rawDescData
}

var file_PushServiceToken_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PushServiceToken_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PushServiceToken_proto_goTypes = []interface{}{
	(PushServiceToken_PushType)(0), // 0: AcFunDanmu.Im.Basic.PushServiceToken.PushType
	(*PushServiceToken)(nil),       // 1: AcFunDanmu.Im.Basic.PushServiceToken
}
var file_PushServiceToken_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.Im.Basic.PushServiceToken.pushType:type_name -> AcFunDanmu.Im.Basic.PushServiceToken.PushType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PushServiceToken_proto_init() }
func file_PushServiceToken_proto_init() {
	if File_PushServiceToken_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PushServiceToken_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushServiceToken); i {
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
			RawDescriptor: file_PushServiceToken_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PushServiceToken_proto_goTypes,
		DependencyIndexes: file_PushServiceToken_proto_depIdxs,
		EnumInfos:         file_PushServiceToken_proto_enumTypes,
		MessageInfos:      file_PushServiceToken_proto_msgTypes,
	}.Build()
	File_PushServiceToken_proto = out.File
	file_PushServiceToken_proto_rawDesc = nil
	file_PushServiceToken_proto_goTypes = nil
	file_PushServiceToken_proto_depIdxs = nil
}
