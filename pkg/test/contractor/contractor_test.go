package contractor_test

import (
	"testing"

	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
	"github.com/frangdelsolar/todo_cli/pkg/test/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateContractor(t *testing.T) {
    ow, _ := helpers.CreateRandomUser()
    name := helpers.RandomName()
    contractor, err := c.CreateContractor(name, ow.GetIDString())
    if err != nil {
        t.Errorf("Failed to create contractor: %v", err)
    }
    assert.Equal(t, contractor.Name, name, "Expected name to be %s, but got %s", name, contractor.Name)
}

func TestUpdateContractorName(t *testing.T) {

    contractor, err := helpers.CreateRandomContractor()
    if err != nil {
        t.Errorf("Failed to create random contractor: %v", err)
    }

    newName := "Contractor 2"
    err = c.UpdateContractorName(contractor.GetIDString(), newName, contractor.CreatedBy.GetIDString())
    if err != nil {
        t.Errorf("Failed to update name: %v", err)
    }

    updated, err := c.GetContractorById(contractor.GetIDString(), contractor.CreatedBy.GetIDString())
    if err != nil {
        t.Errorf("Failed to get updated contractor: %v", err)
    }
    
    assert.Equal(t, updated.Name, newName, "Expected name to be %s, but got %s", newName, updated.Name)

}

func TestListContractor(t *testing.T) {

    c1, _ := helpers.CreateRandomContractor()
    c2, _ := helpers.CreateRandomContractor() // this would be created by a differnt user

    list1:= c.GetAllContractors(c1.CreatedBy.GetIDString())

    // iterate over the list of contractors and check if c1 and c2 are in the list
    c1Present := false
    c2Present := false
    for _, c := range list1 {
        if c.GetIDString() == c1.GetIDString() {
            c1Present = true
        }
        if c.GetIDString() == c2.GetIDString() {
            c2Present = true
        }
    }

    if !c1Present {
        t.Errorf("Contractor %s not found in list", c1.GetIDString())
    }

    if c2Present {
        t.Errorf("Contractor %s found in list", c2.GetIDString())
    }
    
}

func TestGetContractorById(t *testing.T) {

    c1, _ := helpers.CreateRandomContractor()
    ow1 := c1.CreatedBy
    ow2, _ := helpers.CreateRandomUser()

    _, err := c.GetContractorById(c1.GetIDString(), ow1.GetIDString())
    if err != nil {
        t.Errorf("Failed to get contractor: %v", err)
    }

    _, err = c.GetContractorById(c1.GetIDString(), ow2.GetIDString())
    if err == nil {
        t.Errorf("Expected error, but got none")
    }
    
}
