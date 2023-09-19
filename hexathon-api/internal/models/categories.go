package models

// Category is the db model for categories table
type Category struct {
	ID          string `gorm:"default:uuid_generate_v4();primaryKey"`
	Name        string `gorm:"unique;not null"`
	PhotoURL    string
	Description string
	Items       []Item `gorm:"foreignKey:CategoryID"`
}
