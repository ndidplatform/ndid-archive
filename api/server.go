package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ndidplatform/ndid/api/identity"
)

var tendermintAddr string

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/identity/:ns/id/:id", identity.GetIdentifier)
	e.POST("/identity", identity.CreateIdentity)

	e.Logger.Fatal(e.Start(":8000"))
}
