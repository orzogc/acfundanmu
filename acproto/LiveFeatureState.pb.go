// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: LiveFeatureState.proto

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

type LiveFeatureState_FeatureType int32

const (
	LiveFeatureState_FEATURE_UNKNOWN   LiveFeatureState_FeatureType = 0
	LiveFeatureState_LANDSCAPE_COMMENT LiveFeatureState_FeatureType = 1
)

// Enum value maps for LiveFeatureState_FeatureType.
var (
	LiveFeatureState_FeatureType_name = map[int32]string{
		0: "FEATURE_UNKNOWN",
		1: "LANDSCAPE_COMMENT",
	}
	LiveFeatureState_FeatureType_value = map[string]int32{
		"FEATURE_UNKNOWN":   0,
		"LANDSCAPE_COMMENT": 1,
	}
)

func (x LiveFeatureState_FeatureType) Enum() *LiveFeatureState_FeatureType {
	p := new(LiveFeatureState_FeatureType)
	*p = x
	return p
}

func (x LiveFeatureState_FeatureType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LiveFeatureState_FeatureType) Descriptor() protoreflect.EnumDescriptor {
	return file_LiveFeatureState_proto_enumTypes[0].Descriptor()
}

func (LiveFeatureState_FeatureType) Type() protoreflect.EnumType {
	return &file_LiveFeatureState_proto_enumTypes[0]
}

func (x LiveFeatureState_FeatureType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LiveFeatureState_FeatureType.Descriptor instead.
func (LiveFeatureState_FeatureType) EnumDescriptor() ([]byte, []int) {
	return file_LiveFeatureState_proto_rawDescGZIP(), []int{0, 0}
}

type LiveFeatureState_FeatureState int32

const (
	LiveFeatureState_FEATURE_STATE_UNKNOWN LiveFeatureState_FeatureState = 0
	LiveFeatureState_FEATURE_STATE_OPEND   LiveFeatureState_FeatureState = 1
	LiveFeatureState_FEATURE_STATE_CLOSED  LiveFeatureState_FeatureState = 2
)

// Enum value maps for LiveFeatureState_FeatureState.
var (
	LiveFeatureState_FeatureState_name = map[int32]string{
		0: "FEATURE_STATE_UNKNOWN",
		1: "FEATURE_STATE_OPEND",
		2: "FEATURE_STATE_CLOSED",
	}
	LiveFeatureState_FeatureState_value = map[string]int32{
		"FEATURE_STATE_UNKNOWN": 0,
		"FEATURE_STATE_OPEND":   1,
		"FEATURE_STATE_CLOSED":  2,
	}
)

func (x LiveFeatureState_FeatureState) Enum() *LiveFeatureState_FeatureState {
	p := new(LiveFeatureState_FeatureState)
	*p = x
	return p
}

func (x LiveFeatureState_FeatureState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LiveFeatureState_FeatureState) Descriptor() protoreflect.EnumDescriptor {
	return file_LiveFeatureState_proto_enumTypes[1].Descriptor()
}

func (LiveFeatureState_FeatureState) Type() protoreflect.EnumType {
	return &file_LiveFeatureState_proto_enumTypes[1]
}

func (x LiveFeatureState_FeatureState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LiveFeatureState_FeatureState.Descriptor instead.
func (LiveFeatureState_FeatureState) EnumDescriptor() ([]byte, []int) {
	return file_LiveFeatureState_proto_rawDescGZIP(), []int{0, 1}
}

type LiveFeatureState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  LiveFeatureState_FeatureType  `protobuf:"varint,1,opt,name=type,proto3,enum=AcFunDanmu.LiveFeatureState_FeatureType" json:"type,omitempty"`
	State LiveFeatureState_FeatureState `protobuf:"varint,2,opt,name=state,proto3,enum=AcFunDanmu.LiveFeatureState_FeatureState" json:"state,omitempty"`
}

func (x *LiveFeatureState) Reset() {
	*x = LiveFeatureState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LiveFeatureState_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveFeatureState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveFeatureState) ProtoMessage() {}

func (x *LiveFeatureState) ProtoReflect() protoreflect.Message {
	mi := &file_LiveFeatureState_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveFeatureState.ProtoReflect.Descriptor instead.
func (*LiveFeatureState) Descriptor() ([]byte, []int) {
	return file_LiveFeatureState_proto_rawDescGZIP(), []int{0}
}

func (x *LiveFeatureState) GetType() LiveFeatureState_FeatureType {
	if x != nil {
		return x.Type
	}
	return LiveFeatureState_FEATURE_UNKNOWN
}

func (x *LiveFeatureState) GetState() LiveFeatureState_FeatureState {
	if x != nil {
		return x.State
	}
	return LiveFeatureState_FEATURE_STATE_UNKNOWN
}

var File_LiveFeatureState_proto protoreflect.FileDescriptor

var file_LiveFeatureState_proto_rawDesc = []byte{
	0x0a, 0x16, 0x4c, 0x69, 0x76, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44,
	0x61, 0x6e, 0x6d, 0x75, 0x22, 0xaa, 0x02, 0x0a, 0x10, 0x4c, 0x69, 0x76, 0x65, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44,
	0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61,
	0x6e, 0x6d, 0x75, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x39, 0x0a, 0x0b, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x45, 0x41, 0x54, 0x55,
	0x52, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11,
	0x4c, 0x41, 0x4e, 0x44, 0x53, 0x43, 0x41, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x01, 0x22, 0x5c, 0x0a, 0x0c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17,
	0x0a, 0x13, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x4f, 0x50, 0x45, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x45, 0x41, 0x54, 0x55,
	0x52, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10,
	0x02, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d,
	0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_LiveFeatureState_proto_rawDescOnce sync.Once
	file_LiveFeatureState_proto_rawDescData = file_LiveFeatureState_proto_rawDesc
)

func file_LiveFeatureState_proto_rawDescGZIP() []byte {
	file_LiveFeatureState_proto_rawDescOnce.Do(func() {
		file_LiveFeatureState_proto_rawDescData = protoimpl.X.CompressGZIP(file_LiveFeatureState_proto_rawDescData)
	})
	return file_LiveFeatureState_proto_rawDescData
}

var file_LiveFeatureState_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_LiveFeatureState_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LiveFeatureState_proto_goTypes = []interface{}{
	(LiveFeatureState_FeatureType)(0),  // 0: AcFunDanmu.LiveFeatureState.FeatureType
	(LiveFeatureState_FeatureState)(0), // 1: AcFunDanmu.LiveFeatureState.FeatureState
	(*LiveFeatureState)(nil),           // 2: AcFunDanmu.LiveFeatureState
}
var file_LiveFeatureState_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.LiveFeatureState.type:type_name -> AcFunDanmu.LiveFeatureState.FeatureType
	1, // 1: AcFunDanmu.LiveFeatureState.state:type_name -> AcFunDanmu.LiveFeatureState.FeatureState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_LiveFeatureState_proto_init() }
func file_LiveFeatureState_proto_init() {
	if File_LiveFeatureState_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_LiveFeatureState_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveFeatureState); i {
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
			RawDescriptor: file_LiveFeatureState_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LiveFeatureState_proto_goTypes,
		DependencyIndexes: file_LiveFeatureState_proto_depIdxs,
		EnumInfos:         file_LiveFeatureState_proto_enumTypes,
		MessageInfos:      file_LiveFeatureState_proto_msgTypes,
	}.Build()
	File_LiveFeatureState_proto = out.File
	file_LiveFeatureState_proto_rawDesc = nil
	file_LiveFeatureState_proto_goTypes = nil
	file_LiveFeatureState_proto_depIdxs = nil
}
