package contractor_test

import (
	"testing"

	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
	"github.com/stretchr/testify/assert"
)

func TestValidateEffectivePeriod(t *testing.T) {
    epInput := &c.NewEffectivePeriodInput{
        StartDate: "2022-01-01",
        EndDate: "2022-01-01",
        RequestedBy: "",
    }

    err := epInput.Validate()
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    expected := "invalid user id"

    assert.ErrorContains(t, err, expected)
}
