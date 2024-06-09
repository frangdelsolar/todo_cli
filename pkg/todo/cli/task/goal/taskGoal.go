package goal

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/todo/cli/prompt"
	"github.com/spf13/cobra"
)

var TaskGoalCmdActions = []prompt.SelectableItem{
	{Key: "add", Label: "Add"},
}

var TaskGoalCmd = &cobra.Command{
	Use:   "task_goal",
	Short: "Manage Efective Periods",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Select an action")
		pc := prompt.PromptContent{
			Label: "Action",
			Items: TaskGoalCmdActions,
		}
		action := prompt.GetSelectInput(pc)
		switch action.Key {
		case "add":
			CreateTaskGoalCmd.Run(cmd, args)
		}
	},
}

func init() {
	TaskGoalCmd.AddCommand(CreateTaskGoalCmd)
}
