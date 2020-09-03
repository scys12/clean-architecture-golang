package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d *delivery) GetAllCategories(c echo.Context) error {
	ctx := c.Request().Context()
	categories, err := d.usecase.GetAllCategories(ctx)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, categories)
}
