// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mapper.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import empty "github.com/golang/protobuf/ptypes/empty"

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

type Map struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Map) Reset()         { *m = Map{} }
func (m *Map) String() string { return proto.CompactTextString(m) }
func (*Map) ProtoMessage()    {}
func (*Map) Descriptor() ([]byte, []int) {
	return fileDescriptor_mapper_b2b14b8bba718431, []int{0}
}
func (m *Map) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Map.Unmarshal(m, b)
}
func (m *Map) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Map.Marshal(b, m, deterministic)
}
func (dst *Map) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Map.Merge(dst, src)
}
func (m *Map) XXX_Size() int {
	return xxx_messageInfo_Map.Size(m)
}
func (m *Map) XXX_DiscardUnknown() {
	xxx_messageInfo_Map.DiscardUnknown(m)
}

var xxx_messageInfo_Map proto.InternalMessageInfo

type Map_Request struct {
	// args is the list of argument types.
	Args []*any.Any `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	// result is the desired result type.
	Result               string   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Map_Request) Reset()         { *m = Map_Request{} }
func (m *Map_Request) String() string { return proto.CompactTextString(m) }
func (*Map_Request) ProtoMessage()    {}
func (*Map_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_mapper_b2b14b8bba718431, []int{0, 0}
}
func (m *Map_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Map_Request.Unmarshal(m, b)
}
func (m *Map_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Map_Request.Marshal(b, m, deterministic)
}
func (dst *Map_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Map_Request.Merge(dst, src)
}
func (m *Map_Request) XXX_Size() int {
	return xxx_messageInfo_Map_Request.Size(m)
}
func (m *Map_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Map_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Map_Request proto.InternalMessageInfo

func (m *Map_Request) GetArgs() []*any.Any {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Map_Request) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type Map_Response struct {
	// result is the mapped data type that matches the type expected
	// by the MapRequest.result field.
	Result               *any.Any `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Map_Response) Reset()         { *m = Map_Response{} }
func (m *Map_Response) String() string { return proto.CompactTextString(m) }
func (*Map_Response) ProtoMessage()    {}
func (*Map_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_mapper_b2b14b8bba718431, []int{0, 1}
}
func (m *Map_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Map_Response.Unmarshal(m, b)
}
func (m *Map_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Map_Response.Marshal(b, m, deterministic)
}
func (dst *Map_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Map_Response.Merge(dst, src)
}
func (m *Map_Response) XXX_Size() int {
	return xxx_messageInfo_Map_Response.Size(m)
}
func (m *Map_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Map_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Map_Response proto.InternalMessageInfo

func (m *Map_Response) GetResult() *any.Any {
	if m != nil {
		return m.Result
	}
	return nil
}

type Map_ListResponse struct {
	// FuncSpec
	Funcs                []*FuncSpec `protobuf:"bytes,1,rep,name=funcs,proto3" json:"funcs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Map_ListResponse) Reset()         { *m = Map_ListResponse{} }
func (m *Map_ListResponse) String() string { return proto.CompactTextString(m) }
func (*Map_ListResponse) ProtoMessage()    {}
func (*Map_ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_mapper_b2b14b8bba718431, []int{0, 2}
}
func (m *Map_ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Map_ListResponse.Unmarshal(m, b)
}
func (m *Map_ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Map_ListResponse.Marshal(b, m, deterministic)
}
func (dst *Map_ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Map_ListResponse.Merge(dst, src)
}
func (m *Map_ListResponse) XXX_Size() int {
	return xxx_messageInfo_Map_ListResponse.Size(m)
}
func (m *Map_ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Map_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Map_ListResponse proto.InternalMessageInfo

func (m *Map_ListResponse) GetFuncs() []*FuncSpec {
	if m != nil {
		return m.Funcs
	}
	return nil
}

func init() {
	proto.RegisterType((*Map)(nil), "proto.Map")
	proto.RegisterType((*Map_Request)(nil), "proto.Map.Request")
	proto.RegisterType((*Map_Response)(nil), "proto.Map.Response")
	proto.RegisterType((*Map_ListResponse)(nil), "proto.Map.ListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MapperClient is the client API for Mapper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MapperClient interface {
	// ListMappers returns the list of mappers that this plugin supports.
	ListMappers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Map_ListResponse, error)
	// Map executes a mapper.
	Map(ctx context.Context, in *Map_Request, opts ...grpc.CallOption) (*Map_Response, error)
}

type mapperClient struct {
	cc *grpc.ClientConn
}

func NewMapperClient(cc *grpc.ClientConn) MapperClient {
	return &mapperClient{cc}
}

func (c *mapperClient) ListMappers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Map_ListResponse, error) {
	out := new(Map_ListResponse)
	err := c.cc.Invoke(ctx, "/proto.Mapper/ListMappers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapperClient) Map(ctx context.Context, in *Map_Request, opts ...grpc.CallOption) (*Map_Response, error) {
	out := new(Map_Response)
	err := c.cc.Invoke(ctx, "/proto.Mapper/Map", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapperServer is the server API for Mapper service.
type MapperServer interface {
	// ListMappers returns the list of mappers that this plugin supports.
	ListMappers(context.Context, *empty.Empty) (*Map_ListResponse, error)
	// Map executes a mapper.
	Map(context.Context, *Map_Request) (*Map_Response, error)
}

func RegisterMapperServer(s *grpc.Server, srv MapperServer) {
	s.RegisterService(&_Mapper_serviceDesc, srv)
}

func _Mapper_ListMappers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapperServer).ListMappers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Mapper/ListMappers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapperServer).ListMappers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mapper_Map_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Map_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapperServer).Map(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Mapper/Map",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapperServer).Map(ctx, req.(*Map_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mapper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Mapper",
	HandlerType: (*MapperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMappers",
			Handler:    _Mapper_ListMappers_Handler,
		},
		{
			MethodName: "Map",
			Handler:    _Mapper_Map_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mapper.proto",
}

func init() { proto.RegisterFile("mapper.proto", fileDescriptor_mapper_b2b14b8bba718431) }

var fileDescriptor_mapper_b2b14b8bba718431 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x89, 0xb5, 0x55, 0xa7, 0x0b, 0x42, 0x94, 0x5a, 0xe3, 0x65, 0x11, 0x84, 0x3d, 0x48,
	0x0a, 0x15, 0xc1, 0x93, 0xe0, 0x41, 0x2f, 0xba, 0x97, 0xf8, 0x0b, 0xd2, 0x65, 0xba, 0x14, 0xb6,
	0xc9, 0xb8, 0x49, 0xc0, 0xfe, 0x3e, 0xff, 0x98, 0x34, 0xd9, 0x2d, 0x56, 0xf1, 0x14, 0xde, 0xbc,
	0x37, 0xdf, 0x84, 0x07, 0xd9, 0x5a, 0x13, 0x61, 0x2b, 0xa9, 0xb5, 0xde, 0xf2, 0x61, 0x7c, 0xc4,
	0x65, 0x6d, 0x6d, 0xdd, 0xe0, 0x2c, 0xaa, 0x45, 0x58, 0xce, 0xb4, 0xd9, 0xa4, 0x84, 0xb8, 0xfa,
	0x6d, 0xe1, 0x9a, 0x7c, 0x6f, 0x66, 0xd4, 0x84, 0x7a, 0x65, 0x92, 0xba, 0xfe, 0x62, 0x30, 0x28,
	0x35, 0x89, 0x57, 0x38, 0x52, 0xf8, 0x11, 0xd0, 0x79, 0x5e, 0xc0, 0xa1, 0x6e, 0x6b, 0x37, 0x65,
	0xf9, 0xa0, 0x18, 0xcf, 0xcf, 0x65, 0x82, 0xc9, 0x1e, 0x26, 0x9f, 0xcc, 0x46, 0xc5, 0x04, 0x9f,
	0xc0, 0xa8, 0x45, 0x17, 0x1a, 0x3f, 0x3d, 0xc8, 0x59, 0x71, 0xa2, 0x3a, 0x25, 0x1e, 0xe0, 0x58,
	0xa1, 0x23, 0x6b, 0x1c, 0xf2, 0xdb, 0x5d, 0x86, 0xe5, 0xec, 0x5f, 0x5e, 0xbf, 0x79, 0x0f, 0xd9,
	0xdb, 0xca, 0xf9, 0xdd, 0xf6, 0x0d, 0x0c, 0x97, 0xc1, 0x54, 0xfd, 0x67, 0x4e, 0xd3, 0x96, 0x7c,
	0x09, 0xa6, 0x7a, 0x27, 0xac, 0x54, 0x72, 0xe7, 0x9f, 0x30, 0x2a, 0x63, 0x45, 0xfc, 0x11, 0xc6,
	0x5b, 0x40, 0x52, 0x8e, 0x4f, 0xfe, 0x5c, 0x7b, 0xde, 0x56, 0x21, 0x2e, 0x3a, 0x50, 0xa9, 0x49,
	0xee, 0x1d, 0x94, 0xb1, 0x0e, 0xce, 0x7f, 0xf8, 0x5d, 0x2f, 0xe2, 0x6c, 0x6f, 0x96, 0xf2, 0x8b,
	0x51, 0x9c, 0xdd, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x41, 0xd6, 0x7b, 0xcd, 0xa3, 0x01, 0x00,
	0x00,
}