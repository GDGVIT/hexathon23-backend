package models

import "github.com/google/uuid"

type Transaction struct {
	ID     uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Team   Team      `gorm:"foreignKey:TeamID"`
	TeamID uuid.UUID
	Items  []Item `gorm:"many2many:transaction_items;"`
}
