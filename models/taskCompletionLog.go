package models

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

// TaskCompletionLog represents a task completion log.
//
// Fields:
// - ID: the ID of the task completion log.
// - TaskID: the ID of the task associated with the task completion log.
// - CompletedAt: the timestamp when the task was completed.
type TaskCompletionLog struct {
	ID          uuid.UUID `json:"id"`
	TaskID      uuid.UUID `json:"taskId"`
	CompletedAt string    `json:"completedAt"`
}

// String returns a string representation of the TaskCompletionLog.
//
// Returns:
// - string: a string representation of the TaskCompletionLog.
func (t *TaskCompletionLog) String() string {
	return fmt.Sprintf("TaskCompletionLog %s\nTask ID: %s\nCompleted At: %s\n", t.ID, t.TaskID, t.CompletedAt)
}

// Update updates the completedAt field of the TaskCompletionLog.
//
// Parameters:
// - in_completedAt: the new completedAt date in string format.
//
// Returns:
// - error: an error if there was a problem parsing the completedAt date.
func (t *TaskCompletionLog) Update(in_completedAt string) error {
	formatedCompletedAt, err := time.Parse(time.DateOnly, in_completedAt)
	if err != nil {
		log.Err(err).Msg("Error parsing completedAt date")
		return err
	}
	t.CompletedAt = formatedCompletedAt.Format(time.RFC3339)
	return nil
}

// NewTaskCompletionLog creates a new TaskCompletionLog with the given task ID and sets the current time as the CompletedAt field.
//
// Parameters:
// - taskID: the ID of the task associated with the task completion log.
// - completedAt: the date and time when the task was completed. If empty, the current time is used.
//
// Returns:
// - *TaskCompletionLog: a pointer to the newly created TaskCompletionLog.
// - error: an error if the task ID is empty or if there was a problem parsing the completedAt date.
func NewTaskCompletionLog(taskID string, completedAt string) (*TaskCompletionLog, error) {

	if taskID == "" {
		return nil, fmt.Errorf("task ID cannot be empty")
	}
	
	now := time.Now().String()[0:10]
	if completedAt == "" {
		log.Warn().Msg("No completedAt date provided, defaulting to current time")
		completedAt = now
	}

	parsedCompletedAt, err := time.Parse(time.DateOnly, completedAt)
	if err != nil {
		log.Warn().Msg("Invalid completedAt date provided, defaulting to current time")
		parsedCompletedAt, _ = time.Parse(time.DateOnly, now)
	}

	return &TaskCompletionLog{
		ID:          uuid.Must(uuid.NewV4()),
		TaskID:      uuid.Must(uuid.FromString(taskID)),
		CompletedAt: parsedCompletedAt.Format(time.RFC3339),
	}, nil
}