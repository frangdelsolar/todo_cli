package contractor_cli

import (
	"fmt"

	a "github.com/frangdelsolar/todo_cli/cmd/cli/auth_cli"
	c "github.com/frangdelsolar/todo_cli/pkg/contractor"
	"github.com/spf13/cobra"
)

var ContractorListCmd = &cobra.Command{
		Use:   "list",
		Short: "List of Contractors",
		Long: `List of all Contractors.`,
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("List of all Contractors")
            contractors := c.GetAllContractors(a.GetUserId())
            fmt.Printf("Contractors: %d\n", len(contractors))
            for _, c := range contractors {
                fmt.Println(c)
            }
        },
}

