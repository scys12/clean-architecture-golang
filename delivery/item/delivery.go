package item

import (
	"github.com/labstack/echo/v4"
)

type Delivery interface {
	CreateItem(echo.Context) error
	UpdateItem(echo.Context) error
	RemoveItem(echo.Context) error
	GetItem(echo.Context) error
	GetAllUserItems(echo.Context) error
	GetItemsBasedOnCategory(echo.Context) error
	GetItemsBasedOnUserOwner(echo.Context) error
	SearchUserItem(echo.Context) error
}
