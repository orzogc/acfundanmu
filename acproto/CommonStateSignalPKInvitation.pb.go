// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: CommonStateSignalPKInvitation.proto

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

type CommonStateSignalPKInvitation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string        `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B *PkPlayerInfo `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C int64         `protobuf:"varint,3,opt,name=c,proto3" json:"c,omitempty"`
}

func (x *CommonStateSignalPKInvitation) Reset() {
	*x = CommonStateSignalPKInvitation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonStateSignalPKInvitation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonStateSignalPKInvitation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonStateSignalPKInvitation) ProtoMessage() {}

func (x *CommonStateSignalPKInvitation) ProtoReflect() protoreflect.Message {
	mi := &file_CommonStateSignalPKInvitation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonStateSignalPKInvitation.ProtoReflect.Descriptor instead.
func (*CommonStateSignalPKInvitation) Descriptor() ([]byte, []int) {
	return file_CommonStateSignalPKInvitation_proto_rawDescGZIP(), []int{0}
}

func (x *CommonStateSignalPKInvitation) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *CommonStateSignalPKInvitation) GetB() *PkPlayerInfo {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *CommonStateSignalPKInvitation) GetC() int64 {
	if x != nil {
		return x.C
	}
	return 0
}

var File_CommonStateSignalPKInvitation_proto protoreflect.FileDescriptor

var file_CommonStateSignalPKInvitation_proto_rawDesc = []byte{
	0x0a, 0x23, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x50, 0x4b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d,
	0x75, 0x1a, 0x12, 0x50, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x1d, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x50, 0x4b, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x01, 0x61, 0x12, 0x26, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x50, 0x6b, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x63, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f,
	0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonStateSignalPKInvitation_proto_rawDescOnce sync.Once
	file_CommonStateSignalPKInvitation_proto_rawDescData = file_CommonStateSignalPKInvitation_proto_rawDesc
)

func file_CommonStateSignalPKInvitation_proto_rawDescGZIP() []byte {
	file_CommonStateSignalPKInvitation_proto_rawDescOnce.Do(func() {
		file_CommonStateSignalPKInvitation_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonStateSignalPKInvitation_proto_rawDescData)
	})
	return file_CommonStateSignalPKInvitation_proto_rawDescData
}

var file_CommonStateSignalPKInvitation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonStateSignalPKInvitation_proto_goTypes = []any{
	(*CommonStateSignalPKInvitation)(nil), // 0: AcFunDanmu.CommonStateSignalPKInvitation
	(*PkPlayerInfo)(nil),                  // 1: AcFunDanmu.PkPlayerInfo
}
var file_CommonStateSignalPKInvitation_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.CommonStateSignalPKInvitation.b:type_name -> AcFunDanmu.PkPlayerInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CommonStateSignalPKInvitation_proto_init() }
func file_CommonStateSignalPKInvitation_proto_init() {
	if File_CommonStateSignalPKInvitation_proto != nil {
		return
	}
	file_PkPlayerInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CommonStateSignalPKInvitation_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommonStateSignalPKInvitation); i {
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
			RawDescriptor: file_CommonStateSignalPKInvitation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonStateSignalPKInvitation_proto_goTypes,
		DependencyIndexes: file_CommonStateSignalPKInvitation_proto_depIdxs,
		MessageInfos:      file_CommonStateSignalPKInvitation_proto_msgTypes,
	}.Build()
	File_CommonStateSignalPKInvitation_proto = out.File
	file_CommonStateSignalPKInvitation_proto_rawDesc = nil
	file_CommonStateSignalPKInvitation_proto_goTypes = nil
	file_CommonStateSignalPKInvitation_proto_depIdxs = nil
}
