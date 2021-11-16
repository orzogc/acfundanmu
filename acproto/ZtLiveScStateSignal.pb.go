// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ZtLiveScStateSignal.proto

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

type ZtLiveScStateSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item []*ZtLiveStateSignalItem `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
}

func (x *ZtLiveScStateSignal) Reset() {
	*x = ZtLiveScStateSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ZtLiveScStateSignal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZtLiveScStateSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZtLiveScStateSignal) ProtoMessage() {}

func (x *ZtLiveScStateSignal) ProtoReflect() protoreflect.Message {
	mi := &file_ZtLiveScStateSignal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZtLiveScStateSignal.ProtoReflect.Descriptor instead.
func (*ZtLiveScStateSignal) Descriptor() ([]byte, []int) {
	return file_ZtLiveScStateSignal_proto_rawDescGZIP(), []int{0}
}

func (x *ZtLiveScStateSignal) GetItem() []*ZtLiveStateSignalItem {
	if x != nil {
		return x.Item
	}
	return nil
}

var File_ZtLiveScStateSignal_proto protoreflect.FileDescriptor

var file_ZtLiveScStateSignal_proto_rawDesc = []byte{
	0x0a, 0x19, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46,
	0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x1a, 0x1b, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x13, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x35, 0x0a, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x41, 0x63, 0x46, 0x75,
	0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e,
	0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ZtLiveScStateSignal_proto_rawDescOnce sync.Once
	file_ZtLiveScStateSignal_proto_rawDescData = file_ZtLiveScStateSignal_proto_rawDesc
)

func file_ZtLiveScStateSignal_proto_rawDescGZIP() []byte {
	file_ZtLiveScStateSignal_proto_rawDescOnce.Do(func() {
		file_ZtLiveScStateSignal_proto_rawDescData = protoimpl.X.CompressGZIP(file_ZtLiveScStateSignal_proto_rawDescData)
	})
	return file_ZtLiveScStateSignal_proto_rawDescData
}

var file_ZtLiveScStateSignal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ZtLiveScStateSignal_proto_goTypes = []interface{}{
	(*ZtLiveScStateSignal)(nil),   // 0: AcFunDanmu.ZtLiveScStateSignal
	(*ZtLiveStateSignalItem)(nil), // 1: AcFunDanmu.ZtLiveStateSignalItem
}
var file_ZtLiveScStateSignal_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.ZtLiveScStateSignal.item:type_name -> AcFunDanmu.ZtLiveStateSignalItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ZtLiveScStateSignal_proto_init() }
func file_ZtLiveScStateSignal_proto_init() {
	if File_ZtLiveScStateSignal_proto != nil {
		return
	}
	file_ZtLiveStateSignalItem_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ZtLiveScStateSignal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZtLiveScStateSignal); i {
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
			RawDescriptor: file_ZtLiveScStateSignal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ZtLiveScStateSignal_proto_goTypes,
		DependencyIndexes: file_ZtLiveScStateSignal_proto_depIdxs,
		MessageInfos:      file_ZtLiveScStateSignal_proto_msgTypes,
	}.Build()
	File_ZtLiveScStateSignal_proto = out.File
	file_ZtLiveScStateSignal_proto_rawDesc = nil
	file_ZtLiveScStateSignal_proto_goTypes = nil
	file_ZtLiveScStateSignal_proto_depIdxs = nil
}
