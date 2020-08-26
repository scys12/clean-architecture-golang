package book

import "github.com/scys12/simple-api-go/models/book"

//This is the books repository contract
type Repository interface {
	InsertBook(book.Book) error
	FindAllBooks() ([]book.Book, error)
	FindBooksByID(string) (book.Book, error)
	UpdateBook(book.Book) error
	RemoveBook(string) error
}
