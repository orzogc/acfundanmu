// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.17.3
// source: Register.proto

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
	return file_Register_proto_enumTypes[0].Descriptor()
}

func (RegisterRequest_PresenceStatus) Type() protoreflect.EnumType {
	return &file_Register_proto_enumTypes[0]
}

func (x RegisterRequest_PresenceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterRequest_PresenceStatus.Descriptor instead.
func (RegisterRequest_PresenceStatus) EnumDescriptor() ([]byte, []int) {
	return file_Register_proto_rawDescGZIP(), []int{0, 0}
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
	return file_Register_proto_enumTypes[1].Descriptor()
}

func (RegisterRequest_ActiveStatus) Type() protoreflect.EnumType {
	return &file_Register_proto_enumTypes[1]
}

func (x RegisterRequest_ActiveStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegisterRequest_ActiveStatus.Descriptor instead.
func (RegisterRequest_ActiveStatus) EnumDescriptor() ([]byte, []int) {
	return file_Register_proto_rawDescGZIP(), []int{0, 1}
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppInfo              *AppInfo                       `protobuf:"bytes,1,opt,name=appInfo,proto3" json:"appInfo,omitempty"`
	DeviceInfo           *DeviceInfo                    `protobuf:"bytes,2,opt,name=deviceInfo,proto3" json:"deviceInfo,omitempty"`
	EnvInfo              *EnvInfo                       `protobuf:"bytes,3,opt,name=envInfo,proto3" json:"envInfo,omitempty"`
	PresenceStatus       RegisterRequest_PresenceStatus `protobuf:"varint,4,opt,name=presenceStatus,proto3,enum=AcFunDanmu.RegisterRequest_PresenceStatus" json:"presenceStatus,omitempty"`
	AppActiveStatus      RegisterRequest_ActiveStatus   `protobuf:"varint,5,opt,name=appActiveStatus,proto3,enum=AcFunDanmu.RegisterRequest_ActiveStatus" json:"appActiveStatus,omitempty"`
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
		mi := &file_Register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Register_proto_msgTypes[0]
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
	return file_Register_proto_rawDescGZIP(), []int{0}
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

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessPointsConfig         *AccessPointsConfig `protobuf:"bytes,1,opt,name=accessPointsConfig,proto3" json:"accessPointsConfig,omitempty"`
	SessKey                    []byte              `protobuf:"bytes,2,opt,name=sessKey,proto3" json:"sessKey,omitempty"`
	InstanceId                 int64               `protobuf:"varint,3,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	SdkOption                  *SdkOption          `protobuf:"bytes,4,opt,name=sdkOption,proto3" json:"sdkOption,omitempty"`
	AccessPointsConfigIpv6     *AccessPointsConfig `protobuf:"bytes,5,opt,name=accessPointsConfigIpv6,proto3" json:"accessPointsConfigIpv6,omitempty"`
	AccessPointsConfigQUic     *AccessPointsConfig `protobuf:"bytes,6,opt,name=accessPointsConfigQUic,proto3" json:"accessPointsConfigQUic,omitempty"`
	AccessPointsConfigQuicIpv6 *AccessPointsConfig `protobuf:"bytes,7,opt,name=accessPointsConfigQuicIpv6,proto3" json:"accessPointsConfigQuicIpv6,omitempty"`
	CleanAccessPoint           bool                `protobuf:"varint,8,opt,name=cleanAccessPoint,proto3" json:"cleanAccessPoint,omitempty"`
	FlowCostSampleRate         float32             `protobuf:"fixed32,9,opt,name=flowCostSampleRate,proto3" json:"flowCostSampleRate,omitempty"`
	CommandSampleRate          float32             `protobuf:"fixed32,10,opt,name=commandSampleRate,proto3" json:"commandSampleRate,omitempty"`
	AccessPointsConfigWs       *AccessPointsConfig `protobuf:"bytes,11,opt,name=accessPointsConfigWs,proto3" json:"accessPointsConfigWs,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_Register_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetAccessPointsConfig() *AccessPointsConfig {
	if x != nil {
		return x.AccessPointsConfig
	}
	return nil
}

func (x *RegisterResponse) GetSessKey() []byte {
	if x != nil {
		return x.SessKey
	}
	return nil
}

func (x *RegisterResponse) GetInstanceId() int64 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *RegisterResponse) GetSdkOption() *SdkOption {
	if x != nil {
		return x.SdkOption
	}
	return nil
}

func (x *RegisterResponse) GetAccessPointsConfigIpv6() *AccessPointsConfig {
	if x != nil {
		return x.AccessPointsConfigIpv6
	}
	return nil
}

func (x *RegisterResponse) GetAccessPointsConfigQUic() *AccessPointsConfig {
	if x != nil {
		return x.AccessPointsConfigQUic
	}
	return nil
}

func (x *RegisterResponse) GetAccessPointsConfigQuicIpv6() *AccessPointsConfig {
	if x != nil {
		return x.AccessPointsConfigQuicIpv6
	}
	return nil
}

func (x *RegisterResponse) GetCleanAccessPoint() bool {
	if x != nil {
		return x.CleanAccessPoint
	}
	return false
}

func (x *RegisterResponse) GetFlowCostSampleRate() float32 {
	if x != nil {
		return x.FlowCostSampleRate
	}
	return 0
}

func (x *RegisterResponse) GetCommandSampleRate() float32 {
	if x != nil {
		return x.CommandSampleRate
	}
	return 0
}

func (x *RegisterResponse) GetAccessPointsConfigWs() *AccessPointsConfig {
	if x != nil {
		return x.AccessPointsConfigWs
	}
	return nil
}

var File_Register_proto protoreflect.FileDescriptor

var file_Register_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x1a, 0x0d, 0x41, 0x70,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x45,
	0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x5a, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x50, 0x75, 0x73, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x64, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x06, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e,
	0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61,
	0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x41, 0x63, 0x46,
	0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d,
	0x0a, 0x07, 0x65, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x45, 0x6e, 0x76,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x52, 0x0a,
	0x0e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e,
	0x6d, 0x75, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x52, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x41, 0x63, 0x46,
	0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f,
	0x61, 0x70, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x48, 0x0a, 0x10, 0x70, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x41, 0x63, 0x46, 0x75,
	0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x10, 0x70, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x50, 0x0a, 0x14, 0x70, 0x75, 0x73,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44,
	0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x14, 0x70, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x6b,
	0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x53, 0x65, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x6b, 0x65, 0x65, 0x70, 0x61,
	0x6c, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x12,
	0x3c, 0x0a, 0x0c, 0x7a, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e,
	0x6d, 0x75, 0x2e, 0x5a, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0c, 0x7a, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a,
	0x0d, 0x69, 0x70, 0x76, 0x36, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x70, 0x76, 0x36, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x22, 0x3b, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x6b, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x63, 0x65, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x6b,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x01,
	0x22, 0x48, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0c, 0x0a, 0x08, 0x6b, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x6b, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x46, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x6b, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x42, 0x61,
	0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x02, 0x22, 0xbf, 0x05, 0x0a, 0x10, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x12, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x41, 0x63,
	0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x12, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x73, 0x64, 0x6b,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41,
	0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x53, 0x64, 0x6b, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x73, 0x64, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x56,
	0x0a, 0x16, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x49, 0x70, 0x76, 0x36, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x16,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x49, 0x70, 0x76, 0x36, 0x12, 0x56, 0x0a, 0x16, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x51, 0x55, 0x69, 0x63,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61,
	0x6e, 0x6d, 0x75, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x16, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x51, 0x55, 0x69, 0x63, 0x12, 0x5e,
	0x0a, 0x1a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x51, 0x75, 0x69, 0x63, 0x49, 0x70, 0x76, 0x36, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61, 0x6e, 0x6d, 0x75, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x1a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x51, 0x75, 0x69, 0x63, 0x49, 0x70, 0x76, 0x36, 0x12, 0x2a,
	0x0a, 0x10, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x66, 0x6c,
	0x6f, 0x77, 0x43, 0x6f, 0x73, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x73, 0x74,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x52, 0x0a, 0x14, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x57, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x41, 0x63, 0x46, 0x75, 0x6e, 0x44, 0x61,
	0x6e, 0x6d, 0x75, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x14, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x57, 0x73, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x7a, 0x6f, 0x67,
	0x63, 0x2f, 0x61, 0x63, 0x66, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x6d, 0x75, 0x2f, 0x61, 0x63, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Register_proto_rawDescOnce sync.Once
	file_Register_proto_rawDescData = file_Register_proto_rawDesc
)

func file_Register_proto_rawDescGZIP() []byte {
	file_Register_proto_rawDescOnce.Do(func() {
		file_Register_proto_rawDescData = protoimpl.X.CompressGZIP(file_Register_proto_rawDescData)
	})
	return file_Register_proto_rawDescData
}

var file_Register_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_Register_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Register_proto_goTypes = []interface{}{
	(RegisterRequest_PresenceStatus)(0), // 0: AcFunDanmu.RegisterRequest.PresenceStatus
	(RegisterRequest_ActiveStatus)(0),   // 1: AcFunDanmu.RegisterRequest.ActiveStatus
	(*RegisterRequest)(nil),             // 2: AcFunDanmu.RegisterRequest
	(*RegisterResponse)(nil),            // 3: AcFunDanmu.RegisterResponse
	(*AppInfo)(nil),                     // 4: AcFunDanmu.AppInfo
	(*DeviceInfo)(nil),                  // 5: AcFunDanmu.DeviceInfo
	(*EnvInfo)(nil),                     // 6: AcFunDanmu.EnvInfo
	(*PushServiceToken)(nil),            // 7: AcFunDanmu.PushServiceToken
	(*ZtCommonInfo)(nil),                // 8: AcFunDanmu.ZtCommonInfo
	(*AccessPointsConfig)(nil),          // 9: AcFunDanmu.AccessPointsConfig
	(*SdkOption)(nil),                   // 10: AcFunDanmu.SdkOption
}
var file_Register_proto_depIdxs = []int32{
	4,  // 0: AcFunDanmu.RegisterRequest.appInfo:type_name -> AcFunDanmu.AppInfo
	5,  // 1: AcFunDanmu.RegisterRequest.deviceInfo:type_name -> AcFunDanmu.DeviceInfo
	6,  // 2: AcFunDanmu.RegisterRequest.envInfo:type_name -> AcFunDanmu.EnvInfo
	0,  // 3: AcFunDanmu.RegisterRequest.presenceStatus:type_name -> AcFunDanmu.RegisterRequest.PresenceStatus
	1,  // 4: AcFunDanmu.RegisterRequest.appActiveStatus:type_name -> AcFunDanmu.RegisterRequest.ActiveStatus
	7,  // 5: AcFunDanmu.RegisterRequest.pushServiceToken:type_name -> AcFunDanmu.PushServiceToken
	7,  // 6: AcFunDanmu.RegisterRequest.pushServiceTokenList:type_name -> AcFunDanmu.PushServiceToken
	8,  // 7: AcFunDanmu.RegisterRequest.ztCommonInfo:type_name -> AcFunDanmu.ZtCommonInfo
	9,  // 8: AcFunDanmu.RegisterResponse.accessPointsConfig:type_name -> AcFunDanmu.AccessPointsConfig
	10, // 9: AcFunDanmu.RegisterResponse.sdkOption:type_name -> AcFunDanmu.SdkOption
	9,  // 10: AcFunDanmu.RegisterResponse.accessPointsConfigIpv6:type_name -> AcFunDanmu.AccessPointsConfig
	9,  // 11: AcFunDanmu.RegisterResponse.accessPointsConfigQUic:type_name -> AcFunDanmu.AccessPointsConfig
	9,  // 12: AcFunDanmu.RegisterResponse.accessPointsConfigQuicIpv6:type_name -> AcFunDanmu.AccessPointsConfig
	9,  // 13: AcFunDanmu.RegisterResponse.accessPointsConfigWs:type_name -> AcFunDanmu.AccessPointsConfig
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_Register_proto_init() }
func file_Register_proto_init() {
	if File_Register_proto != nil {
		return
	}
	file_AppInfo_proto_init()
	file_DeviceInfo_proto_init()
	file_EnvInfo_proto_init()
	file_ZtCommonInfo_proto_init()
	file_AccessPointsConfig_proto_init()
	file_PushServiceToken_proto_init()
	file_SdkOption_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
			RawDescriptor: file_Register_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Register_proto_goTypes,
		DependencyIndexes: file_Register_proto_depIdxs,
		EnumInfos:         file_Register_proto_enumTypes,
		MessageInfos:      file_Register_proto_msgTypes,
	}.Build()
	File_Register_proto = out.File
	file_Register_proto_rawDesc = nil
	file_Register_proto_goTypes = nil
	file_Register_proto_depIdxs = nil
}
