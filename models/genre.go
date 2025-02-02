package models

import "strings"

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
}

type Genres []*Genre

type KeyGenre struct{}

const (
	CREATE_TABLE_GENRES string = `CREATE TABLE IF NOT EXISTS genres (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL
	);`
	CREATE_TABLE_BOOKGENRES string = `CREATE TABLE IF NOT EXISTS book_genres (
            book_id INTEGER,
            genre_id INTEGER,
            PRIMARY KEY (book_id, genre_id),
            FOREIGN KEY (book_id) REFERENCES books(id),
            FOREIGN KEY (genre_id) REFERENCES genres(id)
        );`
	SELECT_COUNT_GENRES   string = `SELECT COUNT(*) FROM genres`
	SELECT_GENRES         string = `SELECT id, name FROM genres`
	SELECT_LIMITED_GENRES string = `SELECT id, name FROM genres LIMIT ? OFFSET ?`
	SELECT_GENRE          string = `SELECT id, name FROM genres WHERE id = ?`
	INSERT_GENRE          string = `INSERT INTO genres (name) VALUES (?)`
	UPDATE_GENRE          string = `UPDATE genres SET name = ? WHERE id = ?`
	DELETE_GENRE          string = `DELETE FROM genres WHERE id = ?`
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

// GetNames retunrs slice of strings that contains names
func (gs *Genres) GetNames() []string {
	names := make([]string, len(*gs))
	for i, g := range *gs {
		names[i] = g.Name
	}
	return names
}

// GetID returns ID of genre if found in the genres
// if not found will return -1
func (gs *Genres) GetID(name string) int {
	// make sure the name is caseinsensetiv and trimed space
	cleanName := strings.TrimSpace(strings.ToLower(name))

	for _, g := range *gs {
		if g.Name == cleanName {
			return g.ID
		}
	}
	return -1
}

func (gs *Genres) Add(g *Genre) {
	*gs = append(*gs, g)
}
