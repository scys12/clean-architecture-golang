package http

import (
	"net/http"

	"github.com/scys12/clean-architecture-golang/payload/response"

	"github.com/labstack/echo/v4"
)

func (d *delivery) GetAllCategories(c echo.Context) error {
	ctx := c.Request().Context()
	categories, err := d.usecase.GetAllCategories(ctx)
	if err != nil {
		return response.Error(c, http.StatusNotFound, err)
	}
	return response.OK(c, categories)
}
