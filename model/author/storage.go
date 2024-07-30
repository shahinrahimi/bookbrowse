package author

type Storage interface {
	GetAuthors() ([]*Author, error)
	GetAuthor(id string) (*Author, error)
	CreateAuthor(b *Author) error
	UpdateAuthor(id string, b *Author) error
	DeleteAuthor(id string) error
}
