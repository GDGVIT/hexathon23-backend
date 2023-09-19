package models

// Item is the db model for items table
type Item struct {
	ID          string `gorm:"default:uuid_generate_v4();primaryKey"`
	Name        string `gorm:"unique;not null"`
	PhotoURL    string
	Description string
	Price       int
	Category    Category `gorm:"foreignKey:CategoryID"`
	CategoryID  string
}
