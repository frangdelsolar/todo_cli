package helpers

import (
	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
)

func CreateRandomContractor() (*c.Contractor, error) {
    owner, _:= CreateRandomUser()
    name := RandomName()
    return c.CreateContractor(name, owner.GetIDString())
}
