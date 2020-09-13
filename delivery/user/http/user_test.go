package http_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/scys12/clean-architecture-golang/pkg/payload/request"

	"github.com/labstack/echo/v4"
	uHttp "github.com/scys12/clean-architecture-golang/delivery/user/http"
	"github.com/scys12/clean-architecture-golang/model"
	sessMocks "github.com/scys12/clean-architecture-golang/pkg/session/mocks"
	"github.com/scys12/clean-architecture-golang/usecase/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthenticate(t *testing.T) {
	tts := []struct {
		name             string
		expectedMockUser *model.User
		err              error
		resultCode       int
	}{
		{
			name:             "Success Authenticate",
			expectedMockUser: &model.User{},
			err:              nil,
			resultCode:       http.StatusOK,
		},
		{
			name:             "Failed Authenticate",
			expectedMockUser: &model.User{},
			err:              errors.New("failed"),
			resultCode:       http.StatusNotFound,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {

			mockUserUCase := new(mocks.Usecase)
			mockUserUCase.On("AuthenticateUser", mock.Anything, mock.Anything).Return(tt.expectedMockUser, tt.err)
			mockSession := new(sessMocks.SessionStore)
			e := echo.New()
			mockSession.On("CreateSession", mock.Anything, tt.expectedMockUser).Return(tt.err)
			req, err := http.NewRequest(echo.POST, "/auth/signin", strings.NewReader(""))
			assert.NoError(t, err)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			handler := uHttp.New(mockUserUCase, mockSession)
			uHttp.SetRoute(e, handler, mockSession)
			err = handler.AuthenticateUser(c)
			assert.Equal(t, tt.resultCode, rec.Code)
			mockUserUCase.AssertExpectations(t)
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
				Username: "abc",
				Password: "abc",
			},
			requestBody: `{"username": "abc", "password": "abc", "email": "sam@gmail.com"}`,
		},
		{
			name:        "Failed Wrong Request Body",
			err:         errors.New("failed"),
			resultCode:  http.StatusInternalServerError,
			registerReq: nil,
			requestBody: `{"username":"abc",}`,
		},
		{
			name:       "Failed Username/Email exist",
			err:        errors.New("failed"),
			resultCode: http.StatusBadRequest,
			registerReq: &request.RegisterRequest{
				Email:    "sam@gmail.com",
				Username: "abc",
				Password: "abc",
			},
			requestBody: `{"username": "abc", "password": "abc", "email": "sam@gmail.com"}`,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req, err := http.NewRequest(echo.POST, "/auth/register", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			assert.NoError(t, err)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			mockSession := new(sessMocks.SessionStore)
			mockUserUCase := new(mocks.Usecase)
			mockUserUCase.On("RegisterUser", c.Request().Context(), tt.registerReq).Return(tt.err)
			handler := uHttp.New(mockUserUCase, mockSession)
			uHttp.SetRoute(e, handler, mockSession)
			err = handler.RegisterUser(c)
			assert.Equal(t, tt.resultCode, rec.Code)
		})
	}
}
