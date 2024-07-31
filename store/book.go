package store

import "github.com/shahinrahimi/bookbrowse/pkg/book"

func (s *SqliteStore) GetBooks() ([]*book.Book, error) {
	rows, err := s.db.Query(book.SelectAll)
	if err != nil {
		s.logger.Printf("Error scranning rows for books: %v", err)
		return nil, err
	}
	// initiate books as empty slice
	books := []*book.Book{}
	for rows.Next() {
		var b book.Book
		if err := rows.Scan(b.ToFeilds()...); err != nil {
			s.logger.Printf("Error scranning rows for a book: %v", err)
			continue
		}
		books = append(books, &b)
	}
	return books, nil
}

func (s *SqliteStore) GetBook(id int) (*book.Book, error) {
	var b book.Book
	if err := s.db.QueryRow(book.Select, id).Scan(b.ToFeilds()...); err != nil {
		s.logger.Printf("Error scranning row for the book: %v", err)
		return nil, err
	}
	return &b, nil
}

func (s *SqliteStore) CreateBook(b *book.Book) error {
	if _, err := s.db.Exec(book.Insert, b.ToArgs()...); err != nil {
		s.logger.Printf("Error inserting a new book to DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) UpdateBook(id int, b *book.Book) error {
	if _, err := s.db.Exec(book.Update, b.ToUpdatedArgs(id)...); err != nil {
		s.logger.Printf("Error updating book from DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) DeleteBook(id int) error {
	if _, err := s.db.Exec(book.Delete, id); err != nil {
		s.logger.Printf("Error deleting book from DB: %v", err)
		return err
	}
	return nil
}
