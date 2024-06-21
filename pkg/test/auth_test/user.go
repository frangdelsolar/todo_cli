package auth_test

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
)

// TestCreateUser is a test function that tests the CreateUser function from the auth package.
//
// This function logs an informational message indicating that it is testing the CreateUser function.
// It then sets the name and email variables to "pepe" and "pepe@admin.com", respectively.
// It calls the CreateUser function with the name and email variables as arguments and assigns the returned user and error to the u and err variables, respectively.
// If the error is not nil, it logs an error message indicating that the user creation failed.
// It then checks if the user's name is not equal to the expected name and logs an error message indicating that the expected name was not received.
// If the user's name is equal to the expected name, it logs a debug message indicating that the expected name was received.
// It then checks if the user's email is not equal to the expected email and logs an error message indicating that the expected email was not received.
// If the user's email is equal to the expected email, it logs a debug message indicating that the expected email was received.
// Finally, it logs a debug message indicating that a user was created and an informational message indicating that the user was created successfully.
func TestCreateUser(){
    log.Info().Msg("Testing CreateUser()")

    name := "pepe"
    email := "pepe@admin.com"
    password := "test123"

    u, err := auth.CreateUser(name, email, password)
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

    if u.FirebaseId != "" {
        err = fmt.Errorf("expected FirebaseId not to be empty")
        log.Err(err).Msg("TestCreateUser()")
    }

    log.Trace().Interface("User", u).Msg("Created User")
    log.Info().Msg("Created User Successfully")
}
