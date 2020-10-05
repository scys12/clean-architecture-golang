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
}

func New(usecase uUser.Usecase) dUser.Delivery {
	handler := &delivery{
		usecase: usecase,
	}
	return handler
}

func SetRoute(e *echo.Echo, handler dUser.Delivery, redis session.SessionStore) {
	auth := e.Group("/auth")
	auth.POST("/signin", handler.AuthenticateUser)
	auth.POST("/register", handler.RegisterUser)
	user := e.Group("/user", middleware.SessionMiddleware(redis, model.ROLE_USER))
	user.PUT("/profile", handler.EditUserProfile)
	user.POST("/logout", handler.Logout)
	user.GET("/profile/:username", handler.GetUserProfile)
}
