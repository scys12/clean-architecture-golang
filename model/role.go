package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ROLE_USER  = "user"
	ROLE_ADMIN = "admin"
)

type Role struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Name string
}
