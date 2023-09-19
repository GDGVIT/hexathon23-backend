package models

// Team is the db model for teams table
type Team struct {
	ID    string `gorm:"default:uuid_generate_v4();primaryKey"`
	Name  string `gorm:"unique;not null"`
	Logo  string
	Role  string `gorm:"default:participant"`
	Items []Item
}
