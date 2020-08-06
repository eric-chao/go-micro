package client

import (
	"context"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
	gateway "path/to/service/proto/gateway"
)

type gatewayKey struct {}

// FromContext retrieves the client from the Context
func GatewayFromContext(ctx context.Context) (gateway.GatewayService, bool) {
	c, ok := ctx.Value(gatewayKey{}).(gateway.GatewayService)
	return c, ok
}

// Client returns a wrapper for the GatewayClient
func GatewayWrapper(service micro.Service) server.HandlerWrapper {
	client := gateway.NewGatewayService("go.micro.service.template", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, gatewayKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
