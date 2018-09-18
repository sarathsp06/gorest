package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sarathsp06/gorest/handlers"
)

// AddRoutes attaches the routes
func AddRoutes(e *echo.Echo) {
	e.GET("/", func(ctx echo.Context) error { return ctx.JSON(http.StatusOK, Heartbeat()) })
	e.POST("/order", handlers.CreateOrder)
	e.PUT("/order/:id", handlers.UpdateOrder)
	e.PATCH("/order/:id", handlers.UpdateOrder)
	e.GET("/order", handlers.ListOrders)
}
