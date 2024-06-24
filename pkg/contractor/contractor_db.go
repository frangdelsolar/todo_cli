package contractor

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
)

// CreateContractor creates a new Contractor with the given name and requestedBy.
//
// Parameters:
// - name: the name of the Contractor.
// - requestedBy: the ID of the user who requested the creation of the Contractor.
//
// Returns:
// - *Contractor: a pointer to the newly created Contractor.
// - error: an error if the creation of the Contractor fails.
func CreateContractor(name string, requestedBy string) (*Contractor, error) {
    c, err := NewContractor(name, requestedBy)
    if err != nil {
        log.Err(err).Msg("Error creating new contractor")
        return c, err
    }

    db.Create(&c)

    log.Trace().Interface("Contractor", c).Msg("Created new contractor")
    log.Info().Msg("Created new contractor")

    return c, nil
}

// UpdateContractorName updates the name of a Contractor with the given ID and requestedBy.
//
// Parameters:
// - id: the ID of the Contractor.
// - name: the new name for the Contractor.
// - requestedBy: the ID of the user requesting the update.
// Returns:
// - error: an error if there was a problem updating the name.
func UpdateContractorName(id string, name string, requestedBy string) error {
    c, err := GetContractorById(id, requestedBy)
    if err != nil {
        log.Err(err).Msg("Error getting contractor")
        return err
    }

    user, err := auth.GetUserById(requestedBy)
    if err != nil {
        log.Err(err).Msg("Error getting user")
        return err
    }

    if err = c.UpdateName(name, user); err != nil {
        log.Err(err).Msg("Error updating name")
        return err
    }

    db.Save(&c)

    log.Trace().Interface("Contractor", c).Msg("Updated name of contractor")
    log.Info().Msg("Updated name of contractor")

    return nil
}


// GetContractorById retrieves a Contractor by ID and requestedBy.
//
// Parameters:
// - id: the ID of the Contractor to retrieve.
// - requestedBy: the ID of the user requesting the Contractor.
// Returns:
// - Contractor: the retrieved Contractor.
// - error: an error if the Contractor is not found.
func GetContractorById(id string, requestedBy string) (Contractor, error) {

    var c Contractor
    
    db.First(&c, "id = ? AND created_by_id = ?", id, requestedBy)
    
    if c == (Contractor{}) {
        return c, fmt.Errorf("contractor with ID %s not found", fmt.Sprint(id))
    }

    return c, nil
}


// GetAllContractors retrieves all Contractors from the database.
//
// Parameters:
// - requestedBy: the ID of the user requesting the list of Contractors.
//
// Returns:
// - []Contractor: a slice of all the retrieved Contractors.
func GetAllContractors(requestedBy string) []Contractor {
    var cs []Contractor

    db.Find(&cs, "created_by_id = ?", requestedBy)

    if len(cs) == 0 {
        log.Warn().Msg("No contractors found")
    }
    return cs
}

// DeleteContractor deletes a Contractor from the database by its ID and requestedBy.
//
// Parameters:
// - id: the ID of the Contractor to be deleted.
// - requestedBy: the ID of the user who requested the deletion.
//
// Returns:
// - error: an error if there was a problem deleting the Contractor.
func DeleteContractor(id string, requestedBy string) error {
    c, err := GetContractorById(id, requestedBy)
    if err != nil {
        log.Err(err).Msg("Error getting contractor")
        return err
    }

    db.Delete(&c)

    log.Trace().Interface("Contractor", c).Msg("Deleted contractor")
    log.Info().Msg("Deleted contractor")
    
    return nil
}
