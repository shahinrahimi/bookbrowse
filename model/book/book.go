package book

import (
	"time"

	"github.com/shahinrahimi/bookbrowse/model/author"
	"github.com/shahinrahimi/bookbrowse/model/genre"
)

type Book struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	RateScore   float32       `json:"rate_score"`
	RateCount   int           `json:"rate_count"`
	Url         string        `json:"url"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Author      author.Author `json:"author"`
	AuthorID    int           `json:"author_id"`
	Geners      []genre.Genre `json:"genres"`
}

const (
	CreateTable string = `CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		rate_score REAL,
		rate_count INTEGER,
		url TEXT,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL,
		author_id INTEGER,
		FOREIGN KEY (author_id) REFERENCES authors(id)
	);`
	SelectAll string = `SELECT id, title, author, description, rate_scrore, rate_count, url, created_at, updated_at, author_id FROM books`
	Select    string = `SELECT id, title, author, description, rate_scrore, rate_count, url, created_at, updated_at, author_id FROM books WHERE id = ?`
	Insert    string = `INSERT INTO books (id, title, author, description, rate_scrore, rate_count, url, created_at, updated_at, author_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	Update    string = `UPDATE books SET title = ?, author = ?, description = ?, rate_score = ?, rate_count = ?, created_at ,uptaded_at, author_id = ? WHERE id = ?`
	Delete    string = `DELETE FROM books WHERE id = ?`
)

// ToArgs returns id, title, author, description, rate_scrore, rate_count, url, created_at, updated_at, author_id as value
func (b *Book) ToArgs() []interface{} {
	return []interface{}{b.ID, b.Title, b.Description, b.RateScore, b.RateCount, b.Url, b.CreatedAt, b.UpdatedAt, b.AuthorID}
}

// ToUpdatedArgs returns title, author, description, rate_scrore, rate_count, url, created_at, updated_at, author_id and id as value
func (b *Book) ToUpdatedArgs(id int) []interface{} {
	return []interface{}{b.Title, b.Description, b.RateScore, b.RateCount, b.Url, b.CreatedAt, b.UpdatedAt, b.AuthorID, id}
}

// ToFeilds returns id, title, author, description, rate_scrore, rate_count, url, created_at, updated_at, author_id as reference
func (b *Book) ToFeilds() []interface{} {
	return []interface{}{&b.ID, &b.Title, &b.Description, &b.RateScore, &b.RateCount, &b.Url, &b.CreatedAt, &b.UpdatedAt, &b.AuthorID}
}
