package models

import (
	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"gorm.io/gorm"
)

type SystemData struct {
	gorm.Model
	CreatedByID uint
	CreatedBy *auth.User
	UpdatedByID uint
	UpdatedBy *auth.User
}
