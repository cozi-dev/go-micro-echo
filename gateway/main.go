package main

import (
	"go-micro-echo/gateway/handler"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// New Service
	service := web.NewService(
		web.Name("echo.micro.api.gateway"),
		web.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	g := e.Group("/gateway")
	g.GET("/hello", handler.Hello)

	// Register Handler
	service.Handle("/", e)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
