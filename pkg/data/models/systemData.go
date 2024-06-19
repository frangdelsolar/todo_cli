package models

import (
	"gorm.io/gorm"
)

type SystemData struct {
	gorm.Model
	CreatedByID uint
	UpdatedByID uint
}
