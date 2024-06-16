package models

import (
	"fmt"
	"regexp"

	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email"`
}


func NewUser(name string, email string) (*User, error) {

	if err := NameValidator(name); err != nil {
		return nil, err
	}
	if err := EmailValidator(email); err != nil {
		return nil, err
	}

	return &User{Name: name, Email: email}, nil
}


func NameValidator(name string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}

func EmailValidator(email string) error {
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}

    emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`

    match, err := regexp.MatchString(emailRegex, email)
    if err != nil {
        return fmt.Errorf("error compiling email regex: %w", err)
    }

    if !match {
        return fmt.Errorf("invalid email format")
    }

	return nil
}