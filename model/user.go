package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Name     string             `bson:"name" json:"name"`
	Location string             `bson:"location" json:"location"`
	Phone    string             `bson:"phone" json:"phone"`
	Role     Role               `bson:"role" json:"role"`
}
