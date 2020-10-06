package item

import (
	"github.com/scys12/clean-architecture-golang/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       int                `json:"price"`
	UserID      primitive.ObjectID `json:"user_id"`
	Category    model.Category     `json:"category"`
	Image       string             `json:"image"`
}

func NewResponse(item model.Item) *Response {
	return &Response{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		UserID:      item.UserID,
		Category:    item.Category,
		Image:       item.Image,
	}
}

type ItemsResponse struct {
	Items []model.Item `json:"items"`
}
