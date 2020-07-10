package main

import (
	"context"
	proto "github.com/eric-chao/go-micro/proto_gen/proto"
	"github.com/micro/go-micro/v2"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "你好呀！ " + req.Name
	return nil
}

func main() {
	fnc := micro.NewFunction(
		micro.Name("hello.function"),
		micro.Version("1.0.0"),
	)

	fnc.Init()

	fnc.Handle(new(Greeter))

	fnc.Run()
}
