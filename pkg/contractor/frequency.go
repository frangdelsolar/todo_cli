package contractor

import (
	"errors"
	"slices"
	"strconv"

	data "github.com/frangdelsolar/todo_cli/pkg/data/models"
)

type FrequencyType string

const (
    FrequencyTypeDaily FrequencyType = "daily"
    FrequencyTypeWeekly FrequencyType = "weekly"
    FrequencyTypeMonthly FrequencyType = "monthly"
    FrequencyTypeYearly FrequencyType = "yearly"
)

type Frequency struct {
    data.SystemData
    Type FrequencyType `gorm:"not null"`
    Day int
    Month int
    DayOfWeek int
}

type NewFrequencyInput struct {
    FreqType string
    Day string
    Month string
    DayOfWeek string
    RequestedBy string
}

// Validate validates the frequency input.
//
// It checks the frequency type and calls the corresponding validator function.
// If the frequency type is Daily, it returns nil. If it's Weekly, it calls
// DayOfWeekValidator(). If it's Monthly, it calls DayValidator(). If it's
// Yearly, it calls DayOfMonthValidator(). If the frequency type is not valid,
// it returns an error with the message "invalid frequency type".
//
// Parameters:
// - nfi: a pointer to NewFrequencyInput struct.
//
// Return type:
// - error: an error if the frequency type is invalid, otherwise nil.
func (nfi *NewFrequencyInput) Validate() error {

    if ValidateID(nfi.RequestedBy) != nil {
        return errors.New("invalid user id")
    }

    switch nfi.FreqType {
    case string(FrequencyTypeDaily):
        return nil
    case string(FrequencyTypeWeekly):
        return nfi.DayOfWeekValidator()
    case string(FrequencyTypeMonthly):
        return nfi.DayValidator()
    case string(FrequencyTypeYearly):
        return nfi.DayOfMonthValidator()
    default:
        return errors.New("invalid frequency type")
    }
}

// DayOfWeekValidator validates if the given day of the week is valid.
//
// It takes a string parameter `day` representing the day of the week to be validated.
// The function returns an error if the day of the week is invalid, otherwise it returns nil.
//
// Parameters:
// - day: a string representing the day of the week to be validated.
//
// Return type:
// - error: an error if the day of the week is invalid, otherwise nil.
func (nfi *NewFrequencyInput) DayOfWeekValidator() error {
	if (nfi.DayOfWeek == "" ){
		return errors.New("you must specify a day of the week")
	}

	converted, err := strconv.Atoi(nfi.DayOfWeek)
	if err != nil {
		return errors.New("invalid day of the week")
	}

	if converted < 1 || converted > 7 {
		return errors.New("day of the week must be between 1 and 7")
	}

	return nil
}

// DayValidator validates if the given day of the month is valid.
//
// It takes a string parameter `day` representing the day of the month to be validated.
// The function returns an error if the day of the month is invalid, otherwise it returns nil.
//
// Parameters:
// - day: a string representing the day of the month to be validated.
//
// Return type:
// - error: an error if the day of the month is invalid, otherwise nil.
func (nfi *NewFrequencyInput) DayValidator() error {
	if (nfi.Day == "" ){
		return errors.New("you must specify a day of the month")
	}

	converted, err := strconv.Atoi(nfi.Day)
	if err != nil {
		return errors.New("invalid day of the month")
	}

	if converted < 1 || converted > 31 {
		return errors.New("day of the month must be between 1 and 31")
	}

	return nil
}

// MonthValidator validates if the given month is valid.
//
// It takes a string parameter `month` representing the month to be validated.
// The function returns an error if the month is invalid, otherwise it returns nil.
//
// Parameters:
// - month: a string representing the month to be validated.
//
// Return type:
// - error: an error if the month is invalid, otherwise nil.
func (nfi *NewFrequencyInput) MonthValidator() error {
	if (nfi.Month == "" ){
		return errors.New("you must specify a month")
	}

	converted, err := strconv.Atoi(nfi.Month)
	if err != nil {
		return errors.New("invalid month")
	}

	if converted < 1 || converted > 12 {
		return errors.New("month must be between 1 and 12")
	}

	return nil
}

// DayOfMonthValidator validates if the given day of the month is valid for the specified month.
//
// It takes two string parameters: `day` representing the day of the month to be validated,
// and `month` representing the month to be validated.
// The function returns an error if the day of the month is invalid for the specified month,
// otherwise it returns nil.
//
// Parameters:
// - day: a string representing the day of the month to be validated.
// - month: a string representing the month to be validated.
//
// Return type:
// - error: an error if the day of the month is invalid for the specified month,
//          otherwise nil.
func (nfi *NewFrequencyInput) DayOfMonthValidator() error {
	err:= nfi.MonthValidator()
	if err != nil {
		return err
	}
	convertedMonth, _ := strconv.Atoi(nfi.Month)

	err = nfi.DayValidator()
	if err != nil {
		return err
	}
	convertedDay, _ := strconv.Atoi(nfi.Day)

	monthsThirtyOne := []int{1, 3, 5, 7, 8, 10, 12}
	monthsThirty := []int{4, 6, 9, 11}

	if slices.Contains(monthsThirtyOne, convertedMonth) && convertedDay > 31 {
		return errors.New("day of the month must be between 1 and 31")
	}
	if slices.Contains(monthsThirty, convertedMonth) && convertedDay > 30 {
		return errors.New("day of the month must be between 1 and 30")
	}
	if convertedMonth == 2 && convertedDay > 28 {
		return errors.New("day of the month must be between 1 and 28")
	}

	return nil
}

// NewFrequency creates a new Frequency object based on the provided NewFrequencyInput.
//
// Parameters:
// - freqInput: a pointer to a NewFrequencyInput struct containing the frequency information.
//
// Returns:
// - a pointer to a Frequency struct representing the new frequency.
// - an error if the input is invalid.
func NewFrequency(freqInput *NewFrequencyInput)(*Frequency, error) {

    err := freqInput.Validate()
    if err != nil {
        log.Err(err).Msg("Error validating new frequency")
        return nil, err
    }

    freqType := FrequencyType(freqInput.FreqType)
    day, _ := strconv.Atoi(freqInput.Day)
    month, _ := strconv.Atoi(freqInput.Month)
    dayOfWeek, _ := strconv.Atoi(freqInput.DayOfWeek)
    requestedById, _ := strconv.Atoi(freqInput.RequestedBy)

    return &Frequency{
        Type: freqType,
        Day: day,
        Month: month,
        DayOfWeek: dayOfWeek,
        SystemData: data.SystemData{
            CreatedByID: uint(requestedById),
            UpdatedByID: uint(requestedById),
        },
    }, nil

}
