// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: ZtLiveStateSignalItem.proto

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

type ZtLiveStateSignalItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignalType string `protobuf:"bytes,1,opt,name=signalType,proto3" json:"signalType,omitempty"`
	Payload    []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ZtLiveStateSignalItem) Reset() {
	*x = ZtLiveStateSignalItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ZtLiveStateSignalItem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZtLiveStateSignalItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZtLiveStateSignalItem) ProtoMessage() {}

func (x *ZtLiveStateSignalItem) ProtoReflect() protoreflect.Message {
	mi := &file_ZtLiveStateSignalItem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZtLiveStateSignalItem.ProtoReflect.Descriptor instead.
func (*ZtLiveStateSignalItem) Descriptor() ([]byte, []int) {
	return file_ZtLiveStateSignalItem_proto_rawDescGZIP(), []int{0}
}

func (x *ZtLiveStateSignalItem) GetSignalType() string {
	if x != nil {
		return x.SignalType
	}
	return ""
}

func (x *ZtLiveStateSignalItem) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_ZtLiveStateSignalItem_proto protoreflect.FileDescriptor

var file_ZtLiveStateSignalItem_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41,
	0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0x51, 0x0a, 0x15, 0x5a, 0x74, 0x4c,
	0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67,
	0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ZtLiveStateSignalItem_proto_rawDescOnce sync.Once
	file_ZtLiveStateSignalItem_proto_rawDescData = file_ZtLiveStateSignalItem_proto_rawDesc
)

func file_ZtLiveStateSignalItem_proto_rawDescGZIP() []byte {
	file_ZtLiveStateSignalItem_proto_rawDescOnce.Do(func() {
		file_ZtLiveStateSignalItem_proto_rawDescData = protoimpl.X.CompressGZIP(file_ZtLiveStateSignalItem_proto_rawDescData)
	})
	return file_ZtLiveStateSignalItem_proto_rawDescData
}

var file_ZtLiveStateSignalItem_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ZtLiveStateSignalItem_proto_goTypes = []interface{}{
	(*ZtLiveStateSignalItem)(nil), // 0: AcFunDanmu.ZtLiveStateSignalItem
}
var file_ZtLiveStateSignalItem_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ZtLiveStateSignalItem_proto_init() }
func file_ZtLiveStateSignalItem_proto_init() {
	if File_ZtLiveStateSignalItem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ZtLiveStateSignalItem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZtLiveStateSignalItem); i {
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
			RawDescriptor: file_ZtLiveStateSignalItem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ZtLiveStateSignalItem_proto_goTypes,
		DependencyIndexes: file_ZtLiveStateSignalItem_proto_depIdxs,
		MessageInfos:      file_ZtLiveStateSignalItem_proto_msgTypes,
	}.Build()
	File_ZtLiveStateSignalItem_proto = out.File
	file_ZtLiveStateSignalItem_proto_rawDesc = nil
	file_ZtLiveStateSignalItem_proto_goTypes = nil
	file_ZtLiveStateSignalItem_proto_depIdxs = nil
}
