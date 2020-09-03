package module

import (
	"context"
	"time"

	"github.com/scys12/clean-architecture-golang/model"
)

const timeout = 10 * time.Second

func (u *categoryUsecase) GetAllCategories(c context.Context) (cats []*model.Category, err error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	cats, err = u.repo.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}
	return
}
