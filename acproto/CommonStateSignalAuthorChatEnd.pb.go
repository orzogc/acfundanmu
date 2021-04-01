// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: CommonStateSignalAuthorChatEnd.proto

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

type CommonStateSignalAuthorChatEnd_EndType int32

const (
	CommonStateSignalAuthorChatEnd_UNKNOWN                   CommonStateSignalAuthorChatEnd_EndType = 0
	CommonStateSignalAuthorChatEnd_CANCEL_BY_INVITER         CommonStateSignalAuthorChatEnd_EndType = 1
	CommonStateSignalAuthorChatEnd_END_BY_INVITER            CommonStateSignalAuthorChatEnd_EndType = 2
	CommonStateSignalAuthorChatEnd_END_BY_INVITEE            CommonStateSignalAuthorChatEnd_EndType = 3
	CommonStateSignalAuthorChatEnd_INVITEE_REJECT            CommonStateSignalAuthorChatEnd_EndType = 4
	CommonStateSignalAuthorChatEnd_INVITEE_TIMEOUT           CommonStateSignalAuthorChatEnd_EndType = 5
	CommonStateSignalAuthorChatEnd_INVITEE_HEARTBEAT_TIMEOUT CommonStateSignalAuthorChatEnd_EndType = 6
	CommonStateSignalAuthorChatEnd_INVITER_HEARTBEAT_TIMEOUT CommonStateSignalAuthorChatEnd_EndType = 7
	CommonStateSignalAuthorChatEnd_PEER_LIVE_STOPPED         CommonStateSignalAuthorChatEnd_EndType = 8
)

// Enum value maps for CommonStateSignalAuthorChatEnd_EndType.
var (
	CommonStateSignalAuthorChatEnd_EndType_name = map[int32]string{
		0: "UNKNOWN",
		1: "CANCEL_BY_INVITER",
		2: "END_BY_INVITER",
		3: "END_BY_INVITEE",
		4: "INVITEE_REJECT",
		5: "INVITEE_TIMEOUT",
		6: "INVITEE_HEARTBEAT_TIMEOUT",
		7: "INVITER_HEARTBEAT_TIMEOUT",
		8: "PEER_LIVE_STOPPED",
	}
	CommonStateSignalAuthorChatEnd_EndType_value = map[string]int32{
		"UNKNOWN":                   0,
		"CANCEL_BY_INVITER":         1,
		"END_BY_INVITER":            2,
		"END_BY_INVITEE":            3,
		"INVITEE_REJECT":            4,
		"INVITEE_TIMEOUT":           5,
		"INVITEE_HEARTBEAT_TIMEOUT": 6,
		"INVITER_HEARTBEAT_TIMEOUT": 7,
		"PEER_LIVE_STOPPED":         8,
	}
)

func (x CommonStateSignalAuthorChatEnd_EndType) Enum() *CommonStateSignalAuthorChatEnd_EndType {
	p := new(CommonStateSignalAuthorChatEnd_EndType)
	*p = x
	return p
}

func (x CommonStateSignalAuthorChatEnd_EndType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonStateSignalAuthorChatEnd_EndType) Descriptor() protoreflect.EnumDescriptor {
	return file_CommonStateSignalAuthorChatEnd_proto_enumTypes[0].Descriptor()
}

func (CommonStateSignalAuthorChatEnd_EndType) Type() protoreflect.EnumType {
	return &file_CommonStateSignalAuthorChatEnd_proto_enumTypes[0]
}

func (x CommonStateSignalAuthorChatEnd_EndType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonStateSignalAuthorChatEnd_EndType.Descriptor instead.
func (CommonStateSignalAuthorChatEnd_EndType) EnumDescriptor() ([]byte, []int) {
	return file_CommonStateSignalAuthorChatEnd_proto_rawDescGZIP(), []int{0, 0}
}

type CommonStateSignalAuthorChatEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorChatId string                                 `protobuf:"bytes,1,opt,name=authorChatId,proto3" json:"authorChatId,omitempty"`
	EndType      CommonStateSignalAuthorChatEnd_EndType `protobuf:"varint,2,opt,name=endType,proto3,enum=AcFunDanmu.CommonStateSignalAuthorChatEnd_EndType" json:"endType,omitempty"`
	EndLiveId    string                                 `protobuf:"bytes,3,opt,name=endLiveId,proto3" json:"endLiveId,omitempty"`
}

func (x *CommonStateSignalAuthorChatEnd) Reset() {
	*x = CommonStateSignalAuthorChatEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonStateSignalAuthorChatEnd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonStateSignalAuthorChatEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonStateSignalAuthorChatEnd) ProtoMessage() {}

func (x *CommonStateSignalAuthorChatEnd) ProtoReflect() protoreflect.Message {
	mi := &file_CommonStateSignalAuthorChatEnd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonStateSignalAuthorChatEnd.ProtoReflect.Descriptor instead.
func (*CommonStateSignalAuthorChatEnd) Descriptor() ([]byte, []int) {
	return file_CommonStateSignalAuthorChatEnd_proto_rawDescGZIP(), []int{0}
}

func (x *CommonStateSignalAuthorChatEnd) GetAuthorChatId() string {
	if x != nil {
		return x.AuthorChatId
	}
	return ""
}

func (x *CommonStateSignalAuthorChatEnd) GetEndType() CommonStateSignalAuthorChatEnd_EndType {
	if x != nil {
		return x.EndType
	}
	return CommonStateSignalAuthorChatEnd_UNKNOWN
}

func (x *CommonStateSignalAuthorChatEnd) GetEndLiveId() string {
	if x != nil {
		return x.EndLiveId
	}
	return ""
}

var File_CommonStateSignalAuthorChatEnd_proto protoreflect.FileDescriptor

var file_CommonStateSignalAuthorChatEnd_proto_rawDesc = []byte{
	0x0a, 0x24, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x74, 0x45, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e,
	0x6d, 0x75, 0x22, 0x86, 0x03, 0x0a, 0x1e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x43, 0x68,
	0x61, 0x74, 0x45, 0x6e, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x43,
	0x68, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x43, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x41, 0x63, 0x46,
	0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x43,
	0x68, 0x61, 0x74, 0x45, 0x6e, 0x64, 0x2e, 0x45, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x4c, 0x69,
	0x76, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x4c,
	0x69, 0x76, 0x65, 0x49, 0x64, 0x22, 0xd3, 0x01, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x15,
	0x0a, 0x11, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x42, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x49,
	0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x4e, 0x44, 0x5f, 0x42, 0x59, 0x5f,
	0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x4e, 0x44,
	0x5f, 0x42, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x45, 0x10, 0x03, 0x12, 0x12, 0x0a,
	0x0e, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x45, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10,
	0x04, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x45, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x4f, 0x55, 0x54, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45,
	0x45, 0x5f, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x4f, 0x55, 0x54, 0x10, 0x06, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x52,
	0x5f, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f,
	0x55, 0x54, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x45, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x56,
	0x45, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x08, 0x42, 0x26, 0x5a, 0x24, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63,
	0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonStateSignalAuthorChatEnd_proto_rawDescOnce sync.Once
	file_CommonStateSignalAuthorChatEnd_proto_rawDescData = file_CommonStateSignalAuthorChatEnd_proto_rawDesc
)

func file_CommonStateSignalAuthorChatEnd_proto_rawDescGZIP() []byte {
	file_CommonStateSignalAuthorChatEnd_proto_rawDescOnce.Do(func() {
		file_CommonStateSignalAuthorChatEnd_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonStateSignalAuthorChatEnd_proto_rawDescData)
	})
	return file_CommonStateSignalAuthorChatEnd_proto_rawDescData
}

var file_CommonStateSignalAuthorChatEnd_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CommonStateSignalAuthorChatEnd_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonStateSignalAuthorChatEnd_proto_goTypes = []interface{}{
	(CommonStateSignalAuthorChatEnd_EndType)(0), // 0: AcFunDanmu.CommonStateSignalAuthorChatEnd.EndType
	(*CommonStateSignalAuthorChatEnd)(nil),      // 1: AcFunDanmu.CommonStateSignalAuthorChatEnd
}
var file_CommonStateSignalAuthorChatEnd_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.CommonStateSignalAuthorChatEnd.endType:type_name -> AcFunDanmu.CommonStateSignalAuthorChatEnd.EndType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CommonStateSignalAuthorChatEnd_proto_init() }
func file_CommonStateSignalAuthorChatEnd_proto_init() {
	if File_CommonStateSignalAuthorChatEnd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CommonStateSignalAuthorChatEnd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonStateSignalAuthorChatEnd); i {
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
			RawDescriptor: file_CommonStateSignalAuthorChatEnd_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonStateSignalAuthorChatEnd_proto_goTypes,
		DependencyIndexes: file_CommonStateSignalAuthorChatEnd_proto_depIdxs,
		EnumInfos:         file_CommonStateSignalAuthorChatEnd_proto_enumTypes,
		MessageInfos:      file_CommonStateSignalAuthorChatEnd_proto_msgTypes,
	}.Build()
	File_CommonStateSignalAuthorChatEnd_proto = out.File
	file_CommonStateSignalAuthorChatEnd_proto_rawDesc = nil
	file_CommonStateSignalAuthorChatEnd_proto_goTypes = nil
	file_CommonStateSignalAuthorChatEnd_proto_depIdxs = nil
}
