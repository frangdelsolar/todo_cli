package auth_cli

import (
	"fmt"

	"github.com/spf13/cobra"
)


var RegisterCmd = &cobra.Command{
		Use:   "register",
		Short: "Register a new account",
        Long: `Register a new account. If you already have an account, you can login with the "login" command.`,
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Register")
        },
}

func init() {
}
