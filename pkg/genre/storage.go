package genre

type Storage interface {
	GetGenres() ([]*Genre, error)
	GetGenre(id int) (*Genre, error)
	CreateGenre(g *Genre) error
	UpdateGenre(id int, g *Genre) error
	DeleteGenre(id int) error
}
