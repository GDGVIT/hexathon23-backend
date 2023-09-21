package schemas

import "github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"

// CartSerializer is for displaying cart data
func CartSerializer(cart models.Cart) map[string]interface{} {
	return map[string]interface{}{
		"id":          cart.ID,
		"items_added": CartItemsSerializer(cart.Items),
		"cost":        cart.Cost,
		"amount_left": cart.Team.Amount - cart.Cost,
		"checked_out": cart.CheckedOut,
	}
}

// CartItemSerializer is for displaying cart item data
func CartItemSerializer(item models.Item) map[string]interface{} {
	return map[string]interface{}{
		"id":          item.ID,
		"name":        item.Name,
		"price":       item.Price,
		"category_id": item.CategoryID,
	}
}

// CartItemsSerializer is for displaying cart items data
func CartItemsSerializer(items []models.Item) map[string][]map[string]interface{} {
	// Cart items sorted by category
	cartItems := make(map[string][]map[string]interface{})
	for _, item := range items {
		cartItems[item.CategoryID.String()] = append(cartItems[item.CategoryID.String()], CartItemSerializer(item))
	}
	return cartItems
}
