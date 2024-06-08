package data

import (
	"fmt"
	"time"
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
func GetTaskCompletionLogById(id string) (models.TaskCompletionLog, error) {
	var tcl models.TaskCompletionLog
	DB.First(&tcl, "id = ?", id)
	if tcl == (models.TaskCompletionLog{}) {
		return tcl, fmt.Errorf("task with ID %s not found", fmt.Sprint(id))
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
func GetTaskCompletionLogsByTaskId(taskId string) []models.TaskCompletionLog {
	var tasks []models.TaskCompletionLog

	DB.Where("task_id = ?", taskId).Find(&tasks)

	if len(tasks) == 0 {
		log.Warn().Msg("No task completion logs found")
		return tasks
	}

	return tasks
}

type GoalsContract struct {
	TaskGoalID string
	TaskID string
	StartDate time.Time
	EndDate time.Time
	Label string
	Category models.TaskGoalCategory
	FrequencyID string
	FrequencyType string
	FrequencyDay int
	FrequencyMonth int
	FrequencyDayOfWeek int
}

func getGoalsData(refDate time.Time) []GoalsContract {
	// Get goals for reference date
	var taskGoals []GoalsContract

	DB.Table("task_goals").
		Select(`DISTINCT 
					task_goals.id as TaskGoalID,
					task_goals.start_date as StartDate,
					task_goals.end_date as EndDate,
					task_goals.category as Category,
					tasks.id as TaskID, 
					tasks.title as Label,
					task_frequencies.id as FrequencyID,
					task_frequencies.type as FrequencyType,
					task_frequencies.day as FrequencyDay,
					task_frequencies.month as FrequencyMonth,
					task_frequencies.day_of_week as FrequencyDayOfWeek`,
				).
		Joins("join tasks on task_goals.task_id = tasks.id").
		Joins("join task_frequencies on task_goals.frequency_id = task_frequencies.id").
		Where(`
				task_goals.start_date <= ? AND 
					(
						task_goals.end_date >= ? OR 
						task_goals.end_date == ?
					)
			  `, refDate, refDate, time.Time{}).
		Find(&taskGoals)

	return taskGoals
}

type CompletionContract struct {
	TaskGoalID string
	TaskID string
	DueDate time.Time	
}

func getCompletionLogsByTaskGoalAndDate(refDate time.Time, tg GoalsContract) []CompletionContract{
	var completionLogs []CompletionContract

	DB.Table("task_completion_logs").
		Select(`DISTINCT
					task_completion_logs.task_id as TaskID,
					task_completion_logs.task_goal_id as TaskGoalID,
					task_completion_logs.due_date as DueDate`,
				).
		Where(
			`task_completion_logs.task_goal_id = ? AND
				task_completion_logs.due_date <= ? AND
				task_completion_logs.due_date >= ?`,
			tg.TaskGoalID, refDate, tg.StartDate,
		).
		Order("task_completion_logs.due_date desc").
		Find(&completionLogs)

	return completionLogs
}

type PendingTCLContract struct {
	TaskID        string
	TaskGoalID    string
	DueDate       time.Time
	Label         string
}

func getTCLs(
	refDate time.Time, 
	existingLogsDates []time.Time, 
	tg GoalsContract, 
	limit int,
	validation func(time.Time, time.Time) bool,
) []PendingTCLContract {
	var tcls = []PendingTCLContract{}

	for i := 0; i < limit; i++ {
		nextDate := refDate.AddDate(0, -i, 0)
		if nextDate.Before(tg.StartDate) {
			break
		}
		skip:= false
		for _, d := range existingLogsDates {
			if validation(d, nextDate) {
				log.Debug().Msgf("Matched %s. There's a completion log for this date:", d.Format("2006-01-02"))
				skip = true
			}
		} 
		if !skip {
			tcl:= PendingTCLContract{
				TaskID: tg.TaskID,
				TaskGoalID: tg.TaskGoalID,
				DueDate: nextDate,
				Label: tg.Label,
			}
			tcls = append(tcls, tcl)
		}
	}
	return tcls
}

// refDate is the date where the query is done. It will limit the results up to that date.
func GetPendingTaskCompletionLogs(refDate time.Time) []PendingTCLContract {
	const limit = 50
	var tcls = []PendingTCLContract{}

	taskGoals := getGoalsData(refDate)
	for _, tg := range taskGoals {
		
		completionLogs := getCompletionLogsByTaskGoalAndDate(refDate, tg)

		var completedPeriods = []time.Time{}

		for _, cl := range completionLogs {
			/*
				It should use as a reference the date set to be due
				Not the date of the completion
			*/
			completedPeriods = append(completedPeriods, cl.DueDate)
		}

		// Traverse time according to frquency type and create pendign tasks
		if tg.FrequencyType == string(models.Daily) {
			log.Debug().Msg("Do something about daily")

			dailyPeriodValidation := func (period time.Time, completedPeriod time.Time) bool {
				return (
					period.Day() == completedPeriod.Day() && 
					period.Month() == completedPeriod.Month() && 
					period.Year() == completedPeriod.Year() )
			}
			monthlyTCLs := getTCLs(refDate, completedPeriods, tg, limit, dailyPeriodValidation)
			
			tcls = append(tcls, monthlyTCLs...)
		} else if tg.FrequencyType == string(models.Weekly) {
			log.Debug().Msg("Do somehting about weekly")
		} else if tg.FrequencyType == string(models.Monthly) {
			log.Debug().Msg("Do something about monthly")

			monthPeriodValidation := func (period time.Time, completedPeriod time.Time) bool {
				return (
					period.Month() == completedPeriod.Month() && 
					period.Year() == completedPeriod.Year() )
			}
			monthlyTCLs := getTCLs(refDate, completedPeriods, tg, limit, monthPeriodValidation)

			tcls = append(tcls, monthlyTCLs...)
		} else if tg.FrequencyType == string(models.Yearly) {
			log.Debug().Msg("Do somehting about yearly")
		}
	}

	return tcls
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
func CreateTaskCompletionLog(taskId string, completedAt string, taskGoalId string) (*models.TaskCompletionLog, error) {
	tcl, err := models.NewTaskCompletionLog(taskId, completedAt, taskGoalId)
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
func UpdateTaskCompletionLog(id string, completedAt string) (models.TaskCompletionLog, error) {
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
func DeleteTaskCompletionLog(id string) error {
	var tcl models.TaskCompletionLog
	tcl, err := GetTaskCompletionLogById(id)
	if err != nil {
		log.Err(err).Msg("Error retrieving Task Completion Log")
		return err
	}

	DB.Delete(&tcl)

	return nil
}
