package contractor_test

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
	"github.com/frangdelsolar/todo_cli/pkg/test"
)

func TestCreateFrequency(t *test.TestRunner) {
    t.Run("TestCreateFrequency()")

    // Data prep
    owner, _ := auth.CreateUser("owner", "owner@admin.com", "test123")
    userId := fmt.Sprint(owner.ID)

    frqInput := &c.NewFrequencyInput{
        FreqType: "yearly",
        Day: "12",
        Month: "10",
        DayOfWeek: "6",
        RequestedBy: userId,
    }

    frq, err := c.CreateFrequency(frqInput)
    if err != nil {
        log.Err(err).Msg("Failed to create contractor")
    }

    t.AssertEqual(frq.Type, c.FrequencyTypeYearly)
    t.AssertEqual(frq.Day, 12)
    t.AssertEqual(frq.Month, 10)
    t.AssertEqual(frq.DayOfWeek, 6)
    t.AssertEqual(frq.SystemData.CreatedByID, owner.ID)

    t.Stop()
}

func TestFrequencyValidator(t *test.TestRunner) {
    t.Run("TestFrequencyValidator()")

    // User validation
    frqInput := &c.NewFrequencyInput{
        FreqType: "",
        Day: "",
        Month: "",
        DayOfWeek: "",
        RequestedBy: "0",
    }

    err := frqInput.Validate()
    expected := "invalid user id"

    t.AssertErrorContains(err, expected)

    // Daily validation
    frqInput = &c.NewFrequencyInput{
        FreqType: "daily",
        Day: "",
        Month: "",
        DayOfWeek: "",
        RequestedBy: "1",
    }

    err = frqInput.Validate()
    
    t.AssertErrorNil(err)

    // Weekly validation
    frqInput = &c.NewFrequencyInput{
        FreqType: "weekly",
        Day: "",
        Month: "",
        DayOfWeek: "",
        RequestedBy: "1",
    }

    err = frqInput.Validate()
    expected = "you must specify a day of the week"

    t.AssertErrorContains(err, expected)

    // Monthly validation
    frqInput = &c.NewFrequencyInput{
        FreqType: "monthly",
        Day: "",
        Month: "",
        DayOfWeek: "",
        RequestedBy: "1",
    }

    err = frqInput.Validate()
    expected = "you must specify a day of the month"

    t.AssertErrorContains(err, expected)

    // Yearly validation
    frqInput = &c.NewFrequencyInput{
        FreqType: "yearly",
        Day: "12",
        Month: "",
        DayOfWeek: "",
        RequestedBy: "1",
    }

    err = frqInput.Validate()
    expected = "you must specify a month"

    t.AssertErrorContains(err, expected)

    t.Stop()
}
