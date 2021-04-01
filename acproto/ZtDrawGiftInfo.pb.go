// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: ZtDrawGiftInfo.proto

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

type ZtDrawGiftInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScreenWidth  int64                         `protobuf:"varint,1,opt,name=screenWidth,proto3" json:"screenWidth,omitempty"`
	ScreenHeight int64                         `protobuf:"varint,2,opt,name=screenHeight,proto3" json:"screenHeight,omitempty"`
	DrawPoint    []*ZtDrawGiftInfo_ZtDrawPoint `protobuf:"bytes,3,rep,name=drawPoint,proto3" json:"drawPoint,omitempty"`
}

func (x *ZtDrawGiftInfo) Reset() {
	*x = ZtDrawGiftInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ZtDrawGiftInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZtDrawGiftInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZtDrawGiftInfo) ProtoMessage() {}

func (x *ZtDrawGiftInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ZtDrawGiftInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZtDrawGiftInfo.ProtoReflect.Descriptor instead.
func (*ZtDrawGiftInfo) Descriptor() ([]byte, []int) {
	return file_ZtDrawGiftInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ZtDrawGiftInfo) GetScreenWidth() int64 {
	if x != nil {
		return x.ScreenWidth
	}
	return 0
}

func (x *ZtDrawGiftInfo) GetScreenHeight() int64 {
	if x != nil {
		return x.ScreenHeight
	}
	return 0
}

func (x *ZtDrawGiftInfo) GetDrawPoint() []*ZtDrawGiftInfo_ZtDrawPoint {
	if x != nil {
		return x.DrawPoint
	}
	return nil
}

type ZtDrawGiftInfo_ZtDrawPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MarginLeft int64   `protobuf:"varint,1,opt,name=marginLeft,proto3" json:"marginLeft,omitempty"`
	MarginTop  int64   `protobuf:"varint,2,opt,name=marginTop,proto3" json:"marginTop,omitempty"`
	ScaleRatio float64 `protobuf:"fixed64,3,opt,name=scaleRatio,proto3" json:"scaleRatio,omitempty"`
	Handup     bool    `protobuf:"varint,4,opt,name=handup,proto3" json:"handup,omitempty"`
}

func (x *ZtDrawGiftInfo_ZtDrawPoint) Reset() {
	*x = ZtDrawGiftInfo_ZtDrawPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ZtDrawGiftInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZtDrawGiftInfo_ZtDrawPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZtDrawGiftInfo_ZtDrawPoint) ProtoMessage() {}

func (x *ZtDrawGiftInfo_ZtDrawPoint) ProtoReflect() protoreflect.Message {
	mi := &file_ZtDrawGiftInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZtDrawGiftInfo_ZtDrawPoint.ProtoReflect.Descriptor instead.
func (*ZtDrawGiftInfo_ZtDrawPoint) Descriptor() ([]byte, []int) {
	return file_ZtDrawGiftInfo_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ZtDrawGiftInfo_ZtDrawPoint) GetMarginLeft() int64 {
	if x != nil {
		return x.MarginLeft
	}
	return 0
}

func (x *ZtDrawGiftInfo_ZtDrawPoint) GetMarginTop() int64 {
	if x != nil {
		return x.MarginTop
	}
	return 0
}

func (x *ZtDrawGiftInfo_ZtDrawPoint) GetScaleRatio() float64 {
	if x != nil {
		return x.ScaleRatio
	}
	return 0
}

func (x *ZtDrawGiftInfo_ZtDrawPoint) GetHandup() bool {
	if x != nil {
		return x.Handup
	}
	return false
}

var File_ZtDrawGiftInfo_proto protoreflect.FileDescriptor

var file_ZtDrawGiftInfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x5a, 0x74, 0x44, 0x72, 0x61, 0x77, 0x47, 0x69, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e,
	0x6d, 0x75, 0x22, 0xa2, 0x02, 0x0a, 0x0e, 0x5a, 0x74, 0x44, 0x72, 0x61, 0x77, 0x47, 0x69, 0x66,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x57,
	0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x44, 0x0a, 0x09, 0x64,
	0x72, 0x61, 0x77, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x5a, 0x74, 0x44, 0x72,
	0x61, 0x77, 0x47, 0x69, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x5a, 0x74, 0x44, 0x72, 0x61,
	0x77, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x09, 0x64, 0x72, 0x61, 0x77, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x1a, 0x83, 0x01, 0x0a, 0x0b, 0x5a, 0x74, 0x44, 0x72, 0x61, 0x77, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x4c, 0x65, 0x66, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x4c, 0x65, 0x66,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x70, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0a, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x68, 0x61, 0x6e, 0x64, 0x75, 0x70, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66,
	0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ZtDrawGiftInfo_proto_rawDescOnce sync.Once
	file_ZtDrawGiftInfo_proto_rawDescData = file_ZtDrawGiftInfo_proto_rawDesc
)

func file_ZtDrawGiftInfo_proto_rawDescGZIP() []byte {
	file_ZtDrawGiftInfo_proto_rawDescOnce.Do(func() {
		file_ZtDrawGiftInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ZtDrawGiftInfo_proto_rawDescData)
	})
	return file_ZtDrawGiftInfo_proto_rawDescData
}

var file_ZtDrawGiftInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ZtDrawGiftInfo_proto_goTypes = []interface{}{
	(*ZtDrawGiftInfo)(nil),             // 0: AcFunDanmu.ZtDrawGiftInfo
	(*ZtDrawGiftInfo_ZtDrawPoint)(nil), // 1: AcFunDanmu.ZtDrawGiftInfo.ZtDrawPoint
}
var file_ZtDrawGiftInfo_proto_depIdxs = []int32{
	1, // 0: AcFunDanmu.ZtDrawGiftInfo.drawPoint:type_name -> AcFunDanmu.ZtDrawGiftInfo.ZtDrawPoint
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ZtDrawGiftInfo_proto_init() }
func file_ZtDrawGiftInfo_proto_init() {
	if File_ZtDrawGiftInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ZtDrawGiftInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZtDrawGiftInfo); i {
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
		file_ZtDrawGiftInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZtDrawGiftInfo_ZtDrawPoint); i {
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
			RawDescriptor: file_ZtDrawGiftInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ZtDrawGiftInfo_proto_goTypes,
		DependencyIndexes: file_ZtDrawGiftInfo_proto_depIdxs,
		MessageInfos:      file_ZtDrawGiftInfo_proto_msgTypes,
	}.Build()
	File_ZtDrawGiftInfo_proto = out.File
	file_ZtDrawGiftInfo_proto_rawDesc = nil
	file_ZtDrawGiftInfo_proto_goTypes = nil
	file_ZtDrawGiftInfo_proto_depIdxs = nil
}
