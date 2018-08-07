// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moment/moment-activity.service.proto

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

type MomentActivityParam struct {
	Activity []*MomentActivity `protobuf:"bytes,1,rep,name=activity" json:"activity,omitempty"`
}

func (m *MomentActivityParam) Reset()                    { *m = MomentActivityParam{} }
func (m *MomentActivityParam) String() string            { return proto.CompactTextString(m) }
func (*MomentActivityParam) ProtoMessage()               {}
func (*MomentActivityParam) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MomentActivityParam) GetActivity() []*MomentActivity {
	if m != nil {
		return m.Activity
	}
	return nil
}

type MomentActivityReply struct {
	Activity []*MomentActivity `protobuf:"bytes,1,rep,name=activity" json:"activity,omitempty"`
}

func (m *MomentActivityReply) Reset()                    { *m = MomentActivityReply{} }
func (m *MomentActivityReply) String() string            { return proto.CompactTextString(m) }
func (*MomentActivityReply) ProtoMessage()               {}
func (*MomentActivityReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *MomentActivityReply) GetActivity() []*MomentActivity {
	if m != nil {
		return m.Activity
	}
	return nil
}

type DeleteByActionParam struct {
	UserId      string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	MomentId    string `protobuf:"bytes,2,opt,name=momentId" json:"momentId,omitempty"`
	ActorUserId string `protobuf:"bytes,3,opt,name=actorUserId" json:"actorUserId,omitempty"`
	Action      string `protobuf:"bytes,4,opt,name=action" json:"action,omitempty"`
	MessageId   string `protobuf:"bytes,5,opt,name=messageId" json:"messageId,omitempty"`
}

func (m *DeleteByActionParam) Reset()                    { *m = DeleteByActionParam{} }
func (m *DeleteByActionParam) String() string            { return proto.CompactTextString(m) }
func (*DeleteByActionParam) ProtoMessage()               {}
func (*DeleteByActionParam) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *DeleteByActionParam) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *DeleteByActionParam) GetMomentId() string {
	if m != nil {
		return m.MomentId
	}
	return ""
}

func (m *DeleteByActionParam) GetActorUserId() string {
	if m != nil {
		return m.ActorUserId
	}
	return ""
}

func (m *DeleteByActionParam) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *DeleteByActionParam) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

type DeleteByMomentIdParam struct {
	UserId   string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	MomentId string `protobuf:"bytes,2,opt,name=momentId" json:"momentId,omitempty"`
}

func (m *DeleteByMomentIdParam) Reset()                    { *m = DeleteByMomentIdParam{} }
func (m *DeleteByMomentIdParam) String() string            { return proto.CompactTextString(m) }
func (*DeleteByMomentIdParam) ProtoMessage()               {}
func (*DeleteByMomentIdParam) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *DeleteByMomentIdParam) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *DeleteByMomentIdParam) GetMomentId() string {
	if m != nil {
		return m.MomentId
	}
	return ""
}

func init() {
	proto.RegisterType((*MomentActivityParam)(nil), "moment.MomentActivityParam")
	proto.RegisterType((*MomentActivityReply)(nil), "moment.MomentActivityReply")
	proto.RegisterType((*DeleteByActionParam)(nil), "moment.DeleteByActionParam")
	proto.RegisterType((*DeleteByMomentIdParam)(nil), "moment.DeleteByMomentIdParam")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MomentActivityService service

type MomentActivityServiceClient interface {
	InsertActivity(ctx context.Context, in *MomentActivityParam, opts ...grpc.CallOption) (*MomentActivityReply, error)
	FindByUserId(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*MomentActivityReply, error)
	UpdateAsReadByUserId(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*common.BoolValue, error)
	DeleteByUserId(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*common.BoolValue, error)
	DeleteByAction(ctx context.Context, in *DeleteByActionParam, opts ...grpc.CallOption) (*common.BoolValue, error)
	DeleteByMomentId(ctx context.Context, in *DeleteByMomentIdParam, opts ...grpc.CallOption) (*common.BoolValue, error)
	CountUnread(ctx context.Context, in *common.StringValue, opts ...grpc.CallOption) (*common.Int32Value, error)
	UpdateCointers(ctx context.Context, in *common.StringValueArray, opts ...grpc.CallOption) (*common.BoolValue, error)
}

type momentActivityServiceClient struct {
	cc *grpc.ClientConn
}

func NewMomentActivityServiceClient(cc *grpc.ClientConn) MomentActivityServiceClient {
	return &momentActivityServiceClient{cc}
}

func (c *momentActivityServiceClient) InsertActivity(ctx context.Context, in *MomentActivityParam, opts ...grpc.CallOption) (*MomentActivityReply, error) {
	out := new(MomentActivityReply)
	err := grpc.Invoke(ctx, "/moment.MomentActivityService/InsertActivity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentActivityServiceClient) FindByUserId(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*MomentActivityReply, error) {
	out := new(MomentActivityReply)
	err := grpc.Invoke(ctx, "/moment.MomentActivityService/FindByUserId", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentActivityServiceClient) UpdateAsReadByUserId(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*common.BoolValue, error) {
	out := new(common.BoolValue)
	err := grpc.Invoke(ctx, "/moment.MomentActivityService/UpdateAsReadByUserId", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentActivityServiceClient) DeleteByUserId(ctx context.Context, in *common.Condition, opts ...grpc.CallOption) (*common.BoolValue, error) {
	out := new(common.BoolValue)
	err := grpc.Invoke(ctx, "/moment.MomentActivityService/DeleteByUserId", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentActivityServiceClient) DeleteByAction(ctx context.Context, in *DeleteByActionParam, opts ...grpc.CallOption) (*common.BoolValue, error) {
	out := new(common.BoolValue)
	err := grpc.Invoke(ctx, "/moment.MomentActivityService/DeleteByAction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentActivityServiceClient) DeleteByMomentId(ctx context.Context, in *DeleteByMomentIdParam, opts ...grpc.CallOption) (*common.BoolValue, error) {
	out := new(common.BoolValue)
	err := grpc.Invoke(ctx, "/moment.MomentActivityService/DeleteByMomentId", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentActivityServiceClient) CountUnread(ctx context.Context, in *common.StringValue, opts ...grpc.CallOption) (*common.Int32Value, error) {
	out := new(common.Int32Value)
	err := grpc.Invoke(ctx, "/moment.MomentActivityService/CountUnread", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentActivityServiceClient) UpdateCointers(ctx context.Context, in *common.StringValueArray, opts ...grpc.CallOption) (*common.BoolValue, error) {
	out := new(common.BoolValue)
	err := grpc.Invoke(ctx, "/moment.MomentActivityService/UpdateCointers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MomentActivityService service

type MomentActivityServiceServer interface {
	InsertActivity(context.Context, *MomentActivityParam) (*MomentActivityReply, error)
	FindByUserId(context.Context, *common.Condition) (*MomentActivityReply, error)
	UpdateAsReadByUserId(context.Context, *common.Condition) (*common.BoolValue, error)
	DeleteByUserId(context.Context, *common.Condition) (*common.BoolValue, error)
	DeleteByAction(context.Context, *DeleteByActionParam) (*common.BoolValue, error)
	DeleteByMomentId(context.Context, *DeleteByMomentIdParam) (*common.BoolValue, error)
	CountUnread(context.Context, *common.StringValue) (*common.Int32Value, error)
	UpdateCointers(context.Context, *common.StringValueArray) (*common.BoolValue, error)
}

func RegisterMomentActivityServiceServer(s *grpc.Server, srv MomentActivityServiceServer) {
	s.RegisterService(&_MomentActivityService_serviceDesc, srv)
}

func _MomentActivityService_InsertActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MomentActivityParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentActivityServiceServer).InsertActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentActivityService/InsertActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentActivityServiceServer).InsertActivity(ctx, req.(*MomentActivityParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentActivityService_FindByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Condition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentActivityServiceServer).FindByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentActivityService/FindByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentActivityServiceServer).FindByUserId(ctx, req.(*common.Condition))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentActivityService_UpdateAsReadByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Condition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentActivityServiceServer).UpdateAsReadByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentActivityService/UpdateAsReadByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentActivityServiceServer).UpdateAsReadByUserId(ctx, req.(*common.Condition))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentActivityService_DeleteByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Condition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentActivityServiceServer).DeleteByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentActivityService/DeleteByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentActivityServiceServer).DeleteByUserId(ctx, req.(*common.Condition))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentActivityService_DeleteByAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteByActionParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentActivityServiceServer).DeleteByAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentActivityService/DeleteByAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentActivityServiceServer).DeleteByAction(ctx, req.(*DeleteByActionParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentActivityService_DeleteByMomentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteByMomentIdParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentActivityServiceServer).DeleteByMomentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentActivityService/DeleteByMomentId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentActivityServiceServer).DeleteByMomentId(ctx, req.(*DeleteByMomentIdParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentActivityService_CountUnread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentActivityServiceServer).CountUnread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentActivityService/CountUnread",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentActivityServiceServer).CountUnread(ctx, req.(*common.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentActivityService_UpdateCointers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.StringValueArray)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentActivityServiceServer).UpdateCointers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentActivityService/UpdateCointers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentActivityServiceServer).UpdateCointers(ctx, req.(*common.StringValueArray))
	}
	return interceptor(ctx, in, info, handler)
}

var _MomentActivityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moment.MomentActivityService",
	HandlerType: (*MomentActivityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertActivity",
			Handler:    _MomentActivityService_InsertActivity_Handler,
		},
		{
			MethodName: "FindByUserId",
			Handler:    _MomentActivityService_FindByUserId_Handler,
		},
		{
			MethodName: "UpdateAsReadByUserId",
			Handler:    _MomentActivityService_UpdateAsReadByUserId_Handler,
		},
		{
			MethodName: "DeleteByUserId",
			Handler:    _MomentActivityService_DeleteByUserId_Handler,
		},
		{
			MethodName: "DeleteByAction",
			Handler:    _MomentActivityService_DeleteByAction_Handler,
		},
		{
			MethodName: "DeleteByMomentId",
			Handler:    _MomentActivityService_DeleteByMomentId_Handler,
		},
		{
			MethodName: "CountUnread",
			Handler:    _MomentActivityService_CountUnread_Handler,
		},
		{
			MethodName: "UpdateCointers",
			Handler:    _MomentActivityService_UpdateCointers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moment/moment-activity.service.proto",
}

func init() { proto.RegisterFile("moment/moment-activity.service.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x5f, 0x6b, 0xd4, 0x40,
	0x10, 0xe7, 0x6c, 0x3d, 0xda, 0x3d, 0x39, 0x74, 0xcf, 0x2b, 0x21, 0x56, 0x38, 0x8a, 0x0f, 0x7d,
	0xb9, 0x04, 0x53, 0xf1, 0x49, 0xc5, 0xbb, 0x2b, 0x42, 0x94, 0x82, 0x5c, 0x39, 0x1f, 0x7c, 0x9b,
	0x26, 0x43, 0xba, 0x90, 0xcc, 0x86, 0xdd, 0xbd, 0x4a, 0xbe, 0x89, 0xaf, 0x7e, 0x53, 0xb9, 0xdd,
	0x24, 0xe7, 0xd5, 0xb4, 0xd0, 0x3e, 0x85, 0x99, 0xdf, 0x9f, 0x99, 0xcc, 0xcc, 0xb2, 0x37, 0x85,
	0x2c, 0x90, 0x4c, 0xe8, 0x3e, 0x53, 0x48, 0x8c, 0xb8, 0x11, 0xa6, 0x0a, 0x34, 0xaa, 0x1b, 0x91,
	0x60, 0x50, 0x2a, 0x69, 0x24, 0xef, 0x3b, 0xd8, 0x3f, 0xbe, 0x83, 0x6d, 0x59, 0xfe, 0x38, 0x91,
	0x45, 0x21, 0x29, 0xfc, 0xa5, 0xa0, 0x2c, 0x51, 0x69, 0x97, 0x3e, 0x89, 0xd9, 0xe8, 0xc2, 0xf2,
	0x67, 0x35, 0xfd, 0x3b, 0x28, 0x28, 0x78, 0xc4, 0x0e, 0x1a, 0xbd, 0xd7, 0x9b, 0xec, 0x9d, 0x0e,
	0xa2, 0xa3, 0xc0, 0xf9, 0x06, 0xbb, 0xf4, 0x65, 0xcb, 0xfb, 0xdf, 0x6a, 0x89, 0x65, 0x5e, 0x3d,
	0xca, 0xea, 0x4f, 0x8f, 0x8d, 0xce, 0x31, 0x47, 0x83, 0xf3, 0x6a, 0x03, 0x4b, 0x72, 0x6d, 0x1d,
	0xb1, 0xfe, 0x5a, 0xa3, 0x8a, 0x53, 0xaf, 0x37, 0xe9, 0x9d, 0x1e, 0x2e, 0xeb, 0x88, 0xfb, 0xec,
	0xc0, 0x59, 0xc6, 0xa9, 0xf7, 0xc4, 0x22, 0x6d, 0xcc, 0x27, 0x6c, 0x00, 0x89, 0x91, 0x6a, 0xe5,
	0x84, 0x7b, 0x16, 0xfe, 0x37, 0xb5, 0x71, 0x05, 0x5b, 0xc4, 0xdb, 0x77, 0xae, 0x2e, 0xe2, 0xc7,
	0xec, 0xb0, 0x40, 0xad, 0x21, 0xc3, 0x38, 0xf5, 0x9e, 0x5a, 0x68, 0x9b, 0x38, 0xf9, 0xc6, 0xc6,
	0x4d, 0x8b, 0x17, 0x75, 0xad, 0x47, 0x37, 0x19, 0xfd, 0xde, 0x67, 0xe3, 0xdd, 0x69, 0x5c, 0xba,
	0x1d, 0xf3, 0xaf, 0x6c, 0x18, 0x93, 0x46, 0xd5, 0x02, 0xfc, 0x55, 0xf7, 0xf8, 0x6c, 0x71, 0xff,
	0x0e, 0xd0, 0xad, 0xe2, 0x13, 0x7b, 0xf6, 0x45, 0x50, 0x3a, 0xaf, 0xea, 0x1f, 0x7f, 0x11, 0xb8,
	0xa3, 0x08, 0x16, 0x92, 0x52, 0xb1, 0xf9, 0xe7, 0xfb, 0xf5, 0x1f, 0xd8, 0xcb, 0x55, 0x99, 0x82,
	0xc1, 0x99, 0x5e, 0x22, 0xdc, 0xeb, 0xd3, 0xa6, 0xe6, 0x52, 0xe6, 0x3f, 0x20, 0x5f, 0x23, 0x7f,
	0xcf, 0x86, 0xcd, 0xc0, 0x1e, 0xa4, 0xfb, 0xbc, 0xd5, 0xb9, 0x5b, 0xd8, 0x4e, 0xa0, 0xe3, 0x46,
	0xba, 0x1c, 0xce, 0xd9, 0xf3, 0xdb, 0xab, 0xe2, 0xaf, 0x6f, 0x7b, 0xec, 0x2c, 0xb1, 0xbb, 0xff,
	0xc1, 0x42, 0xae, 0xc9, 0xac, 0x48, 0x21, 0xa4, 0x7c, 0xd4, 0x30, 0x2e, 0x8d, 0x12, 0x94, 0x59,
	0x8e, 0xcf, 0x9b, 0x64, 0x4c, 0xe6, 0x2c, 0x72, 0xba, 0x8f, 0x6c, 0xe8, 0xa6, 0xb6, 0x90, 0x82,
	0x0c, 0x2a, 0xcd, 0xbd, 0x0e, 0xe9, 0x4c, 0x29, 0xa8, 0x3a, 0xca, 0xce, 0xdf, 0xfd, 0x8c, 0x32,
	0x61, 0xae, 0xd7, 0x57, 0x1b, 0x28, 0x2c, 0xdf, 0x26, 0x14, 0x1a, 0x20, 0x03, 0x34, 0x4d, 0x65,
	0x01, 0x82, 0xa6, 0x3a, 0xb9, 0xc6, 0x02, 0xc2, 0x4c, 0xe6, 0x40, 0x59, 0xfd, 0xfa, 0xaf, 0xfa,
	0xf6, 0x79, 0x9f, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xa0, 0xb6, 0xc5, 0x43, 0x04, 0x00,
	0x00,
}