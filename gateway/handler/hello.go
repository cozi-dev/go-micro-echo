package handler

import (
	"net/http"

	baz "go-micro-echo/baz/proto/baz"

	"github.com/labstack/echo/v4"
	"github.com/micro/go-micro/v2/client"
)

// Hello handler
func Hello(c echo.Context) error {
	bazService := baz.NewBazService("echo.micro.service.baz", client.DefaultClient)
	res, err := bazService.Call(c.Request().Context(), &baz.Request{
		Name: c.QueryParam("name"),
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
