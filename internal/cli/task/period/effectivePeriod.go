package period

import (
	"fmt"
	"todo_cli/pkg/prompt"

	"github.com/spf13/cobra"
)

var EffectivePeriodCmdActions = []prompt.SelectableItem{
	{Key: "add", Label: "Add"},
}

var EffectivePeriodCmd = &cobra.Command{
	Use:   "period",
	Short: "Manage Efective Periods",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Select an action")
		pc := prompt.PromptContent{
			Label: "Action",
			Items: EffectivePeriodCmdActions,
		}
		action := prompt.GetSelectInput(pc)
		switch action.Key {
		case "add":
			CreateEffectivePeriodCmd.Run(cmd, args)
		}
	},
}

func init() {
	EffectivePeriodCmd.AddCommand(CreateEffectivePeriodCmd)
}
