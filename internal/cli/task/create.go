package task

import (
	"errors"
	"todo_cli/data"
	"todo_cli/pkg/prompt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var CreateTaskCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a task",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("Creating task")

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
		log.Debug().Interface("cmd", cmd.Parent()).Msg("cmd")
	},
}
