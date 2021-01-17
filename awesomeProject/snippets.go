package main

import (
	"database/sql"
	"errors"
)
type SnippetModel struct { DB *sql.DB
}
func (m *SnippetModel) Insert(title, content, expires string) (int, error) { return 0, nil
}
func (m *SnippetModel) Get(id int) (*models.Snippet, error) { return nil, nil
}
func (m *SnippetModel) Latest() ([]*models.Snippet, error) { return nil, nil
}
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
WHERE expires > UTC_TIMESTAMP() AND id = ?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.Snippet{}
		err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord } else {
			return nil, err }
	}
	return s, nil }
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	snippets := []*models.Snippet{}
	for rows.Next() {
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err }
		snippets = append(snippets, s) }

	if err = rows.Err(); err != nil { return nil, err
	}
	return snippets, nil }