package accounts

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/cli/prompt"
	"github.com/frangdelsolar/todo_cli/pkg/logger"
	"github.com/spf13/cobra"
)

var log = logger.GetLogger()

var accountCmdActions = []prompt.SelectableItem{
	{Key: "create", Label: "Add an Account"},
	{Key: "exit", Label: "Exit"},
}

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Account actions",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Print("Select an action")
		pc := prompt.PromptContent{
			Label: "Action",
			Items: accountCmdActions,
		}
		action := prompt.GetSelectInput(pc)
		switch action.Key {
		case "create":
			CreateAccountCmd.Run(cmd, args)
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
	AccountCmd.AddCommand(CreateAccountCmd)
	// AccountCmd.AddCommand(ListTaskCmd)
	// AccountCmd.AddCommand(UpdateTaskCmd)
	// AccountCmd.AddCommand(DeleteTaskCmd)
}
