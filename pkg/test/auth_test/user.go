package auth_test

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
)

func TestCreateUser(){
    log.Info().Msg("Testing CreateUser()")

    name := "pepe"
    email := "pepe@admin.com"

    u, err := auth.CreateUser(name, email)
    if err != nil {
        log.Err(err).Msg("Failed to create user")
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

    log.Debug().Interface("User", u).Msg("Created User")
    log.Info().Msg("Created User Successfully")
}
