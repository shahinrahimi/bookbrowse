package store

import (
	"database/sql"
	"testing"

	"github.com/shahinrahimi/bookbrowse/pkg/author"
)

func TestGetAuthor(t *testing.T) {
	store := SetupTestStore()
	defer store.CloseDB()
	if err := store.Init(); err != nil {
		t.Fatalf("error initilizing DB: %v", err)
	}

	// create test author
	a := &author.Author{
		Name: "test",
	}

	// insert test author to DB
	if err := store.CreateAuthor(a); err != nil {
		t.Fatalf("failed to insert author to DB: %v", err)
	}

	// test author should hava ID 1
	fethcedAuthor, err := store.GetAuthor(1)
	if err != nil {
		t.Fatalf("failed to select author: %v", err)
	}

	// check if stored author is correctly
	if fethcedAuthor.Name != a.Name {
		t.Fatalf("expected author name '%s', got '%s'", a.Name, fethcedAuthor.Name)
	}
}

func TestGetAuthors(t *testing.T) {
	store := SetupTestStore()
	if err := store.Init(); err != nil {
		t.Fatalf("error initilizing DB: %v", err)
	}

	// create author
	as := []*author.Author{
		{
			Name: "test1",
		},
		{
			Name: "test2",
		},
		{
			Name: "test3",
		},
	}

	// insert authors to DB
	for _, a := range as {
		if err := store.CreateAuthor(a); err != nil {
			t.Fatalf("failed to insert a author to DB: %v", err)
		}
	}

	// check authors lenght
	fetchedAuthors, err := store.GetAuthors()
	if err != nil {
		t.Fatalf("failed to select all authors: %v", err)
	}

	// check authors length
	if len(fetchedAuthors) != len(as) {
		t.Fatalf("expected author length '%d', got '%d'", len(as), len(fetchedAuthors))
	}
}

func TestUpdateAuthor(t *testing.T) {
	store := SetupTestStore()
	defer store.CloseDB()
	if err := store.Init(); err != nil {
		t.Fatalf("error initilizing DB: %v", err)
	}

	// create test author
	a := &author.Author{
		Name: "test",
	}

	// insert test author to DB
	if err := store.CreateAuthor(a); err != nil {
		t.Fatalf("failed to insert author to DB: %v", err)
	}

	// update author
	a.Name = "testtest"
	if err := store.UpdateAuthor(1, a); err != nil {
		t.Fatalf("failed to update author from DB: %v", err)
	}

	// get updated author
	fethcedAuthor, err := store.GetAuthor(1)
	if err != nil {
		t.Fatalf("failed to select a author from DB: %v", err)
	}

	// check if author updated
	if fethcedAuthor.Name != a.Name {
		t.Fatalf("expected author name '%s', got '%s'", a.Name, fethcedAuthor.Name)
	}
}

func TestDeleteAuthor(t *testing.T) {
	store := SetupTestStore()
	defer store.CloseDB()
	if err := store.Init(); err != nil {
		t.Fatalf("error initilizing DB: %v", err)
	}

	// create test author
	a := &author.Author{
		Name: "test",
	}

	// insert test author to DB
	if err := store.CreateAuthor(a); err != nil {
		t.Fatalf("failed to insert author to DB: %v", err)
	}

	// delte test author
	if err := store.DeleteAuthor(1); err != nil {
		t.Fatalf("failed to delete author from DB: %v", err)
	}

	if _, err := store.GetAuthor(1); err != nil {
		if err != sql.ErrNoRows {
			t.Fatalf("failed to select author from DB: %v", err)
		}
	}
}
