package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Task represents a task.
//
// Fields:
// - ID: the ID of the task.
// - Title: the title of the task.
// - CreatedAt: the timestamp when the task was created.
type Task struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// String returns a string representation of the Task.
//
// Returns:
// - string: a string representation of the Task.
func (t *Task) String() string {
	return fmt.Sprintf("%d | %s", t.ID, t.Title)
}

// Update updates the title of the Task.
//
// Parameters:
// - title: the new title for the Task.
//
// Returns:
// - None.
func (t *Task) Update(title string) error {

	err := TaskTitleValidator(title)
	if err != nil {
		return err
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
func NewTask(title string) (Task, error) {

	err := TaskTitleValidator(title)
	if err != nil {
		return Task{}, err
	}

	task := Task{
		Title: title,
	}
	return task, nil
}

// TaskTitleValidator validates the title of a task.
//
// Parameters:
// - title: the title of the task.
//
// Returns:
// - error: an error if the title is empty.
func TaskTitleValidator(title string) error {
	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}

	return nil
}


// TaskIDValidator validates a task ID.
//
// Parameters:
// - id: the ID of the task to validate.
//
// Returns:
// - error: an error if the task ID is 0, otherwise nil.
func TaskIDValidator(id string) error {

	var err error
	if id == "" {
		err = errors.New("task ID cannot be empty")
		return err
	}

	return err
}