package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sarathsp06/gorest/config"
	"github.com/sarathsp06/gorest/core"
	"github.com/sarathsp06/gorest/middlewares"
)

var (
	host = flag.String("host", "0.0.0.0", "Host ip")
	port = flag.String("port", "8080", "Host port")
	// ServiceDirectory the directory where the service runs
	ServiceDirectory string
)

func main() {
	flag.Parse()

	// initialize configuration manager
	if err := configmanager.InitConfig(ServiceDirectory); err != nil {
		log.Fatalln("Failed initializing configmanager. Error = ", err)
	}

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("1024KB"))
	e.Use(middleware.Secure())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middlewares.RequestID)
	e.Use(middlewares.Method)
	e.Use(middlewares.FileLogger())
	// adding routes
	AddRoutes(e)
	// try to intialize the core components and exit of failure
	core.Initialize()
	if err := e.Start(fmt.Sprintf("%s:%d", configmanager.Config.Host, configmanager.Config.Port)); err != nil {
		log.Fatalln("Failed to start server!", err)
	}
}
