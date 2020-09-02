package category

import (
	"net/http"
)

type Delivery interface {
	GetAllCategories(w http.ResponseWriter, req *http.Request)
}
