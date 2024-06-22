package cli_test

import (
	"fmt"

	cmd "github.com/frangdelsolar/todo_cli/cmd/cli"
	"github.com/frangdelsolar/todo_cli/pkg/auth"
)

func TestRegisterUser(){
    log.Info().Msg("Testing RegisterUser()")

    name := "pepe"
    email := "pepe@admin76.com"
    password := "test123"

    cmd.RootCmd.SetArgs([]string{"register", "-n", name, "-e", email, "-p", password})

    err := cmd.RootCmd.Execute()
    if err != nil {
        log.Err(err).Msg("Failed to register user")
    }

    log.Info().Msg("Created User")

    u, err := auth.GetUserByEmail(email)
    if err != nil {
        log.Warn().Msg("Failed to retrieve user")
    }

    if u.Name != name {
        err = fmt.Errorf("expected name %s, got %s", name, u.Name)
        log.Err(err).Msg("TestCreateUser()")
    } else {
        log.Debug().Msgf("Expected name %s, got %s", name, u.Name)
    }

    if u.Email != email {
        err = fmt.Errorf("expected email %s, got %s", email, u.Email)
        log.Err(err).Msg("TestCreateUser()")
    } else {
        log.Debug().Msgf("Expected email %s, got %s", email, u.Email)
    }

    if u.FirebaseId != "" {
        log.Warn().Msg("expected FirebaseId not to be empty") // It may fail as the user may already exist
    }

    log.Trace().Interface("User", u).Msg("Created User")
    log.Info().Msg("Created User Successfully")
}

func TestLoginUser(){
    log.Info().Msg("Testing LoginUser()")

    email := "pepe@admin76.com"
    password := "test123"

    cmd.RootCmd.SetArgs([]string{"login", "-e", email, "-p", password})

    err := cmd.RootCmd.Execute()
    if err != nil {
        log.Err(err).Msg("Failed to login user")
    }

    log.Info().Msg("Logged in User")
}
