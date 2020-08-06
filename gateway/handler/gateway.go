package handler

import (
	"context"
	"encoding/json"
	log "github.com/micro/go-micro/v2/logger"

	"gateway/client"
	"github.com/micro/go-micro/v2/errors"
	api "github.com/micro/go-micro/v2/api/proto"
	gateway "path/to/service/proto/gateway"
)

type Gateway struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Gateway.Call is called by the API as /gateway/call with post body {"name": "foo"}
func (e *Gateway) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Received Gateway.Call request")

	// extract the client from the context
	gatewayClient, ok := client.GatewayFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.gateway.gateway.call", "gateway client not found")
	}

	// make request
	response, err := gatewayClient.Call(ctx, &gateway.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.gateway.gateway.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
