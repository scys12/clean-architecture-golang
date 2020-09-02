package imp

import (
	model "github.com/scys12/simple-api-go/models/book"
)

func (r *repository) InsertBook(book model.Book) error {
	r.dbHelper.Collection()
	return err
}

func (r *repository) FindAllBooks() ([]model.Book, error) {
	return nil, nil
}

func (r *repository) FindBooksByID(id string) (model.Book, error) {
	book := model.Book{}
	return book, nil
}

func (r *repository) UpdateBook(book model.Book) error {
	return nil
}

func (r *repository) RemoveBook(id string) error {
	return nil
}
