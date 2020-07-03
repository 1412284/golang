package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

//X ...........................
type X struct {
	TEXT string `json:"text"`
}

func Hello(c echo.Context) error {
	hello := &X{
		TEXT: "hello",
	}
	return c.JSON(http.StatusOK, hello)
}
