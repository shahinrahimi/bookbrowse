package stores

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shahinrahimi/bookbrowse/models"
)

type Storage interface {
	// books methods

	GetBooks() ([]*models.Book, error)
	GetBook(id int) (*models.Book, error)
	CreateBook(b *models.Book) error
	UpdateBook(id int, b *models.Book) error
	DeleteBook(id int) error

	// authors methods
	GetAuthors() ([]*models.Author, error)
	GetAuthor(id int) (*models.Author, error)
	CreateAuthor(a *models.Author) error
	UpdateAuthor(id int, a *models.Author) error
	DeleteAuthor(id int) error

	// genres methods
	GetGenres() ([]*models.Genre, error)
	GetGenre(id int) (*models.Genre, error)
	CreateGenre(g *models.Genre) error
	UpdateGenre(id int, g *models.Genre) error
	DeleteGenre(id int) error
}

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
	if _, err := s.db.Exec(models.CreateTableBooks); err != nil {
		s.logger.Printf("error creating books table: %v", err)
		return err
	}
	if _, err := s.db.Exec(models.CreateTableAuthors); err != nil {
		s.logger.Printf("error creating authors table: %v", err)
		return err
	}
	if _, err := s.db.Exec(models.CreateTableGenres); err != nil {
		s.logger.Printf("error creating genres table: %v", err)
		return err
	}
	if _, err := s.db.Exec(models.CreateTableBookGenres); err != nil {
		s.logger.Printf("error creating book_genres: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) CloseDB() error {
	return s.db.Close()
}