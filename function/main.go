package main

import (
	"context"
	 proto "github.com/eric-chao/go-micro/proto_gen/greeter/proto_gen"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "你好呀！ " + req.Name
	return nil
}

func main() {
	registry := consul.NewRegistry()

	fnc := micro.NewFunction(
		micro.Name("go.micro.greeter.function"),
		micro.Version("1.0.0"),
		micro.Registry(registry),
	)

	fnc.Init()

	fnc.Handle(new(Greeter))

	fnc.Run()
}
