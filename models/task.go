package models

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

// Task represents a task.
//
// Fields:
// - ID: the ID of the task.
// - Title: the title of the task.
// - CreatedAt: the timestamp when the task was created.
type Task struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedAt string    `json:"createdAt"`
}

// String returns a string representation of the Task.
//
// Returns:
// - string: a string representation of the Task.
func (t *Task) String() string {
	return fmt.Sprintf("Task %s\nTitle: %s\nCreated At: %s\n", t.ID, t.Title, t.CreatedAt)
}

// Update updates the title of the Task.
//
// Parameters:
// - title: the new title for the Task.
//
// Returns:
// - None.
func (t *Task) Update(title string) (error) {
	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}

	t.Title = title

	return nil
}

// NewTask creates a new Task with the given title and sets the current time as the CreatedAt field.
//
// Parameters:
// - title: the title of the Task.
//
// Returns:
// - *Task: a pointer to the newly created Task.
func NewTask(title string) (*Task, error) {
	now := time.Now().Format(time.RFC3339)

	if title == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}

	task := &Task{
		ID:        uuid.Must(uuid.NewV4()),
		Title:     title,
		CreatedAt: now,
	}
	return task, nil
}
