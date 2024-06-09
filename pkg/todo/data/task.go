package data

import (
	"fmt"
	"time"

	"github.com/frangdelsolar/todo_cli/pkg/todo/models"
)

// GetTaskById retrieves a task from the database by its ID.
//
// Parameters:
// - id: the ID of the task to retrieve.
//
// Returns:
// - *models.Task: a pointer to the retrieved task, or nil if not found.
// - error: an error if the task retrieval fails.
func GetTaskById(id string) (models.Task, error) {
	var task models.Task

	DB.First(&task, "id = ?", id)
	if task == (models.Task{}) {
		return task, fmt.Errorf("task with ID %s not found", fmt.Sprint(id))
	}
	return task, nil
}

// GetAllTasks retrieves all tasks from the database.
//
// Returns:
// - []models.Task: a slice of all tasks in the database.
func GetAllTasks() []models.Task {
	var tasks []models.Task

	DB.Find(&tasks)

	if len(tasks) == 0 {
		log.Warn().Msg("No tasks found")
		return tasks
	}

	return tasks
}

// GetActiveTasks retrieves all active tasks from the database.
//
// Returns:
// - []models.Task: a slice of active tasks in the database.
func GetActiveTasks(refDate time.Time) []models.Task {
	var tasks []models.Task

	nullDate := time.Time{}

	DB.Table("task_goals").
		Select("DISTINCT tasks.*").
		Joins("join tasks on task_goals.task_id = tasks.id").
		Where(`
				task_goals.start_date <= ? AND 
					(
						task_goals.end_date >= ? OR 
						task_goals.end_date == ?
					)
			  `, refDate, refDate, nullDate).
		Find(&tasks)

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
func CreateTask(title string) (models.Task, error) {
	task, err := models.NewTask(title)

	if err != nil {
		log.Err(err).Msg("Error creating new Task")
		return models.Task{}, err
	}

	DB.Create(&task)
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
func UpdateTask(id string, title string) (models.Task, error) {
	var task models.Task
	task, err := GetTaskById(id)
	if err != nil {
		log.Err(err).Msg("Error retrieving Task")
		return task, err
	}
	err = task.Update(title)
	if err != nil {
		log.Err(err).Msg("Error updating Task")
		return task, err
	}

	DB.Save(&task)

	return task, nil
}

// DeleteTask deletes a task from the database by its ID.
//
// Parameters:
// - id: the ID of the task to delete.
//
// Returns:
// - error: an error if the task retrieval or deletion fails.
func DeleteTask(taskId string) error {
	task, err := GetTaskById(taskId)
	if err != nil {
		log.Err(err).Msg("Error retrieving Task")
		return err
	}

	// Delete all the task goals related to the task
	taskGoals := GetTaskGoalsByTaskId(taskId)
	if len(taskGoals) != 0 {
		for _, taskGoal := range taskGoals {
			DeleteTaskGoal(fmt.Sprint(taskGoal.ID))
		}
	}

	DB.Delete(&task)

	return nil
}

func difference(a []uint, b []uint) []uint {
	mb := make(map[uint]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []uint
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
