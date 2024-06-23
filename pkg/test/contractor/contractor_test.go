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

func TestListContractor(t *testing.T) {

    owner1, _ := auth.CreateUser("owner1", "owner@admin.com", "test123")
    userId1 := fmt.Sprint(owner1.ID)

    name1 := "Contractor ow1"

    _, err := c.CreateContractor(name1, userId1)
    if err != nil {
        t.Errorf("Failed to create contractor: %v", err)
    }

    owner2, _ := auth.CreateUser("owner2", "owner@admin.com", "test123")
    userId2 := fmt.Sprint(owner2.ID)

    name2 := "Contractor ow2"

    _, err = c.CreateContractor(name2, userId2)
    if err != nil {
        t.Errorf("Failed to create contractor: %v", err)
    }

    list1:= c.GetAllContractors(userId1)
    assert.Equal(t, len(list1), 1, "Expected 1 contractor, but got %d", len(list1))
    assert.Equal(t, list1[0].Name, name1, "Expected name to be %s, but got %s", name1, list1[0].Name)

    list2:= c.GetAllContractors(userId2)
    assert.Equal(t, len(list2), 1, "Expected 1 contractor, but got %d", len(list2))
    assert.Equal(t, list2[0].Name, name2, "Expected name to be %s, but got %s", name2, list2[0].Name)
    
}
