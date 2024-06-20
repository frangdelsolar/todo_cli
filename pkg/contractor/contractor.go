package contractor

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	data "github.com/frangdelsolar/todo_cli/pkg/data/models"
)


type Contractor struct{
    data.SystemData
    Name string
}

// UpdateName updates the name of the Contractor and sets the updatedBy field to the provided auth.User.
//
// Parameters:
// - name: the new name for the Contractor.
// - requestedBy: the auth.User who requested the name update.
//
// Returns:
// - error: an error if the name is invalid.
func (c *Contractor) UpdateName (name string, requestedBy *auth.User) error {
    if err := NameValidator(name); err != nil {
        log.Err(err).Msg("Error validating name")
        return err
    }
    c.Name = name
    c.UpdatedBy = requestedBy
    return nil
}

// NewContractor creates a new Contractor with the given name.
//
// Parameters:
// - name: the name of the Contractor.
//
// Returns:
// - *Contractor: a pointer to the newly created Contractor.
// - error: an error if the name is invalid.
func NewContractor(name string, requestedBy string) (*Contractor, error) {
    user, error := auth.GetUserById(requestedBy)
    if error != nil {
        log.Err(error).Msg("Error getting user")
        return nil, error
    }

    if err := NameValidator(name); err != nil {
        log.Err(err).Msg("Error validating name")
        return nil, err
    }

    return &Contractor{
        Name: name,
        SystemData: data.SystemData{
            CreatedBy: user,
            UpdatedBy: user,
        },
    }, nil
}

// NameValidator validates the given name.
//
// Parameters:
// - name: the name to be validated.
//
// Returns:
// - error: an error if the name is empty, otherwise nil.
func NameValidator(name string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}
