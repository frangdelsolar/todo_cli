package auth_cli

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/cmd/cli/utils"
	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/spf13/cobra"
)

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your account",
	Long:  `Login to your account. If you don't have an account, you can register one with the "register" command.`,
	Run: func(cmd *cobra.Command, args []string) {

		var email string
		var password string

		if cmd.Flags().Lookup("email") != nil {
			email = cmd.Flag("email").Value.String()
		}

		if cmd.Flags().Lookup("password") != nil {
			password = cmd.Flag("password").Value.String()
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

		user, err := auth.GetUserByEmail(email)
		if err != nil {
			fmt.Println("Error logging in:", err)
			return
		}
		cfg.SetSession(userKey, fmt.Sprint(user.ID))
	},
}

func init() {
	LoginCmd.Flags().StringP("email", "e", "", "Email address")
	LoginCmd.Flags().StringP("password", "p", "", "Password")
}
