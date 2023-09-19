package v1

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/middleware"
	"github.com/gofiber/fiber/v2"
)

// CategoriesHandler handles all the routes related to categories
func categoriesHandler(r fiber.Router) {
	group := r.Group("/categories")
	group.Use(middleware.JWTAuthMiddleware)

	// Routes
	group.Get("/", getCategories)
	group.Post("/", createCategory)
	group.Get("/:id", getCategory)
	group.Put("/:id", updateCategory)
	group.Delete("/:id", deleteCategory)
}

// Get a list of all categories
func getCategories(c *fiber.Ctx) error {
	return nil
}

// Create a new category
func createCategory(c *fiber.Ctx) error {
	return nil
}

// Get a category by id
func getCategory(c *fiber.Ctx) error {
	return nil
}

// Update a category by id
func updateCategory(c *fiber.Ctx) error {
	return nil
}

// Delete a category by id
func deleteCategory(c *fiber.Ctx) error {
	return nil
}
