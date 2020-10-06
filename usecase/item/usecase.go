package item

import (
	"context"

	"github.com/scys12/clean-architecture-golang/pkg/payload/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usecase interface {
	CreateItem(context.Context, *request.ItemRequest) (*Response, error)
	UpdateItem(context.Context, *request.ItemRequest) error
	RemoveItem(context.Context, primitive.ObjectID) error
	GetAllItems(context.Context) (*ItemsResponse, error)
	GetTenLatestItems() (*ItemsResponse, error)
	GetItem(context.Context, primitive.ObjectID) (*Response, error)
	GetItemBasedOnCategory(context.Context, primitive.ObjectID) (*ItemsResponse, error)
	GetItemBasedOnUserOwner(context.Context, primitive.ObjectID) (*ItemsResponse, error)
}
