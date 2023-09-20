package models

import "github.com/google/uuid"

type ProblemStatement struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"unique;not null"`
	Description string
}
