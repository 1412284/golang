package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// LoginRequest ...
type LoginRequest struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

// LoginResponse .....
type LoginResponse struct {
	Token string `json:"token" form:"token"`
}

// Handlers
func login(c echo.Context) (err error) {
	req := new(LoginRequest)
	c.Bind(req)
	log.Printf("Login req data: %+v ", req)
	if req.Username != "liemlv" || req.Password != "123456" {
		return c.JSON(http.StatusUnauthorized, "username/password invalid")
	}
	return c.JSON(http.StatusOK, &LoginResponse{Token: "12345678"})
}
