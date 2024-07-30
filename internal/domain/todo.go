package domain

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
}

// NewTodo creates a new todo
func NewTodo(title string, description string) *Todo {
	return &Todo{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
}

// Update updates a todo
func (t *Todo) Update(completed bool, title, description string) {
	t.Title = title
	t.Description = description
	t.Completed = completed
}
