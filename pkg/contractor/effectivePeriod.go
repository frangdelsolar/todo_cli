package contractor

import (
	"errors"
	"strconv"
	"time"

	data "github.com/frangdelsolar/todo_cli/pkg/data/models"
)


type NewEffectivePeriodInput struct {
    StartDate string
    EndDate string
    RequestedBy string
}

// Validate validates the NewEffectivePeriodInput struct.
//
// It checks if the RequestedBy field is a valid ID, if the StartDate and EndDate fields are valid dates,
// and if the StartDate is not after the EndDate.
//
// Returns:
// - error: an error if any of the validations fail.
func (nfi *NewEffectivePeriodInput) Validate() error {

    if ValidateID(nfi.RequestedBy) != nil {
        return errors.New("invalid user id")
    }

    err := DateValidator(nfi.StartDate)
    if err != nil {
        return errors.New("invalid start date")
    }

    // end date can be empty
    if nfi.EndDate == "" {
        return nil
    }

    err = DateValidator(nfi.EndDate)
    if err != nil {
        return errors.New("invalid end date")
    }

    st, _ := time.Parse(time.DateOnly, nfi.StartDate)
    ed, _ := time.Parse(time.DateOnly, nfi.EndDate)
    
    if st.After(ed) {
        return errors.New("start date is after end date")
    }
    
    return nil
}

// NewEffectivePeriod creates a new EffectivePeriod based on the input provided.
//
// Parameters:
// - input: a NewEffectivePeriodInput struct containing the necessary information.
// Returns:
// - *EffectivePeriod: the newly created EffectivePeriod.
// - error: an error if the creation process encounters any issues.
func NewEffectivePeriod(input *NewEffectivePeriodInput) (*EffectivePeriod, error) {

    err := input.Validate()
    if err != nil {
        log.Err(err).Msg("Error validating new effective period")
        return nil, err
    }

    st, _ := time.Parse(time.DateOnly, input.StartDate)
    ed, _ := time.Parse(time.DateOnly, input.EndDate)
    rb, _ := strconv.Atoi(input.RequestedBy)
    
    return &EffectivePeriod{
        StartDate: st, 
        EndDate: ed, 
        SystemData: data.SystemData{
            CreatedByID: uint(rb),
            UpdatedByID: uint(rb),
        },    
    }, nil

}

// DateValidator validates a date string in the format "YYYY-MM-DD".
//
// Parameters:
// - date: the date string to validate.
//
// Returns:
// - error: an error if the date string is not in the correct format or if the date is invalid, otherwise nil.
func DateValidator(date string) error {
	_, err := time.Parse(time.DateOnly, date)
	return err
}
