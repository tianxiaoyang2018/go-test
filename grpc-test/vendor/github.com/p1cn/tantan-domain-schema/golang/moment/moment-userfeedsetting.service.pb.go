// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moment/moment-userfeedsetting.service.proto

package moment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/p1cn/tantan-domain-schema/golang/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserFeedSettingParam struct {
	Setting []*MomentUserFeedSetting `protobuf:"bytes,1,rep,name=setting" json:"setting,omitempty"`
}

func (m *UserFeedSettingParam) Reset()                    { *m = UserFeedSettingParam{} }
func (m *UserFeedSettingParam) String() string            { return proto.CompactTextString(m) }
func (*UserFeedSettingParam) ProtoMessage()               {}
func (*UserFeedSettingParam) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *UserFeedSettingParam) GetSetting() []*MomentUserFeedSetting {
	if m != nil {
		return m.Setting
	}
	return nil
}

type UserFeedSettingReply struct {
	Setting []*MomentUserFeedSetting `protobuf:"bytes,1,rep,name=setting" json:"setting,omitempty"`
}

func (m *UserFeedSettingReply) Reset()                    { *m = UserFeedSettingReply{} }
func (m *UserFeedSettingReply) String() string            { return proto.CompactTextString(m) }
func (*UserFeedSettingReply) ProtoMessage()               {}
func (*UserFeedSettingReply) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *UserFeedSettingReply) GetSetting() []*MomentUserFeedSetting {
	if m != nil {
		return m.Setting
	}
	return nil
}

func init() {
	proto.RegisterType((*UserFeedSettingParam)(nil), "moment.UserFeedSettingParam")
	proto.RegisterType((*UserFeedSettingReply)(nil), "moment.UserFeedSettingReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MomentUserFeedSettingService service

type MomentUserFeedSettingServiceClient interface {
	FindByIds(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*UserFeedSettingReply, error)
	FindMuted(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*UserFeedSettingReply, error)
	UpsertMuted(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*UserFeedSettingReply, error)
	DeleteMuted(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*common.BoolValue, error)
	FindHidden(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*UserFeedSettingReply, error)
	UpsertHidden(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*UserFeedSettingReply, error)
	DeleteHidden(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*common.BoolValue, error)
	DeleteUserFeedSetting(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*common.BoolValue, error)
}

type momentUserFeedSettingServiceClient struct {
	cc *grpc.ClientConn
}

func NewMomentUserFeedSettingServiceClient(cc *grpc.ClientConn) MomentUserFeedSettingServiceClient {
	return &momentUserFeedSettingServiceClient{cc}
}

func (c *momentUserFeedSettingServiceClient) FindByIds(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*UserFeedSettingReply, error) {
	out := new(UserFeedSettingReply)
	err := grpc.Invoke(ctx, "/moment.MomentUserFeedSettingService/FindByIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentUserFeedSettingServiceClient) FindMuted(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*UserFeedSettingReply, error) {
	out := new(UserFeedSettingReply)
	err := grpc.Invoke(ctx, "/moment.MomentUserFeedSettingService/FindMuted", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentUserFeedSettingServiceClient) UpsertMuted(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*UserFeedSettingReply, error) {
	out := new(UserFeedSettingReply)
	err := grpc.Invoke(ctx, "/moment.MomentUserFeedSettingService/UpsertMuted", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentUserFeedSettingServiceClient) DeleteMuted(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*common.BoolValue, error) {
	out := new(common.BoolValue)
	err := grpc.Invoke(ctx, "/moment.MomentUserFeedSettingService/DeleteMuted", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentUserFeedSettingServiceClient) FindHidden(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*UserFeedSettingReply, error) {
	out := new(UserFeedSettingReply)
	err := grpc.Invoke(ctx, "/moment.MomentUserFeedSettingService/FindHidden", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentUserFeedSettingServiceClient) UpsertHidden(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*UserFeedSettingReply, error) {
	out := new(UserFeedSettingReply)
	err := grpc.Invoke(ctx, "/moment.MomentUserFeedSettingService/UpsertHidden", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentUserFeedSettingServiceClient) DeleteHidden(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*common.BoolValue, error) {
	out := new(common.BoolValue)
	err := grpc.Invoke(ctx, "/moment.MomentUserFeedSettingService/DeleteHidden", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentUserFeedSettingServiceClient) DeleteUserFeedSetting(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*common.BoolValue, error) {
	out := new(common.BoolValue)
	err := grpc.Invoke(ctx, "/moment.MomentUserFeedSettingService/DeleteUserFeedSetting", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MomentUserFeedSettingService service

type MomentUserFeedSettingServiceServer interface {
	FindByIds(context.Context, *common.Condition) (*UserFeedSettingReply, error)
	FindMuted(context.Context, *common.Condition) (*UserFeedSettingReply, error)
	UpsertMuted(context.Context, *common.Condition) (*UserFeedSettingReply, error)
	DeleteMuted(context.Context, *common.Condition) (*common.BoolValue, error)
	FindHidden(context.Context, *common.Condition) (*UserFeedSettingReply, error)
	UpsertHidden(context.Context, *common.Condition) (*UserFeedSettingReply, error)
	DeleteHidden(context.Context, *common.Condition) (*common.BoolValue, error)
	DeleteUserFeedSetting(context.Context, *common.Condition) (*common.BoolValue, error)
}

func RegisterMomentUserFeedSettingServiceServer(s *grpc.Server, srv MomentUserFeedSettingServiceServer) {
	s.RegisterService(&_MomentUserFeedSettingService_serviceDesc, srv)
}

func _MomentUserFeedSettingService_FindByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Condition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentUserFeedSettingServiceServer).FindByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentUserFeedSettingService/FindByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentUserFeedSettingServiceServer).FindByIds(ctx, req.(*common.Condition))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentUserFeedSettingService_FindMuted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Condition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentUserFeedSettingServiceServer).FindMuted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentUserFeedSettingService/FindMuted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentUserFeedSettingServiceServer).FindMuted(ctx, req.(*common.Condition))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentUserFeedSettingService_UpsertMuted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Condition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentUserFeedSettingServiceServer).UpsertMuted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentUserFeedSettingService/UpsertMuted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentUserFeedSettingServiceServer).UpsertMuted(ctx, req.(*common.Condition))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentUserFeedSettingService_DeleteMuted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Condition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentUserFeedSettingServiceServer).DeleteMuted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentUserFeedSettingService/DeleteMuted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentUserFeedSettingServiceServer).DeleteMuted(ctx, req.(*common.Condition))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentUserFeedSettingService_FindHidden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Condition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentUserFeedSettingServiceServer).FindHidden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentUserFeedSettingService/FindHidden",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentUserFeedSettingServiceServer).FindHidden(ctx, req.(*common.Condition))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentUserFeedSettingService_UpsertHidden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Condition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentUserFeedSettingServiceServer).UpsertHidden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentUserFeedSettingService/UpsertHidden",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentUserFeedSettingServiceServer).UpsertHidden(ctx, req.(*common.Condition))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentUserFeedSettingService_DeleteHidden_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Condition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentUserFeedSettingServiceServer).DeleteHidden(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentUserFeedSettingService/DeleteHidden",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentUserFeedSettingServiceServer).DeleteHidden(ctx, req.(*common.Condition))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentUserFeedSettingService_DeleteUserFeedSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Condition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentUserFeedSettingServiceServer).DeleteUserFeedSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentUserFeedSettingService/DeleteUserFeedSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentUserFeedSettingServiceServer).DeleteUserFeedSetting(ctx, req.(*common.Condition))
	}
	return interceptor(ctx, in, info, handler)
}

var _MomentUserFeedSettingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moment.MomentUserFeedSettingService",
	HandlerType: (*MomentUserFeedSettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByIds",
			Handler:    _MomentUserFeedSettingService_FindByIds_Handler,
		},
		{
			MethodName: "FindMuted",
			Handler:    _MomentUserFeedSettingService_FindMuted_Handler,
		},
		{
			MethodName: "UpsertMuted",
			Handler:    _MomentUserFeedSettingService_UpsertMuted_Handler,
		},
		{
			MethodName: "DeleteMuted",
			Handler:    _MomentUserFeedSettingService_DeleteMuted_Handler,
		},
		{
			MethodName: "FindHidden",
			Handler:    _MomentUserFeedSettingService_FindHidden_Handler,
		},
		{
			MethodName: "UpsertHidden",
			Handler:    _MomentUserFeedSettingService_UpsertHidden_Handler,
		},
		{
			MethodName: "DeleteHidden",
			Handler:    _MomentUserFeedSettingService_DeleteHidden_Handler,
		},
		{
			MethodName: "DeleteUserFeedSetting",
			Handler:    _MomentUserFeedSettingService_DeleteUserFeedSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moment/moment-userfeedsetting.service.proto",
}

func init() { proto.RegisterFile("moment/moment-userfeedsetting.service.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x11, 0xa5, 0x62, 0xda, 0x8b, 0x8b, 0x05, 0x29, 0x15, 0x44, 0x2f, 0x82, 0x34, 0x8b,
	0xdb, 0x82, 0x17, 0xff, 0xc0, 0x2a, 0x45, 0x0f, 0x45, 0x69, 0xa9, 0x07, 0x6f, 0xe9, 0x66, 0xdc,
	0x06, 0x36, 0x93, 0x90, 0x64, 0x95, 0xde, 0xfd, 0xe0, 0xd2, 0xcd, 0xd6, 0x43, 0xd9, 0x8a, 0xdd,
	0xd3, 0xc0, 0xe4, 0xfd, 0xde, 0xbc, 0x84, 0x0c, 0xb9, 0x94, 0x4a, 0x02, 0xba, 0xd0, 0x97, 0x5e,
	0x6e, 0xc1, 0x7c, 0x00, 0x70, 0x0b, 0xce, 0x09, 0x4c, 0xa9, 0x05, 0xf3, 0x29, 0x12, 0xa0, 0xda,
	0x28, 0xa7, 0x82, 0x86, 0x57, 0x75, 0xce, 0xff, 0x86, 0x0a, 0x71, 0xa7, 0x9d, 0x28, 0x29, 0x15,
	0x86, 0x5f, 0x86, 0x69, 0x0d, 0xc6, 0xfa, 0xf6, 0xd9, 0x0b, 0x39, 0x9a, 0x5a, 0x30, 0x43, 0x00,
	0x3e, 0xf1, 0xfa, 0x57, 0x66, 0x98, 0x0c, 0xae, 0xc9, 0x7e, 0xc9, 0x1f, 0xef, 0x9c, 0xee, 0x5e,
	0x34, 0xa3, 0x13, 0xea, 0xed, 0xe9, 0xa8, 0x28, 0x6b, 0xd0, 0x78, 0xa5, 0xae, 0x30, 0x1c, 0x83,
	0xce, 0x16, 0xb5, 0x0d, 0xa3, 0xef, 0x3d, 0xd2, 0xad, 0x94, 0x4c, 0xfc, 0x63, 0x04, 0x37, 0xe4,
	0x60, 0x28, 0x90, 0xc7, 0x8b, 0x67, 0x6e, 0x83, 0x43, 0xea, 0xef, 0x49, 0x1f, 0x14, 0x72, 0xe1,
	0x84, 0xc2, 0x4e, 0x77, 0x35, 0xa8, 0x32, 0x57, 0x49, 0x8f, 0x72, 0x07, 0x7c, 0x7b, 0xfa, 0x8e,
	0x34, 0xa7, 0xda, 0x82, 0x71, 0x35, 0xf9, 0x3e, 0x69, 0x3e, 0x42, 0x06, 0x0e, 0x36, 0xf2, 0xbf,
	0xad, 0x58, 0xa9, 0xec, 0x8d, 0x65, 0x39, 0x04, 0xb7, 0x84, 0x2c, 0x23, 0x3f, 0x09, 0xce, 0x01,
	0xb7, 0x9f, 0x79, 0x4f, 0x5a, 0x3e, 0x73, 0x5d, 0x83, 0x01, 0x69, 0xf9, 0xd0, 0x9b, 0x0d, 0x2a,
	0x53, 0xb7, 0x3d, 0xb5, 0xe6, 0xf9, 0x3f, 0x3c, 0x1e, 0xbc, 0x47, 0xa9, 0x70, 0xf3, 0x7c, 0xb6,
	0x3c, 0x0a, 0xf5, 0x55, 0x82, 0xa1, 0x63, 0xe8, 0x18, 0xf6, 0xb8, 0x92, 0x4c, 0x60, 0xcf, 0x26,
	0x73, 0x90, 0x2c, 0x4c, 0x55, 0xc6, 0x30, 0x2d, 0x77, 0x61, 0xd6, 0x28, 0x7e, 0x79, 0xff, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x45, 0x8e, 0xea, 0x0a, 0x58, 0x03, 0x00, 0x00,
}
