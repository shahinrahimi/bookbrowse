package stores

import (
	"database/sql"
	"testing"

	"github.com/shahinrahimi/bookbrowse/models"
)

func TestGetEmptyGenres(t *testing.T) {
	store := SetupTestStore()
	defer store.CloseDB()
	if err := store.Init(); err != nil {
		t.Fatalf("error initilizing DB: %v", err)
	}

	// check genres lenght
	fetchedGenres, err := store.GetGenres()
	if err != nil {
		t.Fatalf("failed to select all genres: %v", err)
	}

	// check genres length
	if len(*fetchedGenres) != 0 {
		t.Fatalf("expected genre length '%d', got '%d'", 0, len(*fetchedGenres))
	}

}

func TestGetGenre(t *testing.T) {
	store := SetupTestStore()
	defer store.CloseDB()
	if err := store.Init(); err != nil {
		t.Fatalf("error initilizing DB: %v", err)
	}

	// create test genre
	g := &models.Genre{
		Name: "test",
	}

	// insert test genre to DB
	if err := store.CreateGenre(g); err != nil {
		t.Fatalf("failed to insert genre to DB: %v", err)
	}

	// test genre should hava ID 1
	fethcedGenre, err := store.GetGenre(1)
	if err != nil {
		t.Fatalf("failed to select genre: %v", err)
	}

	// check if stored genre is correctly
	if fethcedGenre.Name != g.Name {
		t.Fatalf("expected genre name '%s', got '%s'", g.Name, fethcedGenre.Name)
	}
}

func TestGetGenres(t *testing.T) {
	store := SetupTestStore()
	if err := store.Init(); err != nil {
		t.Fatalf("error initilizing DB: %v", err)
	}

	// create genre
	gs := []*models.Genre{
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

	// insert genres to DB
	for _, g := range gs {
		if err := store.CreateGenre(g); err != nil {
			t.Fatalf("failed to insert a genre to DB: %v", err)
		}
	}

	// check genres lenght
	fetchedGenres, err := store.GetGenres()
	if err != nil {
		t.Fatalf("failed to select all genres: %v", err)
	}

	// check genres length
	if len(*fetchedGenres) != len(gs) {
		t.Fatalf("expected genre length '%d', got '%d'", len(gs), len(*fetchedGenres))
	}
}

func TestUpdateGenre(t *testing.T) {
	store := SetupTestStore()
	defer store.CloseDB()
	if err := store.Init(); err != nil {
		t.Fatalf("error initilizing DB: %v", err)
	}

	// create test genre
	g := &models.Genre{
		Name: "test",
	}

	// insert test genre to DB
	if err := store.CreateGenre(g); err != nil {
		t.Fatalf("failed to insert genre to DB: %v", err)
	}

	// update genre
	g.Name = "testtest"
	if err := store.UpdateGenre(1, g); err != nil {
		t.Fatalf("failed to update genre from DB: %v", err)
	}

	// get updated genre
	fethcedGenre, err := store.GetGenre(1)
	if err != nil {
		t.Fatalf("failed to select a genre from DB: %v", err)
	}

	// check if genre updated
	if fethcedGenre.Name != g.Name {
		t.Fatalf("expected genre name '%s', got '%s'", g.Name, fethcedGenre.Name)
	}
}

func TestDeleteGenre(t *testing.T) {
	store := SetupTestStore()
	defer store.CloseDB()
	if err := store.Init(); err != nil {
		t.Fatalf("error initilizing DB: %v", err)
	}

	// create test genre
	g := &models.Genre{
		Name: "test",
	}

	// insert test genre to DB
	if err := store.CreateGenre(g); err != nil {
		t.Fatalf("failed to insert genre to DB: %v", err)
	}

	// delte test genre
	if err := store.DeleteGenre(1); err != nil {
		t.Fatalf("failed to delete genre from DB: %v", err)
	}

	if _, err := store.GetGenre(1); err != nil {
		if err != sql.ErrNoRows {
			t.Fatalf("failed to select genre from DB: %v", err)
		}
	}
}
