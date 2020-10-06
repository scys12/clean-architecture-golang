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
	fileParam  = "image"
	UserID     = "userID"
	ItemID     = "itemID"
	CategoryID = "categoryID"
)

func (d *delivery) GetTenLatestItems(c echo.Context) error {
	items, err := d.usecase.GetTenLatestItems()
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.OK(c, items)
}

func (d *delivery) GetAllItems(c echo.Context) error {
	ctx := c.Request().Context()
	items, err := d.usecase.GetAllItems(ctx)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.OK(c, items)
}

func (d *delivery) CreateItem(c echo.Context) error {
	ctx := c.Request().Context()
	form, image, err := util.HandlingMultipartForm(c, fileParam)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	req := new(request.ItemRequest)
	config := util.NewWeaklyTypedConfigDecoder()
	config.Result = req
	err = util.DecodeForm(config, form)
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	req.UserID = c.Get(UserID).(primitive.ObjectID)
	req.Image = image
	categoryID := form["category[id]"].(string)
	req.Category.ID, err = primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	req.Category.Name = form["category[name]"].(string)
	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	item, err := d.usecase.CreateItem(ctx, req)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.OK(c, item)
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

	categoryID := form["category[id]"].(string)
	req.Category.ID, err = primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	req.Category.Name = form["category[name]"].(string)

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
func (d *delivery) GetItem(c echo.Context) error {
	ctx := c.Request().Context()
	itemID, err := primitive.ObjectIDFromHex(c.Param(ItemID))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	item, err := d.usecase.GetItem(ctx, itemID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.OK(c, item)
}
func (d *delivery) GetAllUserItems(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := primitive.ObjectIDFromHex(c.Param(UserID))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	items, err := d.usecase.GetItemBasedOnUserOwner(ctx, userID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.OK(c, items)
}
func (d *delivery) GetItemsBasedOnCategory(c echo.Context) error {
	ctx := c.Request().Context()
	categoryID, err := primitive.ObjectIDFromHex(c.Param(CategoryID))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	items, err := d.usecase.GetItemBasedOnCategory(ctx, categoryID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.OK(c, items)
}
