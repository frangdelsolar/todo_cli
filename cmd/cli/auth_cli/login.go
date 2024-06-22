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
		Long: `Login to your account. If you don't have an account, you can register one with the "register" command.`,
        Run: func(cmd *cobra.Command, args []string) {

            email := cmd.Flag("email").Value.String()
            password := cmd.Flag("password").Value.String()

            if email == ""{
                email, _ = utils.Prompt(utils.PromptConfig{
                    Label: "Email",
                    Validate: auth.EmailValidator,
                }) 
            }

            if password == ""{
                password, _ = utils.Prompt(utils.PromptConfig{
                    Label: "Password",
                    Validate: auth.NameValidator,
                    Password: true,
                })
            }

            user, err := auth.GetUserByEmail(email)
            if err != nil {
                fmt.Println("Error logging in:", err)
            } 
            
            fmt.Println("User logged in:", user.ID)

            cfg.SetSession(userKey, fmt.Sprint(user.ID))
        },

}

func init() {
    // add a flag for email
    // add a flag for password
    LoginCmd.Flags().StringP("email", "e", "", "Email address")
    LoginCmd.Flags().StringP("password", "p", "", "Password")
}
