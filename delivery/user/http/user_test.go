package http_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	uHttp "github.com/scys12/clean-architecture-golang/delivery/user/http"
	"github.com/scys12/clean-architecture-golang/model"
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
			e := echo.New()
			req, err := http.NewRequest(echo.POST, "/auth/signin", strings.NewReader(""))
			assert.NoError(t, err)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			handler := uHttp.New(mockUserUCase)
			uHttp.SetRoute(e, handler)
			err = handler.AuthenticateUser(c)
			assert.Equal(t, tt.resultCode, rec.Code)
			mockUserUCase.AssertExpectations(t)
		})
	}
}
