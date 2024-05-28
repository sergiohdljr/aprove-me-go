package models

import (
	"github.com/google/uuid"
)

type Assignor struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Document string    `gorm:"type:string;unique;not null"`
	Email    string    `gorm:"type:string;unique;not null"`
	Phone    string    `gorm:"type:string;unique;not null"`
	Name     string    `gorm:"type:string;not null"`
	Payments []Payment `gorm:"foreignKey:AssignorID"`
}
