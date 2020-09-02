package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	BuyDate   time.Time          `bson:"buy_date" json:"buy_date"`
	TotalCost int                `bson:"total_cost" json:"total_cost"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	ItemID    primitive.ObjectID `bson:"item_id" json:"item_id"`
}
