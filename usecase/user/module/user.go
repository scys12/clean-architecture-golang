package module

import (
	"context"
	"fmt"
	"time"

	"github.com/scys12/clean-architecture-golang/model"

	"golang.org/x/crypto/bcrypt"

	"github.com/scys12/clean-architecture-golang/pkg/payload/request"
)

const timeout = 10 * time.Second

func (u *usecase) AuthenticateUser(c context.Context, req *request.LoginRequest) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	user, err := u.repo.GetUserAuthenticateData(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", user)
	passwordCheck := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if passwordCheck != nil {
		return nil, passwordCheck
	}
	return user, nil

}
func (u *usecase) RegisterUser(c context.Context, req *request.RegisterRequest) error {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(pwd)
	err = u.repo.RegisterUser(ctx, req)
	return err
}
func (u *usecase) EditUserProfile(context.Context) error {
	panic("need implement")
}
