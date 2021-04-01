// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: SyncCookie.proto

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

type SyncCookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SyncOffset int64 `protobuf:"varint,1,opt,name=syncOffset,proto3" json:"syncOffset,omitempty"`
}

func (x *SyncCookie) Reset() {
	*x = SyncCookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SyncCookie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncCookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncCookie) ProtoMessage() {}

func (x *SyncCookie) ProtoReflect() protoreflect.Message {
	mi := &file_SyncCookie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncCookie.ProtoReflect.Descriptor instead.
func (*SyncCookie) Descriptor() ([]byte, []int) {
	return file_SyncCookie_proto_rawDescGZIP(), []int{0}
}

func (x *SyncCookie) GetSyncOffset() int64 {
	if x != nil {
		return x.SyncOffset
	}
	return 0
}

var File_SyncCookie_proto protoreflect.FileDescriptor

var file_SyncCookie_proto_rawDesc = []byte{
	0x0a, 0x10, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0x2c,
	0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x79, 0x6e, 0x63, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67,
	0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SyncCookie_proto_rawDescOnce sync.Once
	file_SyncCookie_proto_rawDescData = file_SyncCookie_proto_rawDesc
)

func file_SyncCookie_proto_rawDescGZIP() []byte {
	file_SyncCookie_proto_rawDescOnce.Do(func() {
		file_SyncCookie_proto_rawDescData = protoimpl.X.CompressGZIP(file_SyncCookie_proto_rawDescData)
	})
	return file_SyncCookie_proto_rawDescData
}

var file_SyncCookie_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SyncCookie_proto_goTypes = []interface{}{
	(*SyncCookie)(nil), // 0: AcFunDanmu.SyncCookie
}
var file_SyncCookie_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SyncCookie_proto_init() }
func file_SyncCookie_proto_init() {
	if File_SyncCookie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SyncCookie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncCookie); i {
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
			RawDescriptor: file_SyncCookie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SyncCookie_proto_goTypes,
		DependencyIndexes: file_SyncCookie_proto_depIdxs,
		MessageInfos:      file_SyncCookie_proto_msgTypes,
	}.Build()
	File_SyncCookie_proto = out.File
	file_SyncCookie_proto_rawDesc = nil
	file_SyncCookie_proto_goTypes = nil
	file_SyncCookie_proto_depIdxs = nil
}
