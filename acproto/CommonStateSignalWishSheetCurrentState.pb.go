// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.17.3
// source: CommonStateSignalWishSheetCurrentState.proto

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

type CommonStateSignalWishSheetCurrentState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string                                                     `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B []*CommonStateSignalWishSheetCurrentState_WishCurrentState `protobuf:"bytes,2,rep,name=b,proto3" json:"b,omitempty"`
}

func (x *CommonStateSignalWishSheetCurrentState) Reset() {
	*x = CommonStateSignalWishSheetCurrentState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonStateSignalWishSheetCurrentState_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonStateSignalWishSheetCurrentState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonStateSignalWishSheetCurrentState) ProtoMessage() {}

func (x *CommonStateSignalWishSheetCurrentState) ProtoReflect() protoreflect.Message {
	mi := &file_CommonStateSignalWishSheetCurrentState_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonStateSignalWishSheetCurrentState.ProtoReflect.Descriptor instead.
func (*CommonStateSignalWishSheetCurrentState) Descriptor() ([]byte, []int) {
	return file_CommonStateSignalWishSheetCurrentState_proto_rawDescGZIP(), []int{0}
}

func (x *CommonStateSignalWishSheetCurrentState) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *CommonStateSignalWishSheetCurrentState) GetB() []*CommonStateSignalWishSheetCurrentState_WishCurrentState {
	if x != nil {
		return x.B
	}
	return nil
}

type CommonStateSignalWishSheetCurrentState_WishCurrentState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C string `protobuf:"bytes,1,opt,name=c,proto3" json:"c,omitempty"`
	D int64  `protobuf:"varint,2,opt,name=d,proto3" json:"d,omitempty"`
	E int64  `protobuf:"varint,3,opt,name=e,proto3" json:"e,omitempty"`
	F int64  `protobuf:"varint,4,opt,name=f,proto3" json:"f,omitempty"`
	G string `protobuf:"bytes,5,opt,name=g,proto3" json:"g,omitempty"`
	H string `protobuf:"bytes,6,opt,name=h,proto3" json:"h,omitempty"`
}

func (x *CommonStateSignalWishSheetCurrentState_WishCurrentState) Reset() {
	*x = CommonStateSignalWishSheetCurrentState_WishCurrentState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonStateSignalWishSheetCurrentState_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonStateSignalWishSheetCurrentState_WishCurrentState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonStateSignalWishSheetCurrentState_WishCurrentState) ProtoMessage() {}

func (x *CommonStateSignalWishSheetCurrentState_WishCurrentState) ProtoReflect() protoreflect.Message {
	mi := &file_CommonStateSignalWishSheetCurrentState_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonStateSignalWishSheetCurrentState_WishCurrentState.ProtoReflect.Descriptor instead.
func (*CommonStateSignalWishSheetCurrentState_WishCurrentState) Descriptor() ([]byte, []int) {
	return file_CommonStateSignalWishSheetCurrentState_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CommonStateSignalWishSheetCurrentState_WishCurrentState) GetC() string {
	if x != nil {
		return x.C
	}
	return ""
}

func (x *CommonStateSignalWishSheetCurrentState_WishCurrentState) GetD() int64 {
	if x != nil {
		return x.D
	}
	return 0
}

func (x *CommonStateSignalWishSheetCurrentState_WishCurrentState) GetE() int64 {
	if x != nil {
		return x.E
	}
	return 0
}

func (x *CommonStateSignalWishSheetCurrentState_WishCurrentState) GetF() int64 {
	if x != nil {
		return x.F
	}
	return 0
}

func (x *CommonStateSignalWishSheetCurrentState_WishCurrentState) GetG() string {
	if x != nil {
		return x.G
	}
	return ""
}

func (x *CommonStateSignalWishSheetCurrentState_WishCurrentState) GetH() string {
	if x != nil {
		return x.H
	}
	return ""
}

var File_CommonStateSignalWishSheetCurrentState_proto protoreflect.FileDescriptor

var file_CommonStateSignalWishSheetCurrentState_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x57, 0x69, 0x73, 0x68, 0x53, 0x68, 0x65, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0xf1, 0x01, 0x0a, 0x26, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x57, 0x69, 0x73, 0x68, 0x53, 0x68, 0x65, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x61, 0x12, 0x51, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43,
	0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x57, 0x69, 0x73,
	0x68, 0x53, 0x68, 0x65, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x57, 0x69, 0x73, 0x68, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x01, 0x62, 0x1a, 0x66, 0x0a, 0x10, 0x57, 0x69, 0x73, 0x68, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x63, 0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x01, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x01, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x01, 0x66, 0x12, 0x0c, 0x0a, 0x01, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x67,
	0x12, 0x0c, 0x0a, 0x01, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x68, 0x42, 0x26,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a,
	0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61,
	0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonStateSignalWishSheetCurrentState_proto_rawDescOnce sync.Once
	file_CommonStateSignalWishSheetCurrentState_proto_rawDescData = file_CommonStateSignalWishSheetCurrentState_proto_rawDesc
)

func file_CommonStateSignalWishSheetCurrentState_proto_rawDescGZIP() []byte {
	file_CommonStateSignalWishSheetCurrentState_proto_rawDescOnce.Do(func() {
		file_CommonStateSignalWishSheetCurrentState_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonStateSignalWishSheetCurrentState_proto_rawDescData)
	})
	return file_CommonStateSignalWishSheetCurrentState_proto_rawDescData
}

var file_CommonStateSignalWishSheetCurrentState_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_CommonStateSignalWishSheetCurrentState_proto_goTypes = []interface{}{
	(*CommonStateSignalWishSheetCurrentState)(nil),                  // 0: AcFunDanmu.CommonStateSignalWishSheetCurrentState
	(*CommonStateSignalWishSheetCurrentState_WishCurrentState)(nil), // 1: AcFunDanmu.CommonStateSignalWishSheetCurrentState.WishCurrentState
}
var file_CommonStateSignalWishSheetCurrentState_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.CommonStateSignalWishSheetCurrentState.b:type_name -> AcFunDanmu.CommonStateSignalWishSheetCurrentState.WishCurrentState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CommonStateSignalWishSheetCurrentState_proto_init() }
func file_CommonStateSignalWishSheetCurrentState_proto_init() {
	if File_CommonStateSignalWishSheetCurrentState_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CommonStateSignalWishSheetCurrentState_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonStateSignalWishSheetCurrentState); i {
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
		file_CommonStateSignalWishSheetCurrentState_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonStateSignalWishSheetCurrentState_WishCurrentState); i {
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
			RawDescriptor: file_CommonStateSignalWishSheetCurrentState_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonStateSignalWishSheetCurrentState_proto_goTypes,
		DependencyIndexes: file_CommonStateSignalWishSheetCurrentState_proto_depIdxs,
		MessageInfos:      file_CommonStateSignalWishSheetCurrentState_proto_msgTypes,
	}.Build()
	File_CommonStateSignalWishSheetCurrentState_proto = out.File
	file_CommonStateSignalWishSheetCurrentState_proto_rawDesc = nil
	file_CommonStateSignalWishSheetCurrentState_proto_goTypes = nil
	file_CommonStateSignalWishSheetCurrentState_proto_depIdxs = nil
}
