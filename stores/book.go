package stores

import (
	"database/sql"

	"github.com/shahinrahimi/bookbrowse/models"
)

func (s *SqliteStore) GetBooks() (*models.Books, error) {
	rows, err := s.db.Query(models.SELECT_BOOKS_JOIN_AUTHORS)
	if err != nil {
		s.logger.Printf("Error scranning rows for books: %v", err)
		return nil, err
	}
	defer rows.Close()

	var books models.Books
	for rows.Next() {
		var b models.Book
		var a models.Author
		if err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.RateScore, &b.RateCount, &b.Url, &a.ID, &a.Name); err != nil {
			s.logger.Printf("Error scranning rows for a book: %v", err)
			return nil, err
		}
		b.AuthorName = a.Name
		b.AuthorID = a.ID
		// get genres
		genreRows, err := s.db.Query(models.SELECT_GENRES_JOIN_BOOKGENRES, b.ID)
		if err != nil {
			s.logger.Printf("Error scranning rows for genres: %v", err)
			return nil, err
		}
		defer genreRows.Close()
		var genres models.Genres
		for genreRows.Next() {
			var g models.Genre
			if err := genreRows.Scan(g.ToFeilds()...); err != nil {
				s.logger.Printf("Error scranning rows for genres: %v", err)
				continue
			}
			genres = append(genres, &g)
		}
		b.Geners = genres.GetNames()
		books = append(books, &b)
	}

	// Check for errors encountered during iteration
	if err := rows.Err(); err != nil {
		s.logger.Printf("Error encountered during row iteration: %v", err)
		return nil, err
	}

	return &books, nil
}

func (s *SqliteStore) GetLimitedBooks(limit int, offset int) (*models.Books, error) {
	rows, err := s.db.Query(models.SELECT_LIMITEDBOOKS_JOIN_AUTHORS, limit, offset)
	if err != nil {
		s.logger.Printf("Error scranning rows for books: %v", err)
		return nil, err
	}
	defer rows.Close()

	var books models.Books
	for rows.Next() {
		var b models.Book
		var a models.Author
		if err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.RateScore, &b.RateCount, &b.Url, &a.ID, &a.Name); err != nil {
			s.logger.Printf("Error scranning rows for a book: %v", err)
			return nil, err
		}
		b.AuthorName = a.Name
		b.AuthorID = a.ID
		// get genres
		genreRows, err := s.db.Query(models.SELECT_GENRES_JOIN_BOOKGENRES, b.ID)
		if err != nil {
			s.logger.Printf("Error scranning rows for genres: %v", err)
			return nil, err
		}
		defer genreRows.Close()
		var genres models.Genres
		for genreRows.Next() {
			var g models.Genre
			if err := genreRows.Scan(g.ToFeilds()...); err != nil {
				s.logger.Printf("Error scranning rows for genres: %v", err)
				continue
			}
			genres = append(genres, &g)
		}
		b.Geners = genres.GetNames()
		books = append(books, &b)
	}

	// Check for errors encountered during iteration
	if err := rows.Err(); err != nil {
		s.logger.Printf("Error encountered during row iteration: %v", err)
		return nil, err
	}

	return &books, nil
}

func (s *SqliteStore) GetBook(id int) (*models.Book, error) {
	var b models.Book
	var a models.Author

	if err := s.db.QueryRow(models.SELECT_BOOK_JOIN_AUTHOR, id).Scan(&b.ID, &b.Title, &b.Description, &b.RateScore, &b.RateCount, &b.Url, &a.ID, &a.Name); err != nil {
		s.logger.Printf("Error scranning row for the book: %v", err)
		return nil, err
	}
	b.AuthorName = a.Name
	b.AuthorID = a.ID

	genreRows, err := s.db.Query(models.SELECT_GENRES_JOIN_BOOKGENRES, b.ID)
	if err != nil {
		s.logger.Printf("Error scranning rows for genres: %v", err)
		return nil, err
	}
	defer genreRows.Close()
	var genres models.Genres
	for genreRows.Next() {
		var g models.Genre
		if err := genreRows.Scan(g.ToFeilds()...); err != nil {
			s.logger.Printf("Error scranning rows for genres: %v", err)
			continue
		}
		genres = append(genres, &g)
	}
	b.Geners = genres.GetNames()

	return &b, nil
}

func (s *SqliteStore) CreateBook(b *models.Book) error {
	if _, err := s.db.Exec(models.INSERT_BOOK, b.ToArgs()...); err != nil {
		s.logger.Printf("Error inserting a new book to DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) UpdateBook(id int, b *models.Book) error {
	if _, err := s.db.Exec(models.UPDATE_BOOK, b.ToUpdatedArgs(id)...); err != nil {
		s.logger.Printf("Error updating book from DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) DeleteBook(id int) error {
	if _, err := s.db.Exec(models.DELETE_BOOK, id); err != nil {
		s.logger.Printf("Error deleting book from DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) GetBooksCount() (int, error) {
	var count int
	if err := s.db.QueryRow(models.SELECT_COUNT_BOOKS).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		s.logger.Printf("error selecting count from books: %v", err)
		return -1, err
	}
	return count, nil
}
