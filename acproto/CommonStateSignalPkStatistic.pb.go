// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: CommonStateSignalPkStatistic.proto

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

type CommonStateSignalPkStatistic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A string                        `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64                         `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	C int64                         `protobuf:"varint,3,opt,name=c,proto3" json:"c,omitempty"`
	D int64                         `protobuf:"varint,4,opt,name=d,proto3" json:"d,omitempty"`
	E bool                          `protobuf:"varint,5,opt,name=e,proto3" json:"e,omitempty"`
	F int64                         `protobuf:"varint,6,opt,name=f,proto3" json:"f,omitempty"`
	G int64                         `protobuf:"varint,7,opt,name=g,proto3" json:"g,omitempty"`
	H int64                         `protobuf:"varint,8,opt,name=h,proto3" json:"h,omitempty"`
	I []*PkAudienceContributionInfo `protobuf:"bytes,9,rep,name=i,proto3" json:"i,omitempty"`
	J []*PkPlayerStatistic          `protobuf:"bytes,10,rep,name=j,proto3" json:"j,omitempty"`
	K *PkRoundInfo                  `protobuf:"bytes,11,opt,name=k,proto3" json:"k,omitempty"`
	L int64                         `protobuf:"varint,12,opt,name=l,proto3" json:"l,omitempty"`
}

func (x *CommonStateSignalPkStatistic) Reset() {
	*x = CommonStateSignalPkStatistic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonStateSignalPkStatistic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonStateSignalPkStatistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonStateSignalPkStatistic) ProtoMessage() {}

func (x *CommonStateSignalPkStatistic) ProtoReflect() protoreflect.Message {
	mi := &file_CommonStateSignalPkStatistic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonStateSignalPkStatistic.ProtoReflect.Descriptor instead.
func (*CommonStateSignalPkStatistic) Descriptor() ([]byte, []int) {
	return file_CommonStateSignalPkStatistic_proto_rawDescGZIP(), []int{0}
}

func (x *CommonStateSignalPkStatistic) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *CommonStateSignalPkStatistic) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *CommonStateSignalPkStatistic) GetC() int64 {
	if x != nil {
		return x.C
	}
	return 0
}

func (x *CommonStateSignalPkStatistic) GetD() int64 {
	if x != nil {
		return x.D
	}
	return 0
}

func (x *CommonStateSignalPkStatistic) GetE() bool {
	if x != nil {
		return x.E
	}
	return false
}

func (x *CommonStateSignalPkStatistic) GetF() int64 {
	if x != nil {
		return x.F
	}
	return 0
}

func (x *CommonStateSignalPkStatistic) GetG() int64 {
	if x != nil {
		return x.G
	}
	return 0
}

func (x *CommonStateSignalPkStatistic) GetH() int64 {
	if x != nil {
		return x.H
	}
	return 0
}

func (x *CommonStateSignalPkStatistic) GetI() []*PkAudienceContributionInfo {
	if x != nil {
		return x.I
	}
	return nil
}

func (x *CommonStateSignalPkStatistic) GetJ() []*PkPlayerStatistic {
	if x != nil {
		return x.J
	}
	return nil
}

func (x *CommonStateSignalPkStatistic) GetK() *PkRoundInfo {
	if x != nil {
		return x.K
	}
	return nil
}

func (x *CommonStateSignalPkStatistic) GetL() int64 {
	if x != nil {
		return x.L
	}
	return 0
}

var File_CommonStateSignalPkStatistic_proto protoreflect.FileDescriptor

var file_CommonStateSignalPkStatistic_proto_rawDesc = []byte{
	0x0a, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x50, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x1a, 0x20, 0x50, 0x6b, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x50, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x6b, 0x52,
	0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6,
	0x02, 0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x50, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12,
	0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a,
	0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x62, 0x12, 0x0c, 0x0a, 0x01, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x63, 0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x01, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x66, 0x12, 0x0c, 0x0a, 0x01, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x67, 0x12, 0x0c, 0x0a, 0x01, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x68, 0x12,
	0x34, 0x0a, 0x01, 0x69, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x41, 0x63, 0x46,
	0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x50, 0x6b, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x01, 0x69, 0x12, 0x2b, 0x0a, 0x01, 0x6a, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x50, 0x6b,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52,
	0x01, 0x6a, 0x12, 0x25, 0x0a, 0x01, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x50, 0x6b, 0x52, 0x6f, 0x75,
	0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x01, 0x6b, 0x12, 0x0c, 0x0a, 0x01, 0x6c, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x6c, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66,
	0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonStateSignalPkStatistic_proto_rawDescOnce sync.Once
	file_CommonStateSignalPkStatistic_proto_rawDescData = file_CommonStateSignalPkStatistic_proto_rawDesc
)

func file_CommonStateSignalPkStatistic_proto_rawDescGZIP() []byte {
	file_CommonStateSignalPkStatistic_proto_rawDescOnce.Do(func() {
		file_CommonStateSignalPkStatistic_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonStateSignalPkStatistic_proto_rawDescData)
	})
	return file_CommonStateSignalPkStatistic_proto_rawDescData
}

var file_CommonStateSignalPkStatistic_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonStateSignalPkStatistic_proto_goTypes = []interface{}{
	(*CommonStateSignalPkStatistic)(nil), // 0: AcFunDanmu.CommonStateSignalPkStatistic
	(*PkAudienceContributionInfo)(nil),   // 1: AcFunDanmu.PkAudienceContributionInfo
	(*PkPlayerStatistic)(nil),            // 2: AcFunDanmu.PkPlayerStatistic
	(*PkRoundInfo)(nil),                  // 3: AcFunDanmu.PkRoundInfo
}
var file_CommonStateSignalPkStatistic_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.CommonStateSignalPkStatistic.i:type_name -> AcFunDanmu.PkAudienceContributionInfo
	2, // 1: AcFunDanmu.CommonStateSignalPkStatistic.j:type_name -> AcFunDanmu.PkPlayerStatistic
	3, // 2: AcFunDanmu.CommonStateSignalPkStatistic.k:type_name -> AcFunDanmu.PkRoundInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_CommonStateSignalPkStatistic_proto_init() }
func file_CommonStateSignalPkStatistic_proto_init() {
	if File_CommonStateSignalPkStatistic_proto != nil {
		return
	}
	file_PkAudienceContributionInfo_proto_init()
	file_PkPlayerStatistic_proto_init()
	file_PkRoundInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CommonStateSignalPkStatistic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonStateSignalPkStatistic); i {
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
			RawDescriptor: file_CommonStateSignalPkStatistic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonStateSignalPkStatistic_proto_goTypes,
		DependencyIndexes: file_CommonStateSignalPkStatistic_proto_depIdxs,
		MessageInfos:      file_CommonStateSignalPkStatistic_proto_msgTypes,
	}.Build()
	File_CommonStateSignalPkStatistic_proto = out.File
	file_CommonStateSignalPkStatistic_proto_rawDesc = nil
	file_CommonStateSignalPkStatistic_proto_goTypes = nil
	file_CommonStateSignalPkStatistic_proto_depIdxs = nil
}
