package main

import (
	"context"
	proto "github.com/eric-chao/go-micro/proto_gen/greeter/proto_gen"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"github.com/micro/go-plugins/registry/consul/v2"
	"testing"
)

func TestClient(t *testing.T) {
	registry := consul.NewRegistry()

	// create GRPC service
	s := grpc.NewService(
		service.Registry(registry),
	)

	// create the greeter client using the service name and client
	greeter := proto.NewGreeterService("go.micro.greeter.function", s.Client())

	// request the Hello method on the Greeter handler
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{
		Name: "John",
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(rsp.Greeting)
}
