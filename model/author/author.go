package author

type Author struct {
	ID         int    `json:"id"`
	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Lastname   string `json:"lastname"`
}

const (
	CreateTable string = `CREATE TABLE IF NOT EXISTS authors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fisrtname TEXT NOT NULL,
		middlename TEXT,
	);`
	SelectAll string = `SELECT id, firstname, middlename, lastname FROM authors`
	Select    string = `SELECT id, firstname, middlename, lastname FROM authors WHERE id = ?`
	Insert    string = `INSERT INTO authors (id, fisrname, middlename, lastname) VALUES (?, ?, ?, ?)`
	Update    string = `UPDATE authors SET firstname = ?, middlename = ?, lastname = ? WHERE id = ?`
	Delete    string = `DELETE FROM authors WHERE id = ?`
)

// ToArgs returns id, firstname, middlename, lastname as value
func (a *Author) ToArgs() []interface{} {
	return []interface{}{a.ID, a.Firstname, a.Middlename, a.Lastname}
}

// ToUpdatedArgs returns firstname, middlename, lastname and id as reference
func (a *Author) ToUpdatedArgs(id int) []interface{} {
	return []interface{}{a.Firstname, a.Middlename, a.Lastname, id}
}

// ToFields returns id, firstname, middlename, lastname as  reference
func (a *Author) ToFeilds() []interface{} {
	return []interface{}{&a.ID, &a.Firstname, &a.Middlename, &a.Lastname}
}
