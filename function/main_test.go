package main

import (
	"context"
	proto "github.com/eric-chao/go-micro/proto_gen/greeter/proto_gen"
	"github.com/micro/go-micro/v2/client"
	gclient "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-plugins/registry/consul/v2"

	"testing"
)

func TestClient(t *testing.T) {
	var registry = consul.NewRegistry()
	var selector = selector.NewSelector(
		selector.Registry(registry),
	)

	var client = gclient.NewClient(
		client.Registry(registry),
		client.Selector(selector),
	)

	// create the greeter client using the service name and client
	greeter := proto.NewGreeterService("go.micro.greeter.function", client)

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
