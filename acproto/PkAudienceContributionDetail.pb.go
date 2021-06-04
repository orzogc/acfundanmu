// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: PkAudienceContributionDetail.proto

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

type PkAudienceContributionDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A *ZtLiveUserInfo `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64           `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *PkAudienceContributionDetail) Reset() {
	*x = PkAudienceContributionDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PkAudienceContributionDetail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PkAudienceContributionDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PkAudienceContributionDetail) ProtoMessage() {}

func (x *PkAudienceContributionDetail) ProtoReflect() protoreflect.Message {
	mi := &file_PkAudienceContributionDetail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PkAudienceContributionDetail.ProtoReflect.Descriptor instead.
func (*PkAudienceContributionDetail) Descriptor() ([]byte, []int) {
	return file_PkAudienceContributionDetail_proto_rawDescGZIP(), []int{0}
}

func (x *PkAudienceContributionDetail) GetA() *ZtLiveUserInfo {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *PkAudienceContributionDetail) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

var File_PkAudienceContributionDetail_proto protoreflect.FileDescriptor

var file_PkAudienceContributionDetail_proto_rawDesc = []byte{
	0x0a, 0x22, 0x50, 0x6b, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x1a, 0x14, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x1c, 0x50, 0x6b, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x28, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x5a,
	0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x01, 0x61,
	0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x62, 0x42, 0x26,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a,
	0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61,
	0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PkAudienceContributionDetail_proto_rawDescOnce sync.Once
	file_PkAudienceContributionDetail_proto_rawDescData = file_PkAudienceContributionDetail_proto_rawDesc
)

func file_PkAudienceContributionDetail_proto_rawDescGZIP() []byte {
	file_PkAudienceContributionDetail_proto_rawDescOnce.Do(func() {
		file_PkAudienceContributionDetail_proto_rawDescData = protoimpl.X.CompressGZIP(file_PkAudienceContributionDetail_proto_rawDescData)
	})
	return file_PkAudienceContributionDetail_proto_rawDescData
}

var file_PkAudienceContributionDetail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PkAudienceContributionDetail_proto_goTypes = []interface{}{
	(*PkAudienceContributionDetail)(nil), // 0: AcFunDanmu.PkAudienceContributionDetail
	(*ZtLiveUserInfo)(nil),               // 1: AcFunDanmu.ZtLiveUserInfo
}
var file_PkAudienceContributionDetail_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.PkAudienceContributionDetail.a:type_name -> AcFunDanmu.ZtLiveUserInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PkAudienceContributionDetail_proto_init() }
func file_PkAudienceContributionDetail_proto_init() {
	if File_PkAudienceContributionDetail_proto != nil {
		return
	}
	file_ZtLiveUserInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PkAudienceContributionDetail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PkAudienceContributionDetail); i {
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
			RawDescriptor: file_PkAudienceContributionDetail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PkAudienceContributionDetail_proto_goTypes,
		DependencyIndexes: file_PkAudienceContributionDetail_proto_depIdxs,
		MessageInfos:      file_PkAudienceContributionDetail_proto_msgTypes,
	}.Build()
	File_PkAudienceContributionDetail_proto = out.File
	file_PkAudienceContributionDetail_proto_rawDesc = nil
	file_PkAudienceContributionDetail_proto_goTypes = nil
	file_PkAudienceContributionDetail_proto_depIdxs = nil
}
