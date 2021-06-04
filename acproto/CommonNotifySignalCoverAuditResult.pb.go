// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: CommonNotifySignalCoverAuditResult.proto

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

type CommonNotifySignalCoverAuditResult_AuditStatus int32

const (
	CommonNotifySignalCoverAuditResult_SUCCESS              CommonNotifySignalCoverAuditResult_AuditStatus = 0
	CommonNotifySignalCoverAuditResult_COVER_AUDIT_FAILED   CommonNotifySignalCoverAuditResult_AuditStatus = 1
	CommonNotifySignalCoverAuditResult_CAPTION_AUDIT_FAILED CommonNotifySignalCoverAuditResult_AuditStatus = 2
	CommonNotifySignalCoverAuditResult_BOTH_FAILED          CommonNotifySignalCoverAuditResult_AuditStatus = 3
)

// Enum value maps for CommonNotifySignalCoverAuditResult_AuditStatus.
var (
	CommonNotifySignalCoverAuditResult_AuditStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "COVER_AUDIT_FAILED",
		2: "CAPTION_AUDIT_FAILED",
		3: "BOTH_FAILED",
	}
	CommonNotifySignalCoverAuditResult_AuditStatus_value = map[string]int32{
		"SUCCESS":              0,
		"COVER_AUDIT_FAILED":   1,
		"CAPTION_AUDIT_FAILED": 2,
		"BOTH_FAILED":          3,
	}
)

func (x CommonNotifySignalCoverAuditResult_AuditStatus) Enum() *CommonNotifySignalCoverAuditResult_AuditStatus {
	p := new(CommonNotifySignalCoverAuditResult_AuditStatus)
	*p = x
	return p
}

func (x CommonNotifySignalCoverAuditResult_AuditStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonNotifySignalCoverAuditResult_AuditStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_CommonNotifySignalCoverAuditResult_proto_enumTypes[0].Descriptor()
}

func (CommonNotifySignalCoverAuditResult_AuditStatus) Type() protoreflect.EnumType {
	return &file_CommonNotifySignalCoverAuditResult_proto_enumTypes[0]
}

func (x CommonNotifySignalCoverAuditResult_AuditStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonNotifySignalCoverAuditResult_AuditStatus.Descriptor instead.
func (CommonNotifySignalCoverAuditResult_AuditStatus) EnumDescriptor() ([]byte, []int) {
	return file_CommonNotifySignalCoverAuditResult_proto_rawDescGZIP(), []int{0, 0}
}

type CommonNotifySignalCoverAuditResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuditStatus CommonNotifySignalCoverAuditResult_AuditStatus `protobuf:"varint,1,opt,name=auditStatus,proto3,enum=AcFunDanmu.CommonNotifySignalCoverAuditResult_AuditStatus" json:"auditStatus,omitempty"`
}

func (x *CommonNotifySignalCoverAuditResult) Reset() {
	*x = CommonNotifySignalCoverAuditResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonNotifySignalCoverAuditResult_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonNotifySignalCoverAuditResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonNotifySignalCoverAuditResult) ProtoMessage() {}

func (x *CommonNotifySignalCoverAuditResult) ProtoReflect() protoreflect.Message {
	mi := &file_CommonNotifySignalCoverAuditResult_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonNotifySignalCoverAuditResult.ProtoReflect.Descriptor instead.
func (*CommonNotifySignalCoverAuditResult) Descriptor() ([]byte, []int) {
	return file_CommonNotifySignalCoverAuditResult_proto_rawDescGZIP(), []int{0}
}

func (x *CommonNotifySignalCoverAuditResult) GetAuditStatus() CommonNotifySignalCoverAuditResult_AuditStatus {
	if x != nil {
		return x.AuditStatus
	}
	return CommonNotifySignalCoverAuditResult_SUCCESS
}

var File_CommonNotifySignalCoverAuditResult_proto protoreflect.FileDescriptor

var file_CommonNotifySignalCoverAuditResult_proto_rawDesc = []byte{
	0x0a, 0x28, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75,
	0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0xe1, 0x01, 0x0a, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x5c, 0x0a,
	0x0b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5d, 0x0a, 0x0b, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x56, 0x45, 0x52,
	0x5f, 0x41, 0x55, 0x44, 0x49, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x18, 0x0a, 0x14, 0x43, 0x41, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x54,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x4f, 0x54,
	0x48, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f,
	0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonNotifySignalCoverAuditResult_proto_rawDescOnce sync.Once
	file_CommonNotifySignalCoverAuditResult_proto_rawDescData = file_CommonNotifySignalCoverAuditResult_proto_rawDesc
)

func file_CommonNotifySignalCoverAuditResult_proto_rawDescGZIP() []byte {
	file_CommonNotifySignalCoverAuditResult_proto_rawDescOnce.Do(func() {
		file_CommonNotifySignalCoverAuditResult_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonNotifySignalCoverAuditResult_proto_rawDescData)
	})
	return file_CommonNotifySignalCoverAuditResult_proto_rawDescData
}

var file_CommonNotifySignalCoverAuditResult_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CommonNotifySignalCoverAuditResult_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonNotifySignalCoverAuditResult_proto_goTypes = []interface{}{
	(CommonNotifySignalCoverAuditResult_AuditStatus)(0), // 0: AcFunDanmu.CommonNotifySignalCoverAuditResult.AuditStatus
	(*CommonNotifySignalCoverAuditResult)(nil),          // 1: AcFunDanmu.CommonNotifySignalCoverAuditResult
}
var file_CommonNotifySignalCoverAuditResult_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.CommonNotifySignalCoverAuditResult.auditStatus:type_name -> AcFunDanmu.CommonNotifySignalCoverAuditResult.AuditStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CommonNotifySignalCoverAuditResult_proto_init() }
func file_CommonNotifySignalCoverAuditResult_proto_init() {
	if File_CommonNotifySignalCoverAuditResult_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CommonNotifySignalCoverAuditResult_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonNotifySignalCoverAuditResult); i {
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
			RawDescriptor: file_CommonNotifySignalCoverAuditResult_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonNotifySignalCoverAuditResult_proto_goTypes,
		DependencyIndexes: file_CommonNotifySignalCoverAuditResult_proto_depIdxs,
		EnumInfos:         file_CommonNotifySignalCoverAuditResult_proto_enumTypes,
		MessageInfos:      file_CommonNotifySignalCoverAuditResult_proto_msgTypes,
	}.Build()
	File_CommonNotifySignalCoverAuditResult_proto = out.File
	file_CommonNotifySignalCoverAuditResult_proto_rawDesc = nil
	file_CommonNotifySignalCoverAuditResult_proto_goTypes = nil
	file_CommonNotifySignalCoverAuditResult_proto_depIdxs = nil
}
