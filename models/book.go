package models

import (
	"strings"
	"time"
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	RateScore   float32   `json:"rate_score" validate:"required, gt=0"`
	RateCount   int       `json:"rate_count" validate:"required, gt=0"`
	Url         string    `json:"url" validate:"required"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	AuthorName  string    `json:"author_name"`
	AuthorID    int       `json:"author_id" validate:"required"`
	Geners      []string  `json:"genres"`
}

type Books []*Book

type KeyBook struct{}

const (
	CREATE_TABLE_BOOKS string = `CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT UNIQUE NOT NULL,
		description TEXT NOT NULL,
		rate_score REAL,
		rate_count INTEGER,
		url TEXT,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL,
		author_id INTEGER,
		FOREIGN KEY (author_id) REFERENCES authors(id)
	);`
	SELECT_COUNT_BOOKS               string = `SELECT COUNT(*) FROM books`
	SELECT_COUNT_BOOKS_WITH_AUTHORID string = `SELECT COUNT(*) FROM books WHERE author_id = ?`
	SELECT_BOOKS                     string = `SELECT id, title, description, rate_score, rate_count, url, created_at, updated_at, author_id FROM books`
	SELECT_BOOK                      string = `SELECT id, title, description, rate_score, rate_count, url, created_at, updated_at, author_id FROM books WHERE id = ?`
	INSERT_BOOK                      string = `INSERT INTO books (title, description, rate_score, rate_count, url, created_at, updated_at, author_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	UPDATE_BOOK                      string = `UPDATE books SET title = ?, description = ?, rate_score = ?, rate_count = ?, created_at, uptaded_at, author_id = ? WHERE id = ?`
	DELETE_BOOK                      string = `DELETE FROM books WHERE id = ?`
)

// ToArgs returns title, description, rate_score, rate_count, url, created_at, updated_at, author_id as value
// use for inserting to DB
func (b *Book) ToArgs() []interface{} {
	return []interface{}{b.Title, b.Description, b.RateScore, b.RateCount, b.Url, b.CreatedAt, b.UpdatedAt, b.AuthorID}
}

// ToUpdatedArgs returns title, description, rate_score, rate_count, url, created_at, updated_at, author_id and id as value
// use for updating record in DB
func (b *Book) ToUpdatedArgs(id int) []interface{} {
	return []interface{}{b.Title, b.Description, b.RateScore, b.RateCount, b.Url, b.CreatedAt, b.UpdatedAt, b.AuthorID, id}
}

// ToFeilds returns id, title, description, rate_score, rate_count, url, created_at, updated_at, author_id as reference
// use for scanning from DB
func (b *Book) ToFeilds() []interface{} {
	return []interface{}{&b.ID, &b.Title, &b.Description, &b.RateScore, &b.RateCount, &b.Url, &b.CreatedAt, &b.UpdatedAt, &b.AuthorID}
}

// GetTitles returns slice of strings that contains Titles
func (bs *Books) GetTitles() []string {
	titles := make([]string, len(*bs))
	for i, b := range *bs {
		titles[i] = b.Title
	}
	return titles
}

// GetID returns ID of genre if found in the genres
// if not found will return -1
func (bs *Books) GetID(title string) int {
	// make sure the name is caseinsensetiv and trimed space
	cleanTitle := strings.TrimSpace(strings.ToLower(title))

	for _, b := range *bs {
		if b.Title == cleanTitle {
			return b.ID
		}
	}
	return -1
}

func (bs *Books) Add(b *Book) {
	*bs = append(*bs, b)
}
