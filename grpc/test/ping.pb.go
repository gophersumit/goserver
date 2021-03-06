// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ping.proto

package test

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type PingReq struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReq) Reset()         { *m = PingReq{} }
func (m *PingReq) String() string { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()    {}
func (*PingReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d51d96c3ad891f5, []int{0}
}

func (m *PingReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReq.Unmarshal(m, b)
}
func (m *PingReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReq.Marshal(b, m, deterministic)
}
func (m *PingReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReq.Merge(m, src)
}
func (m *PingReq) XXX_Size() int {
	return xxx_messageInfo_PingReq.Size(m)
}
func (m *PingReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReq.DiscardUnknown(m)
}

var xxx_messageInfo_PingReq proto.InternalMessageInfo

func (m *PingReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type PingResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResp) Reset()         { *m = PingResp{} }
func (m *PingResp) String() string { return proto.CompactTextString(m) }
func (*PingResp) ProtoMessage()    {}
func (*PingResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d51d96c3ad891f5, []int{1}
}

func (m *PingResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResp.Unmarshal(m, b)
}
func (m *PingResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResp.Marshal(b, m, deterministic)
}
func (m *PingResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResp.Merge(m, src)
}
func (m *PingResp) XXX_Size() int {
	return xxx_messageInfo_PingResp.Size(m)
}
func (m *PingResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResp.DiscardUnknown(m)
}

var xxx_messageInfo_PingResp proto.InternalMessageInfo

func (m *PingResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ErrorReq struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorReq) Reset()         { *m = ErrorReq{} }
func (m *ErrorReq) String() string { return proto.CompactTextString(m) }
func (*ErrorReq) ProtoMessage()    {}
func (*ErrorReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d51d96c3ad891f5, []int{2}
}

func (m *ErrorReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorReq.Unmarshal(m, b)
}
func (m *ErrorReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorReq.Marshal(b, m, deterministic)
}
func (m *ErrorReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorReq.Merge(m, src)
}
func (m *ErrorReq) XXX_Size() int {
	return xxx_messageInfo_ErrorReq.Size(m)
}
func (m *ErrorReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorReq.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorReq proto.InternalMessageInfo

func (m *ErrorReq) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ErrorReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d51d96c3ad891f5, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PingReq)(nil), "test.PingReq")
	proto.RegisterType((*PingResp)(nil), "test.PingResp")
	proto.RegisterType((*ErrorReq)(nil), "test.ErrorReq")
	proto.RegisterType((*Empty)(nil), "test.Empty")
}

func init() { proto.RegisterFile("ping.proto", fileDescriptor_6d51d96c3ad891f5) }

var fileDescriptor_6d51d96c3ad891f5 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc8, 0xcc, 0x4b,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x49, 0x2d, 0x2e, 0x51, 0x92, 0xe6, 0x62,
	0x0f, 0xc8, 0xcc, 0x4b, 0x0f, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0x64, 0xb8, 0x38, 0x20, 0x92, 0xc5, 0x05, 0x58,
	0x64, 0x0d, 0xb8, 0x38, 0x5c, 0x8b, 0x8a, 0xf2, 0x8b, 0x40, 0x7a, 0x85, 0xb8, 0x58, 0x92, 0xf3,
	0x53, 0x52, 0xc1, 0xd2, 0xbc, 0x41, 0x60, 0x36, 0x4c, 0x07, 0x13, 0x42, 0x07, 0x3b, 0x17, 0xab,
	0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x51, 0x2d, 0xc4, 0xe0, 0x80, 0xfc, 0xbc, 0x74, 0x21, 0x55, 0x2e,
	0x16, 0x10, 0x5b, 0x88, 0x57, 0x0f, 0xe4, 0x20, 0x3d, 0xa8, 0x6b, 0xa4, 0xf8, 0x90, 0xb9, 0xc5,
	0x05, 0x42, 0x6a, 0x5c, 0xac, 0x01, 0x89, 0x79, 0x99, 0xc9, 0x84, 0xd4, 0x29, 0x71, 0x31, 0xbb,
	0x16, 0x15, 0x09, 0x41, 0x85, 0x61, 0x0e, 0x94, 0xe2, 0x86, 0xf2, 0x41, 0xd6, 0x27, 0xb1, 0x81,
	0x43, 0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xea, 0x58, 0x8c, 0x0b, 0x0f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingPongClient is the client API for PingPong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingPongClient interface {
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error)
	Panic(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error)
	Err(ctx context.Context, in *ErrorReq, opts ...grpc.CallOption) (*Empty, error)
}

type pingPongClient struct {
	cc *grpc.ClientConn
}

func NewPingPongClient(cc *grpc.ClientConn) PingPongClient {
	return &pingPongClient{cc}
}

func (c *pingPongClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error) {
	out := new(PingResp)
	err := c.cc.Invoke(ctx, "/test.PingPong/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingPongClient) Panic(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error) {
	out := new(PingResp)
	err := c.cc.Invoke(ctx, "/test.PingPong/Panic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingPongClient) Err(ctx context.Context, in *ErrorReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/test.PingPong/Err", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingPongServer is the server API for PingPong service.
type PingPongServer interface {
	Ping(context.Context, *PingReq) (*PingResp, error)
	Panic(context.Context, *PingReq) (*PingResp, error)
	Err(context.Context, *ErrorReq) (*Empty, error)
}

func RegisterPingPongServer(s *grpc.Server, srv PingPongServer) {
	s.RegisterService(&_PingPong_serviceDesc, srv)
}

func _PingPong_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.PingPong/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PingPong_Panic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServer).Panic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.PingPong/Panic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServer).Panic(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PingPong_Err_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ErrorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServer).Err(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.PingPong/Err",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServer).Err(ctx, req.(*ErrorReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingPong_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.PingPong",
	HandlerType: (*PingPongServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PingPong_Ping_Handler,
		},
		{
			MethodName: "Panic",
			Handler:    _PingPong_Panic_Handler,
		},
		{
			MethodName: "Err",
			Handler:    _PingPong_Err_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ping.proto",
}
