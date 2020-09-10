// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: CsAckErrorCode.proto

package acproto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CsAckErrorCode int32

const (
	CsAckErrorCode_SUCCESS_CS_ACK                  CsAckErrorCode = 0
	CsAckErrorCode_LIVE_CLOSED                     CsAckErrorCode = 1
	CsAckErrorCode_TICKET_ILLEGAL                  CsAckErrorCode = 2
	CsAckErrorCode_ATTACH_ILLEGAL                  CsAckErrorCode = 3
	CsAckErrorCode_USER_NOT_IN_ROOM                CsAckErrorCode = 4
	CsAckErrorCode_SERVER_ERROR                    CsAckErrorCode = 5
	CsAckErrorCode_REQUEST_PARAM_INVALID           CsAckErrorCode = 6
	CsAckErrorCode_ROOM_NOT_EXIST_IN_STATE_MANAGER CsAckErrorCode = 7
	CsAckErrorCode_NEW_LIVE_OPENED                 CsAckErrorCode = 8
)

// Enum value maps for CsAckErrorCode.
var (
	CsAckErrorCode_name = map[int32]string{
		0: "SUCCESS_CS_ACK",
		1: "LIVE_CLOSED",
		2: "TICKET_ILLEGAL",
		3: "ATTACH_ILLEGAL",
		4: "USER_NOT_IN_ROOM",
		5: "SERVER_ERROR",
		6: "REQUEST_PARAM_INVALID",
		7: "ROOM_NOT_EXIST_IN_STATE_MANAGER",
		8: "NEW_LIVE_OPENED",
	}
	CsAckErrorCode_value = map[string]int32{
		"SUCCESS_CS_ACK":                  0,
		"LIVE_CLOSED":                     1,
		"TICKET_ILLEGAL":                  2,
		"ATTACH_ILLEGAL":                  3,
		"USER_NOT_IN_ROOM":                4,
		"SERVER_ERROR":                    5,
		"REQUEST_PARAM_INVALID":           6,
		"ROOM_NOT_EXIST_IN_STATE_MANAGER": 7,
		"NEW_LIVE_OPENED":                 8,
	}
)

func (x CsAckErrorCode) Enum() *CsAckErrorCode {
	p := new(CsAckErrorCode)
	*p = x
	return p
}

func (x CsAckErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CsAckErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_CsAckErrorCode_proto_enumTypes[0].Descriptor()
}

func (CsAckErrorCode) Type() protoreflect.EnumType {
	return &file_CsAckErrorCode_proto_enumTypes[0]
}

func (x CsAckErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CsAckErrorCode.Descriptor instead.
func (CsAckErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_CsAckErrorCode_proto_rawDescGZIP(), []int{0}
}

var File_CsAckErrorCode_proto protoreflect.FileDescriptor

var file_CsAckErrorCode_proto_rawDesc = []byte{
	0x0a, 0x14, 0x43, 0x73, 0x41, 0x63, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e,
	0x6d, 0x75, 0x2a, 0xda, 0x01, 0x0a, 0x0e, 0x43, 0x73, 0x41, 0x63, 0x6b, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x5f, 0x43, 0x53, 0x5f, 0x41, 0x43, 0x4b, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x49, 0x56,
	0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x49,
	0x43, 0x4b, 0x45, 0x54, 0x5f, 0x49, 0x4c, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x41, 0x54, 0x54, 0x41, 0x43, 0x48, 0x5f, 0x49, 0x4c, 0x4c, 0x45, 0x47, 0x41, 0x4c,
	0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49,
	0x4e, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x52, 0x56,
	0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x06, 0x12, 0x23, 0x0a, 0x1f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x45,
	0x57, 0x5f, 0x4c, 0x49, 0x56, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x45, 0x44, 0x10, 0x08, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CsAckErrorCode_proto_rawDescOnce sync.Once
	file_CsAckErrorCode_proto_rawDescData = file_CsAckErrorCode_proto_rawDesc
)

func file_CsAckErrorCode_proto_rawDescGZIP() []byte {
	file_CsAckErrorCode_proto_rawDescOnce.Do(func() {
		file_CsAckErrorCode_proto_rawDescData = protoimpl.X.CompressGZIP(file_CsAckErrorCode_proto_rawDescData)
	})
	return file_CsAckErrorCode_proto_rawDescData
}

var file_CsAckErrorCode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CsAckErrorCode_proto_goTypes = []interface{}{
	(CsAckErrorCode)(0), // 0: AcFunDanmu.CsAckErrorCode
}
var file_CsAckErrorCode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CsAckErrorCode_proto_init() }
func file_CsAckErrorCode_proto_init() {
	if File_CsAckErrorCode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CsAckErrorCode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CsAckErrorCode_proto_goTypes,
		DependencyIndexes: file_CsAckErrorCode_proto_depIdxs,
		EnumInfos:         file_CsAckErrorCode_proto_enumTypes,
	}.Build()
	File_CsAckErrorCode_proto = out.File
	file_CsAckErrorCode_proto_rawDesc = nil
	file_CsAckErrorCode_proto_goTypes = nil
	file_CsAckErrorCode_proto_depIdxs = nil
}
