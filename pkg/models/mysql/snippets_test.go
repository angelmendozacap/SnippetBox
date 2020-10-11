package mysql

import "testing"

func TestSnippetModelInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test.")
	}

	db, teardown := newTestDB(t)
	defer teardown()

	m := SnippetModel{db}

	snippet := struct {
		Title     string
		Content   string
		ExpiresAt string
		UserID    int
	}{
		"Title", "Content", "2020-12-23", 1,
	}

	lastID, err := m.Insert(snippet.Title, snippet.Content, snippet.ExpiresAt, snippet.UserID)
	if err != nil {
		t.Fatal(err)
	}

	if lastID != 1 {
		t.Errorf("want %v; got %v", 1, lastID)
	}

	newSnippet, err := m.Get(lastID)
	if err != nil {
		t.Fatal(err)
	}

	if newSnippet.Title != snippet.Title {
		t.Errorf("want %s; got %s", snippet.Title, newSnippet.Title)
	}

	if newSnippet.User.ID != snippet.UserID {
		t.Errorf("want %v; got %v", snippet.UserID, newSnippet.User.ID)
	}
}
