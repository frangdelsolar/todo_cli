package auth_test

import (
	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/frangdelsolar/todo_cli/pkg/test"
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
func TestCreateUser(t *test.TestRunner){
    t.Run("TestCreateUser()")

    name := "pepe"
    email := "pepe@admin.com"
    password := "test123"

    u, err := auth.CreateUser(name, email, password)
    if err != nil {
        log.Warn().Msg("Failed to create user")
    }

    t.AssertEqual(u.Name, name)
    t.AssertEqual(u.Email, email)

    t.Stop()
}
