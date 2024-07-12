package cli_test

import (
	"testing"

	cmd "github.com/frangdelsolar/todo_cli/cmd/cli"
	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	name := "pepe"
	email := "pepe@admin76.com"
	password := "test123"

	cmd.RootCmd.SetArgs([]string{"register", "-n", name, "-e", email, "-p", password})

	err := cmd.RootCmd.Execute()
	if err != nil {
		t.Errorf("Failed to register user: %v", err)
	}

	u, err := auth.GetUserByEmail(email)
	if err != nil {
		t.Errorf("Failed to get user: %v", err)
	}

	assert.Equal(t, u.Name, name, "Expected name to be %s, but got %s", name, u.Name)
	assert.Equal(t, u.Email, email, "Expected email to be %s, but got %s", email, u.Email)
	assert.NotNil(t, u.FirebaseId, "Expected FirebaseId not to be nil")
}

func TestLoginUser(t *testing.T) {
	email := "pepe@admin76.com"
	password := "test123"

	cmd.RootCmd.SetArgs([]string{"login", "-e", email, "-p", password})

	err := cmd.RootCmd.Execute()
	if err != nil {
		t.Errorf("Failed to login user: %v", err)
	}
}
