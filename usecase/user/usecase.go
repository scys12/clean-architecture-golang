package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/scys12/clean-architecture-golang/pkg/payload/request"
)

type Usecase interface {
	AuthenticateUser(context.Context, *request.LoginRequest) (*Response, error)
	RegisterUser(context.Context, *request.RegisterRequest) error
	EditUserProfile(context.Context, *request.ProfileRequest) (*Response, error)
}

type Response struct {
	ID       primitive.ObjectID `json:"-"`
	Email    string             `json:"email"`
	Username string             `json:"username"`
	RoleName string             `json:"role"`
	Name     string             `json:"name"`
	Location string             `json:"location"`
	Phone    string             `json:"phone"`
	Image    string             `json:"image"`
}
