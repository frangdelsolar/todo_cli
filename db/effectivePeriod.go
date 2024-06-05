package db

import (
	"fmt"
	"todo_cli/models"

	"github.com/gofrs/uuid"
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
func (db *DB) GetEffectivePeriodById(id string) (*models.EffectivePeriod, error) {
	ep := db.EffectivePeriods[id]
	if ep.ID == uuid.Nil {
		return nil, fmt.Errorf("EffectivePeriod with ID %s not found", id)
	}
	
	return &ep, nil
}

// GetEffectivePeriodsByTaskId retrieves all EffectivePeriods associated with a given task ID.
//
// Parameters:
// - taskID: the ID of the task to retrieve EffectivePeriods for.
//
// Returns:
// - []models.EffectivePeriod: a slice of EffectivePeriods associated with the task ID.
func (db *DB) GetEffectivePeriodsByTaskId(taskID string) []models.EffectivePeriod {
	eps := []models.EffectivePeriod{}
	for _, effectivePeriod := range db.EffectivePeriods {
		if effectivePeriod.TaskID == uuid.Must(uuid.FromString(taskID)) {
			eps = append(eps, effectivePeriod)
		}
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
func (db *DB) CreateEffectivePeriod(taskID string, startDate string, endDate string) (*models.EffectivePeriod, error) {
	ep, err := models.NewEffectivePeriod(taskID, startDate, endDate)
	if err != nil {
		log.Err(err).Msg("Error creating new EffectivePeriod")
		return nil, err
	}
	db.Save()
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
func (db *DB) UpdateEffectivePeriod(id string, startDate string, endDate string) (*models.EffectivePeriod, error) {
	var err error
	
	// Retrieve effective period
	ep, err := db.GetEffectivePeriodById(id)
	if err != nil {
		log.Err(err).Msg("Error getting EffectivePeriod")
		return nil, err
	}
	
	// Perform update
	err = ep.Update(startDate, endDate)
	if err != nil {
		log.Err(err).Msg("Error updating EffectivePeriod")
		return nil, err
	}

	db.Save()
	
	return ep, nil
}

// DeleteEffectivePeriod deletes an EffectivePeriod from the DB by its ID.
//
// Parameters:
// - id: the ID of the EffectivePeriod to delete.
//
// Returns:
// - error: an error if the EffectivePeriod retrieval or deletion fails.
func (db *DB) DeleteEffectivePeriod(id string) error {
	// verify if taskPeriod exists
	_, err := db.GetEffectivePeriodById(id)
	if err != nil {
		log.Err(err).Msg("Error getting EffectivePeriod")
		return err
	}

	delete(db.EffectivePeriods, id)

	db.Save()

	return nil
}
