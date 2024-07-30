package book

type Storage interface {
	GetBooks() ([]*Book, error)
	GetBook(id int) (*Book, error)
	CreateBook(b *Book) error
	UpdateBook(id int, b *Book) error
	DeleteBook(id int) error
}
