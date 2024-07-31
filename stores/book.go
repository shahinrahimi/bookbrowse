package stores

import "github.com/shahinrahimi/bookbrowse/models"

func (s *SqliteStore) GetBooks() ([]*models.Book, error) {
	rows, err := s.db.Query(models.SelectAllBooks)
	if err != nil {
		s.logger.Printf("Error scranning rows for books: %v", err)
		return nil, err
	}
	// initiate books as empty slice
	books := []*models.Book{}
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(b.ToFeilds()...); err != nil {
			s.logger.Printf("Error scranning rows for a book: %v", err)
			continue
		}
		books = append(books, &b)
	}
	return books, nil
}

func (s *SqliteStore) GetBook(id int) (*models.Book, error) {
	var b models.Book
	if err := s.db.QueryRow(models.SelectBook, id).Scan(b.ToFeilds()...); err != nil {
		s.logger.Printf("Error scranning row for the book: %v", err)
		return nil, err
	}
	return &b, nil
}

func (s *SqliteStore) CreateBook(b *models.Book) error {
	if _, err := s.db.Exec(models.InsertBook, b.ToArgs()...); err != nil {
		s.logger.Printf("Error inserting a new book to DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) UpdateBook(id int, b *models.Book) error {
	if _, err := s.db.Exec(models.UpdateBook, b.ToUpdatedArgs(id)...); err != nil {
		s.logger.Printf("Error updating book from DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) DeleteBook(id int) error {
	if _, err := s.db.Exec(models.DeleteBook, id); err != nil {
		s.logger.Printf("Error deleting book from DB: %v", err)
		return err
	}
	return nil
}
