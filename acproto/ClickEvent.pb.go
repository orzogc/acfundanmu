// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: ClickEvent.proto

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

type ClickEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	UrlType       ZtLiveCommonModelProto `protobuf:"varint,2,opt,name=urlType,proto3,enum=AcFunDanmu.ZtLiveCommonModelProto" json:"urlType,omitempty"`
	HeightPercent float32                `protobuf:"fixed32,3,opt,name=heightPercent,proto3" json:"heightPercent,omitempty"`
}

func (x *ClickEvent) Reset() {
	*x = ClickEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClickEvent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickEvent) ProtoMessage() {}

func (x *ClickEvent) ProtoReflect() protoreflect.Message {
	mi := &file_ClickEvent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickEvent.ProtoReflect.Descriptor instead.
func (*ClickEvent) Descriptor() ([]byte, []int) {
	return file_ClickEvent_proto_rawDescGZIP(), []int{0}
}

func (x *ClickEvent) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ClickEvent) GetUrlType() ZtLiveCommonModelProto {
	if x != nil {
		return x.UrlType
	}
	return ZtLiveCommonModelProto_ZtLiveCommonModelProtoA
}

func (x *ClickEvent) GetHeightPercent() float32 {
	if x != nil {
		return x.HeightPercent
	}
	return 0
}

var File_ClickEvent_proto protoreflect.FileDescriptor

var file_ClickEvent_proto_rawDesc = []byte{
	0x0a, 0x10, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x1a, 0x1c,
	0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a,
	0x0a, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x3c, 0x0a,
	0x07, 0x75, 0x72, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x5a, 0x74, 0x4c, 0x69,
	0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x07, 0x75, 0x72, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0d, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d,
	0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ClickEvent_proto_rawDescOnce sync.Once
	file_ClickEvent_proto_rawDescData = file_ClickEvent_proto_rawDesc
)

func file_ClickEvent_proto_rawDescGZIP() []byte {
	file_ClickEvent_proto_rawDescOnce.Do(func() {
		file_ClickEvent_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClickEvent_proto_rawDescData)
	})
	return file_ClickEvent_proto_rawDescData
}

var file_ClickEvent_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ClickEvent_proto_goTypes = []any{
	(*ClickEvent)(nil),          // 0: AcFunDanmu.ClickEvent
	(ZtLiveCommonModelProto)(0), // 1: AcFunDanmu.ZtLiveCommonModelProto
}
var file_ClickEvent_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.ClickEvent.urlType:type_name -> AcFunDanmu.ZtLiveCommonModelProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ClickEvent_proto_init() }
func file_ClickEvent_proto_init() {
	if File_ClickEvent_proto != nil {
		return
	}
	file_ZtLiveCommonModelProto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ClickEvent_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ClickEvent); i {
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
			RawDescriptor: file_ClickEvent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClickEvent_proto_goTypes,
		DependencyIndexes: file_ClickEvent_proto_depIdxs,
		MessageInfos:      file_ClickEvent_proto_msgTypes,
	}.Build()
	File_ClickEvent_proto = out.File
	file_ClickEvent_proto_rawDesc = nil
	file_ClickEvent_proto_goTypes = nil
	file_ClickEvent_proto_depIdxs = nil
}
