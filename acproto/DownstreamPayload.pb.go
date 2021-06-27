// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.17.3
// source: DownstreamPayload.proto

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

type DownstreamPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command     string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	SeqId       int64  `protobuf:"varint,2,opt,name=seqId,proto3" json:"seqId,omitempty"`
	ErrorCode   int32  `protobuf:"varint,3,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	PayloadData []byte `protobuf:"bytes,4,opt,name=payloadData,proto3" json:"payloadData,omitempty"`
	ErrorMsg    string `protobuf:"bytes,5,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
	ErrorData   []byte `protobuf:"bytes,6,opt,name=errorData,proto3" json:"errorData,omitempty"`
	SubBiz      string `protobuf:"bytes,7,opt,name=subBiz,proto3" json:"subBiz,omitempty"`
	KlinkPushId int64  `protobuf:"varint,8,opt,name=klinkPushId,proto3" json:"klinkPushId,omitempty"`
}

func (x *DownstreamPayload) Reset() {
	*x = DownstreamPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DownstreamPayload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownstreamPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownstreamPayload) ProtoMessage() {}

func (x *DownstreamPayload) ProtoReflect() protoreflect.Message {
	mi := &file_DownstreamPayload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownstreamPayload.ProtoReflect.Descriptor instead.
func (*DownstreamPayload) Descriptor() ([]byte, []int) {
	return file_DownstreamPayload_proto_rawDescGZIP(), []int{0}
}

func (x *DownstreamPayload) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *DownstreamPayload) GetSeqId() int64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *DownstreamPayload) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *DownstreamPayload) GetPayloadData() []byte {
	if x != nil {
		return x.PayloadData
	}
	return nil
}

func (x *DownstreamPayload) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *DownstreamPayload) GetErrorData() []byte {
	if x != nil {
		return x.ErrorData
	}
	return nil
}

func (x *DownstreamPayload) GetSubBiz() string {
	if x != nil {
		return x.SubBiz
	}
	return ""
}

func (x *DownstreamPayload) GetKlinkPushId() int64 {
	if x != nil {
		return x.KlinkPushId
	}
	return 0
}

var File_DownstreamPayload_proto protoreflect.FileDescriptor

var file_DownstreamPayload_proto_rawDesc = []byte{
	0x0a, 0x17, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e,
	0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0xf7, 0x01, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x42, 0x69, 0x7a, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x42, 0x69, 0x7a, 0x12, 0x20, 0x0a,
	0x0b, 0x6b, 0x6c, 0x69, 0x6e, 0x6b, 0x50, 0x75, 0x73, 0x68, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x6b, 0x6c, 0x69, 0x6e, 0x6b, 0x50, 0x75, 0x73, 0x68, 0x49, 0x64, 0x42,
	0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72,
	0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f,
	0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DownstreamPayload_proto_rawDescOnce sync.Once
	file_DownstreamPayload_proto_rawDescData = file_DownstreamPayload_proto_rawDesc
)

func file_DownstreamPayload_proto_rawDescGZIP() []byte {
	file_DownstreamPayload_proto_rawDescOnce.Do(func() {
		file_DownstreamPayload_proto_rawDescData = protoimpl.X.CompressGZIP(file_DownstreamPayload_proto_rawDescData)
	})
	return file_DownstreamPayload_proto_rawDescData
}

var file_DownstreamPayload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DownstreamPayload_proto_goTypes = []interface{}{
	(*DownstreamPayload)(nil), // 0: AcFunDanmu.DownstreamPayload
}
var file_DownstreamPayload_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DownstreamPayload_proto_init() }
func file_DownstreamPayload_proto_init() {
	if File_DownstreamPayload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DownstreamPayload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownstreamPayload); i {
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
			RawDescriptor: file_DownstreamPayload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DownstreamPayload_proto_goTypes,
		DependencyIndexes: file_DownstreamPayload_proto_depIdxs,
		MessageInfos:      file_DownstreamPayload_proto_msgTypes,
	}.Build()
	File_DownstreamPayload_proto = out.File
	file_DownstreamPayload_proto_rawDesc = nil
	file_DownstreamPayload_proto_goTypes = nil
	file_DownstreamPayload_proto_depIdxs = nil
}
