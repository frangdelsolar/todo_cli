package auth_cli

import (
	"fmt"

	"github.com/spf13/cobra"
)


var LoginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login to your account",
		Long: `Login to your account. If you don't have an account, you can register one with the "register" command.`,
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Login")
        },
}

func init() {
}
