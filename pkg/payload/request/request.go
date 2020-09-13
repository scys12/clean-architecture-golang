package request

import (
	"github.com/scys12/clean-architecture-golang/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Name     string             `bson:"name" json:"name"`
	Location string             `bson:"location" json:"location"`
	Phone    string             `bson:"phone" json:"phone"`
	Role     model.Role         `bson:"role" json:"role"`
}
