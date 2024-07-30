package genre

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

const (
	CreateTable string = `CREATE TABLE IF NOT EXISTS genres (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);`
	SelectAll string = `SELECT id, name FROM genres`
	Select    string = `SELECT id, name FROM genres WHERE id = ?`
	Insert    string = `INSERT INTO genres (id, name) VALUES (?, ?)`
	Update    string = `UPDATE genres SET name = ? WHERE id = ?`
	Delete    string = `DELETE FROM genres WHERE id = ?`
)

// ToArgs returns id, name as value
func (g *Genre) ToArgs() []interface{} {
	return []interface{}{g.ID, g.Name}
}

// ToUpdatedArgs returns name, id as value
func (g *Genre) ToUpdatedArgs(id int) []interface{} {
	return []interface{}{g.Name, id}
}

// ToFeilds returns id, name as referece
func (g *Genre) ToFeilds() []interface{} {
	return []interface{}{&g.ID, &g.Name}
}
