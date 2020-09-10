// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: ZtLiveScActionSignal.proto

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

type ZtLiveScActionSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item []*ZtLiveActionSignalItem `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
}

func (x *ZtLiveScActionSignal) Reset() {
	*x = ZtLiveScActionSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ZtLiveScActionSignal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZtLiveScActionSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZtLiveScActionSignal) ProtoMessage() {}

func (x *ZtLiveScActionSignal) ProtoReflect() protoreflect.Message {
	mi := &file_ZtLiveScActionSignal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZtLiveScActionSignal.ProtoReflect.Descriptor instead.
func (*ZtLiveScActionSignal) Descriptor() ([]byte, []int) {
	return file_ZtLiveScActionSignal_proto_rawDescGZIP(), []int{0}
}

func (x *ZtLiveScActionSignal) GetItem() []*ZtLiveActionSignalItem {
	if x != nil {
		return x.Item
	}
	return nil
}

var File_ZtLiveScActionSignal_proto protoreflect.FileDescriptor

var file_ZtLiveScActionSignal_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x63, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63,
	0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x1a, 0x1c, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x14, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65,
	0x53, 0x63, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x36,
	0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x41,
	0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x61, 0x63, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ZtLiveScActionSignal_proto_rawDescOnce sync.Once
	file_ZtLiveScActionSignal_proto_rawDescData = file_ZtLiveScActionSignal_proto_rawDesc
)

func file_ZtLiveScActionSignal_proto_rawDescGZIP() []byte {
	file_ZtLiveScActionSignal_proto_rawDescOnce.Do(func() {
		file_ZtLiveScActionSignal_proto_rawDescData = protoimpl.X.CompressGZIP(file_ZtLiveScActionSignal_proto_rawDescData)
	})
	return file_ZtLiveScActionSignal_proto_rawDescData
}

var file_ZtLiveScActionSignal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ZtLiveScActionSignal_proto_goTypes = []interface{}{
	(*ZtLiveScActionSignal)(nil),   // 0: AcFunDanmu.ZtLiveScActionSignal
	(*ZtLiveActionSignalItem)(nil), // 1: AcFunDanmu.ZtLiveActionSignalItem
}
var file_ZtLiveScActionSignal_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.ZtLiveScActionSignal.item:type_name -> AcFunDanmu.ZtLiveActionSignalItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ZtLiveScActionSignal_proto_init() }
func file_ZtLiveScActionSignal_proto_init() {
	if File_ZtLiveScActionSignal_proto != nil {
		return
	}
	file_ZtLiveActionSignalItem_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ZtLiveScActionSignal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZtLiveScActionSignal); i {
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
			RawDescriptor: file_ZtLiveScActionSignal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ZtLiveScActionSignal_proto_goTypes,
		DependencyIndexes: file_ZtLiveScActionSignal_proto_depIdxs,
		MessageInfos:      file_ZtLiveScActionSignal_proto_msgTypes,
	}.Build()
	File_ZtLiveScActionSignal_proto = out.File
	file_ZtLiveScActionSignal_proto_rawDesc = nil
	file_ZtLiveScActionSignal_proto_goTypes = nil
	file_ZtLiveScActionSignal_proto_depIdxs = nil
}
