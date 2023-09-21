package models

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/database"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

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

// CreateItem creates a new item
func (item *Item) CreateItem() error {
	return database.DB.Create(item).Error
}

// UpdateItem updates an item
func (item *Item) UpdateItem() error {
	return database.DB.Save(item).Error
}

// DeleteItem deletes an item
func (item *Item) DeleteItem() error {
	return database.DB.Delete(item).Error
}

// ValidateItemName validates an item name
func ValidateItemName(name string) bool {
	return len(name) >= 3
}

// GetItemByID returns an item by id
func GetItemByID(id string) (*Item, error) {
	var item Item
	err := database.DB.Preload(clause.Associations).Where("id = ?", id).First(&item).Error
	return &item, err
}

// GetItems returns a list of all items
func GetItems() ([]Item, error) {
	var items []Item
	err := database.DB.Preload(clause.Associations).Find(&items).Error
	return items, err
}

// GetItemsByCategoryID returns a list of all items by category id
func GetItemsByCategoryID(id string) ([]Item, error) {
	var items []Item
	err := database.DB.Preload(clause.Associations).Where("category_id = ?", id).Find(&items).Error
	return items, err
}
