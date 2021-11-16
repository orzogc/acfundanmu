// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: SdkOption.proto

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

type SdkOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportIntervalSeconds        int32    `protobuf:"varint,1,opt,name=reportIntervalSeconds,proto3" json:"reportIntervalSeconds,omitempty"`
	ReportSecurity               string   `protobuf:"bytes,2,opt,name=reportSecurity,proto3" json:"reportSecurity,omitempty"`
	Lz4CompressionThresholdBytes int32    `protobuf:"varint,3,opt,name=lz4CompressionThresholdBytes,proto3" json:"lz4CompressionThresholdBytes,omitempty"`
	NetCheckServers              []string `protobuf:"bytes,4,rep,name=netCheckServers,proto3" json:"netCheckServers,omitempty"`
}

func (x *SdkOption) Reset() {
	*x = SdkOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SdkOption_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SdkOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SdkOption) ProtoMessage() {}

func (x *SdkOption) ProtoReflect() protoreflect.Message {
	mi := &file_SdkOption_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SdkOption.ProtoReflect.Descriptor instead.
func (*SdkOption) Descriptor() ([]byte, []int) {
	return file_SdkOption_proto_rawDescGZIP(), []int{0}
}

func (x *SdkOption) GetReportIntervalSeconds() int32 {
	if x != nil {
		return x.ReportIntervalSeconds
	}
	return 0
}

func (x *SdkOption) GetReportSecurity() string {
	if x != nil {
		return x.ReportSecurity
	}
	return ""
}

func (x *SdkOption) GetLz4CompressionThresholdBytes() int32 {
	if x != nil {
		return x.Lz4CompressionThresholdBytes
	}
	return 0
}

func (x *SdkOption) GetNetCheckServers() []string {
	if x != nil {
		return x.NetCheckServers
	}
	return nil
}

var File_SdkOption_proto protoreflect.FileDescriptor

var file_SdkOption_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x53, 0x64, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0xd7, 0x01,
	0x0a, 0x09, 0x53, 0x64, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x15, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x1c, 0x6c, 0x7a, 0x34,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x1c, 0x6c, 0x7a, 0x34, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x28, 0x0a,
	0x0f, 0x6e, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66,
	0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SdkOption_proto_rawDescOnce sync.Once
	file_SdkOption_proto_rawDescData = file_SdkOption_proto_rawDesc
)

func file_SdkOption_proto_rawDescGZIP() []byte {
	file_SdkOption_proto_rawDescOnce.Do(func() {
		file_SdkOption_proto_rawDescData = protoimpl.X.CompressGZIP(file_SdkOption_proto_rawDescData)
	})
	return file_SdkOption_proto_rawDescData
}

var file_SdkOption_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SdkOption_proto_goTypes = []interface{}{
	(*SdkOption)(nil), // 0: AcFunDanmu.SdkOption
}
var file_SdkOption_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SdkOption_proto_init() }
func file_SdkOption_proto_init() {
	if File_SdkOption_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SdkOption_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SdkOption); i {
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
			RawDescriptor: file_SdkOption_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SdkOption_proto_goTypes,
		DependencyIndexes: file_SdkOption_proto_depIdxs,
		MessageInfos:      file_SdkOption_proto_msgTypes,
	}.Build()
	File_SdkOption_proto = out.File
	file_SdkOption_proto_rawDesc = nil
	file_SdkOption_proto_goTypes = nil
	file_SdkOption_proto_depIdxs = nil
}
