package user

import (
	"github.com/labstack/echo/v4"
)

type Delivery interface {
	AuthenticateUser(echo.Context) error
	RegisterUser(echo.Context) error
	EditUserProfile(echo.Context) error
	GetUserProfile(echo.Context) error
	Logout(echo.Context) error
}
