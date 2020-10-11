package mock

import (
	"time"

	"github.com/angelmendozacap/SnippetBox/pkg/models"
)

var mockSnippet = &models.Snippet{
	ID:        1,
	User:      mockUser,
	Title:     "An old silent pond",
	Content:   "An old silent pond...",
	CreatedAt: time.Now(),
	ExpiresAt: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title, content, expiresAt string, userID int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
