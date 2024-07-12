package contractor_test

import (
	"fmt"
	"testing"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
	"github.com/stretchr/testify/assert"
)

func TestCreateFrequency(t *testing.T) {

	// Data prep
	owner, _ := auth.CreateUser("owner", "owner@admin.com", "test123")
	userId := fmt.Sprint(owner.ID)

	frqInput := &c.NewFrequencyInput{
		FreqType:    "yearly",
		Day:         "12",
		Month:       "10",
		DayOfWeek:   "6",
		RequestedBy: userId,
	}

	frq, err := c.CreateFrequency(frqInput)
	if err != nil {
		t.Errorf("Failed to create frequency: %v", err)
	}
	assert.Equal(t, frq.Type.String(), frqInput.FreqType, "Expected type to be %s, but got %s", frqInput.FreqType, frq.Type)
	assert.Equal(t, fmt.Sprint(frq.Day), frqInput.Day, "Expected day to be %s, but got %s", frqInput.Day, frq.Day)
	assert.Equal(t, fmt.Sprint(frq.Month), frqInput.Month, "Expected month to be %s, but got %s", frqInput.Month, frq.Month)
	assert.Equal(t, fmt.Sprint(frq.DayOfWeek), frqInput.DayOfWeek, "Expected day of week to be %s, but got %s", frqInput.DayOfWeek, frq.DayOfWeek)
	assert.Equal(t, fmt.Sprint(frq.SystemData.CreatedByID), frqInput.RequestedBy, "Expected created by id to be %s, but got %s", frqInput.RequestedBy, frq.SystemData.CreatedByID)

}

func TestFrequencyValidator(t *testing.T) {
	// User validation
	frqInput := &c.NewFrequencyInput{
		FreqType:    "",
		Day:         "",
		Month:       "",
		DayOfWeek:   "",
		RequestedBy: "0",
	}

	validationErrMsg := frqInput.Validate()
	expected := "invalid user id"

	assert.ErrorContains(t, validationErrMsg, expected)

	// Daily validation
	frqInput = &c.NewFrequencyInput{
		FreqType:    "daily",
		Day:         "",
		Month:       "",
		DayOfWeek:   "",
		RequestedBy: "1",
	}

	validationErrMsg = frqInput.Validate()

	assert.Nil(t, validationErrMsg, "Expected validation error to be nil")

	// Weekly validation
	frqInput = &c.NewFrequencyInput{
		FreqType:    "weekly",
		Day:         "",
		Month:       "",
		DayOfWeek:   "",
		RequestedBy: "1",
	}

	validationErrMsg = frqInput.Validate()
	expected = "you must specify a day of the week"

	assert.ErrorContains(t, validationErrMsg, expected, "Expected validation error to be %s, but got %s", expected, validationErrMsg)

	// Monthly validation
	frqInput = &c.NewFrequencyInput{
		FreqType:    "monthly",
		Day:         "",
		Month:       "",
		DayOfWeek:   "",
		RequestedBy: "1",
	}

	validationErrMsg = frqInput.Validate()
	expected = "you must specify a day of the month"

	assert.ErrorContains(t, validationErrMsg, expected, "Expected validation error to be %s, but got %s", expected, validationErrMsg)

	// Yearly validation
	frqInput = &c.NewFrequencyInput{
		FreqType:    "yearly",
		Day:         "12",
		Month:       "",
		DayOfWeek:   "",
		RequestedBy: "1",
	}

	validationErrMsg = frqInput.Validate()
	expected = "you must specify a month"

	assert.ErrorContains(t, validationErrMsg, expected, "Expected validation error to be %s, but got %s", expected, validationErrMsg)

}
