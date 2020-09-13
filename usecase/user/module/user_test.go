package module_test

import (
	"context"
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/labstack/gommon/log"

	"golang.org/x/crypto/bcrypt"

	"github.com/stretchr/testify/assert"

	"github.com/scys12/clean-architecture-golang/usecase/user/module"

	"github.com/scys12/clean-architecture-golang/model"

	"github.com/stretchr/testify/mock"

	"github.com/scys12/clean-architecture-golang/pkg/payload/request"
	rMock "github.com/scys12/clean-architecture-golang/repository/role/mocks"
	uMock "github.com/scys12/clean-architecture-golang/repository/user/mocks"
)

func TestAuthenticateUser(t *testing.T) {
	tts := []struct {
		name               string
		expectedLoginReq   *request.LoginRequest
		err                error
		expectedUser       *model.User
		expectedResultUser *model.User
	}{
		{
			name:               "Failed No account Found",
			expectedLoginReq:   &request.LoginRequest{Password: "abc", Username: "abc"},
			err:                errors.New("Failed"),
			expectedUser:       nil,
			expectedResultUser: nil,
		},
		{
			name:               "Failed Password Is Not Same",
			expectedLoginReq:   &request.LoginRequest{Password: "abc", Username: "abc"},
			err:                nil,
			expectedUser:       &model.User{Password: "abcdef", Username: "abc"},
			expectedResultUser: nil,
		},
		{
			name:               "Success Login",
			expectedLoginReq:   &request.LoginRequest{Password: "abc", Username: "abc"},
			err:                nil,
			expectedUser:       &model.User{Password: "abc", Username: "abc"},
			expectedResultUser: &model.User{Password: "abc", Username: "abc"},
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedUser != nil {
				password, err := bcrypt.GenerateFromPassword([]byte(tt.expectedUser.Password), bcrypt.DefaultCost)
				assert.NoError(t, err)
				tt.expectedUser.Password = string(password)
				if tt.expectedResultUser != nil {
					tt.expectedResultUser.Password = string(password)
				}
			}
			mockUserRepo := new(uMock.Repository)
			mockUserRepo.On("GetUserAuthenticateData", mock.Anything, tt.expectedLoginReq.Username).Return(tt.expectedUser, tt.err)
			mockRoleRepo := new(rMock.Repository)
			userUC := module.New(mockUserRepo, mockRoleRepo)
			actualUser, _ := userUC.AuthenticateUser(context.TODO(), tt.expectedLoginReq)
			assert.Equal(t, tt.expectedResultUser, actualUser)
		})
	}
}

func TestRegisterUser(t *testing.T) {
	tts := []struct {
		name                string
		expectedRegisterReq *request.RegisterRequest
		expectedUser        model.User
		err                 error
	}{
		{
			name:                "Success Register",
			expectedRegisterReq: &request.RegisterRequest{Password: "abc", Username: "abc"},
			expectedUser: model.User{
				Username: "abc",
				Password: "abc",
				Role:     model.Role{ID: primitive.NewObjectID(), Name: model.ROLE_USER},
			},
			err: nil,
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			pwd, err := bcrypt.GenerateFromPassword([]byte(tt.expectedRegisterReq.Password), bcrypt.DefaultCost)
			if err != nil {
				log.Fatal(err)
			}
			tt.expectedUser.Password = string(pwd)
			mockUserRepo := new(uMock.Repository)
			mockRoleRepo := new(rMock.Repository)
			mockUserRepo.On("RegisterUser", mock.Anything, mock.Anything).Return(tt.err)
			mockRoleRepo.On("GetRoleByName", mock.Anything, mock.Anything).Return(&tt.expectedUser.Role, nil)
			userUC := module.New(mockUserRepo, mockRoleRepo)
			err = userUC.RegisterUser(context.TODO(), tt.expectedRegisterReq)
			assert.Equal(t, err, tt.err)
		})
	}
}
