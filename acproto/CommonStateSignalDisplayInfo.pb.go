// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: CommonStateSignalDisplayInfo.proto

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

type CommonStateSignalDisplayInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WatchingCount string `protobuf:"bytes,1,opt,name=watchingCount,proto3" json:"watchingCount,omitempty"`
	LikeCount     string `protobuf:"bytes,2,opt,name=likeCount,proto3" json:"likeCount,omitempty"`
	LikeDelta     int32  `protobuf:"varint,3,opt,name=likeDelta,proto3" json:"likeDelta,omitempty"`
}

func (x *CommonStateSignalDisplayInfo) Reset() {
	*x = CommonStateSignalDisplayInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonStateSignalDisplayInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonStateSignalDisplayInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonStateSignalDisplayInfo) ProtoMessage() {}

func (x *CommonStateSignalDisplayInfo) ProtoReflect() protoreflect.Message {
	mi := &file_CommonStateSignalDisplayInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonStateSignalDisplayInfo.ProtoReflect.Descriptor instead.
func (*CommonStateSignalDisplayInfo) Descriptor() ([]byte, []int) {
	return file_CommonStateSignalDisplayInfo_proto_rawDescGZIP(), []int{0}
}

func (x *CommonStateSignalDisplayInfo) GetWatchingCount() string {
	if x != nil {
		return x.WatchingCount
	}
	return ""
}

func (x *CommonStateSignalDisplayInfo) GetLikeCount() string {
	if x != nil {
		return x.LikeCount
	}
	return ""
}

func (x *CommonStateSignalDisplayInfo) GetLikeDelta() int32 {
	if x != nil {
		return x.LikeDelta
	}
	return 0
}

var File_CommonStateSignalDisplayInfo_proto protoreflect.FileDescriptor

var file_CommonStateSignalDisplayInfo_proto_rawDesc = []byte{
	0x0a, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x22, 0x80, 0x01, 0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x44, 0x65, 0x6c,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x44, 0x65,
	0x6c, 0x74, 0x61, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61,
	0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_CommonStateSignalDisplayInfo_proto_rawDescOnce sync.Once
	file_CommonStateSignalDisplayInfo_proto_rawDescData = file_CommonStateSignalDisplayInfo_proto_rawDesc
)

func file_CommonStateSignalDisplayInfo_proto_rawDescGZIP() []byte {
	file_CommonStateSignalDisplayInfo_proto_rawDescOnce.Do(func() {
		file_CommonStateSignalDisplayInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonStateSignalDisplayInfo_proto_rawDescData)
	})
	return file_CommonStateSignalDisplayInfo_proto_rawDescData
}

var file_CommonStateSignalDisplayInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonStateSignalDisplayInfo_proto_goTypes = []any{
	(*CommonStateSignalDisplayInfo)(nil), // 0: AcFunDanmu.CommonStateSignalDisplayInfo
}
var file_CommonStateSignalDisplayInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CommonStateSignalDisplayInfo_proto_init() }
func file_CommonStateSignalDisplayInfo_proto_init() {
	if File_CommonStateSignalDisplayInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CommonStateSignalDisplayInfo_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommonStateSignalDisplayInfo); i {
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
			RawDescriptor: file_CommonStateSignalDisplayInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonStateSignalDisplayInfo_proto_goTypes,
		DependencyIndexes: file_CommonStateSignalDisplayInfo_proto_depIdxs,
		MessageInfos:      file_CommonStateSignalDisplayInfo_proto_msgTypes,
	}.Build()
	File_CommonStateSignalDisplayInfo_proto = out.File
	file_CommonStateSignalDisplayInfo_proto_rawDesc = nil
	file_CommonStateSignalDisplayInfo_proto_goTypes = nil
	file_CommonStateSignalDisplayInfo_proto_depIdxs = nil
}
