package contractor

import "fmt"

// CreateEffectivePeriod creates a new EffectivePeriod in the database based on the provided NewEffectivePeriodInput.
//
// Parameters:
// - input: a pointer to a NewEffectivePeriodInput struct containing the necessary information to create the EffectivePeriod.
//
// Returns:
// - ep: a pointer to the newly created EffectivePeriod.
// - err: an error if there was a problem creating the EffectivePeriod.
func CreateEffectivePeriod(input *NewEffectivePeriodInput) (*EffectivePeriod, error) {
    ep, err := NewEffectivePeriod(input)
    if err != nil {
        log.Err(err).Msg("Error creating effective period")
        return nil, err
    }

    db.Create(&ep)

    log.Trace().Interface("effective period", ep).Msg("Created effective period")
    log.Info().Msg("Created effective period")

    return ep, nil
}

// GetEffectivePeriodById retrieves an EffectivePeriod by ID and requestedBy.
//
// Parameters:
// - id: the ID of the EffectivePeriod to retrieve.
// - requestedBy: the user who requested the EffectivePeriod.
// Returns:
// - *EffectivePeriod: a pointer to the retrieved EffectivePeriod.
// - error: an error if the EffectivePeriod is not found.
func GetEffectivePeriodById(id string, requestedBy string) (*EffectivePeriod, error) {
    var ep EffectivePeriod
    
    db.First(&ep, "id = ? AND created_by_id = ?", id, requestedBy)
   
    if ep == (EffectivePeriod{}) {
        return nil, fmt.Errorf("effective period with ID %s not found", id)
    }
   
    return &ep, nil
}  

// DeleteEffectivePeriod deletes an EffectivePeriod from the database by its ID.
//
// Parameters:
// - id: the ID of the EffectivePeriod to delete.
// - requestedBy: the user who requested the deletion.
//
// Returns:
// - error: an error if the deletion fails.
func DeleteEffectivePeriod(id string, requestedBy string) error {

    ep, err := GetEffectivePeriodById(id, requestedBy)
    if err != nil {
        log.Err(err).Msg("Error deleting effective period")
        return err
    }

    db.Delete(&ep)

    log.Info().Msgf("Deleted effective period with ID %s", id)

    return nil
}
