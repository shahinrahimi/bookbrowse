package store

import "github.com/shahinrahimi/bookbrowse/pkg/genre"

func (s *SqliteStore) GetGenres() ([]*genre.Genre, error) {
	return nil, nil
}

func (s *SqliteStore) GetGenre(id int) (*genre.Genre, error) {
	return nil, nil
}

func (s *SqliteStore) CreateGenre(g *genre.Genre) error {
	return nil
}

func (s *SqliteStore) UpdateGenre(id int, g *genre.Genre) error {
	return nil
}

func (s *SqliteStore) DeleteGenre(id int) error {
	return nil
}
