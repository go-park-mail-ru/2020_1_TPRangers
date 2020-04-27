// Code generated by protoc-gen-go. DO NOT EDIT.
// source: likes.proto

package likes

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Like struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	DataId               int32    `protobuf:"varint,2,opt,name=dataId,proto3" json:"dataId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Like) Reset()         { *m = Like{} }
func (m *Like) String() string { return proto.CompactTextString(m) }
func (*Like) ProtoMessage()    {}
func (*Like) Descriptor() ([]byte, []int) {
	return fileDescriptor_cff81f36f81c8d8e, []int{0}
}

func (m *Like) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Like.Unmarshal(m, b)
}
func (m *Like) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Like.Marshal(b, m, deterministic)
}
func (m *Like) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Like.Merge(m, src)
}
func (m *Like) XXX_Size() int {
	return xxx_messageInfo_Like.Size(m)
}
func (m *Like) XXX_DiscardUnknown() {
	xxx_messageInfo_Like.DiscardUnknown(m)
}

var xxx_messageInfo_Like proto.InternalMessageInfo

func (m *Like) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Like) GetDataId() int32 {
	if m != nil {
		return m.DataId
	}
	return 0
}

type Dummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dummy) Reset()         { *m = Dummy{} }
func (m *Dummy) String() string { return proto.CompactTextString(m) }
func (*Dummy) ProtoMessage()    {}
func (*Dummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_cff81f36f81c8d8e, []int{1}
}

func (m *Dummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dummy.Unmarshal(m, b)
}
func (m *Dummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dummy.Marshal(b, m, deterministic)
}
func (m *Dummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dummy.Merge(m, src)
}
func (m *Dummy) XXX_Size() int {
	return xxx_messageInfo_Dummy.Size(m)
}
func (m *Dummy) XXX_DiscardUnknown() {
	xxx_messageInfo_Dummy.DiscardUnknown(m)
}

var xxx_messageInfo_Dummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Like)(nil), "likes.Like")
	proto.RegisterType((*Dummy)(nil), "likes.Dummy")
}

func init() { proto.RegisterFile("likes.proto", fileDescriptor_cff81f36f81c8d8e) }

var fileDescriptor_cff81f36f81c8d8e = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc9, 0xcc, 0x4e,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xcc, 0xb8, 0x58, 0x7c,
	0x32, 0xb3, 0x53, 0x85, 0xc4, 0xb8, 0xd8, 0x4a, 0x8b, 0x53, 0x8b, 0x3c, 0x53, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c, 0x90, 0x78, 0x4a, 0x62, 0x49, 0xa2, 0x67, 0x8a, 0x04, 0x13,
	0x44, 0x1c, 0xc2, 0x53, 0x62, 0xe7, 0x62, 0x75, 0x29, 0xcd, 0xcd, 0xad, 0x34, 0xda, 0xc8, 0xc8,
	0xc5, 0x0d, 0x32, 0xc1, 0x39, 0x23, 0x35, 0x39, 0x3b, 0xb5, 0x48, 0x48, 0x8d, 0x8b, 0x13, 0xc4,
	0x0d, 0xc8, 0x00, 0x59, 0xc2, 0xad, 0x07, 0xb1, 0x12, 0x24, 0x22, 0xc5, 0x03, 0xe5, 0x80, 0xf5,
	0x09, 0x69, 0x72, 0xf1, 0xb8, 0x64, 0x16, 0xe7, 0x10, 0xa3, 0x54, 0x95, 0x8b, 0x03, 0x6c, 0x64,
	0x7e, 0x71, 0x09, 0x3e, 0x65, 0x1a, 0x5c, 0xdc, 0x30, 0x13, 0xf1, 0xab, 0x4c, 0x62, 0x03, 0x07,
	0x81, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x01, 0x72, 0x5a, 0xa2, 0x11, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LikeCheckerClient is the client API for LikeChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LikeCheckerClient interface {
	LikePhoto(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Dummy, error)
	DislikePhoto(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Dummy, error)
	LikePost(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Dummy, error)
	DislikePost(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Dummy, error)
}

type likeCheckerClient struct {
	cc *grpc.ClientConn
}

func NewLikeCheckerClient(cc *grpc.ClientConn) LikeCheckerClient {
	return &likeCheckerClient{cc}
}

func (c *likeCheckerClient) LikePhoto(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Dummy, error) {
	out := new(Dummy)
	err := c.cc.Invoke(ctx, "/likes.LikeChecker/LikePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeCheckerClient) DislikePhoto(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Dummy, error) {
	out := new(Dummy)
	err := c.cc.Invoke(ctx, "/likes.LikeChecker/DislikePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeCheckerClient) LikePost(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Dummy, error) {
	out := new(Dummy)
	err := c.cc.Invoke(ctx, "/likes.LikeChecker/LikePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeCheckerClient) DislikePost(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Dummy, error) {
	out := new(Dummy)
	err := c.cc.Invoke(ctx, "/likes.LikeChecker/DislikePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikeCheckerServer is the server API for LikeChecker service.
type LikeCheckerServer interface {
	LikePhoto(context.Context, *Like) (*Dummy, error)
	DislikePhoto(context.Context, *Like) (*Dummy, error)
	LikePost(context.Context, *Like) (*Dummy, error)
	DislikePost(context.Context, *Like) (*Dummy, error)
}

// UnimplementedLikeCheckerServer can be embedded to have forward compatible implementations.
type UnimplementedLikeCheckerServer struct {
}

func (*UnimplementedLikeCheckerServer) LikePhoto(ctx context.Context, req *Like) (*Dummy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikePhoto not implemented")
}
func (*UnimplementedLikeCheckerServer) DislikePhoto(ctx context.Context, req *Like) (*Dummy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DislikePhoto not implemented")
}
func (*UnimplementedLikeCheckerServer) LikePost(ctx context.Context, req *Like) (*Dummy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikePost not implemented")
}
func (*UnimplementedLikeCheckerServer) DislikePost(ctx context.Context, req *Like) (*Dummy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DislikePost not implemented")
}

func RegisterLikeCheckerServer(s *grpc.Server, srv LikeCheckerServer) {
	s.RegisterService(&_LikeChecker_serviceDesc, srv)
}

func _LikeChecker_LikePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Like)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeCheckerServer).LikePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/likes.LikeChecker/LikePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeCheckerServer).LikePhoto(ctx, req.(*Like))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeChecker_DislikePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Like)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeCheckerServer).DislikePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/likes.LikeChecker/DislikePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeCheckerServer).DislikePhoto(ctx, req.(*Like))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeChecker_LikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Like)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeCheckerServer).LikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/likes.LikeChecker/LikePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeCheckerServer).LikePost(ctx, req.(*Like))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeChecker_DislikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Like)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeCheckerServer).DislikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/likes.LikeChecker/DislikePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeCheckerServer).DislikePost(ctx, req.(*Like))
	}
	return interceptor(ctx, in, info, handler)
}

var _LikeChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "likes.LikeChecker",
	HandlerType: (*LikeCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LikePhoto",
			Handler:    _LikeChecker_LikePhoto_Handler,
		},
		{
			MethodName: "DislikePhoto",
			Handler:    _LikeChecker_DislikePhoto_Handler,
		},
		{
			MethodName: "LikePost",
			Handler:    _LikeChecker_LikePost_Handler,
		},
		{
			MethodName: "DislikePost",
			Handler:    _LikeChecker_DislikePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "likes.proto",
}
