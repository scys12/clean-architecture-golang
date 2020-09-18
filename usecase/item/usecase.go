package item

import (
	"context"

	"github.com/scys12/clean-architecture-golang/model"
	"github.com/scys12/clean-architecture-golang/pkg/payload/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usecase interface {
	CreateItem(context.Context, *request.ItemRequest) error
	UpdateItem(context.Context, *request.ItemRequest) error
	RemoveItem(context.Context, primitive.ObjectID) error
	GetAllItems(context.Context, string) (*Response, error)
	GetItem(context.Context, string) (*Response, error)
	GetItemBasedOnCategory(context.Context) error
	GetItemBasedOnUserOwner(context.Context) error
	SearchUserItem(context.Context) error
}

type Response struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       int                `json:"price"`
	Category    model.Category     `json:"category"`
	Image       string             `json:"image"`
}
