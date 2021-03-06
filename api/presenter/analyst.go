package presenter

import (
	"flavioltonon/hmv/domain/entity"
)

// Analyst is a entity.Analyst presenter
type Analyst struct {
	ID        string `json:"_id"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// NewAnalyst returns a presentation for a Analyst
func NewAnalyst(e *entity.Analyst) *Analyst {
	return &Analyst{
		ID:        e.ID,
		UserID:    e.UserID,
		CreatedAt: e.CreatedAt.Format("02/01/2006 - 15:04:05h"),
		UpdatedAt: e.UpdatedAt.Format("02/01/2006 - 15:04:05h"),
	}
}
