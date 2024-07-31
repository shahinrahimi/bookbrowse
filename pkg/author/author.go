package author

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const (
	CreateTable string = `CREATE TABLE IF NOT EXISTS authors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL
	);`
	SelectAll string = `SELECT id, name FROM authors`
	Select    string = `SELECT id, name FROM authors WHERE id = ?`
	Insert    string = `INSERT INTO authors (name) VALUES (?)`
	Update    string = `UPDATE authors SET name = ? WHERE id = ?`
	Delete    string = `DELETE FROM authors WHERE id = ?`
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
