package models

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const (
	CreateTableAuthors string = `CREATE TABLE IF NOT EXISTS authors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL
	);`
	SelectAllAuthors string = `SELECT id, name FROM authors`
	SelectAuthor     string = `SELECT id, name FROM authors WHERE id = ?`
	InsertAuthor     string = `INSERT INTO authors (name) VALUES (?)`
	UpdateAuthor     string = `UPDATE authors SET name = ? WHERE id = ?`
	DeleteAuthor     string = `DELETE FROM authors WHERE id = ?`
)

// ToArgs returns name as value
// use for inserting to DB
func (a *Author) ToArgs() []interface{} {
	return []interface{}{a.Name}
}

// ToUpdatedArgs returns name and id as reference
// use for updating record in DB
func (a *Author) ToUpdatedArgs(id int) []interface{} {
	return []interface{}{a.Name, id}
}

// ToFields returns id, anme as reference
// use for scanning from DB
func (a *Author) ToFeilds() []interface{} {
	return []interface{}{&a.ID, &a.Name}
}
