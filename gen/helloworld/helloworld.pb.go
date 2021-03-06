// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld/helloworld.proto

package helloworld

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/utahta/grpc-go-protoreflect-example/gen/option"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73149fedf49f4319, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_73149fedf49f4319, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

func init() { proto.RegisterFile("helloworld/helloworld.proto", fileDescriptor_73149fedf49f4319) }

var fileDescriptor_73149fedf49f4319 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4a, 0xc7, 0x30,
	0x10, 0x87, 0xfd, 0xa3, 0x58, 0x3d, 0x9c, 0x22, 0x48, 0xa9, 0x8b, 0x64, 0x10, 0x97, 0x36, 0xa0,
	0xa3, 0x20, 0xe8, 0x62, 0x11, 0xa7, 0xba, 0xb9, 0xa5, 0xf5, 0x4c, 0x0b, 0x49, 0x2f, 0xa6, 0x57,
	0xb4, 0x6f, 0xe7, 0xa3, 0x89, 0xa9, 0xa5, 0x1d, 0x9c, 0xf2, 0xe5, 0xee, 0x3b, 0xee, 0xc7, 0xc1,
	0x79, 0x8b, 0xd6, 0xd2, 0x27, 0x05, 0xfb, 0xa6, 0x56, 0x2c, 0x7c, 0x20, 0x26, 0x01, 0x6b, 0x25,
	0x3b, 0x25, 0xcf, 0x1d, 0xf5, 0x6a, 0x7e, 0x66, 0x41, 0x4a, 0x38, 0x29, 0x7f, 0x95, 0x0a, 0x3f,
	0x46, 0x1c, 0x58, 0x08, 0x38, 0xe8, 0xb5, 0xc3, 0x74, 0x77, 0xb1, 0xbb, 0x3a, 0xae, 0x22, 0xcb,
	0x4b, 0x80, 0x3f, 0xc7, 0xdb, 0x49, 0xa4, 0x90, 0x38, 0x1c, 0x06, 0x6d, 0x16, 0x69, 0xf9, 0x5e,
	0x3f, 0x43, 0xf2, 0x18, 0x10, 0x19, 0x83, 0xb8, 0x87, 0xa3, 0x17, 0x3d, 0xc5, 0x29, 0x91, 0x16,
	0x9b, 0x58, 0xdb, 0x65, 0xd9, 0xd9, 0x3f, 0x1d, 0x6f, 0x27, 0xb9, 0xff, 0x7d, 0xb7, 0xf7, 0xf0,
	0xf4, 0x5a, 0x9a, 0x8e, 0xdb, 0xb1, 0x2e, 0x1a, 0x72, 0x6a, 0x64, 0xdd, 0xb2, 0x56, 0x26, 0xf8,
	0x26, 0x37, 0x94, 0xc7, 0xf0, 0x01, 0xdf, 0x2d, 0x36, 0x9c, 0xe3, 0x97, 0x76, 0xde, 0xa2, 0x32,
	0xd8, 0x6f, 0x2e, 0x70, 0xbb, 0x62, 0x7d, 0x18, 0xfd, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x65, 0x0d, 0x47, 0xce, 0x2c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/helloworld.proto",
}
