package user

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/cli/prompt"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	"github.com/spf13/cobra"
)

var log = logger.GetLogger()


var userCmdActions = []prompt.SelectableItem{
	{Key: "create", Label: "Add a user"},
	{Key: "exit", Label: "Exit"},
}

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "User actions",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Print("Select an action")
		pc := prompt.PromptContent{
			Label: "Action",
			Items: userCmdActions,
		}
		action := prompt.GetSelectInput(pc)
		switch action.Key {
		case "create":
			CreateUserCmd.Run(cmd, args)
		// case "update":
		// 	UpdateTaskCmd.Run(cmd, args)
		// case "delete":
		// 	DeleteTaskCmd.Run(cmd, args)
		// case "list":
		// 	ListTaskCmd.Run(cmd, args)
		}
	},
}

func init() {
	UserCmd.AddCommand(CreateUserCmd)
	// TaskCmd.AddCommand(ListTaskCmd)
	// TaskCmd.AddCommand(UpdateTaskCmd)
	// TaskCmd.AddCommand(DeleteTaskCmd)
}
