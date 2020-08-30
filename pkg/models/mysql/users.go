package mysql

import (
	"database/sql"

	"github.com/angelmendozacap/SnippetBox/pkg/models"
)

// UserModel type wraps a sql.DB connection pool.
type UserModel struct {
	DB *sql.DB
}

// Insert method adds a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate method verifies whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get method fetchs details for a specific user based
// on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
