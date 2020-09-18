package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	util "github.com/scys12/clean-architecture-golang/delivery"
	"github.com/scys12/clean-architecture-golang/pkg/payload/request"
	"github.com/scys12/clean-architecture-golang/pkg/payload/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	fileParam = "image"
	UserID    = "userID"
	ItemID    = "itemID"
)

func (d *delivery) CreateItem(c echo.Context) error {
	ctx := c.Request().Context()
	form, image, err := util.HandlingMultipartForm(c, fileParam)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	req := new(request.ItemRequest)
	req.UserID = c.Get(UserID).(primitive.ObjectID)
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

	err = d.usecase.CreateItem(ctx, req)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.OK(c, nil)
}
func (d *delivery) UpdateItem(c echo.Context) error {
	ctx := c.Request().Context()
	itemID, err := primitive.ObjectIDFromHex(c.Param(ItemID))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	form, image, err := util.HandlingMultipartForm(c, fileParam)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}

	req := &request.ItemRequest{
		ID:     itemID,
		UserID: c.Get(UserID).(primitive.ObjectID),
		Image:  image,
	}

	config := util.NewWeaklyTypedConfigDecoder()
	config.Result = req

	err = util.DecodeForm(config, form)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}

	err = d.usecase.UpdateItem(ctx, req)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.OK(c, nil)
}
func (d *delivery) RemoveItem(c echo.Context) error {
	ctx := c.Request().Context()
	itemID, err := primitive.ObjectIDFromHex(c.Param(ItemID))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	err = d.usecase.RemoveItem(ctx, itemID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.OK(c, nil)

}
func (d *delivery) GetItem(echo.Context) error {
	panic("implement")
}
func (d *delivery) GetAllUserItems(echo.Context) error {
	panic("implement")
}
func (d *delivery) GetItemsBasedOnCategory(echo.Context) error {
	panic("implement")
}
func (d *delivery) GetItemsBasedOnUserOwner(echo.Context) error {
	panic("implement")
}
func (d *delivery) SearchUserItem(echo.Context) error {
	panic("implement")
}
