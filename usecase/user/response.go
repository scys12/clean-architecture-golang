package user

import (
	"github.com/scys12/clean-architecture-golang/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func NewResponse(userAuth *model.UserAuth, userProfile *model.UserProfile) *Response {
	return &Response{
		ID:       userAuth.ID,
		Email:    userAuth.Email,
		Username: userAuth.Username,
		RoleName: userAuth.Role.Name,
		Image:    userProfile.Image,
		Location: userProfile.Location,
		Name:     userProfile.Name,
		Phone:    userProfile.Phone,
	}
}

type AuthenticateResponse struct {
	SessionID string
	Response  *Response
}
