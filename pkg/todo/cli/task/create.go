package task

import (
	"errors"
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/todo/cli/prompt"
	"github.com/frangdelsolar/todo_cli/pkg/todo/cli/task/goal"
	"github.com/frangdelsolar/todo_cli/pkg/todo/data"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var CreateTaskCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a task",
	Run: func(cmd *cobra.Command, args []string) {
		var title string
		pc := prompt.PromptContent{
			Label: "Title",
			Validate: func(input string) error {
				if input == "" {
					return errors.New("title cannot be empty")
				}
				return nil
			},
		}
		title = prompt.PromptGetInput(pc)

		task, err := data.CreateTask(title)
		if err != nil {
			log.Err(err).Msg("Error creating task")
			return
		}
		log.Info().Interface("task", task).Msg("Task created")

		goal.CreateTaskGoalCmd.Run(cmd, []string{fmt.Sprint(task.ID)})
	},
}
