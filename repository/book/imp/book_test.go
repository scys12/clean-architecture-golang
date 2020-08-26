package imp_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	repo "github.com/scys12/simple-api-go/repository/book/imp"
	mock "github.com/scys12/simple-api-go/repository/book/mocks"
)

func TestInsert(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	repoMock := mock.NewMockRepository(controller)

	repoMock.EXPECT().InsertBook(gomock.Any())

	u := repo.New(repoMock)

	t.Run("insert is valid", func(t *testing.T) {
		result := u.InsertBook(gomock.Any())
		if result != nil {
			t.Fail()
		}
	})
}
