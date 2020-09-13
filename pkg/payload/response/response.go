package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func OK(c echo.Context, data interface{}) error {
	r := &Response{
		Status: true,
		Data:   data,
	}
	return c.JSON(http.StatusOK, r)
}

func Error(c echo.Context, errCode int, err error) error {
	r := &Response{
		Status: false,
		Data:   err.Error(),
	}
	return c.JSON(errCode, r)
}
