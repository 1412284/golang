package main

import (
	"golang/echo-example-crud/handler"

	"github.com/labstack/echo"
)

func main() {
	server := echo.New()

	server.GET("/", handler.Hello)
	server.POST("/login", handler.Login)
	server.Logger.Fatal(server.Start(":8080"))
}
