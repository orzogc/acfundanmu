// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: CommonNotifySignalRemoveApplyUser.proto

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

type CommonNotifySignalRemoveApplyUserUnknown int32

const (
	CommonNotifySignalRemoveApplyUser_c CommonNotifySignalRemoveApplyUserUnknown = 0
	CommonNotifySignalRemoveApplyUser_d CommonNotifySignalRemoveApplyUserUnknown = 1
	CommonNotifySignalRemoveApplyUser_e CommonNotifySignalRemoveApplyUserUnknown = 2
	CommonNotifySignalRemoveApplyUser_f CommonNotifySignalRemoveApplyUserUnknown = 3
)

// Enum value maps for CommonNotifySignalRemoveApplyUserUnknown.
var (
	CommonNotifySignalRemoveApplyUserUnknown_name = map[int32]string{
		0: "c",
		1: "d",
		2: "e",
		3: "f",
	}
	CommonNotifySignalRemoveApplyUserUnknown_value = map[string]int32{
		"c": 0,
		"d": 1,
		"e": 2,
		"f": 3,
	}
)

func (x CommonNotifySignalRemoveApplyUserUnknown) Enum() *CommonNotifySignalRemoveApplyUserUnknown {
	p := new(CommonNotifySignalRemoveApplyUserUnknown)
	*p = x
	return p
}

func (x CommonNotifySignalRemoveApplyUserUnknown) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonNotifySignalRemoveApplyUserUnknown) Descriptor() protoreflect.EnumDescriptor {
	return file_CommonNotifySignalRemoveApplyUser_proto_enumTypes[0].Descriptor()
}

func (CommonNotifySignalRemoveApplyUserUnknown) Type() protoreflect.EnumType {
	return &file_CommonNotifySignalRemoveApplyUser_proto_enumTypes[0]
}

func (x CommonNotifySignalRemoveApplyUserUnknown) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonNotifySignalRemoveApplyUserUnknown.Descriptor instead.
func (CommonNotifySignalRemoveApplyUserUnknown) EnumDescriptor() ([]byte, []int) {
	return file_CommonNotifySignalRemoveApplyUser_proto_rawDescGZIP(), []int{0, 0}
}

type CommonNotifySignalRemoveApplyUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64                                    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B CommonNotifySignalRemoveApplyUserUnknown `protobuf:"varint,2,opt,name=b,proto3,enum=AcFunDanmu.CommonNotifySignalRemoveApplyUserUnknown" json:"b,omitempty"`
}

func (x *CommonNotifySignalRemoveApplyUser) Reset() {
	*x = CommonNotifySignalRemoveApplyUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonNotifySignalRemoveApplyUser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonNotifySignalRemoveApplyUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonNotifySignalRemoveApplyUser) ProtoMessage() {}

func (x *CommonNotifySignalRemoveApplyUser) ProtoReflect() protoreflect.Message {
	mi := &file_CommonNotifySignalRemoveApplyUser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonNotifySignalRemoveApplyUser.ProtoReflect.Descriptor instead.
func (*CommonNotifySignalRemoveApplyUser) Descriptor() ([]byte, []int) {
	return file_CommonNotifySignalRemoveApplyUser_proto_rawDescGZIP(), []int{0}
}

func (x *CommonNotifySignalRemoveApplyUser) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *CommonNotifySignalRemoveApplyUser) GetB() CommonNotifySignalRemoveApplyUserUnknown {
	if x != nil {
		return x.B
	}
	return CommonNotifySignalRemoveApplyUser_c
}

var File_CommonNotifySignalRemoveApplyUser_proto protoreflect.FileDescriptor

var file_CommonNotifySignalRemoveApplyUser_proto_rawDesc = []byte{
	0x0a, 0x27, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e,
	0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0x9d, 0x01, 0x0a, 0x21, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x61, 0x12, 0x43, 0x0a, 0x01, 0x62, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d,
	0x75, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x2e, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x01, 0x62, 0x22, 0x25,
	0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x05, 0x0a, 0x01, 0x63, 0x10, 0x00,
	0x12, 0x05, 0x0a, 0x01, 0x64, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x65, 0x10, 0x02, 0x12, 0x05,
	0x0a, 0x01, 0x66, 0x10, 0x03, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e,
	0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonNotifySignalRemoveApplyUser_proto_rawDescOnce sync.Once
	file_CommonNotifySignalRemoveApplyUser_proto_rawDescData = file_CommonNotifySignalRemoveApplyUser_proto_rawDesc
)

func file_CommonNotifySignalRemoveApplyUser_proto_rawDescGZIP() []byte {
	file_CommonNotifySignalRemoveApplyUser_proto_rawDescOnce.Do(func() {
		file_CommonNotifySignalRemoveApplyUser_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonNotifySignalRemoveApplyUser_proto_rawDescData)
	})
	return file_CommonNotifySignalRemoveApplyUser_proto_rawDescData
}

var file_CommonNotifySignalRemoveApplyUser_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CommonNotifySignalRemoveApplyUser_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonNotifySignalRemoveApplyUser_proto_goTypes = []interface{}{
	(CommonNotifySignalRemoveApplyUserUnknown)(0), // 0: AcFunDanmu.CommonNotifySignalRemoveApplyUser.unknown
	(*CommonNotifySignalRemoveApplyUser)(nil),     // 1: AcFunDanmu.CommonNotifySignalRemoveApplyUser
}
var file_CommonNotifySignalRemoveApplyUser_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.CommonNotifySignalRemoveApplyUser.b:type_name -> AcFunDanmu.CommonNotifySignalRemoveApplyUser.unknown
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CommonNotifySignalRemoveApplyUser_proto_init() }
func file_CommonNotifySignalRemoveApplyUser_proto_init() {
	if File_CommonNotifySignalRemoveApplyUser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CommonNotifySignalRemoveApplyUser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonNotifySignalRemoveApplyUser); i {
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
			RawDescriptor: file_CommonNotifySignalRemoveApplyUser_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonNotifySignalRemoveApplyUser_proto_goTypes,
		DependencyIndexes: file_CommonNotifySignalRemoveApplyUser_proto_depIdxs,
		EnumInfos:         file_CommonNotifySignalRemoveApplyUser_proto_enumTypes,
		MessageInfos:      file_CommonNotifySignalRemoveApplyUser_proto_msgTypes,
	}.Build()
	File_CommonNotifySignalRemoveApplyUser_proto = out.File
	file_CommonNotifySignalRemoveApplyUser_proto_rawDesc = nil
	file_CommonNotifySignalRemoveApplyUser_proto_goTypes = nil
	file_CommonNotifySignalRemoveApplyUser_proto_depIdxs = nil
}
