package contractor_test

import (
	"fmt"
	"testing"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
	"github.com/stretchr/testify/assert"
)


func TestCreateContractor(t *testing.T) {

    owner, _ := auth.CreateUser("owner", "owner@admin.com", "test123")
    userId := fmt.Sprint(owner.ID)

    name := "Contractor 1"

    contractor, err := c.CreateContractor(name, userId)
    if err != nil {
        t.Errorf("Failed to create contractor: %v", err)
    }

    assert.Equal(t, contractor.Name, name, "Expected name to be %s, but got %s", name, contractor.Name)
}

func TestUpdateContractorName(t *testing.T) {

    // Data prep
    owner, _ := auth.CreateUser("owner", "owner@admin.com", "test123")
    userId := fmt.Sprint(owner.ID)

    name := "Contractor 1"  

    contractor, err := c.CreateContractor(name, userId)
    if err != nil {
        t.Errorf("Failed to create contractor: %v", err)
    }

    newName := "Contractor 2"
    err = c.UpdateContractorName(fmt.Sprint(contractor.ID), newName, userId)
    if err != nil {
        t.Errorf("Failed to update name: %v", err)
    }

    updated, err := c.GetContractorById(fmt.Sprint(contractor.ID), userId)
    if err != nil {
        t.Errorf("Failed to get updated contractor: %v", err)
    }
    
    // assertions
    assert.Equal(t, updated.Name, newName, "Expected name to be %s, but got %s", newName, updated.Name)
}
