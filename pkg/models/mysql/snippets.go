package mysql

import (
	"database/sql"
	"errors"

	"github.com/angelmendozacap/SnippetBox/pkg/models"
)

// SnippetModel type wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content, expiresAt string, userID int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, expires_at, user_id)
	VALUES (?, ?, DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY), ?)`

	res, err := m.DB.Exec(stmt, title, content, expiresAt, userID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get returns a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT
		s.id, s.title, s.content, s.created_at, s.expires_at,
		u.id, u.name, u.email
	FROM snippets s
	INNER JOIN users u
		ON u.id = s.user_id
	WHERE s.expires_at > NOW() AND s.id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Snippet{}

	err := row.Scan(
		&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt,
		&s.User.ID, &s.User.Name, &s.User.Email,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return s, nil
}

// Latest returns the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created_at, expires_at FROM snippets
	WHERE expires_at > NOW() ORDER BY created_at DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
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
