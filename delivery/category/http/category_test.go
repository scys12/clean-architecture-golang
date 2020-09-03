package http_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/scys12/clean-architecture-golang/usecase/category/mocks"

	"github.com/labstack/echo/v4"
	catHttp "github.com/scys12/clean-architecture-golang/delivery/category/http"

	"github.com/scys12/clean-architecture-golang/model"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
	tts := []struct {
		name                   string
		expectedMockCategories []*model.Category
		err                    error
		resultCode             int
	}{
		{
			name: "Success get categories",
			expectedMockCategories: []*model.Category{
				&model.Category{
					ID:   primitive.NewObjectID(),
					Name: "Motherboard",
				},
				&model.Category{
					ID:   primitive.NewObjectID(),
					Name: "Ram",
				},
			},
			err:        nil,
			resultCode: http.StatusOK,
		},
		{
			name:                   "Failed get Categories",
			expectedMockCategories: []*model.Category{},
			err:                    errors.New("failed"),
			resultCode:             http.StatusNotFound,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mockCatUCase := new(mocks.Usecase)
			mockCatUCase.On("GetAllCategories", mock.Anything).Return(tt.expectedMockCategories, tt.err)
			e := echo.New()
			req, err := http.NewRequest(echo.GET, "/categories", strings.NewReader(""))
			assert.NoError(t, err)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			handler := catHttp.New(mockCatUCase)
			catHttp.SetRoute(e, handler)
			err = handler.GetAllCategories(c)
			assert.Equal(t, tt.resultCode, rec.Code)
			mockCatUCase.AssertExpectations(t)
		})
	}
}
