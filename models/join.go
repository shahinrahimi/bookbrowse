package models

const (
	SELECT_BOOKS_JOIN_AUTHORS string = `
        SELECT books.id, books.title, books.description, books.rate_score, books.rate_count, books.url, authors.id, authors.name
		FROM books
		JOIN authors ON books.author_id = authors.id`
	SELECT_BOOK_JOIN_AUTHOR string = `
		SELECT books.id, books.title, books.description, books.rate_score, books.rate_count, books.url, authors.id, authors.name
		FROM books
		JOIN authors ON books.author_id = authors.id
		WHERE books.id = ?`
	SELECT_LIMITEDBOOKS_JOIN_AUTHORS string = `
		SELECT books.id, books.title, books.description, books.rate_score, books.rate_count, books.url, authors.id, authors.name
		FROM books
		JOIN authors ON books.author_id = authors.id
		LIMIT ? OFFSET ?`
	SELECT_GENRES_JOIN_BOOKGENRES string = `
		SELECT g.id, g.name 
		FROM genres g 
		JOIN book_genres bg ON g.id = bg.genre_id 
		WHERE bg.book_id = ?`
)
