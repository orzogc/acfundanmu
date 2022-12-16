// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: ZtLiveUserInfo.proto

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

type ZtLiveUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64               `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Nickname     string              `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar       []*ImageCdnNode     `protobuf:"bytes,3,rep,name=avatar,proto3" json:"avatar,omitempty"`
	Badge        string              `protobuf:"bytes,4,opt,name=badge,proto3" json:"badge,omitempty"` // bizCustomInfo
	UserIdentity *ZtLiveUserIdentity `protobuf:"bytes,5,opt,name=userIdentity,proto3" json:"userIdentity,omitempty"`
	F            bool                `protobuf:"varint,6,opt,name=f,proto3" json:"f,omitempty"`
	G            string              `protobuf:"bytes,7,opt,name=g,proto3" json:"g,omitempty"`
}

func (x *ZtLiveUserInfo) Reset() {
	*x = ZtLiveUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ZtLiveUserInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZtLiveUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZtLiveUserInfo) ProtoMessage() {}

func (x *ZtLiveUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ZtLiveUserInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZtLiveUserInfo.ProtoReflect.Descriptor instead.
func (*ZtLiveUserInfo) Descriptor() ([]byte, []int) {
	return file_ZtLiveUserInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ZtLiveUserInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ZtLiveUserInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ZtLiveUserInfo) GetAvatar() []*ImageCdnNode {
	if x != nil {
		return x.Avatar
	}
	return nil
}

func (x *ZtLiveUserInfo) GetBadge() string {
	if x != nil {
		return x.Badge
	}
	return ""
}

func (x *ZtLiveUserInfo) GetUserIdentity() *ZtLiveUserIdentity {
	if x != nil {
		return x.UserIdentity
	}
	return nil
}

func (x *ZtLiveUserInfo) GetF() bool {
	if x != nil {
		return x.F
	}
	return false
}

func (x *ZtLiveUserInfo) GetG() string {
	if x != nil {
		return x.G
	}
	return ""
}

var File_ZtLiveUserInfo_proto protoreflect.FileDescriptor

var file_ZtLiveUserInfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e,
	0x6d, 0x75, 0x1a, 0x12, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x64, 0x6e, 0x4e, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xec, 0x01, 0x0a, 0x0e, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44,
	0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x64, 0x6e, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x61, 0x64,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x61, 0x64, 0x67, 0x65, 0x12,
	0x42, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e,
	0x6d, 0x75, 0x2e, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x01,
	0x66, 0x12, 0x0c, 0x0a, 0x01, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x67, 0x42,
	0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72,
	0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f,
	0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ZtLiveUserInfo_proto_rawDescOnce sync.Once
	file_ZtLiveUserInfo_proto_rawDescData = file_ZtLiveUserInfo_proto_rawDesc
)

func file_ZtLiveUserInfo_proto_rawDescGZIP() []byte {
	file_ZtLiveUserInfo_proto_rawDescOnce.Do(func() {
		file_ZtLiveUserInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ZtLiveUserInfo_proto_rawDescData)
	})
	return file_ZtLiveUserInfo_proto_rawDescData
}

var file_ZtLiveUserInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ZtLiveUserInfo_proto_goTypes = []interface{}{
	(*ZtLiveUserInfo)(nil),     // 0: AcFunDanmu.ZtLiveUserInfo
	(*ImageCdnNode)(nil),       // 1: AcFunDanmu.ImageCdnNode
	(*ZtLiveUserIdentity)(nil), // 2: AcFunDanmu.ZtLiveUserIdentity
}
var file_ZtLiveUserInfo_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.ZtLiveUserInfo.avatar:type_name -> AcFunDanmu.ImageCdnNode
	2, // 1: AcFunDanmu.ZtLiveUserInfo.userIdentity:type_name -> AcFunDanmu.ZtLiveUserIdentity
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ZtLiveUserInfo_proto_init() }
func file_ZtLiveUserInfo_proto_init() {
	if File_ZtLiveUserInfo_proto != nil {
		return
	}
	file_ImageCdnNode_proto_init()
	file_ZtLiveUserIdentity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ZtLiveUserInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZtLiveUserInfo); i {
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
			RawDescriptor: file_ZtLiveUserInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ZtLiveUserInfo_proto_goTypes,
		DependencyIndexes: file_ZtLiveUserInfo_proto_depIdxs,
		MessageInfos:      file_ZtLiveUserInfo_proto_msgTypes,
	}.Build()
	File_ZtLiveUserInfo_proto = out.File
	file_ZtLiveUserInfo_proto_rawDesc = nil
	file_ZtLiveUserInfo_proto_goTypes = nil
	file_ZtLiveUserInfo_proto_depIdxs = nil
}
