package stores

import (
	"database/sql"

	"github.com/shahinrahimi/bookbrowse/models"
)

func (s *SqliteStore) GetGenres() (*models.Genres, error) {
	rows, err := s.db.Query(models.SELECT_GENRES)
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

func (s *SqliteStore) GetLimitedGenres(limit int, offset int) (*models.Genres, error) {
	rows, err := s.db.Query(models.SELECT_LIMITED_GENRES, limit, offset)
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
	if err := s.db.QueryRow(models.SELECT_GENRE, id).Scan(g.ToFeilds()...); err != nil {
		s.logger.Printf("Error scranning row for the genre: %v", err)
		return nil, err
	}
	return &g, nil
}

func (s *SqliteStore) CreateGenre(g *models.Genre) error {
	if _, err := s.db.Exec(models.INSERT_GENRE, g.ToArgs()...); err != nil {
		s.logger.Printf("Error inserting a new genre to DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) UpdateGenre(id int, g *models.Genre) error {
	if _, err := s.db.Exec(models.UPDATE_GENRE, g.ToUpdatedArgs(id)...); err != nil {
		s.logger.Printf("Error updating genre from DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) DeleteGenre(id int) error {
	if _, err := s.db.Exec(models.DELETE_GENRE, id); err != nil {
		s.logger.Printf("Error deleting genre from DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) GetGenresCount() (int, error) {
	var count int
	if err := s.db.QueryRow(models.SELECT_COUNT_GENRES).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		s.logger.Printf("error selecting count from genres: %v", err)
		return -1, err
	}
	return count, nil
}
