package main

import (
	"go-micro-echo/baz/handler"
	"go-micro-echo/baz/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	baz "go-micro-echo/baz/proto/baz"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("echo.micro.service.baz"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	baz.RegisterBazHandler(service.Server(), new(handler.Baz))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("echo.micro.service.baz", service.Server(), new(subscriber.Baz))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
