package cli

import (
	"fmt"

	auth "github.com/frangdelsolar/todo_cli/cmd/cli/auth_cli"
	"github.com/frangdelsolar/todo_cli/cmd/cli/utils"
	"github.com/spf13/cobra"
)

var rootActions = []utils.SelectableItem{
	{Key: "register", Label: "Register"},
	{Key: "login", Label: "Login"},
	{Key: "exit", Label: "Exit"},
}

var rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "TODO: A Powerful Command-Line Task Management Tool",
		Long: `Take control of your to-do list with TODO, 
a user-friendly and efficient command-line application.  
With TODO, you can ditch the scattered sticky notes and cluttered reminders, 
and organize your tasks in a centralized and accessible way.
`,
        Run: func(cmd *cobra.Command, args []string) {
            selection, err := utils.SelectPrompt("What would you like to do?", rootActions)
            if err != nil {
                fmt.Println("Error:", err)
                return
            }
            switch selection.Key {
            case "register":
                auth.RegisterCmd.Run(cmd, args)
            case "login":
                auth.LoginCmd.Run(cmd, args)
            case "exit":
                return
            }

        },
        PostRun: func(cmd *cobra.Command, args []string) {
            fmt.Println("Bye!")
        },
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(auth.LoginCmd)
	rootCmd.AddCommand(auth.RegisterCmd)
}
