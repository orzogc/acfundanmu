// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.17.3
// source: Ping.proto

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

type PingRequest_PingType int32

const (
	PingRequest_kInvalid       PingRequest_PingType = 0
	PingRequest_kPriorRegister PingRequest_PingType = 1
	PingRequest_kPostRegister  PingRequest_PingType = 2
)

// Enum value maps for PingRequest_PingType.
var (
	PingRequest_PingType_name = map[int32]string{
		0: "kInvalid",
		1: "kPriorRegister",
		2: "kPostRegister",
	}
	PingRequest_PingType_value = map[string]int32{
		"kInvalid":       0,
		"kPriorRegister": 1,
		"kPostRegister":  2,
	}
)

func (x PingRequest_PingType) Enum() *PingRequest_PingType {
	p := new(PingRequest_PingType)
	*p = x
	return p
}

func (x PingRequest_PingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PingRequest_PingType) Descriptor() protoreflect.EnumDescriptor {
	return file_Ping_proto_enumTypes[0].Descriptor()
}

func (PingRequest_PingType) Type() protoreflect.EnumType {
	return &file_Ping_proto_enumTypes[0]
}

func (x PingRequest_PingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PingRequest_PingType.Descriptor instead.
func (PingRequest_PingType) EnumDescriptor() ([]byte, []int) {
	return file_Ping_proto_rawDescGZIP(), []int{0, 0}
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PingType  PingRequest_PingType `protobuf:"varint,1,opt,name=pingType,proto3,enum=AcFunDanmu.PingRequest_PingType" json:"pingType,omitempty"`
	PingRound uint32               `protobuf:"varint,2,opt,name=pingRound,proto3" json:"pingRound,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Ping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_Ping_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetPingType() PingRequest_PingType {
	if x != nil {
		return x.PingType
	}
	return PingRequest_kInvalid
}

func (x *PingRequest) GetPingRound() uint32 {
	if x != nil {
		return x.PingRound
	}
	return 0
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerTimestamp int32  `protobuf:"fixed32,1,opt,name=serverTimestamp,proto3" json:"serverTimestamp,omitempty"`
	ClientIp        uint32 `protobuf:"fixed32,2,opt,name=clientIp,proto3" json:"clientIp,omitempty"`
	RedirectIp      uint32 `protobuf:"fixed32,3,opt,name=redirectIp,proto3" json:"redirectIp,omitempty"`
	RedirectPort    uint32 `protobuf:"varint,4,opt,name=redirectPort,proto3" json:"redirectPort,omitempty"`
	ClientIpV6      []byte `protobuf:"bytes,5,opt,name=clientIpV6,proto3" json:"clientIpV6,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Ping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_Ping_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetServerTimestamp() int32 {
	if x != nil {
		return x.ServerTimestamp
	}
	return 0
}

func (x *PingResponse) GetClientIp() uint32 {
	if x != nil {
		return x.ClientIp
	}
	return 0
}

func (x *PingResponse) GetRedirectIp() uint32 {
	if x != nil {
		return x.RedirectIp
	}
	return 0
}

func (x *PingResponse) GetRedirectPort() uint32 {
	if x != nil {
		return x.RedirectPort
	}
	return 0
}

func (x *PingResponse) GetClientIpV6() []byte {
	if x != nil {
		return x.ClientIpV6
	}
	return nil
}

var File_Ping_proto protoreflect.FileDescriptor

var file_Ping_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x50, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63,
	0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0xaa, 0x01, 0x0a, 0x0b, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x69, 0x6e, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x41, 0x63, 0x46,
	0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x70, 0x69,
	0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x6f,
	0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x69, 0x6e, 0x67, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x22, 0x3f, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0c, 0x0a, 0x08, 0x6b, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x6b, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x6b, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x10, 0x02, 0x22, 0xb8, 0x01, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0f, 0x52,
	0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x07, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x49, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x07,
	0x52, 0x0a, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x49, 0x70, 0x12, 0x22, 0x0a, 0x0c,
	0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x56, 0x36, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x56, 0x36,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75,
	0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Ping_proto_rawDescOnce sync.Once
	file_Ping_proto_rawDescData = file_Ping_proto_rawDesc
)

func file_Ping_proto_rawDescGZIP() []byte {
	file_Ping_proto_rawDescOnce.Do(func() {
		file_Ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_Ping_proto_rawDescData)
	})
	return file_Ping_proto_rawDescData
}

var file_Ping_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Ping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Ping_proto_goTypes = []interface{}{
	(PingRequest_PingType)(0), // 0: AcFunDanmu.PingRequest.PingType
	(*PingRequest)(nil),       // 1: AcFunDanmu.PingRequest
	(*PingResponse)(nil),      // 2: AcFunDanmu.PingResponse
}
var file_Ping_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.PingRequest.pingType:type_name -> AcFunDanmu.PingRequest.PingType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Ping_proto_init() }
func file_Ping_proto_init() {
	if File_Ping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Ping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_Ping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
			RawDescriptor: file_Ping_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Ping_proto_goTypes,
		DependencyIndexes: file_Ping_proto_depIdxs,
		EnumInfos:         file_Ping_proto_enumTypes,
		MessageInfos:      file_Ping_proto_msgTypes,
	}.Build()
	File_Ping_proto = out.File
	file_Ping_proto_rawDesc = nil
	file_Ping_proto_goTypes = nil
	file_Ping_proto_depIdxs = nil
}
