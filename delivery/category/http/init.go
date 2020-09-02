package http

import (
	"net/http"

	dCategory "github.com/scys12/clean-architecture-golang/delivery/category"
	uCategory "github.com/scys12/clean-architecture-golang/usecase/category"
)

type delivery struct {
	usecase uCategory.Usecase
}

func New(usecase uCategory.Usecase) dCategory.Delivery {
	handler := &delivery{
		usecase: usecase,
	}
	http.HandleFunc("/categories", handler.GetAllCategories)
	return handler
}
