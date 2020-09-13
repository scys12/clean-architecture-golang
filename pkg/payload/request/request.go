package request

import (
	"github.com/scys12/clean-architecture-golang/model"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string      `bson:"email" json:"email"`
	Username string      `bson:"username" json:"username"`
	Password string      `bson:"password" json:"password"`
	Role     *model.Role `bson:"role" json:"role"`
}
