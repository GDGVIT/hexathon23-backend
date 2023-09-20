package v1

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/middleware"
	"github.com/gofiber/fiber/v2"
)

// CategoriesHandler handles all the routes related to categories
func categoriesHandler(r fiber.Router) {
	group := r.Group("/categories")
	group.Use(middleware.JWTAuthMiddleware)
	group.Use(middleware.IsAdminMiddleware)

	// Routes
	group.Get("/", getCategories)        // <server-url>/api/v1/categories/
	group.Post("/", createCategory)      // <server-url>/api/v1/categories/
	group.Get("/:id", getCategory)       // <server-url>/api/v1/categories/:id
	group.Put("/:id", updateCategory)    // <server-url>/api/v1/categories/:id
	group.Delete("/:id", deleteCategory) // <server-url>/api/v1/categories/:id
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
