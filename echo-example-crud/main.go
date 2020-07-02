package main

import (
	"github.com/labstack/echo"
)

func main() {
	server := echo.New()

	// server.GET("/", handler)
	// server.POST("/login", handler.login)
	server.Logger.Fatal(server.Start(":8080"))
}
