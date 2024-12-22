package models

import (
	"database/sql"
	"errors"
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

var (
	ErrNoRecord = errors.New("models: no matching record found")
)

// Insert a new snippet into database and returns the id
func (sm *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets 
    (title, content, created, expires) 
    VALUES 
    ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + $3 * INTERVAL '1 DAY') 
    RETURNING id`

	var id int

	err := sm.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get a snippet corresponding to an id
func (sm *SnippetModel) Get(id int) (Snippet, error) {
	stmt := `SELECT id, title, content, created, expires 
    FROM snippets 
    WHERE expires > CURRENT_TIMESTAMP 
    AND 
    id = $1`

	row := sm.DB.QueryRow(stmt, id)

	var s Snippet

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil
}

// Returns 10 most recent created snippets
func (sm *SnippetModel) Latest() ([]Snippet, error) {
	stmt := `SELECT id, title, content, created, expires 
    FROM snippets
    WHERE expires > CURRENT_TIMESTAMP
    ORDER BY DESC
    LIMIT 10`

	rows, err := sm.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snippets []Snippet

	for rows.Next() {
		var s Snippet

		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
