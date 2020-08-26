package models

import (
	"errors"
	"time"
)

// ErrNoRecord is the error when the models are not matching
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet is the DB model snippet
type Snippet struct {
	ID        uint
	Title     string
	Content   string
	CreatedAt time.Time
	ExpiresAt time.Time
}
