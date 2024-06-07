package data

import (
	"fmt"
	"todo_cli/models"

	"github.com/rs/zerolog/log"
)

// GetEffectivePeriodById retrieves an EffectivePeriod object by its ID.
//
// Parameters:
// - id: the ID of the EffectivePeriod to retrieve.
//
// Returns:
// - *models.EffectivePeriod: a pointer to the retrieved EffectivePeriod, or nil if not found.
// - error: an error if the EffectivePeriod with the provided ID is not found.
func GetEffectivePeriodById(id uint) (models.EffectivePeriod, error) {
	var ep models.EffectivePeriod
	DB.First(&ep, "id = ?", id)
	if ep == (models.EffectivePeriod{}) {
		return ep, fmt.Errorf("effective period with ID %s not found", fmt.Sprint(id))
	}

	return ep, nil
}

// GetEffectivePeriodsByTaskId retrieves all EffectivePeriods associated with a given task ID.
//
// Parameters:
// - taskID: the ID of the task to retrieve EffectivePeriods for.
//
// Returns:
// - []models.EffectivePeriod: a slice of EffectivePeriods associated with the task ID.
func GetEffectivePeriodsByTaskId(taskID uint) []models.EffectivePeriod {
	var eps []models.EffectivePeriod

	DB.Where("task_id = ?", taskID).Find(&eps)

	if len(eps) == 0 {
		log.Warn().Msg("No Effective Periods found for task ID: " + fmt.Sprint(taskID))
		return eps
	}
	return eps
}

// NewEffectivePeriod creates a new EffectivePeriod in the database.
//
// Parameters:
// - taskID: the ID of the task associated with the EffectivePeriod.
// - startDate: the start date of the EffectivePeriod.
// - endDate: the end date of the EffectivePeriod.
//
// Returns:
// - *models.EffectivePeriod: the newly created EffectivePeriod.
// - error: an error if the EffectivePeriod creation fails.
func CreateEffectivePeriod(taskID uint, startDate string, endDate string, frequency string, category string) (*models.EffectivePeriod, error) {
	ep, err := models.NewEffectivePeriod(taskID, startDate, endDate, frequency, category)
	if err != nil {
		log.Err(err).Msg("Error creating new EffectivePeriod")
		return nil, err
	}

	DB.Create(ep)

	return ep, nil
}

// UpdateEffectivePeriod updates an EffectivePeriod in the database.
//
// Parameters:
// - id: the ID of the EffectivePeriod to update.
// - startDate: the new start date of the EffectivePeriod.
// - endDate: the new end date of the EffectivePeriod.
//
// Returns:
// - *models.EffectivePeriod: the updated EffectivePeriod.
// - error: an error if the EffectivePeriod retrieval or update fails.
func UpdateEffectivePeriod(id uint, startDate string, endDate string) (models.EffectivePeriod, error) {
	var err error

	// Retrieve effective period
	ep, err := GetEffectivePeriodById(id)
	if err != nil {
		log.Err(err).Msg("Error getting EffectivePeriod")
		return ep, err
	}

	// Perform update
	err = ep.Update(startDate, endDate)
	if err != nil {
		log.Err(err).Msg("Error updating EffectivePeriod")
		return ep, err
	}

	DB.Save(&ep)

	return ep, nil
}

// DeleteEffectivePeriod deletes an EffectivePeriod from the DB by its ID.
//
// Parameters:
// - id: the ID of the EffectivePeriod to delete.
//
// Returns:
// - error: an error if the EffectivePeriod retrieval or deletion fails.
func DeleteEffectivePeriod(id uint) error {
	// verify if taskPeriod exists
	ep, err := GetEffectivePeriodById(id)
	if err != nil {
		log.Err(err).Msg("Error getting EffectivePeriod")
		return err
	}

	DB.Delete(ep)

	return nil
}
