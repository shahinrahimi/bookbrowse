package stores

import "github.com/shahinrahimi/bookbrowse/models"

func (s *SqliteStore) GetAuthors() ([]*models.Author, error) {
	rows, err := s.db.Query(models.SelectAllAuthors)
	if err != nil {
		s.logger.Printf("Error scranning rows for authors: %v", err)
		return nil, err
	}
	// initiate authors as empty slice
	authors := []*models.Author{}
	for rows.Next() {
		var a models.Author
		if err := rows.Scan(a.ToFeilds()...); err != nil {
			s.logger.Printf("Error scranning rows for a author: %v", err)
			continue
		}
		authors = append(authors, &a)
	}
	return authors, nil
}

func (s *SqliteStore) GetAuthor(id int) (*models.Author, error) {
	var a models.Author
	if err := s.db.QueryRow(models.SelectAuthor, id).Scan(a.ToFeilds()...); err != nil {
		s.logger.Printf("Error scranning row for the author: %v", err)
		return nil, err
	}
	return &a, nil
}

func (s *SqliteStore) CreateAuthor(a *models.Author) error {
	if _, err := s.db.Exec(models.InsertAuthor, a.ToArgs()...); err != nil {
		s.logger.Printf("Error inserting a new author to DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) UpdateAuthor(id int, a *models.Author) error {
	if _, err := s.db.Exec(models.UpdateAuthor, a.ToUpdatedArgs(id)...); err != nil {
		s.logger.Printf("Error updating author from DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) DeleteAuthor(id int) error {
	if _, err := s.db.Exec(models.DeleteAuthor, id); err != nil {
		s.logger.Printf("Error deleting author from DB: %v", err)
		return err
	}
	return nil
}
