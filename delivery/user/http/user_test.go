package http_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/scys12/clean-architecture-golang/pkg/validator"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/scys12/clean-architecture-golang/usecase/user"

	"github.com/scys12/clean-architecture-golang/pkg/payload/request"

	"github.com/labstack/echo/v4"
	uHttp "github.com/scys12/clean-architecture-golang/delivery/user/http"
	sessMocks "github.com/scys12/clean-architecture-golang/pkg/session/mocks"
	"github.com/scys12/clean-architecture-golang/usecase/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthenticate(t *testing.T) {
	tts := []struct {
		name                 string
		loginRequest         *request.LoginRequest
		expectedMockResponse *user.AuthenticateResponse
		err                  error
		resultCode           int
		requestBody          string
	}{
		{
			name: "Success Authenticate",
			requestBody: `
				{
					"username":"def",
					"password":"abc"
				}
			`,
			loginRequest: &request.LoginRequest{
				Password: "abc",
				Username: "def",
			},
			expectedMockResponse: &user.AuthenticateResponse{
				Response: &user.Response{
					Username: "def",
				},
			},
			err:        nil,
			resultCode: http.StatusOK,
		},
		{
			name: "Bad Request Body",
			requestBody: `
				{
					"username":"def",
					"password":
				}
			`,
			loginRequest: &request.LoginRequest{
				Username: "def",
			},
			expectedMockResponse: &user.AuthenticateResponse{
				Response: &user.Response{
					Username: "def",
				},
			},
			err:        errors.New("Bad Request Body"),
			resultCode: http.StatusInternalServerError,
		},
		{
			name: "Failed Authenticate",
			requestBody: `
				{
					"username":"def",
					"password":"klm"
				}
			`,
			loginRequest: &request.LoginRequest{
				Username: "def",
				Password: "klm",
			},
			expectedMockResponse: &user.AuthenticateResponse{
				Response: &user.Response{
					Username: "def",
				},
			},
			err:        errors.New("Account Not Found"),
			resultCode: http.StatusNotFound,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mockUserUCase := new(mocks.Usecase)
			mockUserUCase.On("AuthenticateUser", mock.Anything, tt.loginRequest).Return(tt.expectedMockResponse, tt.err)

			mockSession := new(sessMocks.SessionStore)

			req, err := http.NewRequest(echo.POST, "/auth/signin", strings.NewReader(tt.requestBody))
			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e := echo.New()
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			handler := uHttp.New(mockUserUCase)
			uHttp.SetRoute(e, handler, mockSession)

			_ = handler.AuthenticateUser(c)
			assert.Equal(t, tt.resultCode, rec.Code)
		})
	}
}

func TestRegisterUser(t *testing.T) {
	tts := []struct {
		name        string
		err         error
		resultCode  int
		registerReq *request.RegisterRequest
		requestBody string
	}{
		{
			name:       "Success Authenticate",
			err:        nil,
			resultCode: http.StatusOK,
			registerReq: &request.RegisterRequest{
				Email:    "sam@gmail.com",
				Username: "abcdef",
				Password: "abcdef",
			},
			requestBody: `{"username": "abcdef", "password": "abcdef", "email": "sam@gmail.com"}`,
		},
		{
			name:        "Failed Wrong Request Body",
			err:         errors.New("failed"),
			resultCode:  http.StatusInternalServerError,
			registerReq: nil,
			requestBody: `{"username":"abc",}`,
		},
		{
			name:        "Request Body Failed on Validation",
			err:         errors.New("failed"),
			resultCode:  http.StatusInternalServerError,
			registerReq: nil,
			requestBody: `{"username":"abcdef", "password":"ab"}`,
		},
		{
			name:       "Failed Username/Email exist",
			err:        errors.New("failed"),
			resultCode: http.StatusBadRequest,
			registerReq: &request.RegisterRequest{
				Email:    "sam@gmail.com",
				Username: "abcdef",
				Password: "abcdef",
			},
			requestBody: `{"username": "abcdef", "password": "abcdef", "email": "sam@gmail.com"}`,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			e.Validator = validator.New()
			req, err := http.NewRequest(echo.POST, "/auth/register", strings.NewReader(tt.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			assert.NoError(t, err)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			mockSession := new(sessMocks.SessionStore)
			mockUserUCase := new(mocks.Usecase)
			mockUserUCase.On("RegisterUser", c.Request().Context(), tt.registerReq).Return(tt.err)

			handler := uHttp.New(mockUserUCase)
			uHttp.SetRoute(e, handler, mockSession)

			_ = handler.RegisterUser(c)
			assert.Equal(t, tt.resultCode, rec.Code)
		})
	}
}

func TestEditProfile(t *testing.T) {
	tts := []struct {
		name             string
		err              error
		resultCode       int
		profileReq       *request.ProfileRequest
		requestBody      url.Values
		expectedResponse *user.Response
	}{
		{
			name:       "Success Edit Profile",
			err:        nil,
			resultCode: http.StatusOK,
			profileReq: &request.ProfileRequest{
				ID:    primitive.NewObjectID(),
				Name:  "foo",
				Phone: "bar",
			},
			requestBody: url.Values{
				"name":  []string{"foo"},
				"phone": []string{"bar"},
			},
			expectedResponse: &user.Response{
				Name:  "foo",
				Phone: "bar",
			},
		},
		{
			name:             "No Form Params, Failed On Validation",
			err:              errors.New("failed"),
			resultCode:       http.StatusInternalServerError,
			profileReq:       &request.ProfileRequest{},
			requestBody:      nil,
			expectedResponse: nil,
		},
		{
			name:       "Failed Update Profile",
			err:        errors.New("failed"),
			resultCode: http.StatusInternalServerError,
			profileReq: &request.ProfileRequest{
				Name:  "foo",
				Phone: "bar",
			},
			requestBody: url.Values{
				"name":  []string{"foo"},
				"phone": []string{"bar"},
			},
			expectedResponse: nil,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			e.Validator = validator.New()

			req, err := http.NewRequest(echo.PUT, "/user/profile", strings.NewReader(tt.requestBody.Encode()))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
			assert.NoError(t, err)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set("userID", tt.profileReq.ID)

			mockSession := new(sessMocks.SessionStore)
			mockUserUCase := new(mocks.Usecase)
			mockUserUCase.On("EditUserProfile", c.Request().Context(), tt.profileReq).Return(tt.expectedResponse, tt.err)

			handler := uHttp.New(mockUserUCase)
			uHttp.SetRoute(e, handler, mockSession)

			_ = handler.EditUserProfile(c)
			assert.Equal(t, tt.resultCode, rec.Code)
		})
	}
}

func TestGetUserProfile(t *testing.T) {
	tts := []struct {
		name               string
		username           string
		response           *user.Response
		expectedResultCode int
		err                error
	}{
		{
			name:     "success get another user profile",
			username: "testing",
			response: &user.Response{
				Username: "testing",
			},
			expectedResultCode: http.StatusOK,
			err:                nil,
		},
		{
			name:     "failed get another user profile",
			username: "",
			response: &user.Response{
				Username: "",
			},
			expectedResultCode: http.StatusInternalServerError,
			err:                errors.New("failed"),
		},
	}
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req, err := http.NewRequest(echo.GET, "/user/profile/:username", strings.NewReader(""))
			assert.NoError(t, err)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("username")
			c.SetParamValues(tt.username)
			mockSession := new(sessMocks.SessionStore)
			mockUserUCase := new(mocks.Usecase)
			mockUserUCase.On("GetUserProfile", c.Request().Context(), tt.username).Return(tt.response, tt.err)
			handler := uHttp.New(mockUserUCase)
			uHttp.SetRoute(e, handler, mockSession)

			err = handler.GetUserProfile(c)
			assert.Equal(t, tt.expectedResultCode, rec.Code)
		})
	}
}

func TestLogout(t *testing.T) {
	tts := []struct {
		name       string
		err        error
		resultCode int
		sessionID  string
		userID     string
	}{
		{
			name:       "Success Logout",
			err:        nil,
			resultCode: http.StatusOK,
			sessionID:  uuid.New().String(),
			userID:     uuid.New().String(),
		},
		{
			name:       "Failed Logout",
			err:        errors.New("failed"),
			resultCode: http.StatusInternalServerError,
			sessionID:  uuid.New().String(),
			userID:     uuid.New().String(),
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			e.Validator = validator.New()

			req, err := http.NewRequest(echo.POST, "/user/logout", strings.NewReader(""))
			assert.NoError(t, err)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set("userID", tt.userID)
			c.Set("sessionID", tt.sessionID)

			mockSession := new(sessMocks.SessionStore)
			mockUserUCase := new(mocks.Usecase)
			mockUserUCase.On("Logout", c.Request().Context(), tt.sessionID, tt.userID).Return(tt.err)

			handler := uHttp.New(mockUserUCase)
			uHttp.SetRoute(e, handler, mockSession)

			_ = handler.Logout(c)
			assert.Equal(t, tt.resultCode, rec.Code)
		})
	}
}
