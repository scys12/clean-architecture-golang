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
	redis   session.SessionStore
}

func New(usecase uItem.Usecase, redis session.SessionStore) dItem.Delivery {
	handler := &delivery{
		usecase: usecase,
		redis:   redis,
	}
	return handler
}

func SetRoute(e *echo.Echo, handler dItem.Delivery, redis session.SessionStore) {
	user := e.Group("/user", middleware.SessionMiddleware(redis, model.ROLE_USER))
	user.GET("/items/", handler.GetAllUserItems)
	user.GET(":itemID", handler.GetItem)
	user.POST("/item", handler.CreateItem)
	user.PUT("/item/:itemID", handler.UpdateItem)
	user.DELETE("/item/:itemID", handler.RemoveItem)
}
