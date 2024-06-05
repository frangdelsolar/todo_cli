package db

import (
	"fmt"
	"todo_cli/models"

	"github.com/gofrs/uuid"
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
func (db *DB) GetTaskCompletionLogById(id string) (*models.TaskCompletionLog, error) {
	tcl := db.TaskCompletionLogs[id]
	if tcl.ID == uuid.Nil {
		return nil, fmt.Errorf("task with ID %s not found", id)
	}
	return &tcl, nil
}

// GetTaskCompletionLogsByTaskId retrieves all TaskCompletionLogs associated with the given task ID.
//
// Parameters:
// - taskId: the ID of the task associated with the TaskCompletionLogs.
//
// Returns:
// - []models.TaskCompletionLog: a slice of TaskCompletionLogs associated with the given task ID.
func (db *DB) GetTaskCompletionLogsByTaskId(taskId string) []models.TaskCompletionLog {
	tasks := []models.TaskCompletionLog{}
	for _, tcl := range db.TaskCompletionLogs {
		if tcl.TaskID.String() == taskId {
			tasks = append(tasks, tcl)
		}
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
func (db *DB) CreateTaskCompletionLog(taskId string, completedAt string) (*models.TaskCompletionLog, error) {
	tcl, err := models.NewTaskCompletionLog(taskId, completedAt)
	if err != nil {
		log.Err(err).Msg("Error creating new Task Completion Log")
		return nil, err
	}
	db.TaskCompletionLogs[tcl.ID.String()] = *tcl
	db.Save()
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
func (db *DB) UpdateTaskCompletionLog(id string, completedAt string) (*models.TaskCompletionLog, error) {
	tcl, err := db.GetTaskCompletionLogById(id)
	if err != nil {
		log.Err(err).Msg("Error retrieving Task")
		return nil, err
	}
	err = tcl.Update(completedAt)
	if err != nil {
		log.Err(err).Msg("Error updating Task Completion Log")
		return nil, err
	}
	db.TaskCompletionLogs[tcl.ID.String()] = *tcl
	db.Save()
	return tcl, nil
}

// DeleteTaskCompletionLog deletes a TaskCompletionLog from the database by its ID.
//
// Parameters:
// - id: the ID of the TaskCompletionLog to delete.
//
// Returns:
// - error: an error if the TaskCompletionLog retrieval or deletion fails.
func (db *DB) DeleteTaskCompletionLog(id string) error {
	_, err := db.GetTaskCompletionLogById(id)
	if err != nil {
		log.Err(err).Msg("Error retrieving Task Completion Log")
		return err
	}

	delete(db.TaskCompletionLogs, id)
	db.Save()

	return nil
}
