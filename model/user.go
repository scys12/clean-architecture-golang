package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAuth struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Role     Role               `bson:"role" json:"role"`
}

type UserProfile struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Location string             `bson:"location" json:"location"`
	Phone    string             `bson:"phone" json:"phone"`
	Image    string             `bson:"image" json:"image"`
}
