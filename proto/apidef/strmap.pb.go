// Code generated by protoc-gen-go. DO NOT EDIT.
// source: strmap.proto

package apidef // import "proto/apidef"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StrMapReq struct {
	Payload              string   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StrMapReq) Reset()         { *m = StrMapReq{} }
func (m *StrMapReq) String() string { return proto.CompactTextString(m) }
func (*StrMapReq) ProtoMessage()    {}
func (*StrMapReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_strmap_233e2f9353693ca4, []int{0}
}
func (m *StrMapReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StrMapReq.Unmarshal(m, b)
}
func (m *StrMapReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StrMapReq.Marshal(b, m, deterministic)
}
func (dst *StrMapReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StrMapReq.Merge(dst, src)
}
func (m *StrMapReq) XXX_Size() int {
	return xxx_messageInfo_StrMapReq.Size(m)
}
func (m *StrMapReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StrMapReq.DiscardUnknown(m)
}

var xxx_messageInfo_StrMapReq proto.InternalMessageInfo

func (m *StrMapReq) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *StrMapReq) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type StrMapResp struct {
	Ret                  int32    `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Result               string   `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StrMapResp) Reset()         { *m = StrMapResp{} }
func (m *StrMapResp) String() string { return proto.CompactTextString(m) }
func (*StrMapResp) ProtoMessage()    {}
func (*StrMapResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_strmap_233e2f9353693ca4, []int{1}
}
func (m *StrMapResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StrMapResp.Unmarshal(m, b)
}
func (m *StrMapResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StrMapResp.Marshal(b, m, deterministic)
}
func (dst *StrMapResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StrMapResp.Merge(dst, src)
}
func (m *StrMapResp) XXX_Size() int {
	return xxx_messageInfo_StrMapResp.Size(m)
}
func (m *StrMapResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StrMapResp.DiscardUnknown(m)
}

var xxx_messageInfo_StrMapResp proto.InternalMessageInfo

func (m *StrMapResp) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *StrMapResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *StrMapResp) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*StrMapReq)(nil), "apidef.StrMapReq")
	proto.RegisterType((*StrMapResp)(nil), "apidef.StrMapResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StringMapClient is the client API for StringMap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StringMapClient interface {
	Map(ctx context.Context, in *StrMapReq, opts ...grpc.CallOption) (*StrMapResp, error)
}

type stringMapClient struct {
	cc *grpc.ClientConn
}

func NewStringMapClient(cc *grpc.ClientConn) StringMapClient {
	return &stringMapClient{cc}
}

func (c *stringMapClient) Map(ctx context.Context, in *StrMapReq, opts ...grpc.CallOption) (*StrMapResp, error) {
	out := new(StrMapResp)
	err := c.cc.Invoke(ctx, "/apidef.StringMap/Map", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StringMapServer is the server API for StringMap service.
type StringMapServer interface {
	Map(context.Context, *StrMapReq) (*StrMapResp, error)
}

func RegisterStringMapServer(s *grpc.Server, srv StringMapServer) {
	s.RegisterService(&_StringMap_serviceDesc, srv)
}

func _StringMap_Map_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StrMapReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringMapServer).Map(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apidef.StringMap/Map",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringMapServer).Map(ctx, req.(*StrMapReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _StringMap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apidef.StringMap",
	HandlerType: (*StringMapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Map",
			Handler:    _StringMap_Map_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmap.proto",
}

func init() { proto.RegisterFile("strmap.proto", fileDescriptor_strmap_233e2f9353693ca4) }

var fileDescriptor_strmap_233e2f9353693ca4 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x29, 0xca,
	0x4d, 0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2c, 0xc8, 0x4c, 0x49, 0x4d,
	0x53, 0x72, 0xe6, 0xe2, 0x0c, 0x2e, 0x29, 0xf2, 0x4d, 0x2c, 0x08, 0x4a, 0x2d, 0x14, 0x92, 0xe0,
	0x62, 0x2f, 0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82,
	0x71, 0x85, 0x64, 0xb8, 0x38, 0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12, 0x73, 0x0b, 0x24, 0x98,
	0xc0, 0x72, 0x08, 0x01, 0x25, 0x0f, 0x2e, 0x2e, 0x98, 0x21, 0xc5, 0x05, 0x42, 0x02, 0x5c, 0xcc,
	0x45, 0xa9, 0x25, 0x60, 0x13, 0x58, 0x83, 0x40, 0x4c, 0x90, 0x48, 0x6e, 0x71, 0x3a, 0x54, 0x1f,
	0x88, 0x29, 0x24, 0xc6, 0xc5, 0x56, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x22, 0xc1, 0x0c, 0x16, 0x84,
	0xf2, 0x8c, 0xac, 0xc1, 0xce, 0xc9, 0xcc, 0x4b, 0xf7, 0x4d, 0x2c, 0x10, 0xd2, 0xe3, 0x62, 0x06,
	0x51, 0x82, 0x7a, 0x10, 0xb7, 0xea, 0xc1, 0x1d, 0x2a, 0x25, 0x84, 0x2e, 0x54, 0x5c, 0xa0, 0xc4,
	0xe0, 0xc4, 0x17, 0xc5, 0x03, 0xf6, 0x9c, 0x3e, 0x44, 0x32, 0x89, 0x0d, 0xcc, 0x33, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xde, 0x5e, 0xb4, 0x5e, 0xfa, 0x00, 0x00, 0x00,
}
