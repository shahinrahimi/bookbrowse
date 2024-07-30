package genre

type Storage interface {
	GetGenres() ([]*Genre, error)
	GetGenre(id string) (*Genre, error)
	CreateGenre(g *Genre) error
	UpdateGenre(id string, g *Genre) error
	DeleteGenre(id string) error
}
