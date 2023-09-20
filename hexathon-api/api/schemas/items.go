package schemas

import "github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"

// Item Serializer for displaying item data
func ItemSerializer(item models.Item) map[string]interface{} {
	return map[string]interface{}{
		"id":            item.ID,
		"name":          item.Name,
		"photo_url":     item.PhotoURL,
		"description":   item.Description,
		"price":         item.Price,
		"category_id":   item.CategoryID,
		"category_name": item.Category.Name,
	}
}

// ItemListSerializer for displaying list of items
func ItemListSerializer(items []models.Item) []map[string]interface{} {
	var result []map[string]interface{}

	for _, item := range items {
		result = append(result, ItemSerializer(item))
	}

	return result
}
