package models

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/database"
	"github.com/google/uuid"
)

// Category is the db model for categories table
type Category struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"unique;not null"`
	PhotoURL    string
	Description string
	Items       []Item
}

// CreateCategory creates a new category
func (category *Category) CreateCategory() error {
	return database.DB.Create(category).Error
}

// UpdateCategory updates a category
func (category *Category) UpdateCategory() error {
	return database.DB.Save(category).Error
}

// DeleteCategory deletes a category
func (category *Category) DeleteCategory() error {
	return database.DB.Delete(category).Error
}

// CheckCategoryExists checks if a category exists
func CheckCategoryExists(id uuid.UUID) bool {
	var count int64
	database.DB.Model(&Category{}).Where("id = ?", id).Count(&count)
	return count > 0
}

// ValidateCategoryName validates a category name
func ValidateCategoryName(name string) bool {
	return len(name) >= 3
}

// GetCategoryByID returns a category by id
func GetCategoryByID(id string) (*Category, error) {
	var category Category
	err := database.DB.Where("id = ?", id).First(&category).Error
	return &category, err
}

// GetCategories returns a list of all categories
func GetCategories() ([]Category, error) {
	var categories []Category
	err := database.DB.Find(&categories).Error
	return categories, err
}
