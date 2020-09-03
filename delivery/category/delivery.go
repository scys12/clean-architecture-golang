package category

import (
	"github.com/labstack/echo/v4"
)

type Delivery interface {
	GetAllCategories(echo.Context) error
}
