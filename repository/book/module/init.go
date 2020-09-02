package imp

import (
	bookRepo "github.com/scys12/simple-api-go/repository/book"
	db "github.com/scys12/simple-api-go/repository/mongodb"
)

type repository struct {
	dbHelper db.DatabaseHelper
}

func New(db db.DatabaseHelper) bookRepo.Repository {
	return &repository{
		dbHelper: db,
	}
}
