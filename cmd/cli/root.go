package cli

import (
	"fmt"

	auth "github.com/frangdelsolar/todo_cli/cmd/cli/auth_cli"
	contractor "github.com/frangdelsolar/todo_cli/cmd/cli/contractor_cli"
	"github.com/frangdelsolar/todo_cli/cmd/cli/utils"
	"github.com/spf13/cobra"
)

var anonymousActions = []utils.SelectableItem{
	{Key: "register", Label: "Register"},
	{Key: "login", Label: "Login"},
	{Key: "exit", Label: "Exit"},
}

var rootActions = []utils.SelectableItem{
    {Key: "contractor", Label: "Contractor"},
    {Key: "exit", Label: "Exit"},
}

var RootCmd = &cobra.Command{
		Use:   "todo",
		Short: "TODO: A Powerful Command-Line Task Management Tool",
		Long: `Take control of your to-do list with TODO, 
a user-friendly and efficient command-line application.  
With TODO, you can ditch the scattered sticky notes and cluttered reminders, 
and organize your tasks in a centralized and accessible way.
`,
        Run: func(cmd *cobra.Command, args []string) {
            var actions utils.SelectableList
            userIsLoggedIn := auth.IsLoggedIn()
            if !userIsLoggedIn {
                actions = anonymousActions
            } else {
                actions = rootActions
            }

            selection, err := utils.SelectPrompt("What would you like to do?", actions)
            if err != nil {
                fmt.Println("Error:", err)
                return
            }
            switch selection.Key {
            case "register":
                auth.RegisterCmd.Run(cmd, args)
            case "login":
                auth.LoginCmd.Run(cmd, args)
            case "contractor":
                contractor.ContractorCmd.Run(cmd, args)
            case "exit":
                return
            }

        },
        PostRun: func(cmd *cobra.Command, args []string) {
            // wait for user to press any key and then execute again
            fmt.Print("\nPress any key to continue...")
            fmt.Scanln()
            cmd.Execute()
        },
}

// Execute executes the root command.
func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(auth.LoginCmd)
	RootCmd.AddCommand(auth.RegisterCmd)
	RootCmd.AddCommand(contractor.ContractorCmd)
}
