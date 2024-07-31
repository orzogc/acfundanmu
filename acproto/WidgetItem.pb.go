// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: WidgetItem.proto

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

type WidgetItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64               `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B string              `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C int32               `protobuf:"varint,3,opt,name=c,proto3" json:"c,omitempty"`
	D int64               `protobuf:"varint,4,opt,name=d,proto3" json:"d,omitempty"`
	E int64               `protobuf:"varint,5,opt,name=e,proto3" json:"e,omitempty"`
	F *WidgetPictureInfo  `protobuf:"bytes,6,opt,name=f,proto3" json:"f,omitempty"`
	G string              `protobuf:"bytes,7,opt,name=g,proto3" json:"g,omitempty"`
	H ZtLiveWidgetProtoA  `protobuf:"varint,8,opt,name=h,proto3,enum=AcFunDanmu.ZtLiveWidgetProtoA" json:"h,omitempty"`
	I ZtLiveWidgetProtoB  `protobuf:"varint,9,opt,name=i,proto3,enum=AcFunDanmu.ZtLiveWidgetProtoB" json:"i,omitempty"`
	J *WidgetDisplayStyle `protobuf:"bytes,10,opt,name=j,proto3" json:"j,omitempty"`
	K ZtLiveWidgetProtoC  `protobuf:"varint,11,opt,name=k,proto3,enum=AcFunDanmu.ZtLiveWidgetProtoC" json:"k,omitempty"`
	L []string            `protobuf:"bytes,12,rep,name=l,proto3" json:"l,omitempty"`
	M []string            `protobuf:"bytes,13,rep,name=m,proto3" json:"m,omitempty"`
	N int64               `protobuf:"varint,14,opt,name=n,proto3" json:"n,omitempty"`
}

func (x *WidgetItem) Reset() {
	*x = WidgetItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WidgetItem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidgetItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidgetItem) ProtoMessage() {}

func (x *WidgetItem) ProtoReflect() protoreflect.Message {
	mi := &file_WidgetItem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidgetItem.ProtoReflect.Descriptor instead.
func (*WidgetItem) Descriptor() ([]byte, []int) {
	return file_WidgetItem_proto_rawDescGZIP(), []int{0}
}

func (x *WidgetItem) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *WidgetItem) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

func (x *WidgetItem) GetC() int32 {
	if x != nil {
		return x.C
	}
	return 0
}

func (x *WidgetItem) GetD() int64 {
	if x != nil {
		return x.D
	}
	return 0
}

func (x *WidgetItem) GetE() int64 {
	if x != nil {
		return x.E
	}
	return 0
}

func (x *WidgetItem) GetF() *WidgetPictureInfo {
	if x != nil {
		return x.F
	}
	return nil
}

func (x *WidgetItem) GetG() string {
	if x != nil {
		return x.G
	}
	return ""
}

func (x *WidgetItem) GetH() ZtLiveWidgetProtoA {
	if x != nil {
		return x.H
	}
	return ZtLiveWidgetProtoA_ZtLiveWidgetProtoAa
}

func (x *WidgetItem) GetI() ZtLiveWidgetProtoB {
	if x != nil {
		return x.I
	}
	return ZtLiveWidgetProtoB_ZtLiveWidgetProtoBa
}

func (x *WidgetItem) GetJ() *WidgetDisplayStyle {
	if x != nil {
		return x.J
	}
	return nil
}

func (x *WidgetItem) GetK() ZtLiveWidgetProtoC {
	if x != nil {
		return x.K
	}
	return ZtLiveWidgetProtoC_ZtLiveWidgetProtoCa
}

func (x *WidgetItem) GetL() []string {
	if x != nil {
		return x.L
	}
	return nil
}

func (x *WidgetItem) GetM() []string {
	if x != nil {
		return x.M
	}
	return nil
}

func (x *WidgetItem) GetN() int64 {
	if x != nil {
		return x.N
	}
	return 0
}

var File_WidgetItem_proto protoreflect.FileDescriptor

var file_WidgetItem_proto_rawDesc = []byte{
	0x0a, 0x10, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x1a, 0x18,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x79,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x02, 0x0a, 0x0a, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x01, 0x63, 0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x64, 0x12, 0x0c, 0x0a, 0x01, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x65, 0x12,
	0x2b, 0x0a, 0x01, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x41, 0x63, 0x46,
	0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x01, 0x66, 0x12, 0x0c, 0x0a, 0x01,
	0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x67, 0x12, 0x2c, 0x0a, 0x01, 0x68, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e,
	0x6d, 0x75, 0x2e, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x41, 0x52, 0x01, 0x68, 0x12, 0x2c, 0x0a, 0x01, 0x69, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x2e, 0x5a, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x52, 0x01, 0x69, 0x12, 0x2c, 0x0a, 0x01, 0x6a, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x79, 0x6c,
	0x65, 0x52, 0x01, 0x6a, 0x12, 0x2c, 0x0a, 0x01, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x5a, 0x74, 0x4c,
	0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x43, 0x52,
	0x01, 0x6b, 0x12, 0x0c, 0x0a, 0x01, 0x6c, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x01, 0x6c,
	0x12, 0x0c, 0x0a, 0x01, 0x6d, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x01, 0x6d, 0x12, 0x0c,
	0x0a, 0x01, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x6e, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67,
	0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WidgetItem_proto_rawDescOnce sync.Once
	file_WidgetItem_proto_rawDescData = file_WidgetItem_proto_rawDesc
)

func file_WidgetItem_proto_rawDescGZIP() []byte {
	file_WidgetItem_proto_rawDescOnce.Do(func() {
		file_WidgetItem_proto_rawDescData = protoimpl.X.CompressGZIP(file_WidgetItem_proto_rawDescData)
	})
	return file_WidgetItem_proto_rawDescData
}

var file_WidgetItem_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WidgetItem_proto_goTypes = []any{
	(*WidgetItem)(nil),         // 0: AcFunDanmu.WidgetItem
	(*WidgetPictureInfo)(nil),  // 1: AcFunDanmu.WidgetPictureInfo
	(ZtLiveWidgetProtoA)(0),    // 2: AcFunDanmu.ZtLiveWidgetProtoA
	(ZtLiveWidgetProtoB)(0),    // 3: AcFunDanmu.ZtLiveWidgetProtoB
	(*WidgetDisplayStyle)(nil), // 4: AcFunDanmu.WidgetDisplayStyle
	(ZtLiveWidgetProtoC)(0),    // 5: AcFunDanmu.ZtLiveWidgetProtoC
}
var file_WidgetItem_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.WidgetItem.f:type_name -> AcFunDanmu.WidgetPictureInfo
	2, // 1: AcFunDanmu.WidgetItem.h:type_name -> AcFunDanmu.ZtLiveWidgetProtoA
	3, // 2: AcFunDanmu.WidgetItem.i:type_name -> AcFunDanmu.ZtLiveWidgetProtoB
	4, // 3: AcFunDanmu.WidgetItem.j:type_name -> AcFunDanmu.WidgetDisplayStyle
	5, // 4: AcFunDanmu.WidgetItem.k:type_name -> AcFunDanmu.ZtLiveWidgetProtoC
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_WidgetItem_proto_init() }
func file_WidgetItem_proto_init() {
	if File_WidgetItem_proto != nil {
		return
	}
	file_WidgetDisplayStyle_proto_init()
	file_WidgetPictureInfo_proto_init()
	file_ZtLiveWidgetProto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WidgetItem_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*WidgetItem); i {
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
			RawDescriptor: file_WidgetItem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WidgetItem_proto_goTypes,
		DependencyIndexes: file_WidgetItem_proto_depIdxs,
		MessageInfos:      file_WidgetItem_proto_msgTypes,
	}.Build()
	File_WidgetItem_proto = out.File
	file_WidgetItem_proto_rawDesc = nil
	file_WidgetItem_proto_goTypes = nil
	file_WidgetItem_proto_depIdxs = nil
}
