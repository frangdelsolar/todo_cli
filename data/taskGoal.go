package data

import (
	"fmt"
	"todo_cli/models"

	"github.com/rs/zerolog/log"
)

// GetTaskGoalById retrieves an TaskGoal object by its ID.
//
// Parameters:
// - id: the ID of the TaskGoal to retrieve.
//
// Returns:
// - *models.TaskGoal: a pointer to the retrieved TaskGoal, or nil if not found.
// - error: an error if the TaskGoal with the provided ID is not found.
func GetTaskGoalById(id string) (models.TaskGoal, error) {
	var ep models.TaskGoal
	DB.First(&ep, "id = ?", id)
	if ep == (models.TaskGoal{}) {
		return ep, fmt.Errorf("task goal with ID %s not found", fmt.Sprint(id))
	}

	return ep, nil
}

// GetTaskGoalsByTaskId retrieves all TaskGoals associated with a given task ID.
//
// Parameters:
// - taskID: the ID of the task to retrieve TaskGoals for.
//
// Returns:
// - []models.TaskGoal: a slice of TaskGoals associated with the task ID.
func GetTaskGoalsByTaskId(taskID string) []models.TaskGoal {
	var eps []models.TaskGoal

	DB.Where("task_id = ?", taskID).Find(&eps)

	if len(eps) == 0 {
		log.Warn().Msg("No Task Goals found for task ID: " + fmt.Sprint(taskID))
		return eps
	}
	return eps
}

// NewTaskGoal creates a new TaskGoal in the database.
//
// Parameters:
// - taskID: the ID of the task associated with the TaskGoal.
// - startDate: the start date of the TaskGoal.
// - endDate: the end date of the TaskGoal.
//
// Returns:
// - *models.TaskGoal: the newly created TaskGoal.
// - error: an error if the TaskGoal creation fails.
func CreateTaskGoal(
	taskID string, 
	startDate string, 
	endDate string, 
	frequency string, 
	category string,
) (*models.TaskGoal, error) {
	ep, err := models.NewTaskGoal(
		taskID, 
		startDate, 
		endDate, 
		frequency, 
		category,
	)
	
	if err != nil {
		log.Err(err).Msg("Error creating new TaskGoal")
		return nil, err
	}

	DB.Create(ep)

	return ep, nil
}

// UpdateTaskGoal updates an TaskGoal in the database.
//
// Parameters:
// - id: the ID of the TaskGoal to update.
// - startDate: the new start date of the TaskGoal.
// - endDate: the new end date of the TaskGoal.
//
// Returns:
// - *models.TaskGoal: the updated TaskGoal.
// - error: an error if the TaskGoal retrieval or update fails.
func UpdateTaskGoal(id string, startDate string, endDate string) (models.TaskGoal, error) {
	var err error

	// Retrieve task goal
	ep, err := GetTaskGoalById(id)
	if err != nil {
		log.Err(err).Msg("Error getting TaskGoal")
		return ep, err
	}

	// Perform update
	err = ep.Update(startDate, endDate)
	if err != nil {
		log.Err(err).Msg("Error updating TaskGoal")
		return ep, err
	}

	DB.Save(&ep)

	return ep, nil
}

// DeleteTaskGoal deletes an TaskGoal from the DB by its ID.
//
// Parameters:
// - id: the ID of the TaskGoal to delete.
//
// Returns:
// - error: an error if the TaskGoal retrieval or deletion fails.
func DeleteTaskGoal(id string) error {
	// verify if taskPeriod exists
	ep, err := GetTaskGoalById(id)
	if err != nil {
		log.Err(err).Msg("Error getting TaskGoal")
		return err
	}

	DB.Delete(ep)

	return nil
}
