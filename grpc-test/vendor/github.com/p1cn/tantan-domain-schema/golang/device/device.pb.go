// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device/device.proto

/*
Package device is a generated protocol buffer package.

It is generated from these files:
	device/device.proto
	device/device.service.proto

It has these top-level messages:
	Device
	DevicePushNotifications
	PushNotificationToken
	DeviceIdentifier
	DeviceApp
	DeviceOperatingSystem
	DeviceHardware
	BuildInfo
	DevicesReply
	DeviceIdentifiersReply
	DeviceNameFilter
	DeviceIdentifierTokensFilter
	FindDevicesByIdsRequest
	FindDevicesByIdsParams
	FindDevicesByUserIdRequest
	FindDevicesByUserIdParams
	FindDevicesByUserIdsRequest
	FindDevicesByUserIdsParams
	FindDevicesByDeviceIdentifierTokensRequest
	FindDevicesByDeviceIdentifierTokensParams
	FindExistsByUserIdRequest
	FindExistsByUserIdParams
	FindDeviceIdentifiersByDeviceIdentifierTokensRequest
	FindDeviceIdentifiersByDeviceIdentifierTokensParams
	InsertDeviceRequest
	InsertDeviceParams
	UpdateDeviceRequest
	UpdateDeviceParams
	RemoveDevicesByIdsRequest
	RemoveDevicesByIdsParams
	RemoveDevicesByUserIdsRequest
	RemoveDevicesByUserIdsParams
	InvalidateTokenRequest
	InvalidateTokenParams
	InvalidateTokenReply
*/
package device

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Device_DeviceStatusEnum int32

const (
	Device_STATUS_DEFAULT Device_DeviceStatusEnum = 0
	Device_STATUS_DELETED Device_DeviceStatusEnum = 1
)

var Device_DeviceStatusEnum_name = map[int32]string{
	0: "STATUS_DEFAULT",
	1: "STATUS_DELETED",
}
var Device_DeviceStatusEnum_value = map[string]int32{
	"STATUS_DEFAULT": 0,
	"STATUS_DELETED": 1,
}

func (x Device_DeviceStatusEnum) String() string {
	return proto.EnumName(Device_DeviceStatusEnum_name, int32(x))
}
func (Device_DeviceStatusEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type PushNotificationToken_PushNotificationTokenStatusEnum int32

const (
	PushNotificationToken_STATUS_DEFAULT PushNotificationToken_PushNotificationTokenStatusEnum = 0
	PushNotificationToken_STATUS_DELETED PushNotificationToken_PushNotificationTokenStatusEnum = 1
)

var PushNotificationToken_PushNotificationTokenStatusEnum_name = map[int32]string{
	0: "STATUS_DEFAULT",
	1: "STATUS_DELETED",
}
var PushNotificationToken_PushNotificationTokenStatusEnum_value = map[string]int32{
	"STATUS_DEFAULT": 0,
	"STATUS_DELETED": 1,
}

func (x PushNotificationToken_PushNotificationTokenStatusEnum) String() string {
	return proto.EnumName(PushNotificationToken_PushNotificationTokenStatusEnum_name, int32(x))
}
func (PushNotificationToken_PushNotificationTokenStatusEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

type DeviceIdentifier_DeviceIdentifierStatusEnum int32

const (
	DeviceIdentifier_STATUS_DEFAULT DeviceIdentifier_DeviceIdentifierStatusEnum = 0
	DeviceIdentifier_STATUS_DELETED DeviceIdentifier_DeviceIdentifierStatusEnum = 1
)

var DeviceIdentifier_DeviceIdentifierStatusEnum_name = map[int32]string{
	0: "STATUS_DEFAULT",
	1: "STATUS_DELETED",
}
var DeviceIdentifier_DeviceIdentifierStatusEnum_value = map[string]int32{
	"STATUS_DEFAULT": 0,
	"STATUS_DELETED": 1,
}

func (x DeviceIdentifier_DeviceIdentifierStatusEnum) String() string {
	return proto.EnumName(DeviceIdentifier_DeviceIdentifierStatusEnum_name, int32(x))
}
func (DeviceIdentifier_DeviceIdentifierStatusEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type BuildInfo_BuildInfoStatusEnum int32

const (
	BuildInfo_STATUS_DEFAULT BuildInfo_BuildInfoStatusEnum = 0
	BuildInfo_STATUS_DELETED BuildInfo_BuildInfoStatusEnum = 1
)

var BuildInfo_BuildInfoStatusEnum_name = map[int32]string{
	0: "STATUS_DEFAULT",
	1: "STATUS_DELETED",
}
var BuildInfo_BuildInfoStatusEnum_value = map[string]int32{
	"STATUS_DEFAULT": 0,
	"STATUS_DELETED": 1,
}

func (x BuildInfo_BuildInfoStatusEnum) String() string {
	return proto.EnumName(BuildInfo_BuildInfoStatusEnum_name, int32(x))
}
func (BuildInfo_BuildInfoStatusEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0}
}

type Device struct {
	Id                string                   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OwnerId           string                   `protobuf:"bytes,2,opt,name=ownerId" json:"ownerId,omitempty"`
	PushNotifications *DevicePushNotifications `protobuf:"bytes,3,opt,name=pushNotifications" json:"pushNotifications,omitempty"`
	DeviceIdentifiers []*DeviceIdentifier      `protobuf:"bytes,4,rep,name=deviceIdentifiers" json:"deviceIdentifiers,omitempty"`
	App               *DeviceApp               `protobuf:"bytes,5,opt,name=app" json:"app,omitempty"`
	OperatingSystem   *DeviceOperatingSystem   `protobuf:"bytes,6,opt,name=operatingSystem" json:"operatingSystem,omitempty"`
	Hardware          *DeviceHardware          `protobuf:"bytes,7,opt,name=hardware" json:"hardware,omitempty"`
	CreatedTime       int64                    `protobuf:"varint,8,opt,name=createdTime" json:"createdTime,omitempty"`
	Status            Device_DeviceStatusEnum  `protobuf:"varint,9,opt,name=status,enum=device.Device_DeviceStatusEnum" json:"status,omitempty"`
	Ringtone          string                   `protobuf:"bytes,10,opt,name=ringtone" json:"ringtone,omitempty"`
	// for 兼容
	DeviceIdentifier string `protobuf:"bytes,11,opt,name=deviceIdentifier" json:"deviceIdentifier,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Device) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Device) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *Device) GetPushNotifications() *DevicePushNotifications {
	if m != nil {
		return m.PushNotifications
	}
	return nil
}

func (m *Device) GetDeviceIdentifiers() []*DeviceIdentifier {
	if m != nil {
		return m.DeviceIdentifiers
	}
	return nil
}

func (m *Device) GetApp() *DeviceApp {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *Device) GetOperatingSystem() *DeviceOperatingSystem {
	if m != nil {
		return m.OperatingSystem
	}
	return nil
}

func (m *Device) GetHardware() *DeviceHardware {
	if m != nil {
		return m.Hardware
	}
	return nil
}

func (m *Device) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *Device) GetStatus() Device_DeviceStatusEnum {
	if m != nil {
		return m.Status
	}
	return Device_STATUS_DEFAULT
}

func (m *Device) GetRingtone() string {
	if m != nil {
		return m.Ringtone
	}
	return ""
}

func (m *Device) GetDeviceIdentifier() string {
	if m != nil {
		return m.DeviceIdentifier
	}
	return ""
}

type DevicePushNotifications struct {
	Service string                   `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Token   string                   `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Tokens  []*PushNotificationToken `protobuf:"bytes,3,rep,name=tokens" json:"tokens,omitempty"`
}

func (m *DevicePushNotifications) Reset()                    { *m = DevicePushNotifications{} }
func (m *DevicePushNotifications) String() string            { return proto.CompactTextString(m) }
func (*DevicePushNotifications) ProtoMessage()               {}
func (*DevicePushNotifications) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DevicePushNotifications) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *DevicePushNotifications) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DevicePushNotifications) GetTokens() []*PushNotificationToken {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type PushNotificationToken struct {
	Service string                                                `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Token   string                                                `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Status  PushNotificationToken_PushNotificationTokenStatusEnum `protobuf:"varint,3,opt,name=status,enum=device.PushNotificationToken_PushNotificationTokenStatusEnum" json:"status,omitempty"`
}

func (m *PushNotificationToken) Reset()                    { *m = PushNotificationToken{} }
func (m *PushNotificationToken) String() string            { return proto.CompactTextString(m) }
func (*PushNotificationToken) ProtoMessage()               {}
func (*PushNotificationToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PushNotificationToken) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *PushNotificationToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PushNotificationToken) GetStatus() PushNotificationToken_PushNotificationTokenStatusEnum {
	if m != nil {
		return m.Status
	}
	return PushNotificationToken_STATUS_DEFAULT
}

type DeviceIdentifier struct {
	Name        string                                      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Token       string                                      `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Id          string                                      `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	DeviceId    string                                      `protobuf:"bytes,4,opt,name=deviceId" json:"deviceId,omitempty"`
	CreatedTime int64                                       `protobuf:"varint,5,opt,name=createdTime" json:"createdTime,omitempty"`
	UpdatedTime int64                                       `protobuf:"varint,6,opt,name=updatedTime" json:"updatedTime,omitempty"`
	Status      DeviceIdentifier_DeviceIdentifierStatusEnum `protobuf:"varint,7,opt,name=status,enum=device.DeviceIdentifier_DeviceIdentifierStatusEnum" json:"status,omitempty"`
}

func (m *DeviceIdentifier) Reset()                    { *m = DeviceIdentifier{} }
func (m *DeviceIdentifier) String() string            { return proto.CompactTextString(m) }
func (*DeviceIdentifier) ProtoMessage()               {}
func (*DeviceIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeviceIdentifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceIdentifier) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DeviceIdentifier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeviceIdentifier) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *DeviceIdentifier) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *DeviceIdentifier) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

func (m *DeviceIdentifier) GetStatus() DeviceIdentifier_DeviceIdentifierStatusEnum {
	if m != nil {
		return m.Status
	}
	return DeviceIdentifier_STATUS_DEFAULT
}

type DeviceApp struct {
	Version   string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Build     string `protobuf:"bytes,2,opt,name=build" json:"build,omitempty"`
	UiVersion string `protobuf:"bytes,3,opt,name=uiVersion" json:"uiVersion,omitempty"`
	Language  string `protobuf:"bytes,4,opt,name=language" json:"language,omitempty"`
}

func (m *DeviceApp) Reset()                    { *m = DeviceApp{} }
func (m *DeviceApp) String() string            { return proto.CompactTextString(m) }
func (*DeviceApp) ProtoMessage()               {}
func (*DeviceApp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeviceApp) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DeviceApp) GetBuild() string {
	if m != nil {
		return m.Build
	}
	return ""
}

func (m *DeviceApp) GetUiVersion() string {
	if m != nil {
		return m.UiVersion
	}
	return ""
}

func (m *DeviceApp) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

type DeviceOperatingSystem struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Locale  string `protobuf:"bytes,3,opt,name=locale" json:"locale,omitempty"`
}

func (m *DeviceOperatingSystem) Reset()                    { *m = DeviceOperatingSystem{} }
func (m *DeviceOperatingSystem) String() string            { return proto.CompactTextString(m) }
func (*DeviceOperatingSystem) ProtoMessage()               {}
func (*DeviceOperatingSystem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeviceOperatingSystem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceOperatingSystem) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DeviceOperatingSystem) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

type DeviceHardware struct {
	Name       string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	MacAddress string `protobuf:"bytes,2,opt,name=macAddress" json:"macAddress,omitempty"`
	IpAddress  string `protobuf:"bytes,3,opt,name=ipAddress" json:"ipAddress,omitempty"`
}

func (m *DeviceHardware) Reset()                    { *m = DeviceHardware{} }
func (m *DeviceHardware) String() string            { return proto.CompactTextString(m) }
func (*DeviceHardware) ProtoMessage()               {}
func (*DeviceHardware) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeviceHardware) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceHardware) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *DeviceHardware) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

type BuildInfo struct {
	Id            string                        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OwnerId       string                        `protobuf:"bytes,2,opt,name=ownerId" json:"ownerId,omitempty"`
	BuildInfoHash string                        `protobuf:"bytes,3,opt,name=buildInfoHash" json:"buildInfoHash,omitempty"`
	BuildInfoJson string                        `protobuf:"bytes,4,opt,name=buildInfoJson" json:"buildInfoJson,omitempty"`
	CreatedTime   int64                         `protobuf:"varint,5,opt,name=createdTime" json:"createdTime,omitempty"`
	UpdatedTime   int64                         `protobuf:"varint,6,opt,name=updatedTime" json:"updatedTime,omitempty"`
	Status        BuildInfo_BuildInfoStatusEnum `protobuf:"varint,7,opt,name=status,enum=device.BuildInfo_BuildInfoStatusEnum" json:"status,omitempty"`
}

func (m *BuildInfo) Reset()                    { *m = BuildInfo{} }
func (m *BuildInfo) String() string            { return proto.CompactTextString(m) }
func (*BuildInfo) ProtoMessage()               {}
func (*BuildInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BuildInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BuildInfo) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *BuildInfo) GetBuildInfoHash() string {
	if m != nil {
		return m.BuildInfoHash
	}
	return ""
}

func (m *BuildInfo) GetBuildInfoJson() string {
	if m != nil {
		return m.BuildInfoJson
	}
	return ""
}

func (m *BuildInfo) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *BuildInfo) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

func (m *BuildInfo) GetStatus() BuildInfo_BuildInfoStatusEnum {
	if m != nil {
		return m.Status
	}
	return BuildInfo_STATUS_DEFAULT
}

func init() {
	proto.RegisterType((*Device)(nil), "device.Device")
	proto.RegisterType((*DevicePushNotifications)(nil), "device.DevicePushNotifications")
	proto.RegisterType((*PushNotificationToken)(nil), "device.PushNotificationToken")
	proto.RegisterType((*DeviceIdentifier)(nil), "device.DeviceIdentifier")
	proto.RegisterType((*DeviceApp)(nil), "device.DeviceApp")
	proto.RegisterType((*DeviceOperatingSystem)(nil), "device.DeviceOperatingSystem")
	proto.RegisterType((*DeviceHardware)(nil), "device.DeviceHardware")
	proto.RegisterType((*BuildInfo)(nil), "device.BuildInfo")
	proto.RegisterEnum("device.Device_DeviceStatusEnum", Device_DeviceStatusEnum_name, Device_DeviceStatusEnum_value)
	proto.RegisterEnum("device.PushNotificationToken_PushNotificationTokenStatusEnum", PushNotificationToken_PushNotificationTokenStatusEnum_name, PushNotificationToken_PushNotificationTokenStatusEnum_value)
	proto.RegisterEnum("device.DeviceIdentifier_DeviceIdentifierStatusEnum", DeviceIdentifier_DeviceIdentifierStatusEnum_name, DeviceIdentifier_DeviceIdentifierStatusEnum_value)
	proto.RegisterEnum("device.BuildInfo_BuildInfoStatusEnum", BuildInfo_BuildInfoStatusEnum_name, BuildInfo_BuildInfoStatusEnum_value)
}

func init() { proto.RegisterFile("device/device.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x7d, 0x8e, 0x53, 0x37, 0xb9, 0xd1, 0xcb, 0x4b, 0xa7, 0xaf, 0x7d, 0xa3, 0xea, 0x41, 0x2d,
	0x03, 0x52, 0x84, 0xd4, 0x46, 0xa4, 0x20, 0x24, 0xa4, 0x2e, 0x52, 0x25, 0xa5, 0x81, 0xf2, 0x21,
	0x27, 0x65, 0x81, 0x84, 0xd0, 0xc4, 0x9e, 0x26, 0x23, 0xe2, 0x19, 0xcb, 0x1f, 0xad, 0xd8, 0xf1,
	0xcf, 0x58, 0xb0, 0xe3, 0x9f, 0xf0, 0x2f, 0x90, 0xed, 0xb1, 0x13, 0x3b, 0x0e, 0xb4, 0x12, 0x2b,
	0xcf, 0x3d, 0xf7, 0xdc, 0xeb, 0x99, 0x73, 0x8f, 0xc7, 0xb0, 0x6d, 0xd3, 0x2b, 0x66, 0xd1, 0x4e,
	0xf2, 0x38, 0x74, 0x3d, 0x11, 0x08, 0xa4, 0x25, 0x91, 0xf1, 0xbd, 0x0a, 0x5a, 0x3f, 0x5e, 0xa2,
	0x26, 0x54, 0x98, 0x8d, 0x15, 0x5d, 0x69, 0xd7, 0xcd, 0x0a, 0xb3, 0x11, 0x86, 0x4d, 0x71, 0xcd,
	0xa9, 0x37, 0xb4, 0x71, 0x25, 0x06, 0xd3, 0x10, 0xbd, 0x82, 0x2d, 0x37, 0xf4, 0x67, 0xaf, 0x45,
	0xc0, 0x2e, 0x99, 0x45, 0x02, 0x26, 0xb8, 0x8f, 0x55, 0x5d, 0x69, 0x37, 0xba, 0xfb, 0x87, 0xf2,
	0x35, 0x49, 0xd3, 0xb7, 0x45, 0x9a, 0xb9, 0x5a, 0x89, 0x4e, 0x61, 0x2b, 0x29, 0x1a, 0xda, 0x94,
	0x47, 0x09, 0xea, 0xf9, 0xb8, 0xaa, 0xab, 0xed, 0x46, 0x17, 0xe7, 0xdb, 0x2d, 0x08, 0xe6, 0x6a,
	0x09, 0xba, 0x07, 0x2a, 0x71, 0x5d, 0xbc, 0x11, 0x6f, 0x64, 0x2b, 0x5f, 0xd9, 0x73, 0x5d, 0x33,
	0xca, 0xa2, 0xe7, 0xf0, 0x8f, 0x70, 0xa9, 0x47, 0x02, 0xc6, 0xa7, 0xa3, 0xcf, 0x7e, 0x40, 0x1d,
	0xac, 0xc5, 0x05, 0x77, 0xf2, 0x05, 0x6f, 0xf2, 0x24, 0xb3, 0x58, 0x85, 0xba, 0x50, 0x9b, 0x11,
	0xcf, 0xbe, 0x26, 0x1e, 0xc5, 0x9b, 0x71, 0x87, 0xdd, 0x7c, 0x87, 0x33, 0x99, 0x35, 0x33, 0x1e,
	0xd2, 0xa1, 0x61, 0x79, 0x94, 0x04, 0xd4, 0x1e, 0x33, 0x87, 0xe2, 0x9a, 0xae, 0xb4, 0x55, 0x73,
	0x19, 0x42, 0x4f, 0x41, 0xf3, 0x03, 0x12, 0x84, 0x3e, 0xae, 0xeb, 0x4a, 0xbb, 0x59, 0xd4, 0x53,
	0x3e, 0x46, 0x31, 0x65, 0xc0, 0x43, 0xc7, 0x94, 0x74, 0xb4, 0x07, 0x35, 0x8f, 0xf1, 0x69, 0x20,
	0x38, 0xc5, 0x10, 0x8f, 0x2b, 0x8b, 0xd1, 0x43, 0x68, 0x15, 0xd5, 0xc2, 0x8d, 0x98, 0xb3, 0x82,
	0x1b, 0xcf, 0xa0, 0x55, 0x7c, 0x07, 0x42, 0xd0, 0x1c, 0x8d, 0x7b, 0xe3, 0x8b, 0xd1, 0xc7, 0xfe,
	0xe0, 0xb4, 0x77, 0x71, 0x3e, 0x6e, 0xfd, 0x95, 0xc3, 0xce, 0x07, 0xe3, 0x41, 0xbf, 0xa5, 0x18,
	0x5f, 0x14, 0xf8, 0x6f, 0xcd, 0xdc, 0x23, 0x37, 0xf9, 0xd4, 0x8b, 0x72, 0xd2, 0x62, 0x69, 0x88,
	0xfe, 0x85, 0x8d, 0x40, 0x7c, 0xa2, 0x5c, 0xba, 0x2c, 0x09, 0xd0, 0x13, 0xd0, 0xe2, 0x45, 0x64,
	0x2c, 0x75, 0x79, 0x3c, 0xc5, 0xd6, 0xe3, 0x88, 0x65, 0x4a, 0xb2, 0xf1, 0x43, 0x81, 0x9d, 0x52,
	0xc6, 0xad, 0x37, 0x70, 0x91, 0x4d, 0x42, 0x8d, 0x27, 0x71, 0xfc, 0xcb, 0x0d, 0x94, 0xa3, 0xab,
	0x73, 0x32, 0x86, 0xb0, 0xff, 0x1b, 0xea, 0x8d, 0xe5, 0xfe, 0x5a, 0x49, 0x67, 0xb5, 0x98, 0x1f,
	0x42, 0x50, 0xe5, 0xc4, 0x49, 0xcf, 0x18, 0xaf, 0xd7, 0x1c, 0x30, 0xf9, 0xde, 0xd5, 0xec, 0x7b,
	0xdf, 0x83, 0x5a, 0xea, 0x06, 0x5c, 0x4d, 0x1c, 0x94, 0xc6, 0x45, 0xe3, 0x6e, 0xac, 0x1a, 0x57,
	0x87, 0x46, 0xe8, 0xda, 0x19, 0x43, 0x4b, 0x18, 0x4b, 0x10, 0x7a, 0x99, 0x09, 0xba, 0x19, 0x0b,
	0x7a, 0xb4, 0xee, 0xdb, 0x5e, 0x01, 0x4a, 0x64, 0xec, 0xc3, 0xde, 0x7a, 0xd6, 0x8d, 0x15, 0x0c,
	0xa1, 0x9e, 0x5d, 0x0f, 0x91, 0x41, 0xae, 0xa8, 0xe7, 0x33, 0xc1, 0x53, 0x83, 0xc8, 0x30, 0xd2,
	0x6f, 0x12, 0xb2, 0x79, 0x7a, 0x0f, 0x26, 0x01, 0xfa, 0x1f, 0xea, 0x21, 0x7b, 0x27, 0x2b, 0x12,
	0x19, 0x17, 0x40, 0xa4, 0xe6, 0x9c, 0xf0, 0x69, 0x48, 0xa6, 0x34, 0x55, 0x33, 0x8d, 0x8d, 0x0f,
	0xb0, 0x53, 0x7a, 0xc9, 0x94, 0x0e, 0x6f, 0x69, 0x5b, 0x95, 0xfc, 0xb6, 0x76, 0x41, 0x9b, 0x0b,
	0x8b, 0xcc, 0xa9, 0x7c, 0xbb, 0x8c, 0x8c, 0x09, 0x34, 0xf3, 0x37, 0x50, 0x69, 0xdf, 0xbb, 0x00,
	0x0e, 0xb1, 0x7a, 0xb6, 0xed, 0x51, 0xdf, 0x97, 0xad, 0x97, 0x90, 0xe8, 0x78, 0xcc, 0x4d, 0xd3,
	0xf2, 0x78, 0x19, 0x60, 0x7c, 0xab, 0x40, 0xfd, 0x24, 0x92, 0x61, 0xc8, 0x2f, 0xc5, 0x2d, 0x7e,
	0x1d, 0xf7, 0xe1, 0xef, 0x49, 0x5a, 0x76, 0x46, 0xfc, 0x99, 0xec, 0x9c, 0x07, 0x73, 0xac, 0x17,
	0xbe, 0xe0, 0x52, 0xc1, 0x3c, 0xf8, 0x47, 0x4c, 0x79, 0x5c, 0x30, 0xe5, 0x83, 0xd4, 0x94, 0xd9,
	0xe1, 0x16, 0xab, 0x12, 0x1b, 0x1e, 0xc3, 0x76, 0x49, 0xfa, 0xa6, 0xfe, 0x3b, 0x79, 0xfc, 0xbe,
	0x3b, 0x65, 0xc1, 0x2c, 0x9c, 0x1c, 0x5a, 0xc2, 0xe9, 0xb8, 0x8f, 0x2c, 0xde, 0x09, 0x08, 0x0f,
	0x08, 0x3f, 0xb0, 0x85, 0x43, 0x18, 0x3f, 0xf0, 0xad, 0x19, 0x75, 0x48, 0x67, 0x2a, 0x22, 0xef,
	0xc8, 0x3f, 0xf8, 0x44, 0x8b, 0x7f, 0xe1, 0x47, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x64, 0xaf,
	0x8a, 0x69, 0xd9, 0x07, 0x00, 0x00,
}
