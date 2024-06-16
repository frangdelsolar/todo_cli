package user

import (
	"github.com/frangdelsolar/todo_cli/cli/prompt"
	d "github.com/frangdelsolar/todo_cli/pkg/auth/data"
	m "github.com/frangdelsolar/todo_cli/pkg/auth/models"
	"github.com/spf13/cobra"
)


var CreateUserCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a user",
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		pc := prompt.PromptContent{
			Label: "Name",
			Validate:m.NameValidator,
		}
		name = prompt.PromptGetInput(pc)

		var email string
		pc = prompt.PromptContent{
			Label: "Email",
			Validate: m.EmailValidator,
		}
		email = prompt.PromptGetInput(pc)

		user, err := d.CreateUser(name, email)
		if err != nil {
			log.Err(err).Msg("Error creating user")
			return
		}
		log.Info().Interface("user", user).Msg("User created")

	},
}
