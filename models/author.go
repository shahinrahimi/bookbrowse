package models

import "strings"

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
}

type Authors []*Author

type KeyAuthor struct{}

const (
	CREATE_TABLE_AUTHORS string = `
		CREATE TABLE IF NOT EXISTS authors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL);`
	SELECT_COUNT_AUHTORS   string = `SELECT COUNT(*) FROM authors`
	SELECT_AUTHORS         string = `SELECT id, name FROM authors`
	SELECT_LIMITED_AUTHORS string = `SELECT id, name FROM authors LIMIT ? OFFSET ?`
	SELECT_AUTHOR          string = `SELECT id, name FROM authors WHERE id = ?`
	INSERT_AUTHOR          string = `INSERT INTO authors (name) VALUES (?)`
	UPDATE_AUTHOR          string = `UPDATE authors SET name = ? WHERE id = ?`
	DELETE_AUTHOR          string = `DELETE FROM authors WHERE id = ?`
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

// GetNames retunrs slice of strings that contains names
func (as *Authors) GetNames() []string {
	names := make([]string, len(*as))
	for i, a := range *as {
		names[i] = a.Name
	}
	return names
}

// GetID returns ID of author if found in the authors
// if not found will return -1
func (as *Authors) GetID(name string) int {
	// make sure the name is caseinsensetiv and trimed space
	cleanName := strings.TrimSpace(strings.ToLower(name))

	for _, a := range *as {
		if a.Name == cleanName {
			return a.ID
		}
	}
	return -1
}

func (as *Authors) Add(a *Author) {
	*as = append(*as, a)
}
