package module

import (
	"context"
	"fmt"
	"time"

	"github.com/scys12/clean-architecture-golang/pkg/session"

	"github.com/scys12/clean-architecture-golang/usecase/user"

	"github.com/scys12/clean-architecture-golang/model"

	"golang.org/x/crypto/bcrypt"

	awsS3 "github.com/scys12/clean-architecture-golang/pkg/aws"
	"github.com/scys12/clean-architecture-golang/pkg/payload/request"
)

const timeout = 10 * time.Second

func (u *usecase) AuthenticateUser(c context.Context, req *request.LoginRequest) (*user.AuthenticateResponse, error) {
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
	sessID, err := u.session.CreateSession(userAuth)
	if err != nil {
		return nil, err
	}
	return &user.AuthenticateResponse{
		Response:  user.NewResponse(userAuth, userProfile),
		SessionID: sessID,
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
	return user.NewResponse(userAuth, userProfile), err
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
	return user.NewResponse(userAuth, userProfile), err
}

func (u *usecase) Logout(ctx context.Context, sessionID, userID string) error {
	field := []string{fmt.Sprintf("user:%v", userID)}
	data := make([]interface{}, len(field))
	for i, str := range field {
		data[i] = str
	}
	err := u.session.Del(session.FieldStore{TypeField: "DEL", DataField: data})
	if err != nil {
		return err
	}
	field = []string{"session", "sessionID"}
	data = make([]interface{}, len(field))
	for i, str := range field {
		data[i] = str
	}
	err = u.session.Del(session.FieldStore{TypeField: "HDEL", DataField: data})
	return err
}
