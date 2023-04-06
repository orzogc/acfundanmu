// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: DeviceInfo.proto

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

type DeviceInfo_PlatformType int32

const (
	DeviceInfo_kInvalid       DeviceInfo_PlatformType = 0
	DeviceInfo_kAndroid       DeviceInfo_PlatformType = 1
	DeviceInfo_kiOS           DeviceInfo_PlatformType = 2
	DeviceInfo_kWindows       DeviceInfo_PlatformType = 3
	DeviceInfo_WECHAT_ANDROID DeviceInfo_PlatformType = 4
	DeviceInfo_WECHAT_IOS     DeviceInfo_PlatformType = 5
	DeviceInfo_H5             DeviceInfo_PlatformType = 6
	DeviceInfo_H5_ANDROID     DeviceInfo_PlatformType = 7
	DeviceInfo_H5_IOS         DeviceInfo_PlatformType = 8
	DeviceInfo_H5_WINDOWS     DeviceInfo_PlatformType = 9
	DeviceInfo_H5_MAC         DeviceInfo_PlatformType = 10
	DeviceInfo_kPlatformNum   DeviceInfo_PlatformType = 11
)

// Enum value maps for DeviceInfo_PlatformType.
var (
	DeviceInfo_PlatformType_name = map[int32]string{
		0:  "kInvalid",
		1:  "kAndroid",
		2:  "kiOS",
		3:  "kWindows",
		4:  "WECHAT_ANDROID",
		5:  "WECHAT_IOS",
		6:  "H5",
		7:  "H5_ANDROID",
		8:  "H5_IOS",
		9:  "H5_WINDOWS",
		10: "H5_MAC",
		11: "kPlatformNum",
	}
	DeviceInfo_PlatformType_value = map[string]int32{
		"kInvalid":       0,
		"kAndroid":       1,
		"kiOS":           2,
		"kWindows":       3,
		"WECHAT_ANDROID": 4,
		"WECHAT_IOS":     5,
		"H5":             6,
		"H5_ANDROID":     7,
		"H5_IOS":         8,
		"H5_WINDOWS":     9,
		"H5_MAC":         10,
		"kPlatformNum":   11,
	}
)

func (x DeviceInfo_PlatformType) Enum() *DeviceInfo_PlatformType {
	p := new(DeviceInfo_PlatformType)
	*p = x
	return p
}

func (x DeviceInfo_PlatformType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceInfo_PlatformType) Descriptor() protoreflect.EnumDescriptor {
	return file_DeviceInfo_proto_enumTypes[0].Descriptor()
}

func (DeviceInfo_PlatformType) Type() protoreflect.EnumType {
	return &file_DeviceInfo_proto_enumTypes[0]
}

func (x DeviceInfo_PlatformType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceInfo_PlatformType.Descriptor instead.
func (DeviceInfo_PlatformType) EnumDescriptor() ([]byte, []int) {
	return file_DeviceInfo_proto_rawDescGZIP(), []int{0, 0}
}

type DeviceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlatformType DeviceInfo_PlatformType `protobuf:"varint,1,opt,name=platformType,proto3,enum=AcFunDanmu.Im.Basic.DeviceInfo_PlatformType" json:"platformType,omitempty"`
	OsVersion    string                  `protobuf:"bytes,2,opt,name=osVersion,proto3" json:"osVersion,omitempty"`
	DeviceModel  string                  `protobuf:"bytes,3,opt,name=deviceModel,proto3" json:"deviceModel,omitempty"`
	ImeiMd5      []byte                  `protobuf:"bytes,4,opt,name=imeiMd5,proto3" json:"imeiMd5,omitempty"`
	DeviceId     string                  `protobuf:"bytes,5,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	SoftDid      string                  `protobuf:"bytes,6,opt,name=softDid,proto3" json:"softDid,omitempty"`
	KwaiDid      string                  `protobuf:"bytes,7,opt,name=kwaiDid,proto3" json:"kwaiDid,omitempty"`
	Manufacturer string                  `protobuf:"bytes,8,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	DeviceName   string                  `protobuf:"bytes,9,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
}

func (x *DeviceInfo) Reset() {
	*x = DeviceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DeviceInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceInfo) ProtoMessage() {}

func (x *DeviceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_DeviceInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceInfo.ProtoReflect.Descriptor instead.
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return file_DeviceInfo_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceInfo) GetPlatformType() DeviceInfo_PlatformType {
	if x != nil {
		return x.PlatformType
	}
	return DeviceInfo_kInvalid
}

func (x *DeviceInfo) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

func (x *DeviceInfo) GetDeviceModel() string {
	if x != nil {
		return x.DeviceModel
	}
	return ""
}

func (x *DeviceInfo) GetImeiMd5() []byte {
	if x != nil {
		return x.ImeiMd5
	}
	return nil
}

func (x *DeviceInfo) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *DeviceInfo) GetSoftDid() string {
	if x != nil {
		return x.SoftDid
	}
	return ""
}

func (x *DeviceInfo) GetKwaiDid() string {
	if x != nil {
		return x.KwaiDid
	}
	return ""
}

func (x *DeviceInfo) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *DeviceInfo) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

var File_DeviceInfo_proto protoreflect.FileDescriptor

var file_DeviceInfo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x49,
	0x6d, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x22, 0x87, 0x04, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x50, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x41,
	0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x49, 0x6d, 0x2e, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x73, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x73, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d, 0x65, 0x69,
	0x4d, 0x64, 0x35, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x69, 0x6d, 0x65, 0x69, 0x4d,
	0x64, 0x35, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x6f, 0x66, 0x74, 0x44, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x6f, 0x66, 0x74, 0x44, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x77, 0x61, 0x69,
	0x44, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x77, 0x61, 0x69, 0x44,
	0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x6b, 0x49, 0x6e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x6b, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69,
	0x64, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x6b, 0x69, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x6b, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x57,
	0x45, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x04, 0x12,
	0x0e, 0x0a, 0x0a, 0x57, 0x45, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x49, 0x4f, 0x53, 0x10, 0x05, 0x12,
	0x06, 0x0a, 0x02, 0x48, 0x35, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x35, 0x5f, 0x41, 0x4e,
	0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x35, 0x5f, 0x49, 0x4f,
	0x53, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x35, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57,
	0x53, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x35, 0x5f, 0x4d, 0x41, 0x43, 0x10, 0x0a, 0x12,
	0x10, 0x0a, 0x0c, 0x6b, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4e, 0x75, 0x6d, 0x10,
	0x0b, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d,
	0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_DeviceInfo_proto_rawDescOnce sync.Once
	file_DeviceInfo_proto_rawDescData = file_DeviceInfo_proto_rawDesc
)

func file_DeviceInfo_proto_rawDescGZIP() []byte {
	file_DeviceInfo_proto_rawDescOnce.Do(func() {
		file_DeviceInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_DeviceInfo_proto_rawDescData)
	})
	return file_DeviceInfo_proto_rawDescData
}

var file_DeviceInfo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DeviceInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DeviceInfo_proto_goTypes = []interface{}{
	(DeviceInfo_PlatformType)(0), // 0: AcFunDanmu.Im.Basic.DeviceInfo.PlatformType
	(*DeviceInfo)(nil),           // 1: AcFunDanmu.Im.Basic.DeviceInfo
}
var file_DeviceInfo_proto_depIdxs = []int32{
	0, // 0: AcFunDanmu.Im.Basic.DeviceInfo.platformType:type_name -> AcFunDanmu.Im.Basic.DeviceInfo.PlatformType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DeviceInfo_proto_init() }
func file_DeviceInfo_proto_init() {
	if File_DeviceInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DeviceInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceInfo); i {
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
			RawDescriptor: file_DeviceInfo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DeviceInfo_proto_goTypes,
		DependencyIndexes: file_DeviceInfo_proto_depIdxs,
		EnumInfos:         file_DeviceInfo_proto_enumTypes,
		MessageInfos:      file_DeviceInfo_proto_msgTypes,
	}.Build()
	File_DeviceInfo_proto = out.File
	file_DeviceInfo_proto_rawDesc = nil
	file_DeviceInfo_proto_goTypes = nil
	file_DeviceInfo_proto_depIdxs = nil
}
