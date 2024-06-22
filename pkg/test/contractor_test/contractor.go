package contractor_test

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
	"github.com/frangdelsolar/todo_cli/pkg/test"
)


func TestCreateContractor(t *test.TestRunner) {
    t.Run("TestCreateContractor()")

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
    t.AssertEqual(contractor.Name, name)

    t.Stop()
}

func TestUpdateContractorName(t *test.TestRunner) {
    t.Run("TestUpdateContractorName()")

    // Data prep
    owner, _ := auth.CreateUser("owner", "owner@admin.com", "test123")
    userId := fmt.Sprint(owner.ID)

    name := "Contractor 1"  

    contractor, err := c.CreateContractor(name, userId)
    if err != nil {
        log.Err(err).Msg("Failed to create contractor")
    }

    newName := "Contractor 2"
    err = c.UpdateContractorName(fmt.Sprint(contractor.ID), newName, userId)
    if err != nil {
        log.Err(err).Msg("Failed to update contractor name")
    }

    updated, _ := c.GetContractorById(fmt.Sprint(contractor.ID), userId)
    
    // assertions
    t.AssertEqual(updated.Name, newName)

    t.Stop()
}
