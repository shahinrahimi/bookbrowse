package store

import "github.com/shahinrahimi/bookbrowse/pkg/genre"

func (s *SqliteStore) GetGenres() ([]*genre.Genre, error) {
	rows, err := s.db.Query(genre.SelectAll)
	if err != nil {
		s.logger.Printf("Error scranning rows for genres: %v", err)
		return nil, err
	}
	// initiate books as empty slice
	genres := []*genre.Genre{}
	for rows.Next() {
		var g genre.Genre
		if err := rows.Scan(&g); err != nil {
			s.logger.Printf("Error scranning rows for a genre: %v", err)
			continue
		}
		genres = append(genres, &g)
	}
	return genres, nil
}

func (s *SqliteStore) GetGenre(id int) (*genre.Genre, error) {
	var g genre.Genre
	if err := s.db.QueryRow(genre.Select, id).Scan(g.ToFeilds()...); err != nil {
		s.logger.Printf("Error scranning row for the genre: %v", err)
		return nil, err
	}
	return &g, nil
}

func (s *SqliteStore) CreateGenre(g *genre.Genre) error {
	if _, err := s.db.Exec(genre.Insert, g.ToArgs()...); err != nil {
		s.logger.Printf("Error inserting a new genre to DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) UpdateGenre(id int, g *genre.Genre) error {
	if _, err := s.db.Exec(genre.Update, g.ToUpdatedArgs(id)...); err != nil {
		s.logger.Printf("Error updating genre from DB: %v", err)
		return err
	}
	return nil
}

func (s *SqliteStore) DeleteGenre(id int) error {
	if _, err := s.db.Exec(genre.Delete, id); err != nil {
		s.logger.Printf("Error deleting genre from DB: %v", err)
		return err
	}
	return nil
}
