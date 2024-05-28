package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Value        float64   `gorm:"type:float;not null"`
	EmissionDate time.Time `gorm:"type:time;not null"`
	AssignorID   uuid.UUID `gorm:"type:uuid;not null"`
}
