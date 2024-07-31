package models

import "time"

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	RateScore   float32   `json:"rate_score"`
	RateCount   int       `json:"rate_count"`
	Url         string    `json:"url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Author      Author    `json:"author"`
	AuthorID    int       `json:"author_id"`
	Geners      []Genre   `json:"genres"`
}

const (
	CreateTableBooks string = `CREATE TABLE IF NOT EXISTS books (
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
	SelectAllBooks string = `SELECT id, title, description, rate_scrore, rate_count, url, created_at, updated_at, author_id FROM books`
	SelectBook     string = `SELECT id, title, description, rate_scrore, rate_count, url, created_at, updated_at, author_id FROM books WHERE id = ?`
	InsertBook     string = `INSERT INTO books (title, description, rate_scrore, rate_count, url, created_at, updated_at, author_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	UpdateBook     string = `UPDATE books SET title = ?, description = ?, rate_score = ?, rate_count = ?, created_at, uptaded_at, author_id = ? WHERE id = ?`
	DeleteBook     string = `DELETE FROM books WHERE id = ?`
)

// ToArgs returns title, description, rate_scrore, rate_count, url, created_at, updated_at, author_id as value
// use for inserting to DB
func (b *Book) ToArgs() []interface{} {
	return []interface{}{b.Title, b.Description, b.RateScore, b.RateCount, b.Url, b.CreatedAt, b.UpdatedAt, b.AuthorID}
}

// ToUpdatedArgs returns title, description, rate_scrore, rate_count, url, created_at, updated_at, author_id and id as value
// use for updating record in DB
func (b *Book) ToUpdatedArgs(id int) []interface{} {
	return []interface{}{b.Title, b.Description, b.RateScore, b.RateCount, b.Url, b.CreatedAt, b.UpdatedAt, b.AuthorID, id}
}

// ToFeilds returns id, title, description, rate_scrore, rate_count, url, created_at, updated_at, author_id as reference
// use for scanning from DB
func (b *Book) ToFeilds() []interface{} {
	return []interface{}{&b.ID, &b.Title, &b.Description, &b.RateScore, &b.RateCount, &b.Url, &b.CreatedAt, &b.UpdatedAt, &b.AuthorID}
}