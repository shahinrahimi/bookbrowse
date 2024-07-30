package author

type Storage interface {
	GetAuthors() ([]*Author, error)
	GetAuthor(id int) (*Author, error)
	CreateAuthor(a *Author) error
	UpdateAuthor(id int, a *Author) error
	DeleteAuthor(id int) error
}
