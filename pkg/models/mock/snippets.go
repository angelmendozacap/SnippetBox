package mock

import (
	"time"

	"github.com/angelmendozacap/SnippetBox/pkg/models"
)

var tm = time.Now()

var mockSnippet = &models.Snippet{
	ID:        1,
	User:      mockUser,
	Title:     "An old silent pond",
	Content:   "An old silent pond...",
	CreatedAt: &tm,
	ExpiresAt: &tm,
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
