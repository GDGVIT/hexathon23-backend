package models

import "github.com/google/uuid"

// Item is the db model for items table
type Item struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"unique;not null"`
	PhotoURL    string
	Description string
	Price       int
	Category    Category `gorm:"foreignKey:CategoryID"`
	CategoryID  uuid.UUID
}
