// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package grhello

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

type Sentiment int32

const (
	Sentiment_HAPPY  Sentiment = 0
	Sentiment_SLEEPY Sentiment = 1
	Sentiment_ANGRY  Sentiment = 2
)

var Sentiment_name = map[int32]string{
	0: "HAPPY",
	1: "SLEEPY",
	2: "ANGRY",
}
var Sentiment_value = map[string]int32{
	"HAPPY":  0,
	"SLEEPY": 1,
	"ANGRY":  2,
}

func (x Sentiment) String() string {
	return proto.EnumName(Sentiment_name, int32(x))
}
func (Sentiment) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_hello_fb5a4e38fed629af, []int{0}
}

// The request message containing the user's name.
type Request struct {
	FirstName            string            `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string            `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Age                  int64             `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Sentiment            Sentiment         `protobuf:"varint,4,opt,name=sentiment,proto3,enum=com.ytel.common.grpc.info.Sentiment" json:"sentiment,omitempty"`
	Hobbies              []string          `protobuf:"bytes,5,rep,name=hobbies,proto3" json:"hobbies,omitempty"`
	BagOfTricks          map[string]string `protobuf:"bytes,6,rep,name=bagOfTricks,proto3" json:"bagOfTricks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_fb5a4e38fed629af, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Request) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Request) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Request) GetSentiment() Sentiment {
	if m != nil {
		return m.Sentiment
	}
	return Sentiment_HAPPY
}

func (m *Request) GetHobbies() []string {
	if m != nil {
		return m.Hobbies
	}
	return nil
}

func (m *Request) GetBagOfTricks() map[string]string {
	if m != nil {
		return m.BagOfTricks
	}
	return nil
}

type Response struct {
	Greeting             string    `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	Sentiment            Sentiment `protobuf:"varint,2,opt,name=sentiment,proto3,enum=com.ytel.common.grpc.info.Sentiment" json:"sentiment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_hello_fb5a4e38fed629af, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func (m *Response) GetSentiment() Sentiment {
	if m != nil {
		return m.Sentiment
	}
	return Sentiment_HAPPY
}

func init() {
	proto.RegisterType((*Request)(nil), "com.ytel.common.grpc.info.Request")
	proto.RegisterMapType((map[string]string)(nil), "com.ytel.common.grpc.info.Request.BagOfTricksEntry")
	proto.RegisterType((*Response)(nil), "com.ytel.common.grpc.info.Response")
	proto.RegisterEnum("com.ytel.common.grpc.info.Sentiment", Sentiment_name, Sentiment_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	// Sends a greeting
	Greet(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Greet(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/com.ytel.common.grpc.info.HelloService/greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	// Sends a greeting
	Greet(context.Context, *Request) (*Response, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ytel.common.grpc.info.HelloService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Greet(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.ytel.common.grpc.info.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "greet",
			Handler:    _HelloService_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_hello_fb5a4e38fed629af) }

var fileDescriptor_hello_fb5a4e38fed629af = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x6d, 0x6b, 0x81, 0x0e, 0xc6, 0x34, 0x1b, 0x0f, 0x95, 0x78, 0x68, 0xd0, 0x43, 0x63,
	0xe2, 0x1e, 0xe0, 0x62, 0x3c, 0x98, 0x40, 0x42, 0xe4, 0x60, 0xb0, 0x59, 0xf4, 0xc0, 0xcd, 0xb6,
	0x19, 0x4a, 0xa5, 0xed, 0xe2, 0xee, 0x42, 0xc2, 0x2b, 0xfa, 0x54, 0xa6, 0xa5, 0x80, 0x9a, 0x88,
	0x89, 0xb7, 0xf9, 0xe6, 0x5f, 0xe6, 0xfb, 0x65, 0xa0, 0x39, 0xc3, 0x34, 0xe5, 0x74, 0x21, 0xb8,
	0xe2, 0xe4, 0x3c, 0xe2, 0x19, 0x5d, 0x2b, 0x4c, 0x69, 0xc4, 0xb3, 0x8c, 0xe7, 0x34, 0x16, 0x8b,
	0x88, 0x26, 0xf9, 0x94, 0xb7, 0x3f, 0x74, 0xa8, 0x33, 0x7c, 0x5f, 0xa2, 0x54, 0xe4, 0x02, 0xac,
	0x69, 0x22, 0xa4, 0x1a, 0x05, 0x19, 0x3a, 0x9a, 0xab, 0x79, 0x16, 0xdb, 0x27, 0x48, 0x0b, 0x1a,
	0x69, 0x50, 0x15, 0xf5, 0xb2, 0xb8, 0xd3, 0xc4, 0x06, 0x23, 0x88, 0xd1, 0x31, 0x5c, 0xcd, 0x33,
	0x58, 0x11, 0x92, 0x3e, 0x58, 0x12, 0x73, 0x95, 0x64, 0x98, 0x2b, 0xe7, 0xd8, 0xd5, 0xbc, 0xd3,
	0xce, 0x15, 0xfd, 0xf5, 0x0c, 0x3a, 0xde, 0xf6, 0xb2, 0xfd, 0x18, 0x71, 0xa0, 0x3e, 0xe3, 0x61,
	0x98, 0xa0, 0x74, 0x4c, 0xd7, 0xf0, 0x2c, 0xb6, 0x95, 0xe4, 0x05, 0x9a, 0x61, 0x10, 0x3f, 0x4d,
	0x9f, 0x45, 0x12, 0xcd, 0xa5, 0x53, 0x73, 0x0d, 0xaf, 0xd9, 0xe9, 0x1e, 0xd8, 0x5f, 0x59, 0xa4,
	0xfd, 0xfd, 0xd4, 0x20, 0x57, 0x62, 0xcd, 0xbe, 0xee, 0x69, 0xdd, 0x83, 0xfd, 0xb3, 0xa1, 0xb0,
	0x36, 0xc7, 0x75, 0x85, 0xa3, 0x08, 0xc9, 0x19, 0x98, 0xab, 0x20, 0x5d, 0x6e, 0x29, 0x6c, 0xc4,
	0x9d, 0x7e, 0xab, 0xb5, 0xdf, 0xa0, 0xc1, 0x50, 0x2e, 0x78, 0x2e, 0x4b, 0x5c, 0xb1, 0x40, 0x54,
	0x49, 0x1e, 0x57, 0xc3, 0x3b, 0xfd, 0x1d, 0x8e, 0xfe, 0x2f, 0x38, 0xd7, 0x37, 0x60, 0xed, 0xf2,
	0xc4, 0x02, 0x73, 0xd8, 0xf3, 0xfd, 0x89, 0x7d, 0x44, 0x00, 0x6a, 0xe3, 0xc7, 0xc1, 0xc0, 0x9f,
	0xd8, 0x5a, 0x91, 0xee, 0x8d, 0x1e, 0xd8, 0xc4, 0xd6, 0x3b, 0xaf, 0x70, 0x32, 0x2c, 0x3e, 0x62,
	0x8c, 0x62, 0x95, 0x44, 0x48, 0x7c, 0x30, 0xcb, 0x73, 0x48, 0xfb, 0x6f, 0x6a, 0xad, 0xcb, 0x83,
	0x3d, 0x1b, 0xc3, 0x7d, 0xdd, 0xd7, 0xc2, 0x5a, 0xf9, 0x6f, 0xdd, 0xcf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x18, 0x03, 0x29, 0xdb, 0x7e, 0x02, 0x00, 0x00,
}
