package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/scys12/clean-architecture-golang/pkg/payload/request"
	"github.com/scys12/clean-architecture-golang/pkg/payload/response"
)

func (d *delivery) AuthenticateUser(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.LoginRequest)
	if err := c.Bind(req); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	user, err := d.usecase.AuthenticateUser(ctx, req)
	if err != nil {
		return response.Error(c, http.StatusNotFound, err)
	}
	d.redis.CreateSession(c, user)
	return response.OK(c, user)
}

func (d *delivery) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	err := d.usecase.RegisterUser(ctx, req)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	return response.OK(c, nil)
}

func (d *delivery) EditUserProfile(c echo.Context) error {
	panic("need implement")
}
