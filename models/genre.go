package models

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const (
	CreateTableGenres string = `CREATE TABLE IF NOT EXISTS genres (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL
	);`
	CreateTableBookGenres string = `CREATE TABLE IF NOT EXISTS book_genres (
            book_id INTEGER,
            genre_id INTEGER,
            PRIMARY KEY (book_id, genre_id),
            FOREIGN KEY (book_id) REFERENCES books(id),
            FOREIGN KEY (genre_id) REFERENCES genres(id)
        );`
	SelectAllGenres string = `SELECT id, name FROM genres`
	SelectGenre     string = `SELECT id, name FROM genres WHERE id = ?`
	InsertGenre     string = `INSERT INTO genres (name) VALUES (?)`
	UpdateGenre     string = `UPDATE genres SET name = ? WHERE id = ?`
	DeleteGenre     string = `DELETE FROM genres WHERE id = ?`
)

// ToArgs returns name as value
// use for inserting to DB
func (g *Genre) ToArgs() []interface{} {
	return []interface{}{g.Name}
}

// ToUpdatedArgs returns name, id as value
// use for updating record in DB
func (g *Genre) ToUpdatedArgs(id int) []interface{} {
	return []interface{}{g.Name, id}
}

// ToFeilds returns id, name as referece
// use for scanning from DB
func (g *Genre) ToFeilds() []interface{} {
	return []interface{}{&g.ID, &g.Name}
}
