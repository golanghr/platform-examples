// Code generated by protoc-gen-go.
// source: protos/hello.proto
// DO NOT EDIT!

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	protos/hello.proto

It has these top-level messages:
	HelloWorld
*/
package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
import platform "github.com/golanghr/platform/protos"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HelloWorld struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloWorld) Reset()         { *m = HelloWorld{} }
func (m *HelloWorld) String() string { return proto.CompactTextString(m) }
func (*HelloWorld) ProtoMessage()    {}

func init() {
	proto.RegisterType((*HelloWorld)(nil), "hello.HelloWorld")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Hello service

type HelloClient interface {
	HelloWorld(ctx context.Context, in *platform.Request, opts ...grpc.CallOption) (*HelloWorld, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) HelloWorld(ctx context.Context, in *platform.Request, opts ...grpc.CallOption) (*HelloWorld, error) {
	out := new(HelloWorld)
	err := grpc.Invoke(ctx, "/hello.Hello/HelloWorld", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	HelloWorld(context.Context, *platform.Request) (*HelloWorld, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(platform.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(HelloServer).HelloWorld(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _Hello_HelloWorld_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
