package models

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// TaskCompletionLog represents a task completion log.
//
// Fields:
// - ID: the ID of the task completion log.
// - TaskID: the ID of the task associated with the task completion log.
// - CompletedAt: the timestamp when the task was completed.
type TaskCompletionLog struct {
	gorm.Model
	ID          string      `json:"id" gorm:"primaryKey"`
	TaskID      string      `json:"taskId"`
	Task        *Task     `json:"task" gorm:"foreignKey:TaskID"`
	TaskGoalID  string      `json:"taskGoalId"`
	TaskGoal    *TaskGoal `json:"taskGoal" gorm:"foreignKey:TaskGoalID"`
	CompletedAt time.Time `json:"completedAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// String returns a string representation of the TaskCompletionLog.
//
// Returns:
// - string: a string representation of the TaskCompletionLog.
func (t *TaskCompletionLog) String() string {
	return fmt.Sprintf("TaskCompletionLog %d\nTask ID: %d\nCompleted At: %s\n", t.ID, t.TaskID, t.CompletedAt)
}

// Update updates the completedAt field of the TaskCompletionLog.
//
// Parameters:
// - in_completedAt: the new completedAt date in string format.
//
// Returns:
// - error: an error if there was a problem parsing the completedAt date.
func (t *TaskCompletionLog) Update(in_completedAt string) error {

	err := DateValidator(in_completedAt)
	if err != nil {
		log.Err(err).Msg("Error validating completedAt date")
		return err
	}

	formatedCompletedAt, _ := time.Parse(time.DateOnly, in_completedAt)
	t.CompletedAt = formatedCompletedAt

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
func NewTaskCompletionLog(taskID string, completedAt string, taskGoalID string) (*TaskCompletionLog, error) {

	err := TaskIDValidator(taskID)
	if err != nil {
		log.Err(err).Msg("Error validating Task ID")
		return nil, err
	}

    err = TaskGoalIDValidator(taskGoalID)
	if err != nil {
		log.Err(err).Msg("Error validating Task Goal ID")
		return nil, err
	}

	now := time.Now()
	var cd time.Time

	err = DateValidator(completedAt)
	if err != nil {
		log.Warn().Msg("Error validating completedAt date. Defaulting to current time")
		cd = now
	} else {
		cd, _ = time.Parse(time.DateOnly, completedAt)
	}

	return &TaskCompletionLog{
		TaskID:      taskID,
		CompletedAt: cd,
		TaskGoalID: taskGoalID,
	}, nil
}
