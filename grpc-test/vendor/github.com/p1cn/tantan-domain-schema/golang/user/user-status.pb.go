// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user-status.proto

package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserStatus struct {
	UserID string         `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	Status UserStatusEnum `protobuf:"varint,2,opt,name=status,enum=user.UserStatusEnum" json:"status,omitempty"`
}

func (m *UserStatus) Reset()                    { *m = UserStatus{} }
func (m *UserStatus) String() string            { return proto.CompactTextString(m) }
func (*UserStatus) ProtoMessage()               {}
func (*UserStatus) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *UserStatus) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserStatus) GetStatus() UserStatusEnum {
	if m != nil {
		return m.Status
	}
	return UserStatusEnum_default
}

func init() {
	proto.RegisterType((*UserStatus)(nil), "user.UserStatus")
}

func init() { proto.RegisterFile("user/user-status.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x2d, 0x4e, 0x2d,
	0xd2, 0x07, 0x11, 0xba, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x2c, 0x20, 0x21, 0x29, 0x41, 0xb0, 0x6c, 0x72, 0x7e, 0x6e, 0x6e, 0x7e, 0x1e, 0x44, 0x42,
	0x29, 0x88, 0x8b, 0x2b, 0xb4, 0x38, 0xb5, 0x28, 0x18, 0xac, 0x58, 0x48, 0x8c, 0x8b, 0x0d, 0xa4,
	0xc4, 0xd3, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xca, 0x13, 0xd2, 0xe1, 0x62, 0x83,
	0x18, 0x27, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x67, 0x24, 0xa2, 0x07, 0x92, 0xd0, 0x43, 0xe8, 0x74,
	0xcd, 0x2b, 0xcd, 0x0d, 0x82, 0xaa, 0x71, 0x32, 0x8a, 0x32, 0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x30, 0x4c, 0xce, 0xd3, 0x2f, 0x49, 0xcc, 0x2b, 0x49, 0xcc,
	0xd3, 0x4d, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0xd3, 0x2d, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x4f,
	0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0x07, 0xbb, 0x36, 0x89, 0x0d, 0xec, 0x1c, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x0f, 0xff, 0x60, 0xea, 0xc1, 0x00, 0x00, 0x00,
}