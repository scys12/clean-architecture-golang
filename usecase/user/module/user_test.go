package module_test

import (
	"context"
	"errors"
	"log"
	"testing"

	"github.com/google/uuid"

	"github.com/scys12/clean-architecture-golang/pkg/aws/mocks"
	sessMocks "github.com/scys12/clean-architecture-golang/pkg/session/mocks"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/scys12/clean-architecture-golang/usecase/user"

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
	session := uuid.New().String()
	tts := []struct {
		name                string
		expectedLoginReq    *request.LoginRequest
		err                 error
		session             string
		errSession          error
		expectedUserAuth    *model.UserAuth
		expectedUserProfile *model.UserProfile
		expectedResponse    *user.AuthenticateResponse
	}{
		{
			name:                "Failed No account Found",
			expectedLoginReq:    &request.LoginRequest{Password: "abc", Username: "abc"},
			err:                 errors.New("Failed"),
			errSession:          nil,
			expectedUserAuth:    nil,
			session:             "",
			expectedUserProfile: nil,
			expectedResponse:    nil,
		},
		{
			name:                "Failed Password Is Not Same",
			expectedLoginReq:    &request.LoginRequest{Password: "abc", Username: "abc"},
			err:                 nil,
			errSession:          nil,
			session:             "",
			expectedUserAuth:    &model.UserAuth{Password: "abcdef", Username: "abc"},
			expectedUserProfile: &model.UserProfile{},
			expectedResponse:    nil,
		},
		{
			name:                "Success Login",
			expectedLoginReq:    &request.LoginRequest{Password: "abc", Username: "abc"},
			err:                 nil,
			errSession:          nil,
			session:             session,
			expectedUserAuth:    &model.UserAuth{Password: "abc", Username: "abc"},
			expectedUserProfile: &model.UserProfile{},
			expectedResponse: &user.AuthenticateResponse{
				Response: &user.Response{
					Username: "abc",
				},
				SessionID: session,
			},
		},
		{
			name:                "Error Create Session",
			expectedLoginReq:    &request.LoginRequest{Password: "abc", Username: "abc"},
			err:                 nil,
			errSession:          errors.New("failed"),
			session:             session,
			expectedUserAuth:    &model.UserAuth{Password: "abc", Username: "abc"},
			expectedUserProfile: &model.UserProfile{},
			expectedResponse:    nil,
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			filter := make(map[string]interface{})
			filter["username"] = tt.expectedLoginReq.Username

			if tt.expectedUserAuth != nil {
				password, err := bcrypt.GenerateFromPassword([]byte(tt.expectedUserAuth.Password), bcrypt.DefaultCost)
				assert.NoError(t, err)
				tt.expectedUserAuth.Password = string(password)
			}

			mockAWS := new(mocks.S3Store)
			mockUserRepo := new(uMock.Repository)
			mockUserRepo.On("GetUserAuthenticateData", mock.Anything, filter).Return(tt.expectedUserAuth, tt.expectedUserProfile, tt.err)
			mockRoleRepo := new(rMock.Repository)
			mockSession := new(sessMocks.SessionStore)
			mockSession.On("CreateSession", tt.expectedUserAuth).Return(tt.session, tt.errSession)

			userUC := module.New(mockUserRepo, mockRoleRepo, mockAWS, mockSession)

			userResponse, _ := userUC.AuthenticateUser(context.TODO(), tt.expectedLoginReq)
			assert.Equal(t, tt.expectedResponse, userResponse)
		})
	}
}

func TestRegisterUser(t *testing.T) {
	tts := []struct {
		name                string
		expectedRegisterReq *request.RegisterRequest
		expectedUser        model.UserAuth
		err                 error
	}{
		{
			name:                "Success Register",
			expectedRegisterReq: &request.RegisterRequest{Password: "abc", Username: "abc"},
			expectedUser: model.UserAuth{
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

			mockAWS := new(mocks.S3Store)
			mockUserRepo := new(uMock.Repository)
			mockRoleRepo := new(rMock.Repository)
			mockSession := new(sessMocks.SessionStore)

			mockUserRepo.On("RegisterUser", mock.Anything, mock.Anything).Return(tt.err)
			mockRoleRepo.On("GetRoleByName", mock.Anything, mock.Anything).Return(&tt.expectedUser.Role, nil)

			userUC := module.New(mockUserRepo, mockRoleRepo, mockAWS, mockSession)
			err = userUC.RegisterUser(context.TODO(), tt.expectedRegisterReq)
			assert.Equal(t, err, tt.err)
		})
	}
}

func TestEditProfile(t *testing.T) {
	ID := primitive.NewObjectID()
	tts := []struct {
		name                string
		expectedProfileReq  *request.ProfileRequest
		expectedUserAuth    *model.UserAuth
		expectedUserProfile *model.UserProfile
		errNoUserData       error
		errUserProfile      error
	}{
		{
			name:               "Success Register",
			expectedProfileReq: &request.ProfileRequest{ID: ID, Location: "abc", Name: "abc"},
			expectedUserAuth: &model.UserAuth{
				Username: "abc",
				Password: "abc",
			},
			expectedUserProfile: &model.UserProfile{
				ID:       ID,
				Location: "abc",
				Name:     "abc",
			},
			errNoUserData:  nil,
			errUserProfile: nil,
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mockAWS := new(mocks.S3Store)
			mockUserRepo := new(uMock.Repository)
			mockRoleRepo := new(rMock.Repository)
			mockSession := new(sessMocks.SessionStore)

			filter := make(map[string]interface{})
			filter["_id"] = tt.expectedProfileReq.ID
			mockUserRepo.On("GetUserAuthenticateData", mock.Anything, filter).Return(tt.expectedUserAuth, tt.expectedUserProfile, tt.errNoUserData)
			mockUserRepo.On("EditUserProfile", mock.Anything, *tt.expectedUserProfile).Return(tt.errUserProfile)
			mockAWS.On("UploadFileToS3", mock.Anything).Return("", nil)

			userUC := module.New(mockUserRepo, mockRoleRepo, mockAWS, mockSession)
			actualResponse, _ := userUC.EditUserProfile(context.TODO(), tt.expectedProfileReq)
			assert.Equal(t, actualResponse.Location, tt.expectedUserProfile.Location)
		})
	}
}

func TestGetUserProfile(t *testing.T) {
	tts := []struct {
		name                string
		username            string
		response            *user.Response
		err                 error
		expectedUserAuth    *model.UserAuth
		expectedUserProfile *model.UserProfile
	}{
		{
			name:     "Success get user profile",
			username: "testing",
			response: &user.Response{
				Username: "testing",
			},
			err: nil,
			expectedUserAuth: &model.UserAuth{
				Username: "testing",
			},
			expectedUserProfile: &model.UserProfile{},
		},
		{
			name:                "Failed get user profile",
			username:            "",
			response:            nil,
			err:                 errors.New("failed"),
			expectedUserAuth:    &model.UserAuth{},
			expectedUserProfile: &model.UserProfile{},
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mockAWS := new(mocks.S3Store)
			mockUserRepo := new(uMock.Repository)
			mockRoleRepo := new(rMock.Repository)
			mockSession := new(sessMocks.SessionStore)

			filter := make(map[string]interface{})
			filter["username"] = tt.username
			mockUserRepo.On("GetUserAuthenticateData", mock.Anything, filter).Return(tt.expectedUserAuth, tt.expectedUserProfile, tt.err)
			userUC := module.New(mockUserRepo, mockRoleRepo, mockAWS, mockSession)
			actualResponse, _ := userUC.GetUserProfile(context.TODO(), tt.username)
			assert.Equal(t, actualResponse, tt.response)
		})
	}
}

func TestLogout(t *testing.T) {
	session := uuid.New().String()
	userID := uuid.New().String()

	tts := []struct {
		name      string
		err       error
		sessionID string
		userID    string
	}{
		{
			name:      "Success Logout",
			err:       nil,
			sessionID: session,
			userID:    userID,
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mockAWS := new(mocks.S3Store)
			mockUserRepo := new(uMock.Repository)
			mockRoleRepo := new(rMock.Repository)
			mockSession := new(sessMocks.SessionStore)
			mockSession.On("Del", mock.Anything).Return(tt.err)
			userUC := module.New(mockUserRepo, mockRoleRepo, mockAWS, mockSession)
			err := userUC.Logout(context.TODO(), tt.sessionID, tt.userID)
			assert.Equal(t, err, tt.err)
		})
	}
}
