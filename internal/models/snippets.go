package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB sql.DB
}

// Insert a new snippet into database
func (sm *SnippetModel) Insert(title string, content string, expires time.Time) (int, error) {
	return 0, nil
}

// Get a snippet corresponding to an id
func (sm *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Returns 10 most recent created snippets
func (sm *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
