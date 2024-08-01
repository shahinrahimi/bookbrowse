package stores

import "github.com/shahinrahimi/bookbrowse/models"

func (s *SqliteStore) GetGenres() (*models.Genres, error) {
	rows, err := s.db.Query(models.SelectAllGenres)
	if err != nil {
		s.logger.Printf("Error scranning rows for genres: %v", err)
		return nil, err
	}
	defer rows.Close()
	// initiate books as empty slice
	var genres models.Genres
	for rows.Next() {
		var g models.Genre
		if err := rows.Scan(g.ToFeilds()...); err != nil {
			s.logger.Printf("Error scranning rows for a genre: %v", err)
			continue
		}
		genres = append(genres, &g)
	}
	// Check for errors encountered during iteration
	if err := rows.Err(); err != nil {
		s.logger.Printf("Error encountered during row iteration: %v", err)
		return nil, err
	}
	return &genres, nil
}

func (s *SqliteStore) GetGenre(id int) (*models.Genre, error) {
	var g models.Genre
	if err := s.db.QueryRow(models.SelectGenre, id).Scan(g.ToFeilds()...); err != nil {
		s.logger.Printf("Error scranning row for the genre: %v", err)
		return nil, err
	}
	return &g, nil
}

func (s *SqliteStore) CreateGenre(g *models.Genre) error {
	if _, err := s.db.Exec(models.InsertGenre, g.ToArgs()...); err != nil {
		s.logger.Printf("Error inserting a new genre to DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) UpdateGenre(id int, g *models.Genre) error {
	if _, err := s.db.Exec(models.UpdateGenre, g.ToUpdatedArgs(id)...); err != nil {
		s.logger.Printf("Error updating genre from DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) DeleteGenre(id int) error {
	if _, err := s.db.Exec(models.DeleteGenre, id); err != nil {
		s.logger.Printf("Error deleting genre from DB: %v", err)
		return err
	}
	return nil
}
