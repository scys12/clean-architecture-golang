package module_test

import (
	"context"
	"errors"
	"testing"

	"github.com/scys12/clean-architecture-golang/usecase/category/module"

	"github.com/scys12/clean-architecture-golang/repository/category/mocks"

	"github.com/scys12/clean-architecture-golang/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetCategories(t *testing.T) {
	tts := []struct {
		name                   string
		expectedMockCategories []*model.Category
		err                    error
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
			err: nil,
		},
		{
			name:                   "Failed get Categories",
			expectedMockCategories: nil,
			err:                    errors.New("failed"),
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mockCatRepo := new(mocks.Repository)
			mockCatRepo.On("GetAllCategories", mock.Anything).Return(tt.expectedMockCategories, tt.err)
			u := module.New(mockCatRepo)
			cats, err := u.GetAllCategories(context.TODO())
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.expectedMockCategories, cats)
			mockCatRepo.AssertExpectations(t)
		})
	}
}
