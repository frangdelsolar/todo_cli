package contractor_cli

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/cmd/cli/utils"
	"github.com/spf13/cobra"
)

var contractorActions = []utils.SelectableItem{
	{Key: "create", Label: "Create"},
	{Key: "list", Label: "List"},
	{Key: "exit", Label: "Exit"},
}

var ContractorCmd = &cobra.Command{
	Use:   "contractor",
	Short: "Contractor actions",
	Long:  `Contractor actions.`,
	Run: func(cmd *cobra.Command, args []string) {
		selection, err := utils.SelectPrompt("What would you like to do?", contractorActions)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		switch selection.Key {
		case "list":
			ContractorListCmd.Run(cmd, args)
		case "create":
			ContractorCreateCmd.Run(cmd, args)
		case "exit":
			return
		}
	},
}

func init() {
	ContractorCmd.AddCommand(ContractorListCmd)
	ContractorCmd.AddCommand(ContractorCreateCmd)
}
