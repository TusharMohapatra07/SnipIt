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
	DB *sql.DB
}

// Insert a new snippet into database and returns the id
func (sm *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) 
    VALUES ($1, $2, CURRENT_TIMESTAMP, DATE_ADD(CURRENT_TIMESTAMP, '$3 DAYS'::INTERVAL)) RETURNING id`

	var id int

	err := sm.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

// Get a snippet corresponding to an id
func (sm *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// Returns 10 most recent created snippets
func (sm *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
