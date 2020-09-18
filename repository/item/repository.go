package item

import (
	"context"

	"github.com/scys12/clean-architecture-golang/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	CreateItem(context.Context, model.Item) error
	UpdateItem(context.Context, model.Item) error
	RemoveItem(context.Context, primitive.ObjectID) error
	GetItemData(context.Context, map[string]interface{}) (*model.Item, error)
}
