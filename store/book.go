package store

import "github.com/shahinrahimi/bookbrowse/pkg/book"

func (s *SqliteStore) GetBooks() ([]*book.Book, error) {
	return nil, nil
}

func (s *SqliteStore) GetBook(id int) (*book.Book, error) {
	return nil, nil
}

func (s *SqliteStore) CreateBook(b *book.Book) error {
	return nil
}

func (s *SqliteStore) UpdateBook(id int, b *book.Book) error {
	return nil
}

func (s *SqliteStore) DeleteBook(id int) error {
	return nil
}
