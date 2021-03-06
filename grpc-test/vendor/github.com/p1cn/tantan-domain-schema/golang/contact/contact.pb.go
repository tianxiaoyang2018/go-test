// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contact/contact.proto

/*
Package contact is a generated protocol buffer package.

It is generated from these files:
	contact/contact.proto
	contact/contact.service.proto

It has these top-level messages:
	Contact
	GetBidirectionSecretCrushUserIdsRequest
	GetBidirectionSecretCrushUserIdsReply
	UncrushCrushesRequest
	UncrushCrushesReply
	UpsertRequest
	UpsertReply
	SelectMobileContactHashesCountRequest
	SelectMobileContactHashesCountReply
	FindSecretCrushesRequest
	FindSecretCrushesReply
	FindAsUserIdsRequest
	FindAsUserIdsReply
	FindInversedAsUserIdsRequest
	FindInversedAsUserIdsReply
	FindByIdsRequest
	FindByIdsReply
	FindMobileHomeLocationRequest
	FindMobileHomeLocationReply
	FindContactUserIdsRequest
	FindContactUserIdsReply
*/
package contact

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import relationship "github.com/p1cn/tantan-domain-schema/golang/relationship"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Contact struct {
	UserId       string                     `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	Name         string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Id           string                     `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Hashes       []string                   `protobuf:"bytes,4,rep,name=hashes" json:"hashes,omitempty"`
	PhoneNumbers []string                   `protobuf:"bytes,5,rep,name=phoneNumbers" json:"phoneNumbers,omitempty"`
	SecretCrush  bool                       `protobuf:"varint,6,opt,name=secretCrush" json:"secretCrush,omitempty"`
	Relationship *relationship.Relationship `protobuf:"bytes,7,opt,name=relationship" json:"relationship,omitempty"`
	OtherUserIds []string                   `protobuf:"bytes,8,rep,name=otherUserIds" json:"otherUserIds,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Contact) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Contact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contact) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Contact) GetHashes() []string {
	if m != nil {
		return m.Hashes
	}
	return nil
}

func (m *Contact) GetPhoneNumbers() []string {
	if m != nil {
		return m.PhoneNumbers
	}
	return nil
}

func (m *Contact) GetSecretCrush() bool {
	if m != nil {
		return m.SecretCrush
	}
	return false
}

func (m *Contact) GetRelationship() *relationship.Relationship {
	if m != nil {
		return m.Relationship
	}
	return nil
}

func (m *Contact) GetOtherUserIds() []string {
	if m != nil {
		return m.OtherUserIds
	}
	return nil
}

func init() {
	proto.RegisterType((*Contact)(nil), "contact.Contact")
}

func init() { proto.RegisterFile("contact/contact.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x86, 0xb1, 0x93, 0xda, 0xa9, 0x12, 0x3a, 0x08, 0x5a, 0x44, 0x96, 0x8a, 0x4c, 0x5e, 0x62,
	0xd1, 0x86, 0xae, 0x1d, 0x9a, 0xa9, 0x4b, 0x07, 0x43, 0x97, 0x6e, 0xb2, 0x7c, 0x58, 0x82, 0x48,
	0x32, 0x92, 0xfc, 0x0c, 0x7d, 0xed, 0x62, 0xd9, 0x05, 0x79, 0xb2, 0xbf, 0x4f, 0xff, 0xc1, 0x7f,
	0x87, 0x1e, 0x85, 0x35, 0x81, 0x8b, 0xc0, 0x96, 0x6f, 0x3d, 0x38, 0x1b, 0x2c, 0x2e, 0x17, 0x3c,
	0x3e, 0x3b, 0xb8, 0xf1, 0xa0, 0xac, 0xf1, 0x52, 0x0d, 0x2c, 0x85, 0x39, 0x79, 0xfa, 0xcd, 0x51,
	0x79, 0x9d, 0xc3, 0xf8, 0x09, 0x15, 0xa3, 0x07, 0xf7, 0xd9, 0x91, 0x8c, 0x66, 0xd5, 0x7d, 0xb3,
	0x10, 0xc6, 0x68, 0x6b, 0xb8, 0x06, 0x92, 0x47, 0x1b, 0xff, 0xf1, 0x03, 0xca, 0x55, 0x47, 0x36,
	0xd1, 0xe4, 0xaa, 0x9b, 0x66, 0x25, 0xf7, 0x12, 0x3c, 0xd9, 0xd2, 0xcd, 0x34, 0x3b, 0x13, 0x3e,
	0xa1, 0xc3, 0x20, 0xad, 0x81, 0xaf, 0x51, 0xb7, 0xe0, 0x3c, 0xb9, 0x8b, 0xaf, 0x2b, 0x87, 0x29,
	0xda, 0x7b, 0x10, 0x0e, 0xc2, 0xd5, 0x8d, 0x5e, 0x92, 0x82, 0x66, 0xd5, 0xae, 0x49, 0x15, 0x7e,
	0x47, 0x87, 0xb4, 0x3b, 0x29, 0x69, 0x56, 0xed, 0x5f, 0x8f, 0xf5, 0x6a, 0xa1, 0x26, 0x81, 0x66,
	0x95, 0x9f, 0x5a, 0xd8, 0x20, 0xc1, 0x7d, 0xc7, 0x85, 0x3c, 0xd9, 0xcd, 0x2d, 0x52, 0xf7, 0xf1,
	0xf6, 0x73, 0xe9, 0x55, 0x90, 0x63, 0x5b, 0x0b, 0xab, 0xd9, 0xf0, 0x22, 0x0c, 0x0b, 0xdc, 0x04,
	0x6e, 0xce, 0x9d, 0xd5, 0x5c, 0x99, 0xb3, 0x17, 0x12, 0x34, 0x67, 0xbd, 0xbd, 0x71, 0xd3, 0xff,
	0x1f, 0xbc, 0x2d, 0xe2, 0x1d, 0x2f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x85, 0x3a, 0x78, 0xa1,
	0x8a, 0x01, 0x00, 0x00,
}
