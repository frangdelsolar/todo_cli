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
        Long: `Register a new account. If you already have an account, you can login with the "login" command.`,
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Register")

            email, _ := utils.Prompt(utils.PromptConfig{
                Label: "Email",
                Validate: auth.EmailValidator,
            })

            name, _ := utils.Prompt(utils.PromptConfig{
                Label: "Name",
                Validate: auth.NameValidator,
            })

            user, err := auth.CreateUser(name, email)
            if err != nil {
                fmt.Println("Error creating user:", err)
            } else {
                fmt.Println("User created:", user.ID)
            }
        },
}

func init() {
}
