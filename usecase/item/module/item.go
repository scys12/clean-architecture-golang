package module

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/scys12/clean-architecture-golang/model"

	awsS3 "github.com/scys12/clean-architecture-golang/pkg/aws"
	"github.com/scys12/clean-architecture-golang/pkg/payload/request"

	"github.com/scys12/clean-architecture-golang/usecase/item"
)

const timeout = 10 * time.Second

func (u *usecase) CreateItem(c context.Context, request *request.ItemRequest) (*item.Response, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	newItem := &model.Item{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Category:    request.Category,
		UserID:      request.UserID,
	}
	if request.Image != nil {
		var err error
		newItem.Image, err = u.awsStore.UploadFileToS3(awsS3.FileParam{FileURL: "", FileHeader: request.Image, UserID: request.UserID, FolderName: "item"})
		if err != nil {
			return nil, err
		}
	}
	err := u.itemRepo.CreateItem(ctx, newItem)
	if err != nil {
		return nil, err
	}
	err = u.inMemory.InsertLatestItem(*newItem)
	return item.NewResponse(*newItem), err
}
func (u *usecase) UpdateItem(c context.Context, request *request.ItemRequest) error {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	filter := make(map[string]interface{})
	filter["_id"] = request.ID
	filter["user_id"] = request.UserID

	item, err := u.itemRepo.GetItemData(ctx, filter)
	if err != nil {
		return err
	}
	updatedItem := new(model.Item)
	updatedItem.ID = request.ID
	updatedItem.Description = request.Description
	updatedItem.Name = request.Name
	updatedItem.Price = request.Price
	updatedItem.UserID = request.UserID
	updatedItem.Category = request.Category
	if request.Image != nil {
		updatedItem.Image, err = u.awsStore.UploadFileToS3(awsS3.FileParam{
			FileURL:    item.Image,
			FileHeader: request.Image,
			UserID:     request.ID,
			FolderName: "item",
		})
		if err != nil {
			return err
		}
	}
	err = u.itemRepo.UpdateItem(ctx, *updatedItem)
	if err != nil {
		return err
	}
	err = u.inMemory.UpdateItem(*item, *updatedItem)
	return err
}
func (u *usecase) RemoveItem(c context.Context, itemID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	err := u.itemRepo.RemoveItem(ctx, itemID)
	return err
}
func (u *usecase) GetAllItems(c context.Context) (*item.ItemsResponse, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	items, err := u.itemRepo.GetAllItems(ctx)
	if err != nil {
		return nil, err
	}
	return &item.ItemsResponse{Items: *items}, nil
}
func (u *usecase) GetItem(c context.Context, itemID primitive.ObjectID) (*item.Response, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	filter := make(map[string]interface{})
	filter["_id"] = itemID
	it, err := u.itemRepo.GetItemData(ctx, filter)
	if err != nil {
		return nil, err
	}
	resp := item.NewResponse(*it)
	return resp, nil
}
func (u *usecase) GetItemBasedOnCategory(c context.Context, categoryID primitive.ObjectID) (*item.ItemsResponse, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	filter := make(map[string]interface{})
	filter["category._id"] = categoryID
	items, err := u.itemRepo.GetItemBasedOnCategory(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &item.ItemsResponse{Items: *items}, err

}
func (u *usecase) GetItemBasedOnUserOwner(c context.Context, userID primitive.ObjectID) (*item.ItemsResponse, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	filter := make(map[string]interface{})
	filter["user_id"] = userID
	items, err := u.itemRepo.GetItemsBasedOnUserOwner(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &item.ItemsResponse{Items: *items}, err
}

func (u *usecase) GetTenLatestItems() (*item.ItemsResponse, error) {
	items, err := u.inMemory.GetItems()
	return &item.ItemsResponse{Items: *items}, err
}
