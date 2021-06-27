// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.17.3
// source: CommonActionSignalGift.proto

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

type CommonActionSignalGift struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User                  *ZtLiveUserInfo `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"` //userInfo
	SendTimeMs            int64           `protobuf:"varint,2,opt,name=sendTimeMs,proto3" json:"sendTimeMs,omitempty"`
	GiftId                int64           `protobuf:"varint,3,opt,name=giftId,proto3" json:"giftId,omitempty"`
	Count                 int32           `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`    //batchSize
	Combo                 int32           `protobuf:"varint,5,opt,name=combo,proto3" json:"combo,omitempty"`    //comboCount
	Value                 int64           `protobuf:"varint,6,opt,name=value,proto3" json:"value,omitempty"`    //rank
	ComboId               string          `protobuf:"bytes,7,opt,name=comboId,proto3" json:"comboId,omitempty"` //comboKey
	SlotDisplayDurationMs int64           `protobuf:"varint,8,opt,name=slotDisplayDurationMs,proto3" json:"slotDisplayDurationMs,omitempty"`
	ExpireDurationMs      int64           `protobuf:"varint,9,opt,name=expireDurationMs,proto3" json:"expireDurationMs,omitempty"`
	DrawGiftInfo          *ZtDrawGiftInfo `protobuf:"bytes,10,opt,name=drawGiftInfo,proto3" json:"drawGiftInfo,omitempty"`
}

func (x *CommonActionSignalGift) Reset() {
	*x = CommonActionSignalGift{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CommonActionSignalGift_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonActionSignalGift) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonActionSignalGift) ProtoMessage() {}

func (x *CommonActionSignalGift) ProtoReflect() protoreflect.Message {
	mi := &file_CommonActionSignalGift_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonActionSignalGift.ProtoReflect.Descriptor instead.
func (*CommonActionSignalGift) Descriptor() ([]byte, []int) {
	return file_CommonActionSignalGift_proto_rawDescGZIP(), []int{0}
}

func (x *CommonActionSignalGift) GetUser() *ZtLiveUserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CommonActionSignalGift) GetSendTimeMs() int64 {
	if x != nil {
		return x.SendTimeMs
	}
	return 0
}

func (x *CommonActionSignalGift) GetGiftId() int64 {
	if x != nil {
		return x.GiftId
	}
	return 0
}

func (x *CommonActionSignalGift) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CommonActionSignalGift) GetCombo() int32 {
	if x != nil {
		return x.Combo
	}
	return 0
}

func (x *CommonActionSignalGift) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *CommonActionSignalGift) GetComboId() string {
	if x != nil {
		return x.ComboId
	}
	return ""
}

func (x *CommonActionSignalGift) GetSlotDisplayDurationMs() int64 {
	if x != nil {
		return x.SlotDisplayDurationMs
	}
	return 0
}

func (x *CommonActionSignalGift) GetExpireDurationMs() int64 {
	if x != nil {
		return x.ExpireDurationMs
	}
	return 0
}

func (x *CommonActionSignalGift) GetDrawGiftInfo() *ZtDrawGiftInfo {
	if x != nil {
		return x.DrawGiftInfo
	}
	return nil
}

var File_CommonActionSignalGift_proto protoreflect.FileDescriptor

var file_CommonActionSignalGift_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x47, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x1a, 0x14, 0x5a, 0x74, 0x4c, 0x69,
	0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x5a, 0x74, 0x44, 0x72, 0x61, 0x77, 0x47, 0x69, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x47, 0x69, 0x66,
	0x74, 0x12, 0x2e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x5a, 0x74, 0x4c,
	0x69, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x69, 0x66, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x67, 0x69, 0x66, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6d, 0x62, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x6d, 0x62, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x62, 0x6f, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x62, 0x6f, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x6c, 0x6f, 0x74, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x73, 0x6c, 0x6f, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x3e, 0x0a, 0x0c, 0x64, 0x72, 0x61, 0x77, 0x47,
	0x69, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x5a, 0x74, 0x44, 0x72, 0x61,
	0x77, 0x47, 0x69, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x64, 0x72, 0x61, 0x77, 0x47,
	0x69, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66,
	0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CommonActionSignalGift_proto_rawDescOnce sync.Once
	file_CommonActionSignalGift_proto_rawDescData = file_CommonActionSignalGift_proto_rawDesc
)

func file_CommonActionSignalGift_proto_rawDescGZIP() []byte {
	file_CommonActionSignalGift_proto_rawDescOnce.Do(func() {
		file_CommonActionSignalGift_proto_rawDescData = protoimpl.X.CompressGZIP(file_CommonActionSignalGift_proto_rawDescData)
	})
	return file_CommonActionSignalGift_proto_rawDescData
}

var file_CommonActionSignalGift_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CommonActionSignalGift_proto_goTypes = []interface{}{
	(*CommonActionSignalGift)(nil), // 0: AcFunDanmu.CommonActionSignalGift
	(*ZtLiveUserInfo)(nil),         // 1: AcFunDanmu.ZtLiveUserInfo
	(*ZtDrawGiftInfo)(nil),         // 2: AcFunDanmu.ZtDrawGiftInfo
}
var file_CommonActionSignalGift_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.CommonActionSignalGift.user:type_name -> AcFunDanmu.ZtLiveUserInfo
	2, // 1: AcFunDanmu.CommonActionSignalGift.drawGiftInfo:type_name -> AcFunDanmu.ZtDrawGiftInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_CommonActionSignalGift_proto_init() }
func file_CommonActionSignalGift_proto_init() {
	if File_CommonActionSignalGift_proto != nil {
		return
	}
	file_ZtLiveUserInfo_proto_init()
	file_ZtDrawGiftInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CommonActionSignalGift_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonActionSignalGift); i {
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
			RawDescriptor: file_CommonActionSignalGift_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CommonActionSignalGift_proto_goTypes,
		DependencyIndexes: file_CommonActionSignalGift_proto_depIdxs,
		MessageInfos:      file_CommonActionSignalGift_proto_msgTypes,
	}.Build()
	File_CommonActionSignalGift_proto = out.File
	file_CommonActionSignalGift_proto_rawDesc = nil
	file_CommonActionSignalGift_proto_goTypes = nil
	file_CommonActionSignalGift_proto_depIdxs = nil
}
