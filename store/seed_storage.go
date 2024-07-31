package store

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strings"

	"github.com/shahinrahimi/bookbrowse/pkg/book"
)

func (s *SqliteStore) MainSeed() {
	path := "books.csv"
	authors := s.getUniqueAuthors(path)
	genres := s.getUniqueGenres(path)
	s.logger.Println("authors length is", len(authors))
	s.logger.Println("genres length is", len(genres))

	// for _, item := range genres {
	// 	var g genre.Genre
	// 	g.Name = item
	// 	if err := s.CreateGenre(&g); err != nil {
	// 		s.logger.Fatal(err)
	// 	}
	// }

	// for _, item := range authors {
	// 	var a author.Author
	// 	a.Name = item
	// 	if err := s.CreateAuthor(&a); err != nil {
	// 		s.logger.Fatal(err)
	// 	}
	// }
}

func (s *SqliteStore) Seed() {

	file, err := os.Open("books.csv")
	if err != nil {
		s.logger.Printf("error opening the file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var allgenres []string
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				s.logger.Println("end of file reached")
				break
			}
			s.logger.Fatalf("error reading a csv: %v", err)
		}

		bookrow := record[0]
		// bookTitle := record[1]
		// authorName := record[2]
		// bookDesc := record[3]
		bookCatg := record[4]
		// bookRate := record[5]
		// bookcount := record[6]
		// bookurl := record[7]
		var genres []string

		if bookrow != "" {
			s.logger.Println(bookCatg)
			s.logger.Printf("%v", bookCatg)
			cleanGen := strings.ReplaceAll(bookCatg, "'", "\"")
			s.logger.Println(cleanGen)
			if err := json.Unmarshal([]byte(cleanGen), &genres); err != nil {
				s.logger.Fatal(err)
			}
			// s.logger.Println(bookrow, bookTitle, authorName, bookCatg)
			s.logger.Println(genres)
			for _, g := range genres {
				if !contains(allgenres, g) {
					allgenres = append(allgenres, g)
				}
			}
		}

	}
	s.logger.Println(allgenres)
	s.logger.Println("end")

}

func (s *SqliteStore) getUniqueGenres(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		s.logger.Printf("error opening the file: %v", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	var genres []string
	for {
		var rowGenres []string
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				s.logger.Println("end of file reached")
				break
			}
			s.logger.Fatalf("error reading a csv: %v", err)
		}
		rawRow := record[0]
		rawGenres := record[4]
		if rawRow != "" {
			cleanRawGenres := strings.ReplaceAll(rawGenres, "'", "\"")
			if err := json.Unmarshal([]byte(cleanRawGenres), &rowGenres); err != nil {
				s.logger.Fatal(err)
			}
			for _, rg := range rowGenres {
				g := strings.ToLower(rg)
				if !contains(genres, g) {
					genres = append(genres, g)
				}
			}
		}
	}
	return genres
}

func (s *SqliteStore) getUniqueAuthors(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		s.logger.Printf("error opening the file: %v", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	var authors []string
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				s.logger.Println("end of file reached")
				break
			}
			s.logger.Fatalf("error reading a csv: %v", err)
		}
		rawRow := record[0]
		rawAuthor := record[2]
		if rawRow != "" {
			a := strings.ToLower(rawAuthor)
			if !contains(authors, a) {
				authors = append(authors, a)
			}
		}

	}
	return authors
}

func (s *SqliteStore) getBooks(path string) []*book.Book {
	file, err := os.Open(path)
	if err != nil {
		s.logger.Printf("error opening the file: %v", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	var books []*book.Book
	for {
		_, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				s.logger.Println("end of file reached")
				break
			}
			s.logger.Fatalf("error reading a csv: %v", err)
		}

		// rowRaw := record[0]
		// title := record[1]
		// authorName := record[2]
		// description := record[3]
		// rawGenres := record[4]
		// rateScoreStr := record[5]
		// rateCountStr := record[6]
		// url := record[7]
		// if rowRaw != "" {
		// 	var b book.Book
		// 	b.Title = title
		// 	b.Au = authorName
		// }

	}

	return books
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
