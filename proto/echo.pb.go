// Code generated by protoc-gen-go.
// source: echo.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	echo.proto

It has these top-level messages:
	Hello
	EchoHello
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Hello struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Hello) Reset()                    { *m = Hello{} }
func (m *Hello) String() string            { return proto1.CompactTextString(m) }
func (*Hello) ProtoMessage()               {}
func (*Hello) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EchoHello struct {
	Echomsg string `protobuf:"bytes,2,opt,name=echomsg" json:"echomsg,omitempty"`
}

func (m *EchoHello) Reset()                    { *m = EchoHello{} }
func (m *EchoHello) String() string            { return proto1.CompactTextString(m) }
func (*EchoHello) ProtoMessage()               {}
func (*EchoHello) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto1.RegisterType((*Hello)(nil), "proto.Hello")
	proto1.RegisterType((*EchoHello)(nil), "proto.EchoHello")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for EchoService service

type EchoServiceClient interface {
	Echo(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoClient, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, opts ...grpc.CallOption) (EchoService_EchoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EchoService_serviceDesc.Streams[0], c.cc, "/proto.EchoService/Echo", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceEchoClient{stream}
	return x, nil
}

type EchoService_EchoClient interface {
	Send(*Hello) error
	Recv() (*EchoHello, error)
	grpc.ClientStream
}

type echoServiceEchoClient struct {
	grpc.ClientStream
}

func (x *echoServiceEchoClient) Send(m *Hello) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceEchoClient) Recv() (*EchoHello, error) {
	m := new(EchoHello)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	Echo(EchoService_EchoServer) error
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).Echo(&echoServiceEchoServer{stream})
}

type EchoService_EchoServer interface {
	Send(*EchoHello) error
	Recv() (*Hello, error)
	grpc.ServerStream
}

type echoServiceEchoServer struct {
	grpc.ServerStream
}

func (x *echoServiceEchoServer) Send(m *EchoHello) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceEchoServer) Recv() (*Hello, error) {
	m := new(Hello)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Echo",
			Handler:       _EchoService_Echo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto1.RegisterFile("echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x92, 0x5c, 0xac, 0x1e, 0xa9,
	0x39, 0x39, 0xf9, 0x42, 0x02, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x20, 0xa6, 0x92, 0x2a, 0x17, 0xa7, 0x2b, 0x50, 0x3d, 0x44, 0x5a, 0x82, 0x8b, 0x1d, 0xa4,
	0x19, 0xa4, 0x84, 0x09, 0xac, 0x04, 0xc6, 0x35, 0xb2, 0xe6, 0xe2, 0x06, 0x29, 0x0b, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xd2, 0xe1, 0x62, 0x01, 0x71, 0x85, 0x78, 0x20, 0xf6, 0xe8, 0x81,
	0xb5, 0x4b, 0x09, 0x40, 0x79, 0x70, 0x03, 0x95, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x93, 0xd8, 0xc0,
	0xc2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x62, 0x9e, 0xe6, 0xf9, 0x9a, 0x00, 0x00, 0x00,
}
