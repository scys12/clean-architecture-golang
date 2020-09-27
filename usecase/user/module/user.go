package module

import (
	"context"
	"time"

	"github.com/scys12/clean-architecture-golang/usecase/user"

	"github.com/scys12/clean-architecture-golang/model"

	"golang.org/x/crypto/bcrypt"

	awsS3 "github.com/scys12/clean-architecture-golang/pkg/aws"
	"github.com/scys12/clean-architecture-golang/pkg/payload/request"
)

const timeout = 10 * time.Second

func (u *usecase) AuthenticateUser(c context.Context, req *request.LoginRequest) (*user.Response, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	filter := make(map[string]interface{})
	filter["username"] = req.Username

	userAuth, userProfile, err := u.userRepo.GetUserAuthenticateData(ctx, filter)
	if err != nil {
		return nil, err
	}
	passwordCheck := bcrypt.CompareHashAndPassword([]byte(userAuth.Password), []byte(req.Password))
	if passwordCheck != nil {
		return nil, passwordCheck
	}
	u.session.CreateSession()
	return &user.Response{
		ID:       userAuth.ID,
		Email:    userAuth.Email,
		Username: userAuth.Username,
		RoleName: userAuth.Role.Name,
		Image:    userProfile.Image,
		Location: userProfile.Location,
		Name:     userProfile.Name,
		Phone:    userProfile.Phone,
	}, nil

}
func (u *usecase) RegisterUser(c context.Context, req *request.RegisterRequest) error {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(pwd)
	req.Role, err = u.roleRepo.GetRoleByName(ctx, model.ROLE_USER)
	if err != nil {
		return err
	}
	user := model.UserAuth{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		Role:     *req.Role,
	}
	err = u.userRepo.RegisterUser(ctx, user)
	return err
}

func (u *usecase) EditUserProfile(c context.Context, req *request.ProfileRequest) (*user.Response, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	filter := make(map[string]interface{})
	filter["_id"] = req.ID

	_, userProfile, err := u.userRepo.GetUserAuthenticateData(ctx, filter)
	if err != nil {
		return nil, err
	}
	userProfile.ID = req.ID
	userProfile.Location = req.Location
	userProfile.Phone = req.Phone
	userProfile.Name = req.Name
	if req.Image != nil {
		userProfile.Image, err = u.awsStore.UploadFileToS3(awsS3.FileParam{
			FileURL:    userProfile.Image,
			FileHeader: req.Image,
			UserID:     req.ID,
			FolderName: "user",
		})
		if err != nil {
			return nil, err
		}
	}
	err = u.userRepo.EditUserProfile(ctx, *userProfile)
	if err != nil {
		return nil, err
	}
	userAuth, userProfile, err := u.userRepo.GetUserAuthenticateData(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &user.Response{
		ID:       userAuth.ID,
		Email:    userAuth.Email,
		Username: userAuth.Username,
		RoleName: userAuth.Role.Name,
		Image:    userProfile.Image,
		Location: userProfile.Location,
		Name:     userProfile.Name,
		Phone:    userProfile.Phone,
	}, err
}

func (u *usecase) GetUserProfile(ctx context.Context, username string) (*user.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	filter := make(map[string]interface{})
	filter["username"] = username
	userAuth, userProfile, err := u.userRepo.GetUserAuthenticateData(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &user.Response{
		ID:       userAuth.ID,
		Email:    userAuth.Email,
		Username: userAuth.Username,
		RoleName: userAuth.Role.Name,
		Image:    userProfile.Image,
		Location: userProfile.Location,
		Name:     userProfile.Name,
		Phone:    userProfile.Phone,
	}, err
}
