package http

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/labstack/echo/v4"
	util "github.com/scys12/clean-architecture-golang/delivery"
	"github.com/scys12/clean-architecture-golang/pkg/payload/request"
	"github.com/scys12/clean-architecture-golang/pkg/payload/response"
)

const (
	fileParam = "image"
	sessionID = "sessionID"
	userID    = "userID"
	userName  = "username"
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
	return response.OK(c, user)
}

func (d *delivery) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	err := d.usecase.RegisterUser(ctx, req)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	return response.OK(c, nil)
}

func (d *delivery) EditUserProfile(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.ProfileRequest)
	req.ID = c.Get(userID).(primitive.ObjectID)
	form, image, err := util.HandlingMultipartForm(c, fileParam)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	req.Image = image
	config := util.NewWeaklyTypedConfigDecoder()
	config.Result = req
	err = util.DecodeForm(config, form)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	user, err := d.usecase.EditUserProfile(ctx, req)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.OK(c, user)
}

func (d *delivery) GetUserProfile(c echo.Context) error {
	ctx := c.Request().Context()
	username := c.Param(userName)
	user, err := d.usecase.GetUserProfile(ctx, username)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.OK(c, user)
}

func (d *delivery) Logout(c echo.Context) error {
	ID := c.Get(sessionID).(string)
	err := d.redis.Del(ID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.OK(c, nil)
}
