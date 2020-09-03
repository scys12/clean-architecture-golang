package http

import (
	"github.com/labstack/echo/v4"

	dCategory "github.com/scys12/clean-architecture-golang/delivery/category"
	uCategory "github.com/scys12/clean-architecture-golang/usecase/category"
)

type delivery struct {
	usecase uCategory.Usecase
}

func New(usecase uCategory.Usecase) dCategory.Delivery {
	handler := &delivery{
		usecase: usecase,
	}
	return handler
}

func SetRoute(e *echo.Echo, handler dCategory.Delivery) {
	e.GET("/categories", handler.GetAllCategories)
}
