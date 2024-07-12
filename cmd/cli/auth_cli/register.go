package auth_cli

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/cmd/cli/utils"
	auth "github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/spf13/cobra"
)

var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new account",
	Long:  `Register a new account. If you already have an account, you can login with the "login" command.`,
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		var email string
		var password string
		var err error

		if cmd.Flags().Lookup("name") != nil {
			name = cmd.Flag("name").Value.String()
		}

		if cmd.Flags().Lookup("email") != nil {
			email = cmd.Flag("email").Value.String()
		}

		if cmd.Flags().Lookup("password") != nil {
			password = cmd.Flag("password").Value.String()
		}

		if name == "" {
			name, _ = utils.Prompt(utils.PromptConfig{
				Label:    "Name",
				Validate: auth.NameValidator,
			})
		}
		if email == "" {
			email, _ = utils.Prompt(utils.PromptConfig{
				Label:    "Email",
				Validate: auth.EmailValidator,
			})
		}
		if password == "" {
			password, _ = utils.Prompt(utils.PromptConfig{
				Label:    "Password",
				Validate: auth.NameValidator,
				Password: true,
			})
		}
		user, err := auth.CreateUser(name, email, password)
		if err != nil {
			fmt.Println("Error creating user:", err)
		} else {
			fmt.Println("User created ID:", user.ID)
		}

		cfg.SetSession(userKey, fmt.Sprint(user.ID))
	},
}

func init() {
	RegisterCmd.Flags().StringP("email", "e", "", "Email address")
	RegisterCmd.Flags().StringP("name", "n", "", "Name")
	RegisterCmd.Flags().StringP("password", "p", "", "Password")
}
