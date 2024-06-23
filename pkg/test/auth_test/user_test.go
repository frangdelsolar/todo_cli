package auth_test

import (
	"testing"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T){
    name := "pepe"
    email := "pepe@admin.com"
    password := "test123"

    u, err := auth.CreateUser(name, email, password)
    if err != nil {
        log.Warn().Msg("Failed to create user")
    }

    assert.Equal(t, u.Name, name, "Expected name to be %s, but got %s", name, u.Name)
    assert.Equal(t, u.Email, email, "Expected email to be %s, but got %s", email, u.Email)

}
