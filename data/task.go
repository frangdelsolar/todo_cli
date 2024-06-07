package data

import (
	"fmt"
	"time"
	"todo_cli/models"

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
func GetTaskById(id uint) (models.Task, error) {
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
func GetActiveTasks() []models.Task {
	var tasks []models.Task

	now := time.Now()
	nullDate := time.Time{}

	DB.Table("effective_periods").
		Select("DISTINCT tasks.*").
		Joins("join tasks on effective_periods.task_id = tasks.id").
		Where(`
				effective_periods.start_date <= ? AND 
					(
						effective_periods.end_date >= ? OR 
						effective_periods.end_date == ?
					)
			  `, now, now, nullDate).
		Find(&tasks)

	return tasks
}

type PendingTaskContract struct {
	TaskID              uint
	EffectivePeriodId   uint
	TaskCompletionLogId uint
}

func GetPendingTasksTodoMonthly(date time.Time) []models.Task {

	// Auxiliary variables
	firstDayOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)
	nullDate := time.Time{}

	// Are there any effective period for the time?
	var activeTasksIds []uint
	DB.Table("effective_periods").
		Select("DISTINCT task_id").
		Joins("join tasks on effective_periods.task_id = tasks.id").
		Where(`
				effective_periods.start_date <= ? AND 
					(
						effective_periods.end_date >= ? OR 
						effective_periods.end_date == ?
					) AND
				effective_periods.frequency = ?
			  `, date, date, nullDate, models.Monthly).
		Find(&activeTasksIds)

	// todos that only happen monthly
	// Is there any completion log for the time?
	var tasksWithCompletionLog []uint

	DB.
		Table("task_completion_logs").
		Select("task_id").
		Where("task_id IN ?", activeTasksIds).
		Where("completed_at BETWEEN ? AND ?", firstDayOfMonth, lastDayOfMonth).
		Find(&tasksWithCompletionLog)

	// find those tasks that don't have a completion log
	dueTasksIds := difference(activeTasksIds, tasksWithCompletionLog)

	var tasks []models.Task

	DB.
		Table("tasks").
		Select("tasks.*").
		Where("id IN ?", dueTasksIds).
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
func UpdateTask(id uint, title string) (models.Task, error) {
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
func DeleteTask(taskId uint) error {
	task, err := GetTaskById(taskId)
	if err != nil {
		log.Err(err).Msg("Error retrieving Task")
		return err
	}

	// Delete all the effective periods related to the task
	effectivePeriods := GetEffectivePeriodsByTaskId(taskId)
	if len(effectivePeriods) != 0 {
		for _, effectivePeriod := range effectivePeriods {
			DeleteEffectivePeriod(effectivePeriod.ID)
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
