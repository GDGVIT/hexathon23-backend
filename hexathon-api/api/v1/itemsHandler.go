package v1

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func itemsHandler(r fiber.Router) {
	group := r.Group("/items")
	group.Use(middleware.JWTAuthMiddleware)
	group.Use(middleware.IsAdminMiddleware)

	// Routes
	group.Post("/", createItem)      // <server-url>/api/v1/items/
	group.Get("/", getItems)         // <server-url>/api/v1/items/
	group.Get("/:id", getItem)       // <server-url>/api/v1/items/:id
	group.Put("/:id", updateItem)    // <server-url>/api/v1/items/:id
	group.Delete("/:id", deleteItem) // <server-url>/api/v1/items/:id
}

// Create a new item
func createItem(c *fiber.Ctx) error {
	return nil
}

// Get a list of all items
func getItems(c *fiber.Ctx) error {
	return nil
}

// Get an item by id
func getItem(c *fiber.Ctx) error {
	return nil
}

// Update an item by id
func updateItem(c *fiber.Ctx) error {
	return nil
}

// Delete an item by id
func deleteItem(c *fiber.Ctx) error {
	return nil
}
