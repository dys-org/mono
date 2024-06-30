package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/dys-org/oapi-todo/api"
)

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.RegisterHandlers(e, server)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(":6969"))
}
