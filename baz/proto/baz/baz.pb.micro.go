// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/baz/baz.proto

package echo_micro_service_baz

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Baz service

func NewBazEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Baz service

type BazService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Baz_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Baz_PingPongService, error)
}

type bazService struct {
	c    client.Client
	name string
}

func NewBazService(name string, c client.Client) BazService {
	return &bazService{
		c:    c,
		name: name,
	}
}

func (c *bazService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Baz.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bazService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Baz_StreamService, error) {
	req := c.c.NewRequest(c.name, "Baz.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &bazServiceStream{stream}, nil
}

type Baz_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type bazServiceStream struct {
	stream client.Stream
}

func (x *bazServiceStream) Close() error {
	return x.stream.Close()
}

func (x *bazServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bazServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bazServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bazServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bazService) PingPong(ctx context.Context, opts ...client.CallOption) (Baz_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Baz.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &bazServicePingPong{stream}, nil
}

type Baz_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type bazServicePingPong struct {
	stream client.Stream
}

func (x *bazServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *bazServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *bazServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bazServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bazServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *bazServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Baz service

type BazHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Baz_StreamStream) error
	PingPong(context.Context, Baz_PingPongStream) error
}

func RegisterBazHandler(s server.Server, hdlr BazHandler, opts ...server.HandlerOption) error {
	type baz interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Baz struct {
		baz
	}
	h := &bazHandler{hdlr}
	return s.Handle(s.NewHandler(&Baz{h}, opts...))
}

type bazHandler struct {
	BazHandler
}

func (h *bazHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.BazHandler.Call(ctx, in, out)
}

func (h *bazHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.BazHandler.Stream(ctx, m, &bazStreamStream{stream})
}

type Baz_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type bazStreamStream struct {
	stream server.Stream
}

func (x *bazStreamStream) Close() error {
	return x.stream.Close()
}

func (x *bazStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bazStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bazStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bazStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *bazHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.BazHandler.PingPong(ctx, &bazPingPongStream{stream})
}

type Baz_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type bazPingPongStream struct {
	stream server.Stream
}

func (x *bazPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *bazPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bazPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bazPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *bazPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *bazPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
