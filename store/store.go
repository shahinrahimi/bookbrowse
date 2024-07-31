package store

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shahinrahimi/bookbrowse/pkg/author"
	"github.com/shahinrahimi/bookbrowse/pkg/book"
	"github.com/shahinrahimi/bookbrowse/pkg/genre"
)

type SqliteStore struct {
	logger *log.Logger
	db     *sql.DB
}

// NewSqliteStore creates 'db' dir if not exist
// and create connection to sqlite database
func NewSqliteStore(logger *log.Logger) *SqliteStore {
	// create directory if not exists
	if err := os.MkdirAll("db", 0755); err != nil {
		logger.Fatalf("Unable to create a directory for DB: %v", err)
	}

	// create connection to db
	db, err := sql.Open("sqlite3", "./db/bookbrowse.db")
	if err != nil {
		logger.Fatalf("Unable to connect to DB: %v", err)
	}
	logger.Println("DB Connected!")

	return &SqliteStore{
		logger: logger,
		db:     db,
	}
}

// NewTestSqliteStore will open connection to DB @ memory
func NewTestSqliteStore(logger *log.Logger) *SqliteStore {
	// create connection to db
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		logger.Fatalf("Unable to connect to DB: %v", err)
	}
	logger.Println("DB Connected!")

	return &SqliteStore{
		logger: logger,
		db:     db,
	}
}

// Init create tables for books, authors and genre if not exists
// if error raised the function will panic
func (s *SqliteStore) Init() error {
	if _, err := s.db.Exec(book.CreateTable); err != nil {
		s.logger.Printf("error creating books table: %v", err)
		return err
	}
	if _, err := s.db.Exec(author.CreateTable); err != nil {
		s.logger.Printf("error creating authors table: %v", err)
		return err
	}
	if _, err := s.db.Exec(genre.CreateTable); err != nil {
		s.logger.Printf("error creating genres table: %v", err)
		return err
	}
	if _, err := s.db.Exec(genre.CreateTableBookGenre); err != nil {
		s.logger.Printf("error creating book_genres: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) CloseDB() error {
	return s.db.Close()
}
