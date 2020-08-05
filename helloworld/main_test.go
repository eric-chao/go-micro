package main

import (
	"context"
	"github.com/micro/go-micro/v2/client"
	gclient "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-plugins/registry/consul/v2"
	"helloworld/proto/helloworld"
	"testing"
	"time"
)

var service helloworld.HelloworldService

func TestCall(t *testing.T) {
	var ctx = context.Background()
	var req = &helloworld.Request{
		Name: "eric chao",
	}
	resp, err := service.Call(ctx, req)
	if err != nil {
		t.Error(err)
	}
	t.Logf("call msg: %s", resp.GetMsg())
}

func TestStream(t *testing.T) {
	var count = int64(10)
	var ctx = context.Background()
	var req = &helloworld.StreamingRequest{
		Count: count,
	}
	stream, err := service.Stream(ctx, req)
	if err != nil {
		t.Error(err)
	}

	for {
		if count > 0 {
			resp, err := stream.Recv()
			if err != nil {
				t.Error(err)
			}
			t.Logf("stream msg: %d", resp.Count)
		}
		if count <= 0 {
			break
		}
		count--
	}

}

func TestPingPong(t *testing.T) {
	var ctx = context.Background()
	var stream, err = service.PingPong(ctx)
	if err != nil {
		t.Error(err)
	}
	go func() {
		for {
			pong, _ := stream.Recv()
			t.Logf("pong: %d", pong.Stroke)
		}
	}()

	go func() {
		var i = 1
		for {
			stream.Send(&helloworld.Ping{Stroke: int64(i)})
			if i >= 10 {
				break
			}
			i++
		}
	}()

	// wait
	//select {
	//}
	time.Sleep(time.Second * 3)
}

func TestMain(m *testing.M) {
	var registry = consul.NewRegistry()
	var selector = selector.NewSelector(
		selector.Registry(registry),
	)

	var client = gclient.NewClient(
		client.Registry(registry),
		client.Selector(selector),
	)

	// create the greeter client using the service name and client
	service = helloworld.NewHelloworldService("go.micro.service.helloworld", client)

	// run
	m.Run()
}
