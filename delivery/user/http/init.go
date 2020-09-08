package http

import (
	"github.com/labstack/echo/v4"
	"github.com/scys12/clean-architecture-golang/delivery/middleware"
	"github.com/scys12/clean-architecture-golang/session"

	dUser "github.com/scys12/clean-architecture-golang/delivery/user"
	uUser "github.com/scys12/clean-architecture-golang/usecase/user"
)

type delivery struct {
	usecase uUser.Usecase
	redis   *session.RedisClient
}

func New(usecase uUser.Usecase, redis *session.RedisClient) dUser.Delivery {
	handler := &delivery{
		usecase: usecase,
		redis:   redis,
	}
	return handler
}

func SetRoute(e *echo.Echo, handler dUser.Delivery, redis *session.RedisClient) {
	e.POST("/auth/signin", handler.AuthenticateUser)
	e.POST("/auth/register", handler.RegisterUser)
	e.PUT("/auth/register", middleware.SessionMiddleware(redis)(handler.EditUserProfile))
}
