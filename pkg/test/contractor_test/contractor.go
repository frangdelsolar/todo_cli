package contractor_test

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
)


func TestCreateContractor() {
    log.Info().Msg("Testing CreateContractor()")

    // Data prep
    owner, _ := auth.CreateUser("owner", "owner@admin.com", "test123")
    userId := fmt.Sprint(owner.ID)

    name := "Contractor 1"

    contractor, err := c.CreateContractor(name, userId)
    if err != nil {
        log.Err(err).Msg("Failed to create contractor")
    }

    log.Info().Msg("Created Contractor Successfully")

    // assertions

    if contractor.Name != name {
        err = fmt.Errorf("expected name %s, got %s", name, contractor.Name)
        log.Err(err).Msg("TestCreateContractor()")
    } else {
        log.Debug().Msgf("Expected name %s, got %s", name, contractor.Name)
    }
}

func TestUpdateContractorName() {
    log.Info().Msg("Testing UpdateContractorName()")

    // Data prep
    owner, _ := auth.CreateUser("owner", "owner@admin.com", "test123")
    userId := fmt.Sprint(owner.ID)

    name := "Contractor 1"  

    contractor, err := c.CreateContractor(name, userId)
    if err != nil {
        log.Err(err).Msg("Failed to create contractor")
    }

    log.Trace().Interface("Contractor", contractor).Msg("Created Contractor Successfully")   

    newName := "Contractor 2"
    err = c.UpdateContractorName(fmt.Sprint(contractor.ID), newName, userId)
    if err != nil {
        log.Err(err).Msg("Failed to update contractor name")
    }

    log.Info().Msg("Updated Contractor Name Successfully")

    // assertions
    updated, _ := c.GetContractorById(fmt.Sprint(contractor.ID), userId)
    if updated.Name != newName {
        err = fmt.Errorf("expected name %s, got %s", newName, updated.Name)
        log.Err(err).Msg("TestUpdateContractorName()")
    } else {
        log.Debug().Msgf("Expected name %s, got %s", newName, updated.Name)
    }
}
