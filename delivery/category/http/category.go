package http

import (
	"fmt"
	"net/http"
)

func (d *delivery) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := d.usecase.GetAllCategories()
	fmt.Printf("%v\n", categories)
}
