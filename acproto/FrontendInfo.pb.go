// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: FrontendInfo.proto

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

type FrontendInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *FrontendInfo) Reset() {
	*x = FrontendInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FrontendInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontendInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendInfo) ProtoMessage() {}

func (x *FrontendInfo) ProtoReflect() protoreflect.Message {
	mi := &file_FrontendInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendInfo.ProtoReflect.Descriptor instead.
func (*FrontendInfo) Descriptor() ([]byte, []int) {
	return file_FrontendInfo_proto_rawDescGZIP(), []int{0}
}

func (x *FrontendInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *FrontendInfo) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_FrontendInfo_proto protoreflect.FileDescriptor

var file_FrontendInfo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x22, 0x32, 0x0a, 0x0c, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64,
	0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FrontendInfo_proto_rawDescOnce sync.Once
	file_FrontendInfo_proto_rawDescData = file_FrontendInfo_proto_rawDesc
)

func file_FrontendInfo_proto_rawDescGZIP() []byte {
	file_FrontendInfo_proto_rawDescOnce.Do(func() {
		file_FrontendInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_FrontendInfo_proto_rawDescData)
	})
	return file_FrontendInfo_proto_rawDescData
}

var file_FrontendInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FrontendInfo_proto_goTypes = []interface{}{
	(*FrontendInfo)(nil), // 0: AcFunDanmu.FrontendInfo
}
var file_FrontendInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FrontendInfo_proto_init() }
func file_FrontendInfo_proto_init() {
	if File_FrontendInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FrontendInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontendInfo); i {
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
			RawDescriptor: file_FrontendInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FrontendInfo_proto_goTypes,
		DependencyIndexes: file_FrontendInfo_proto_depIdxs,
		MessageInfos:      file_FrontendInfo_proto_msgTypes,
	}.Build()
	File_FrontendInfo_proto = out.File
	file_FrontendInfo_proto_rawDesc = nil
	file_FrontendInfo_proto_goTypes = nil
	file_FrontendInfo_proto_depIdxs = nil
}
