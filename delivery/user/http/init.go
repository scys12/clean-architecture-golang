package http

import (
	"github.com/labstack/echo/v4"
	"github.com/scys12/clean-architecture-golang/delivery/middleware"
	"github.com/scys12/clean-architecture-golang/model"
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
	auth := e.Group("/auth")
	auth.POST("/signin", handler.AuthenticateUser)
	auth.POST("/register", handler.RegisterUser)
	user := e.Group("/user/profile", middleware.SessionMiddleware(redis, model.ROLE_USER))
	user.PUT("", handler.EditUserProfile)
	user.GET("/:username", handler.GetUserProfile)
}
