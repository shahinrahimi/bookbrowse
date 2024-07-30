package store

import "github.com/shahinrahimi/bookbrowse/pkg/author"

func (s *SqliteStore) GetAuthors() ([]*author.Author, error) {
	return nil, nil
}

func (s *SqliteStore) GetAuthor(id int) (*author.Author, error) {
	return nil, nil
}

func (s *SqliteStore) CreateAuthor(a *author.Author) error {
	return nil
}

func (s *SqliteStore) UpdateAuthor(id int, a *author.Author) error {
	return nil
}

func (s *SqliteStore) DeleteAuthor(id int) error {
	return nil
}
