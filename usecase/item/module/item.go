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

func (u *usecase) CreateItem(c context.Context, request *request.ItemRequest) error {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	image, err := u.awsStore.UploadFileToS3(awsS3.FileParam{
		FileURL:    "",
		FileHeader: request.Image,
		UserID:     request.UserID,
		FolderName: "item",
	})
	if err != nil {
		return err
	}
	err = u.itemRepo.CreateItem(ctx, model.Item{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Category:    request.Category,
		Image:       image,
	})
	return err
}
func (u *usecase) UpdateItem(c context.Context, request *request.ItemRequest) error {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	filter := make(map[string]interface{})
	filter["_id"] = request.ID
	item, err := u.itemRepo.GetItemData(ctx, filter)
	if err != nil {
		return err
	}
	item.ID = request.ID
	item.Description = request.Description
	item.Name = request.Name
	item.Price = request.Price
	item.UserID = request.UserID
	item.Category = request.Category
	if request.Image != nil {
		item.Image, err = u.awsStore.UploadFileToS3(awsS3.FileParam{
			FileURL:    item.Image,
			FileHeader: request.Image,
			UserID:     request.ID,
			FolderName: "item",
		})
		if err != nil {
			return err
		}
	}
	err = u.itemRepo.UpdateItem(ctx, *item)
	return err
}
func (u *usecase) RemoveItem(c context.Context, itemID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	err := u.itemRepo.RemoveItem(ctx, itemID)
	return err
}
func (u *usecase) GetAllItems(context.Context, string) (*item.Response, error) {
	panic("implement")
}
func (u *usecase) GetItem(context.Context, string) (*item.Response, error) {
	panic("implement")
}
func (u *usecase) GetItemBasedOnCategory(context.Context) error {
	panic("implement")
}
func (u *usecase) GetItemBasedOnUserOwner(context.Context) error {
	panic("implement")
}
func (u *usecase) SearchUserItem(context.Context) error {
	panic("implement")
}
