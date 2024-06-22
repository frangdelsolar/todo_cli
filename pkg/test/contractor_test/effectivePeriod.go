package contractor_test

import (
	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
	"github.com/frangdelsolar/todo_cli/pkg/test"
)

func TestValidateEffectivePeriod(t *test.TestRunner) {
    t.Run("TestValidateEffectivePeriod()")

    epInput := &c.NewEffectivePeriodInput{
        StartDate: "2022-01-01",
        EndDate: "2022-01-01",
        RequestedBy: "",
    }

    err := epInput.Validate()
    expected := "invalid user id"
    
    t.AssertErrorContains(err, expected)
    t.Stop()
}
