package store

import "github.com/shahinrahimi/bookbrowse/pkg/author"

func (s *SqliteStore) GetAuthors() ([]*author.Author, error) {
	rows, err := s.db.Query(author.SelectAll)
	if err != nil {
		s.logger.Printf("Error scranning rows for authors: %v", err)
		return nil, err
	}
	// initiate authors as empty slice
	authors := []*author.Author{}
	for rows.Next() {
		var a author.Author
		if err := rows.Scan(a.ToFeilds()...); err != nil {
			s.logger.Printf("Error scranning rows for a author: %v", err)
			continue
		}
		authors = append(authors, &a)
	}
	return authors, nil
}

func (s *SqliteStore) GetAuthor(id int) (*author.Author, error) {
	var a author.Author
	if err := s.db.QueryRow(author.Select, id).Scan(a.ToFeilds()...); err != nil {
		s.logger.Printf("Error scranning row for the author: %v", err)
		return nil, err
	}
	return &a, nil
}

func (s *SqliteStore) CreateAuthor(a *author.Author) error {
	if _, err := s.db.Exec(author.Insert, a.ToArgs()...); err != nil {
		s.logger.Printf("Error inserting a new author to DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) UpdateAuthor(id int, a *author.Author) error {
	if _, err := s.db.Exec(author.Update, a.ToUpdatedArgs(id)...); err != nil {
		s.logger.Printf("Error updating author from DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) DeleteAuthor(id int) error {
	if _, err := s.db.Exec(author.Delete, id); err != nil {
		s.logger.Printf("Error deleting author from DB: %v", err)
		return err
	}
	return nil
}
