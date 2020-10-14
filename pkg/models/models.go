package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord is the error when the models are not matching
	ErrNoRecord = errors.New("models: no matching record found")

	// ErrInvalidCredentials error. If a user
	// tries to login with an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// ErrDuplicateEmail error. If a user
	// tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

// Snippet is the DB model snippet
type Snippet struct {
	ID        uint       `json:"id"`
	User      *User      `json:"user"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// User is the DB model user
type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  []byte     `json:"-"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	IsActive  bool       `json:"is_active,omitempty"`
}
