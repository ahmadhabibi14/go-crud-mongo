package main

import (
	"go-crud-mongo/controller"

	"github.com/labstack/echo/v4"
)

func startEchoServer() {
	e := echo.New()
	e.POST("/user", controller.AddUser)

	e.Logger.Fatal(e.Start(":1323"))
}
