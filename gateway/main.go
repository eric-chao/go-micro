package main

import (
	log "github.com/micro/go-micro/v2/logger"

	"github.com/micro/go-micro/v2"
	"gateway/handler"
	"gateway/client"

	gateway "gateway/proto/gateway"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.gateway"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Gateway service client
		micro.WrapHandler(client.GatewayWrapper(service)),
	)

	// Register Handler
	gateway.RegisterGatewayHandler(service.Server(), new(handler.Gateway))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
