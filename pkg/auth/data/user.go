package data

import (
	m "github.com/frangdelsolar/todo_cli/pkg/auth/models"
)

func CreateUser(name string, email string) (*m.User, error) {

	u, err := m.NewUser(name, email)
	if err != nil {
		return u, err
	}

	db.Create(&u)

	return u, nil
}