package models

import (
	"fmt"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
	"gorm.io/gorm"
)

type SystemData struct {
	gorm.Model
	CreatedByID uint `gorm:"not null" json:"createdById"`
	CreatedBy *auth.User `gorm:"foreignKey:CreatedByID" json:"createdBy"`
	UpdatedByID uint  `gorm:"not null" json:"updatedById"`
	UpdatedBy *auth.User `gorm:"foreignKey:UpdatedByID" json:"updatedBy"`
}

// ID returns the ID of the SystemData as a string.
//
// Returns:
// - string: the ID of the SystemData.
func (s *SystemData) GetIDString() string {
    return fmt.Sprint(s.ID)
}
