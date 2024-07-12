package models

import (
	"errors"
	"slices"
	"strconv"

	"gorm.io/gorm"
)

type FrequencyTypeName string

type TaskFrequency struct {
	gorm.Model
	ID        uint              `json:"id" gorm:"primaryKey"`
	Type      FrequencyTypeName `json:"type"`
	Day       int               `json:"day"`
	Month     int               `json:"month"`
	DayOfWeek int               `json:"dayOfWeek"`
}

const (
	Daily   FrequencyTypeName = "daily"
	Weekly  FrequencyTypeName = "weekly"
	Monthly FrequencyTypeName = "monthly"
	Yearly  FrequencyTypeName = "yearly"
)

func (ft *TaskFrequency) String() string {
	return string(ft.Type)
}

func NewTaskFrequency(
	freqType string,
	day string,
	month string,
	dayOfWeek string,
) (*TaskFrequency, error) {

	// var frequency TaskFrequency
	err := FrequencyValidator(
		freqType,
		day,
		month,
		dayOfWeek,
	)

	if err != nil {
		return nil, err
	}

	cday, _ := strconv.Atoi(day)
	cmonth, _ := strconv.Atoi(month)
	cdayOfWeek, _ := strconv.Atoi(dayOfWeek)

	return &TaskFrequency{
		Type:      FrequencyTypeName(freqType),
		Day:       cday,
		Month:     cmonth,
		DayOfWeek: cdayOfWeek,
	}, nil

}

// FrequencyValidator validates the given frequency string.
//
// It checks if the frequency is one of the following: Daily, Weekly, Monthly, Yearly.
// If the frequency is not valid, it returns an error with the message "invalid frequency".
// If the frequency is valid, it returns nil.
//
// Parameters:
// - frequency: a string representing the frequency to be validated.
//
// Returns:
// - error: an error if the frequency is invalid, otherwise nil.
func FrequencyValidator(
	in_frequency_name string,
	in_frequency_day string,
	in_frequency_month string,
	in_frequency_day_of_week string,
) error {
	switch in_frequency_name {
	case string(Daily):
		return nil

	case string(Weekly):
		return DayOfWeekValidator(in_frequency_day_of_week)

	case string(Monthly):
		return DayValidator(in_frequency_day)

	case string(Yearly):
		return DayOfMonthValidator(in_frequency_day, in_frequency_month)

	default:
		return errors.New("invalid frequency")
	}
}

// CategoryValidator validates the given category string.
//
// It checks if the category is equal to the string representation of the Todo constant.
// If the category is valid, it returns nil. Otherwise, it returns an error with the message "invalid category".
//
// Parameters:
// - category: a string representing the category to be validated.
//
// Returns:
// - error: an error if the category is invalid, otherwise nil.
func CategoryValidator(category string) error {
	if category == string(Todo) {
		return nil
	}
	return errors.New("invalid category")
}

func DayOfWeekValidator(day string) error {
	if day == "" {
		return errors.New("you must specify a day of the week")
	}

	converted, err := strconv.Atoi(day)
	if err != nil {
		return errors.New("invalid day of the week")
	}

	if converted < 1 || converted > 7 {
		return errors.New("day of the week must be between 1 and 7")
	}

	return nil
}

func DayValidator(day string) error {
	if day == "" {
		return errors.New("you must specify a day of the month")
	}

	converted, err := strconv.Atoi(day)
	if err != nil {
		return errors.New("invalid day of the month")
	}

	if converted < 1 || converted > 31 {
		return errors.New("day of the month must be between 1 and 31")
	}

	return nil
}

func MonthValidator(month string) error {
	if month == "" {
		return errors.New("you must specify a month")
	}

	converted, err := strconv.Atoi(month)
	if err != nil {
		return errors.New("invalid month")
	}

	if converted < 1 || converted > 12 {
		return errors.New("month must be between 1 and 12")
	}

	return nil
}

func DayOfMonthValidator(day string, month string) error {
	err := MonthValidator(month)
	if err != nil {
		return err
	}
	convertedMonth, _ := strconv.Atoi(month)

	err = DayValidator(day)
	if err != nil {
		return err
	}
	convertedDay, _ := strconv.Atoi(day)

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
