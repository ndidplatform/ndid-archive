package main

import (
	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ndidplatform/ndid/api/client/tendermint"
	"github.com/ndidplatform/ndid/api/identity"
	"github.com/ndidplatform/ndid/api/rp"
)

var port string

func initFlag() {
	flag.StringVar(&port, "port", ":8000", "port for start service example :8000 or 127.0.0.1:8000")
	flag.StringVar(&tendermint.TendermintAddr, "tenderm", "127.0.0.1:46657", "tendermint port example 127.0.0.1:46657 (please add ip address and port)")

	flag.Parse()
}

func main() {
	initFlag()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/identity/:ns/:id", identity.GetIdentifier)
	e.POST("/identity", identity.CreateIdentity)

	e.POST("/rp/requests/:ns/:id", rp.CreateRequest)

	e.Logger.Fatal(e.Start(port))
}
