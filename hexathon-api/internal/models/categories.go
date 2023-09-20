package models

import "github.com/google/uuid"

// Category is the db model for categories table
type Category struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"unique;not null"`
	PhotoURL    string
	Description string
	Items       []Item
}
