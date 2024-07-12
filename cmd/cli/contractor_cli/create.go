package contractor_cli

import (
	"fmt"

	auth "github.com/frangdelsolar/todo_cli/cmd/cli/auth_cli"
	"github.com/frangdelsolar/todo_cli/cmd/cli/utils"
	"github.com/frangdelsolar/todo_cli/pkg/contractor"
	"github.com/spf13/cobra"
)

var ContractorCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Contractor",
	Long:  `Create a new Contractor.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create a new Contractor")
		var name string
		var err error

		if cmd.Flags().Lookup("name") != nil {
			name = cmd.Flag("name").Value.String()
		}

		if name == "" {
			name, err = utils.Prompt(utils.PromptConfig{
				Label:    "Name",
				Validate: contractor.NameValidator,
			})
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		userId := auth.GetUserId()

		_, err = contractor.CreateContractor(name, userId)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	},
}

func init() {
	ContractorCreateCmd.Flags().StringP("name", "n", "", "Contractor Name")
}
