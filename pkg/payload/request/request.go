package request

import (
	"mime/multipart"

	"github.com/scys12/clean-architecture-golang/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type RegisterRequest struct {
	Email    string      `json:"email,omitempty" validate:"required,email,max=25"`
	Username string      `json:"username,omitempty" validate:"required,max=25,min=4"`
	Password string      `json:"password,omitempty" validate:"required,max=25,min=4"`
	Role     *model.Role `json:"role"`
}

type ProfileRequest struct {
	ID       primitive.ObjectID    `json:"id"`
	Name     string                `json:"name,omitempty" validate:"required,max=25"`
	Location string                `json:"location,omitempty" validate:"max=25"`
	Phone    string                `json:"phone,omitempty" validate:"max=14"`
	Image    *multipart.FileHeader `json:"image,omitempty"`
}
