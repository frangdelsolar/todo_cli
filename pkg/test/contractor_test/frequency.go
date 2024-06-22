package contractor_test

import (
	"fmt"
	"strings"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
)

func TestCreateFrequency() {
    log.Info().Msg("Testing CreateFrequency()")

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

    log.Info().Msg("Created Frequency Successfully")

    // assertions

    if frq.Type != c.FrequencyTypeYearly {
        err = fmt.Errorf("expected type %s, got %s", c.FrequencyTypeYearly, frq.Type)
        log.Err(err).Msg("TestCreateFrequency()")
    } else {
        log.Debug().Msgf("Expected type %s, got %s", c.FrequencyTypeYearly, frq.Type)
    }

    if frq.Day != 12 {
        err = fmt.Errorf("expected day %d, got %d", 12, frq.Day)
        log.Err(err).Msg("TestCreateFrequency()")
    } else {
        log.Debug().Msgf("Expected day %d, got %d", 12, frq.Day)
    }

    if frq.Month != 10 {
        err = fmt.Errorf("expected month %d, got %d", 10, frq.Month)
        log.Err(err).Msg("TestCreateFrequency()")
    } else {
        log.Debug().Msgf("Expected month %d, got %d", 10, frq.Month)
    }

    if frq.DayOfWeek != 6 {
        err = fmt.Errorf("expected day of week %d, got %d", 6, frq.DayOfWeek)
        log.Err(err).Msg("TestCreateFrequency()")
    } else {
        log.Debug().Msgf("Expected day of week %d, got %d", 6, frq.DayOfWeek)
    }

    if frq.SystemData.CreatedByID != owner.ID {
        err = fmt.Errorf("expected requested by %d, got %d", owner.ID, frq.SystemData.CreatedByID)
        log.Err(err).Msg("TestCreateFrequency()")
    } else {
        log.Debug().Msgf("Expected requested by %d, got %d", owner.ID, frq.SystemData.CreatedByID)
    }
}

func TestFrequencyValidator() {

    log.Info().Msg("Testing FrequencyValidator()")

    // User validation
    frqInput := &c.NewFrequencyInput{
        FreqType: "",
        Day: "",
        Month: "",
        DayOfWeek: "",
        RequestedBy: "0",
    }

    err := frqInput.Validate()
    if err != nil {
        expected := "invalid user id"
        if strings.Contains(err.Error(), expected) {
            log.Info().Msgf("expected error %s, got %s", expected, err.Error())
        } else {
            log.Err(err).Msg("TestFrequencyValidator()")
        }
    } else {
        err := fmt.Errorf("expected error, got nil")
        log.Err(err).Msg("TestFrequencyValidator()")
    }

    // Daily validation
    frqInput = &c.NewFrequencyInput{
        FreqType: "daily",
        Day: "",
        Month: "",
        DayOfWeek: "",
        RequestedBy: "1",
    }

    err = frqInput.Validate()
    if err != nil {
        log.Err(err).Msg("Expected nil error, got error")
    } 

    // Weekly validation
    frqInput = &c.NewFrequencyInput{
        FreqType: "weekly",
        Day: "",
        Month: "",
        DayOfWeek: "",
        RequestedBy: "1",
    }

    err = frqInput.Validate()
    expected := "you must specify a day of the week"
    if err != nil {
        if strings.Contains(err.Error(), expected) {
            log.Info().Msgf("expected error %s, got %s", expected, err.Error())
        } else {
            log.Err(err).Msg("TestFrequencyValidator()")
        }
    } else {
        err := fmt.Errorf("expected error, got nil")
        log.Err(err).Msg("TestFrequencyValidator()")
    }

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
    if err != nil {
        if strings.Contains(err.Error(), expected) {
            log.Info().Msgf("expected error %s, got %s", expected, err.Error())
        } else {
            log.Err(err).Msg("TestFrequencyValidator()")
        }
    } else {
        err := fmt.Errorf("expected error, got nil")
        log.Err(err).Msg("TestFrequencyValidator()")
    }

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
    if err != nil {
        if strings.Contains(err.Error(), expected) {
            log.Info().Msgf("expected error %s, got %s", expected, err.Error())
        } else {
            log.Err(err).Msg("TestFrequencyValidator()")
        }
    } else {
        err := fmt.Errorf("expected error, got nil")
        log.Err(err).Msg("TestFrequencyValidator()")
    }
}
