package stores

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/shahinrahimi/bookbrowse/models"
)

func (s *SqliteStore) Seed() {
	path := "books.csv"
	s.createBooks(path)
}

func (s *SqliteStore) createBooks(path string) {
	file, err := os.Open(path)
	if err != nil {
		s.logger.Printf("error opening the file: %v", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	books, err := s.GetBooks()
	if err != nil {
		s.logger.Fatalf("error getting all books: %v", err)
	}

	authors, err := s.GetAuthors()
	if err != nil {
		s.logger.Fatalf("error getting all authors: %v", err)
	}

	genres, err := s.GetGenres()
	if err != nil {
		s.logger.Fatalf("error getting all genres: %v", err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				s.logger.Println("end of file reached")
				break
			}
			s.logger.Fatalf("error reading a csv: %v", err)
		}

		rowRaw := record[0]
		title := record[1]
		authorName := record[2]
		description := record[3]
		rawGenres := strings.ReplaceAll(record[4], "'", "\"")
		rateScoreStr := record[5]
		rateCountStr := strings.ReplaceAll(strings.ReplaceAll(record[6], ",", ""), "\"", "")
		url := record[7]
		// check book title
		bookID := books.GetID(title)
		if rowRaw != "" && bookID < 0 {
			var b models.Book
			s.logger.Printf("Book title not found in DB: %s", title)
			// check author name
			authorID := authors.GetID(authorName)
			if authorID < 0 {
				var a models.Author
				s.logger.Printf("Author name not found in DB: %s", authorName)
				a.Name = strings.TrimSpace(strings.ToLower(authorName))
				result, err := s.db.Exec(models.InsertAuthor, a.ToArgs()...)
				if err != nil {
					s.logger.Fatalf("error inserting a new author: %v", err)
				}
				// get new ID
				id, err := result.LastInsertId()
				if err != nil {
					s.logger.Fatalf("error getting a new author ID: %v", err)
				}
				a.ID = int(id)
				s.logger.Printf("New author created: %s with ID: %d", authorName, a.ID)
				authors.Add(&a)
			}
			b.AuthorID = authorID

			b.Title = strings.TrimSpace(strings.ToLower(title))
			b.Description = description
			rateScore, err := strconv.ParseFloat(rateScoreStr, 32)
			if err != nil {
				s.logger.Fatalf("error parsing string to float for rateScore: %v", err)
			}
			b.RateScore = float32(rateScore)
			rateCount, err := strconv.Atoi(rateCountStr)
			if err != nil {
				s.logger.Fatalf("error converting string to int for rate count")
			}
			b.RateCount = rateCount
			b.Url = url
			result, err := s.db.Exec(models.InsertBook, b.ToArgs()...)
			if err != nil {
				s.logger.Fatalf("error inserting a new book: %v", err)
			}
			// get new ID
			id, err := result.LastInsertId()
			if err != nil {
				s.logger.Fatalf("error getting a new book ID: %v", err)
			}
			b.ID = int(id)
			bookID = b.ID
			s.logger.Printf("New book created with ID: %d", b.ID)

			var genresString []string
			if err := json.Unmarshal([]byte(rawGenres), &genresString); err != nil {
				s.logger.Fatalf("error parsing rawGenres to slice of string: %v", err)
			}

			for _, genreName := range genresString {
				var g models.Genre
				cleanGenreName := strings.TrimSpace(strings.ToLower(genreName))
				genreID := genres.GetID(cleanGenreName)
				if genreID < 0 {
					s.logger.Printf("Genre name not found in DB: %s", genreName)
					g.Name = cleanGenreName
					result, err := s.db.Exec(models.InsertGenre, g.ToArgs()...)
					if err != nil {
						s.logger.Fatalf("error inserting a new genre: %v", err)
					}
					// get new ID
					id, err := result.LastInsertId()
					if err != nil {
						s.logger.Fatalf("error getting a new genre ID: %v", err)
					}
					g.ID = int(id)
					genreID = g.ID
					s.logger.Printf("New genre created: %s with ID: %d", g.Name, g.ID)
					genres.Add(&g)
				}

				// s.logger.Println(bookID, genreID)
				_, err = s.db.Exec(`INSERT INTO book_genres (book_id, genre_id) VALUES (?, ?)`, bookID, genreID)
				if err != nil {
					s.logger.Fatalf("Error inserting records to book_genres: %v", err)
				}
			}
			books.Add(&b)
		}

	}
}
