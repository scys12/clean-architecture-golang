package module

import (
	"github.com/scys12/clean-architecture-golang/repository/category"
)

type repository struct {
}

func New() category.Repository {
	return &repository{}
}
