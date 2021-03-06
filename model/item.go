package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Price       int                `bson:"price" json:"price"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	Category    Category           `bson:"category" json:"category"`
	Image       string             `bson:"image" json:"image"`
}
