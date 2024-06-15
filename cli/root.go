package cli

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/cli/prompt"
	taskcmd "github.com/frangdelsolar/todo_cli/cli/task"
	"github.com/spf13/cobra"
)

var APP_VERSION string
var rootCmdActions = []prompt.SelectableItem{
	{Key: "task", Label: "Task actions"},
}

var rootCmd = &cobra.Command{
	PreRun: func(cmd *cobra.Command, args []string) {
		// clear screen
		fmt.Print("\033[2J\033[1;1H")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("TODO: A Powerful Command-Line Task Management Tool v%s\n\n", APP_VERSION)
		fmt.Print(`Take control of your to-do list with TODO, 
a user-friendly and efficient command-line application.  
With TODO, you can ditch the scattered sticky notes and cluttered reminders, 
and organize your tasks in a centralized and accessible way.
		`)

		pc := prompt.PromptContent{
			Label: "What do you want to do?",
			Items: rootCmdActions,
		}
		selection := prompt.GetSelectInput(pc)

		switch selection.Key {
		case "task":
			taskcmd.TaskCmd.Run(cmd, args)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		// wait for user to press any key and then execute again
		fmt.Print("\nPress any key to continue...")
		fmt.Scanln()
		cmd.Execute()
	},
}


func init() {
	rootCmd.AddCommand(taskcmd.TaskCmd)
}

type CLI struct {
	*cobra.Command
}

// Execute executes the root command.
func NewCLI(version string) *CLI {

	APP_VERSION = version
	// log =models.Logger{lg}

	return &CLI{rootCmd}
}

