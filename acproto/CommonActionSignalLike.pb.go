// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: CommonActionSignalLike.proto

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

type CommonActionSignalLike struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo   *ZtLiveUserInfo `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	SendTimeMs int64           `protobuf:"varint,2,opt,name=sendTimeMs,proto3" json:"sendTimeMs,omitempty"`
}

func (x *CommonActionSignalLike) Reset() {
	*x = CommonActionSignalLike{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonActionSignalLike_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonActionSignalLike) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonActionSignalLike) ProtoMessage() {}

func (x *CommonActionSignalLike) ProtoReflect() protoreflect.Message {
	mi := &file_CommonActionSignalLike_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonActionSignalLike.ProtoReflect.Descriptor instead.
func (*CommonActionSignalLike) Descriptor() ([]byte, []int) {
	return file_CommonActionSignalLike_proto_rawDescGZIP(), []int{0}
}

func (x *CommonActionSignalLike) GetUserInfo() *ZtLiveUserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *CommonActionSignalLike) GetSendTimeMs() int64 {
	if x != nil {
		return x.SendTimeMs
	}
	return 0
}

var File_CommonActionSignalLike_proto protoreflect.FileDescriptor

var file_CommonActionSignalLike_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x1a, 0x14, 0x5a, 0x74, 0x4c, 0x69,
	0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x70, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x41,
	0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x4d, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e,
	0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_CommonActionSignalLike_proto_rawDescOnce sync.Once
	file_CommonActionSignalLike_proto_rawDescData = file_CommonActionSignalLike_proto_rawDesc
)

func file_CommonActionSignalLike_proto_rawDescGZIP() []byte {
	file_CommonActionSignalLike_proto_rawDescOnce.Do(func() {
		file_CommonActionSignalLike_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonActionSignalLike_proto_rawDescData)
	})
	return file_CommonActionSignalLike_proto_rawDescData
}

var file_CommonActionSignalLike_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonActionSignalLike_proto_goTypes = []interface{}{
	(*CommonActionSignalLike)(nil), // 0: AcFunDanmu.CommonActionSignalLike
	(*ZtLiveUserInfo)(nil),         // 1: AcFunDanmu.ZtLiveUserInfo
}
var file_CommonActionSignalLike_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.CommonActionSignalLike.userInfo:type_name -> AcFunDanmu.ZtLiveUserInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CommonActionSignalLike_proto_init() }
func file_CommonActionSignalLike_proto_init() {
	if File_CommonActionSignalLike_proto != nil {
		return
	}
	file_ZtLiveUserInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CommonActionSignalLike_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonActionSignalLike); i {
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
			RawDescriptor: file_CommonActionSignalLike_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonActionSignalLike_proto_goTypes,
		DependencyIndexes: file_CommonActionSignalLike_proto_depIdxs,
		MessageInfos:      file_CommonActionSignalLike_proto_msgTypes,
	}.Build()
	File_CommonActionSignalLike_proto = out.File
	file_CommonActionSignalLike_proto_rawDesc = nil
	file_CommonActionSignalLike_proto_goTypes = nil
	file_CommonActionSignalLike_proto_depIdxs = nil
}
