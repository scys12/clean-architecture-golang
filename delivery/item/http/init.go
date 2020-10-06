package http

import (
	"github.com/labstack/echo/v4"
	"github.com/scys12/clean-architecture-golang/delivery/middleware"
	"github.com/scys12/clean-architecture-golang/model"
	"github.com/scys12/clean-architecture-golang/pkg/session"

	dItem "github.com/scys12/clean-architecture-golang/delivery/item"
	uItem "github.com/scys12/clean-architecture-golang/usecase/item"
)

type delivery struct {
	usecase uItem.Usecase
}

func New(usecase uItem.Usecase) dItem.Delivery {
	handler := &delivery{
		usecase: usecase,
	}
	return handler
}

func SetRoute(e *echo.Echo, handler dItem.Delivery, redis session.SessionStore) {
	items := e.Group("/items")
	items.GET("", handler.GetAllItems)
	items.GET("/category/:categoryID", handler.GetItemsBasedOnCategory)
	e.GET("/latest", handler.GetTenLatestItems)
	user := e.Group("/user", middleware.SessionMiddleware(redis, model.ROLE_USER))
	user.GET("/:userID/items", handler.GetAllUserItems)
	userItem := user.Group("/item")
	userItem.GET("/:itemID", handler.GetItem)
	userItem.POST("", handler.CreateItem)
	userItem.PUT("/:itemID", handler.UpdateItem)
	userItem.DELETE("/:itemID", handler.RemoveItem)
}
