// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: CommonActionSignalUserFollowAuthor.proto

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

type CommonActionSignalUserFollowAuthor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo   *ZtLiveUserInfo `protobuf:"bytes,1,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	SendTimeMs int64           `protobuf:"varint,2,opt,name=sendTimeMs,proto3" json:"sendTimeMs,omitempty"`
}

func (x *CommonActionSignalUserFollowAuthor) Reset() {
	*x = CommonActionSignalUserFollowAuthor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonActionSignalUserFollowAuthor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonActionSignalUserFollowAuthor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonActionSignalUserFollowAuthor) ProtoMessage() {}

func (x *CommonActionSignalUserFollowAuthor) ProtoReflect() protoreflect.Message {
	mi := &file_CommonActionSignalUserFollowAuthor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonActionSignalUserFollowAuthor.ProtoReflect.Descriptor instead.
func (*CommonActionSignalUserFollowAuthor) Descriptor() ([]byte, []int) {
	return file_CommonActionSignalUserFollowAuthor_proto_rawDescGZIP(), []int{0}
}

func (x *CommonActionSignalUserFollowAuthor) GetUserInfo() *ZtLiveUserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *CommonActionSignalUserFollowAuthor) GetSendTimeMs() int64 {
	if x != nil {
		return x.SendTimeMs
	}
	return 0
}

var File_CommonActionSignalUserFollowAuthor_proto protoreflect.FileDescriptor

var file_CommonActionSignalUserFollowAuthor_proto_rawDesc = []byte{
	0x0a, 0x28, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75,
	0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x1a, 0x14, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x22,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d,
	0x75, 0x2e, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f,
	0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonActionSignalUserFollowAuthor_proto_rawDescOnce sync.Once
	file_CommonActionSignalUserFollowAuthor_proto_rawDescData = file_CommonActionSignalUserFollowAuthor_proto_rawDesc
)

func file_CommonActionSignalUserFollowAuthor_proto_rawDescGZIP() []byte {
	file_CommonActionSignalUserFollowAuthor_proto_rawDescOnce.Do(func() {
		file_CommonActionSignalUserFollowAuthor_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonActionSignalUserFollowAuthor_proto_rawDescData)
	})
	return file_CommonActionSignalUserFollowAuthor_proto_rawDescData
}

var file_CommonActionSignalUserFollowAuthor_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonActionSignalUserFollowAuthor_proto_goTypes = []any{
	(*CommonActionSignalUserFollowAuthor)(nil), // 0: AcFunDanmu.CommonActionSignalUserFollowAuthor
	(*ZtLiveUserInfo)(nil),                     // 1: AcFunDanmu.ZtLiveUserInfo
}
var file_CommonActionSignalUserFollowAuthor_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.CommonActionSignalUserFollowAuthor.userInfo:type_name -> AcFunDanmu.ZtLiveUserInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CommonActionSignalUserFollowAuthor_proto_init() }
func file_CommonActionSignalUserFollowAuthor_proto_init() {
	if File_CommonActionSignalUserFollowAuthor_proto != nil {
		return
	}
	file_ZtLiveUserInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CommonActionSignalUserFollowAuthor_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommonActionSignalUserFollowAuthor); i {
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
			RawDescriptor: file_CommonActionSignalUserFollowAuthor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonActionSignalUserFollowAuthor_proto_goTypes,
		DependencyIndexes: file_CommonActionSignalUserFollowAuthor_proto_depIdxs,
		MessageInfos:      file_CommonActionSignalUserFollowAuthor_proto_msgTypes,
	}.Build()
	File_CommonActionSignalUserFollowAuthor_proto = out.File
	file_CommonActionSignalUserFollowAuthor_proto_rawDesc = nil
	file_CommonActionSignalUserFollowAuthor_proto_goTypes = nil
	file_CommonActionSignalUserFollowAuthor_proto_depIdxs = nil
}
