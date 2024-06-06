package data

import (
	"fmt"
	"todo_cli/models"

	"github.com/rs/zerolog/log"
)

// GetTaskCompletionLogById retrieves a TaskCompletionLog from the database by its ID.
//
// Parameters:
// - id: the ID of the TaskCompletionLog to retrieve.
//
// Returns:
// - *models.TaskCompletionLog: a pointer to the retrieved TaskCompletionLog, or nil if not found.
// - error: an error if the TaskCompletionLog retrieval fails.
func GetTaskCompletionLogById(id uint) (models.TaskCompletionLog, error) {
	var tcl models.TaskCompletionLog
	DB.First(&tcl, "id = ?", id)
	if tcl == (models.TaskCompletionLog{}) {
		return tcl, fmt.Errorf("task with ID %s not found", id)
	}
	return tcl, nil
}

// GetTaskCompletionLogsByTaskId retrieves all TaskCompletionLogs associated with the given task ID.
//
// Parameters:
// - taskId: the ID of the task associated with the TaskCompletionLogs.
//
// Returns:
// - []models.TaskCompletionLog: a slice of TaskCompletionLogs associated with the given task ID.
func GetTaskCompletionLogsByTaskId(taskId uint) []models.TaskCompletionLog {
	var tasks []models.TaskCompletionLog

	DB.Where("task_id = ?", taskId).Find(&tasks)

	if len(tasks) == 0 {
		log.Warn().Msg("No task completion logs found")
		return tasks
	}

	return tasks
}

// CreateTaskCompletionLog creates a new TaskCompletionLog in the database.
//
// Parameters:
// - taskId: the ID of the task associated with the TaskCompletionLog.
// - completedAt: the date and time when the task was completed.
//
// Returns:
// - *models.TaskCompletionLog: a pointer to the newly created TaskCompletionLog.
// - error: an error if the TaskCompletionLog creation fails.
func CreateTaskCompletionLog(taskId uint, completedAt string) (*models.TaskCompletionLog, error) {
	tcl, err := models.NewTaskCompletionLog(taskId, completedAt)
	if err != nil {
		log.Err(err).Msg("Error creating new Task Completion Log")
		return nil, err
	}
	DB.Create(&tcl)
	return tcl, nil
}

// UpdateTaskCompletionLog updates the completedAt field of a TaskCompletionLog in the database.
//
// Parameters:
// - id: the ID of the TaskCompletionLog to update.
// - completedAt: the date and time when the task was completed.
//
// Returns:
// - *models.TaskCompletionLog: a pointer to the updated TaskCompletionLog.
// - error: an error if the TaskCompletionLog retrieval or update fails.
func UpdateTaskCompletionLog(id uint, completedAt string) (models.TaskCompletionLog, error) {
	var tcl models.TaskCompletionLog
	tcl, err := GetTaskCompletionLogById(id)
	if err != nil {
		log.Err(err).Msg("Error retrieving Task")
		return tcl, err
	}
	err = tcl.Update(completedAt)
	if err != nil {
		log.Err(err).Msg("Error updating Task Completion Log")
		return tcl, err
	}

	DB.Save(&tcl)

	return tcl, nil
}

// DeleteTaskCompletionLog deletes a TaskCompletionLog from the database by its ID.
//
// Parameters:
// - id: the ID of the TaskCompletionLog to delete.
//
// Returns:
// - error: an error if the TaskCompletionLog retrieval or deletion fails.
func DeleteTaskCompletionLog(id uint) error {
	var tcl models.TaskCompletionLog
	tcl, err := GetTaskCompletionLogById(id)
	if err != nil {
		log.Err(err).Msg("Error retrieving Task Completion Log")
		return err
	}

	DB.Delete(&tcl)

	return nil
}
