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
	ID        uint
	Title     string
	Content   string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// User is the DB model user
type User struct {
	ID        int
	Name      string
	Email     string
	Password  []byte
	CreatedAt time.Time
	IsActive  bool
}
