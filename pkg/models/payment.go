package models

import "time"

type Payment struct {
	ID           string    `json:"id"`
	Value        float64   `json:"value"`
	EmissionDate time.Time `json:"emissionDate"`
	AssignorID   string    `json:"assignor_id"`
}
