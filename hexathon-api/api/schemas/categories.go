package schemas

import "github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"

// Category Serializer for displaying category data
func CategorySerializer(category models.Category) map[string]interface{} {
	return map[string]interface{}{
		"id":          category.ID,
		"name":        category.Name,
		"photo_url":   category.PhotoURL,
		"description": category.Description,
		"max_items":   category.MaxItems,
	}
}

// CategoryListSerializer for displaying list of categories
func CategoryListSerializer(categories []models.Category) []map[string]interface{} {
	var result []map[string]interface{}

	for _, category := range categories {
		result = append(result, CategorySerializer(category))
	}

	return result
}
