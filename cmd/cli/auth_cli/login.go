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
            email, _ := utils.Prompt(utils.PromptConfig{
                Label: "Email",
                Validate: auth.EmailValidator,
            }) 

            user, err := auth.GetUserByEmail(email)
            if err != nil {
                fmt.Println("Error logging in:", err)
            } 
            
            fmt.Println("User logged in:", user.ID)

            cfg.SetSession(userKey, fmt.Sprint(user.ID))
        },
}

func init() {
}
