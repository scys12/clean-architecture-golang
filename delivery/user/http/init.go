package http

import (
	"github.com/labstack/echo/v4"
	"github.com/scys12/clean-architecture-golang/delivery/middleware"
	"github.com/scys12/clean-architecture-golang/pkg/session"

	dUser "github.com/scys12/clean-architecture-golang/delivery/user"
	uUser "github.com/scys12/clean-architecture-golang/usecase/user"
)

type delivery struct {
	usecase uUser.Usecase
	redis   session.SessionStore
}

func New(usecase uUser.Usecase, redis session.SessionStore) dUser.Delivery {
	handler := &delivery{
		usecase: usecase,
		redis:   redis,
	}
	return handler
}

func SetRoute(e *echo.Echo, handler dUser.Delivery, redis session.SessionStore) {
	e.POST("/auth/signin", handler.AuthenticateUser)
	e.POST("/auth/register", handler.RegisterUser)
	e.PUT("/user/profile", middleware.SessionMiddleware(redis)(handler.EditUserProfile))
}
