// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/gateway/gateway.proto

package go_micro_api_gateway

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/micro/go-micro/v2/api/proto"
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

// Api Endpoints for Gateway service

func NewGatewayEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Gateway service

type GatewayService interface {
	Call(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
}

type gatewayService struct {
	c    client.Client
	name string
}

func NewGatewayService(name string, c client.Client) GatewayService {
	return &gatewayService{
		c:    c,
		name: name,
	}
}

func (c *gatewayService) Call(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "Gateway.Call", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gateway service

type GatewayHandler interface {
	Call(context.Context, *proto1.Request, *proto1.Response) error
}

func RegisterGatewayHandler(s server.Server, hdlr GatewayHandler, opts ...server.HandlerOption) error {
	type gateway interface {
		Call(ctx context.Context, in *proto1.Request, out *proto1.Response) error
	}
	type Gateway struct {
		gateway
	}
	h := &gatewayHandler{hdlr}
	return s.Handle(s.NewHandler(&Gateway{h}, opts...))
}

type gatewayHandler struct {
	GatewayHandler
}

func (h *gatewayHandler) Call(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.GatewayHandler.Call(ctx, in, out)
}
