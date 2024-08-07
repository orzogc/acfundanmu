// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: RegisterRequest.proto

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

type RegisterRequest_PresenceStatus int32

const (
	RegisterRequest_kPresenceOffline RegisterRequest_PresenceStatus = 0
	RegisterRequest_kPresenceOnline  RegisterRequest_PresenceStatus = 1
)

// Enum value maps for RegisterRequest_PresenceStatus.
var (
	RegisterRequest_PresenceStatus_name = map[int32]string{
		0: "kPresenceOffline",
		1: "kPresenceOnline",
	}
	RegisterRequest_PresenceStatus_value = map[string]int32{
		"kPresenceOffline": 0,
		"kPresenceOnline":  1,
	}
)

func (x RegisterRequest_PresenceStatus) Enum() *RegisterRequest_PresenceStatus {
	p := new(RegisterRequest_PresenceStatus)
	*p = x
	return p
}

func (x RegisterRequest_PresenceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegisterRequest_PresenceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_RegisterRequest_proto_enumTypes[0].Descriptor()
}

func (RegisterRequest_PresenceStatus) Type() protoreflect.EnumType {
	return &file_RegisterRequest_proto_enumTypes[0]
}

func (x RegisterRequest_PresenceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterRequest_PresenceStatus.Descriptor instead.
func (RegisterRequest_PresenceStatus) EnumDescriptor() ([]byte, []int) {
	return file_RegisterRequest_proto_rawDescGZIP(), []int{0, 0}
}

type RegisterRequest_ActiveStatus int32

const (
	RegisterRequest_kInvalid         RegisterRequest_ActiveStatus = 0
	RegisterRequest_kAppInForeground RegisterRequest_ActiveStatus = 1
	RegisterRequest_kAppInBackground RegisterRequest_ActiveStatus = 2
)

// Enum value maps for RegisterRequest_ActiveStatus.
var (
	RegisterRequest_ActiveStatus_name = map[int32]string{
		0: "kInvalid",
		1: "kAppInForeground",
		2: "kAppInBackground",
	}
	RegisterRequest_ActiveStatus_value = map[string]int32{
		"kInvalid":         0,
		"kAppInForeground": 1,
		"kAppInBackground": 2,
	}
)

func (x RegisterRequest_ActiveStatus) Enum() *RegisterRequest_ActiveStatus {
	p := new(RegisterRequest_ActiveStatus)
	*p = x
	return p
}

func (x RegisterRequest_ActiveStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegisterRequest_ActiveStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_RegisterRequest_proto_enumTypes[1].Descriptor()
}

func (RegisterRequest_ActiveStatus) Type() protoreflect.EnumType {
	return &file_RegisterRequest_proto_enumTypes[1]
}

func (x RegisterRequest_ActiveStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterRequest_ActiveStatus.Descriptor instead.
func (RegisterRequest_ActiveStatus) EnumDescriptor() ([]byte, []int) {
	return file_RegisterRequest_proto_rawDescGZIP(), []int{0, 1}
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppInfo              *AppInfo                       `protobuf:"bytes,1,opt,name=appInfo,proto3" json:"appInfo,omitempty"`
	DeviceInfo           *DeviceInfo                    `protobuf:"bytes,2,opt,name=deviceInfo,proto3" json:"deviceInfo,omitempty"`
	EnvInfo              *EnvInfo                       `protobuf:"bytes,3,opt,name=envInfo,proto3" json:"envInfo,omitempty"`
	PresenceStatus       RegisterRequest_PresenceStatus `protobuf:"varint,4,opt,name=presenceStatus,proto3,enum=AcFunDanmu.Im.Basic.RegisterRequest_PresenceStatus" json:"presenceStatus,omitempty"`
	AppActiveStatus      RegisterRequest_ActiveStatus   `protobuf:"varint,5,opt,name=appActiveStatus,proto3,enum=AcFunDanmu.Im.Basic.RegisterRequest_ActiveStatus" json:"appActiveStatus,omitempty"`
	AppCustomStatus      []byte                         `protobuf:"bytes,6,opt,name=appCustomStatus,proto3" json:"appCustomStatus,omitempty"`
	PushServiceToken     *PushServiceToken              `protobuf:"bytes,7,opt,name=pushServiceToken,proto3" json:"pushServiceToken,omitempty"`
	InstanceId           int64                          `protobuf:"varint,8,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	PushServiceTokenList []*PushServiceToken            `protobuf:"bytes,9,rep,name=pushServiceTokenList,proto3" json:"pushServiceTokenList,omitempty"`
	KeepaliveIntervalSec int32                          `protobuf:"varint,10,opt,name=keepaliveIntervalSec,proto3" json:"keepaliveIntervalSec,omitempty"`
	ZtCommonInfo         *ZtCommonInfo                  `protobuf:"bytes,11,opt,name=ztCommonInfo,proto3" json:"ztCommonInfo,omitempty"`
	Ipv6Available        bool                           `protobuf:"varint,12,opt,name=ipv6Available,proto3" json:"ipv6Available,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RegisterRequest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_RegisterRequest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_RegisterRequest_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetAppInfo() *AppInfo {
	if x != nil {
		return x.AppInfo
	}
	return nil
}

func (x *RegisterRequest) GetDeviceInfo() *DeviceInfo {
	if x != nil {
		return x.DeviceInfo
	}
	return nil
}

func (x *RegisterRequest) GetEnvInfo() *EnvInfo {
	if x != nil {
		return x.EnvInfo
	}
	return nil
}

func (x *RegisterRequest) GetPresenceStatus() RegisterRequest_PresenceStatus {
	if x != nil {
		return x.PresenceStatus
	}
	return RegisterRequest_kPresenceOffline
}

func (x *RegisterRequest) GetAppActiveStatus() RegisterRequest_ActiveStatus {
	if x != nil {
		return x.AppActiveStatus
	}
	return RegisterRequest_kInvalid
}

func (x *RegisterRequest) GetAppCustomStatus() []byte {
	if x != nil {
		return x.AppCustomStatus
	}
	return nil
}

func (x *RegisterRequest) GetPushServiceToken() *PushServiceToken {
	if x != nil {
		return x.PushServiceToken
	}
	return nil
}

func (x *RegisterRequest) GetInstanceId() int64 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *RegisterRequest) GetPushServiceTokenList() []*PushServiceToken {
	if x != nil {
		return x.PushServiceTokenList
	}
	return nil
}

func (x *RegisterRequest) GetKeepaliveIntervalSec() int32 {
	if x != nil {
		return x.KeepaliveIntervalSec
	}
	return 0
}

func (x *RegisterRequest) GetZtCommonInfo() *ZtCommonInfo {
	if x != nil {
		return x.ZtCommonInfo
	}
	return nil
}

func (x *RegisterRequest) GetIpv6Available() bool {
	if x != nil {
		return x.Ipv6Available
	}
	return false
}

var File_RegisterRequest_proto protoreflect.FileDescriptor

var file_RegisterRequest_proto_rawDesc = []byte{
	0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61,
	0x6e, 0x6d, 0x75, 0x2e, 0x49, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x1a, 0x0d, 0x41, 0x70,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x45,
	0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x50, 0x75,
	0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x5a, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x07, 0x0a, 0x0f, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x07,
	0x61, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x49, 0x6d, 0x2e, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x70, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e,
	0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x49, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x07, 0x65, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61,
	0x6e, 0x6d, 0x75, 0x2e, 0x49, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x45, 0x6e, 0x76,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5b, 0x0a,
	0x0e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e,
	0x6d, 0x75, 0x2e, 0x49, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5b, 0x0a, 0x0f, 0x61, 0x70,
	0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75,
	0x2e, 0x49, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0f, 0x61, 0x70, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x51, 0x0a, 0x10, 0x70, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x41, 0x63,
	0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x49, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x10, 0x70, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x59, 0x0a, 0x14, 0x70, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e,
	0x49, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x14, 0x70, 0x75, 0x73, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x32, 0x0a, 0x14, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x6b,
	0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x53, 0x65, 0x63, 0x12, 0x45, 0x0a, 0x0c, 0x7a, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x41, 0x63, 0x46, 0x75,
	0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x49, 0x6d, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x5a, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x7a, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x70,
	0x76, 0x36, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x69, 0x70, 0x76, 0x36, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x22, 0x3b, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x6b, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x4f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x6b, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x01, 0x22, 0x48, 0x0a,
	0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a,
	0x08, 0x6b, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x6b,
	0x41, 0x70, 0x70, 0x49, 0x6e, 0x46, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x10,
	0x01, 0x12, 0x14, 0x0a, 0x10, 0x6b, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x02, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67, 0x63, 0x2f, 0x61, 0x63, 0x66,
	0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RegisterRequest_proto_rawDescOnce sync.Once
	file_RegisterRequest_proto_rawDescData = file_RegisterRequest_proto_rawDesc
)

func file_RegisterRequest_proto_rawDescGZIP() []byte {
	file_RegisterRequest_proto_rawDescOnce.Do(func() {
		file_RegisterRequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_RegisterRequest_proto_rawDescData)
	})
	return file_RegisterRequest_proto_rawDescData
}

var file_RegisterRequest_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_RegisterRequest_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RegisterRequest_proto_goTypes = []any{
	(RegisterRequest_PresenceStatus)(0), // 0: AcFunDanmu.Im.Basic.RegisterRequest.PresenceStatus
	(RegisterRequest_ActiveStatus)(0),   // 1: AcFunDanmu.Im.Basic.RegisterRequest.ActiveStatus
	(*RegisterRequest)(nil),             // 2: AcFunDanmu.Im.Basic.RegisterRequest
	(*AppInfo)(nil),                     // 3: AcFunDanmu.Im.Basic.AppInfo
	(*DeviceInfo)(nil),                  // 4: AcFunDanmu.Im.Basic.DeviceInfo
	(*EnvInfo)(nil),                     // 5: AcFunDanmu.Im.Basic.EnvInfo
	(*PushServiceToken)(nil),            // 6: AcFunDanmu.Im.Basic.PushServiceToken
	(*ZtCommonInfo)(nil),                // 7: AcFunDanmu.Im.Basic.ZtCommonInfo
}
var file_RegisterRequest_proto_depIdxs = []int32{
	3, // 0: AcFunDanmu.Im.Basic.RegisterRequest.appInfo:type_name -> AcFunDanmu.Im.Basic.AppInfo
	4, // 1: AcFunDanmu.Im.Basic.RegisterRequest.deviceInfo:type_name -> AcFunDanmu.Im.Basic.DeviceInfo
	5, // 2: AcFunDanmu.Im.Basic.RegisterRequest.envInfo:type_name -> AcFunDanmu.Im.Basic.EnvInfo
	0, // 3: AcFunDanmu.Im.Basic.RegisterRequest.presenceStatus:type_name -> AcFunDanmu.Im.Basic.RegisterRequest.PresenceStatus
	1, // 4: AcFunDanmu.Im.Basic.RegisterRequest.appActiveStatus:type_name -> AcFunDanmu.Im.Basic.RegisterRequest.ActiveStatus
	6, // 5: AcFunDanmu.Im.Basic.RegisterRequest.pushServiceToken:type_name -> AcFunDanmu.Im.Basic.PushServiceToken
	6, // 6: AcFunDanmu.Im.Basic.RegisterRequest.pushServiceTokenList:type_name -> AcFunDanmu.Im.Basic.PushServiceToken
	7, // 7: AcFunDanmu.Im.Basic.RegisterRequest.ztCommonInfo:type_name -> AcFunDanmu.Im.Basic.ZtCommonInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_RegisterRequest_proto_init() }
func file_RegisterRequest_proto_init() {
	if File_RegisterRequest_proto != nil {
		return
	}
	file_AppInfo_proto_init()
	file_DeviceInfo_proto_init()
	file_EnvInfo_proto_init()
	file_PushServiceToken_proto_init()
	file_ZtCommonInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RegisterRequest_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterRequest); i {
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
			RawDescriptor: file_RegisterRequest_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RegisterRequest_proto_goTypes,
		DependencyIndexes: file_RegisterRequest_proto_depIdxs,
		EnumInfos:         file_RegisterRequest_proto_enumTypes,
		MessageInfos:      file_RegisterRequest_proto_msgTypes,
	}.Build()
	File_RegisterRequest_proto = out.File
	file_RegisterRequest_proto_rawDesc = nil
	file_RegisterRequest_proto_goTypes = nil
	file_RegisterRequest_proto_depIdxs = nil
}
