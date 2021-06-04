// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: TokenInfo.proto

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

type TokenInfo_TokenType int32

const (
	TokenInfo_kInvalid      TokenInfo_TokenType = 0
	TokenInfo_kServiceToken TokenInfo_TokenType = 1
)

// Enum value maps for TokenInfo_TokenType.
var (
	TokenInfo_TokenType_name = map[int32]string{
		0: "kInvalid",
		1: "kServiceToken",
	}
	TokenInfo_TokenType_value = map[string]int32{
		"kInvalid":      0,
		"kServiceToken": 1,
	}
)

func (x TokenInfo_TokenType) Enum() *TokenInfo_TokenType {
	p := new(TokenInfo_TokenType)
	*p = x
	return p
}

func (x TokenInfo_TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenInfo_TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_TokenInfo_proto_enumTypes[0].Descriptor()
}

func (TokenInfo_TokenType) Type() protoreflect.EnumType {
	return &file_TokenInfo_proto_enumTypes[0]
}

func (x TokenInfo_TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenInfo_TokenType.Descriptor instead.
func (TokenInfo_TokenType) EnumDescriptor() ([]byte, []int) {
	return file_TokenInfo_proto_rawDescGZIP(), []int{0, 0}
}

type TokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenType TokenInfo_TokenType `protobuf:"varint,1,opt,name=tokenType,proto3,enum=AcFunDanmu.TokenInfo_TokenType" json:"tokenType,omitempty"`
	Token     []byte              `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TokenInfo) Reset() {
	*x = TokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInfo) ProtoMessage() {}

func (x *TokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_TokenInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenInfo.ProtoReflect.Descriptor instead.
func (*TokenInfo) Descriptor() ([]byte, []int) {
	return file_TokenInfo_proto_rawDescGZIP(), []int{0}
}

func (x *TokenInfo) GetTokenType() TokenInfo_TokenType {
	if x != nil {
		return x.TokenType
	}
	return TokenInfo_kInvalid
}

func (x *TokenInfo) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_TokenInfo_proto protoreflect.FileDescriptor

var file_TokenInfo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x22, 0x8e, 0x01,
	0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x09, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x2c, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a,
	0x08, 0x6b, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0x01, 0x42, 0x26,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a,
	0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61,
	0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TokenInfo_proto_rawDescOnce sync.Once
	file_TokenInfo_proto_rawDescData = file_TokenInfo_proto_rawDesc
)

func file_TokenInfo_proto_rawDescGZIP() []byte {
	file_TokenInfo_proto_rawDescOnce.Do(func() {
		file_TokenInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_TokenInfo_proto_rawDescData)
	})
	return file_TokenInfo_proto_rawDescData
}

var file_TokenInfo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TokenInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TokenInfo_proto_goTypes = []interface{}{
	(TokenInfo_TokenType)(0), // 0: AcFunDanmu.TokenInfo.TokenType
	(*TokenInfo)(nil),        // 1: AcFunDanmu.TokenInfo
}
var file_TokenInfo_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.TokenInfo.tokenType:type_name -> AcFunDanmu.TokenInfo.TokenType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TokenInfo_proto_init() }
func file_TokenInfo_proto_init() {
	if File_TokenInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TokenInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenInfo); i {
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
			RawDescriptor: file_TokenInfo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TokenInfo_proto_goTypes,
		DependencyIndexes: file_TokenInfo_proto_depIdxs,
		EnumInfos:         file_TokenInfo_proto_enumTypes,
		MessageInfos:      file_TokenInfo_proto_msgTypes,
	}.Build()
	File_TokenInfo_proto = out.File
	file_TokenInfo_proto_rawDesc = nil
	file_TokenInfo_proto_goTypes = nil
	file_TokenInfo_proto_depIdxs = nil
}
