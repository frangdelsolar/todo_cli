package db

import (
	"fmt"
	"time"
	"todo_cli/models"

	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

// GetTaskById retrieves a task from the database by its ID.
//
// Parameters:
// - id: the ID of the task to retrieve.
//
// Returns:
// - *models.Task: a pointer to the retrieved task, or nil if not found.
// - error: an error if the task retrieval fails.
func (db *DB) GetTaskById(id string) (*models.Task, error) {
	task := db.Tasks[id]
	if task.ID == uuid.Nil {
		return nil, fmt.Errorf("task with ID %s not found", id)
	}
	return &task, nil
}

// GetAllTasks retrieves all tasks from the database.
//
// Returns:
// - []models.Task: a slice of all tasks in the database.
func (db *DB) GetAllTasks() []models.Task {
	tasks := make([]models.Task, 0, len(db.Tasks))
	for _, task := range db.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// GetActiveTasks retrieves all active tasks from the database.
//
// Returns:
// - []models.Task: a slice of active tasks in the database.
func (db *DB) GetActiveTasks() []models.Task {
	tasks := []models.Task{}
	now := time.Now().Format(time.RFC3339)
	for _, effectivePeriod := range db.EffectivePeriods {
		if effectivePeriod.EndDate > now  || effectivePeriod.EndDate==""{
			task := db.Tasks[effectivePeriod.TaskID.String()]
			tasks = append(tasks, task)
		}
	}
	return tasks
}

// CreateTask creates a new task in the database.
//
// Parameters:
// - title: the title of the task.
//
// Returns:
// - *models.Task: the newly created task.
// - error: an error if the task creation fails.
func (db *DB) CreateTask(title string) (*models.Task, error) {
	task, err := models.NewTask(title)
	if err != nil {
		log.Err(err).Msg("Error creating new Task")
		return nil, err
	}
	db.Tasks[task.ID.String()] = *task
	db.Save()
	return task, nil
}

// UpdateTask updates a task in the database with the given ID and title.
//
// Parameters:
// - id: the ID of the task to update.
// - title: the new title for the task.
//
// Returns:
// - *models.Task: the updated task.
// - error: an error if the task retrieval or update fails.
func (db *DB) UpdateTask(id string, title string) (*models.Task, error) {
	task, err := db.GetTaskById(id)
	if err != nil {
		log.Err(err).Msg("Error retrieving Task")
		return nil, err
	}
	err = task.Update(title)
	if err != nil {
		log.Err(err).Msg("Error updating Task")
		return nil, err
	}
	db.Tasks[task.ID.String()] = *task
	db.Save()
	return task, nil
}

// DeleteTask deletes a task from the database by its ID.
//
// Parameters:
// - id: the ID of the task to delete.
//
// Returns:
// - error: an error if the task retrieval or deletion fails.
func (db *DB) DeleteTask(taskId string) error {
	_, err := db.GetTaskById(taskId)
	if err != nil {
		log.Err(err).Msg("Error retrieving Task")
		return err
	}

	// Delete all the effective periods related to the task
	effectivePeriods := db.GetEffectivePeriodsByTaskId(taskId)
	for _, effectivePeriod := range effectivePeriods {
		db.DeleteEffectivePeriod(effectivePeriod.ID.String())
	}

	delete(db.Tasks, taskId)
	db.Save()

	return nil
}
