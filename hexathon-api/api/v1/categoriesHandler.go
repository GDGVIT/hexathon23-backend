package v1

import (
	"fmt"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/middleware"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/schemas"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

// CategoriesHandler handles all the routes related to categories
func categoriesHandler(r fiber.Router) {
	group := r.Group("/categories")

	// Routes
	group.Use(middleware.JWTAuthMiddleware)
	group.Get("/", getCategories)  // <server-url>/api/v1/categories/
	group.Get("/:id", getCategory) // <server-url>/api/v1/categories/:id

	group.Use(middleware.IsAdminMiddleware)
	group.Post("/", createCategory)      // <server-url>/api/v1/categories/
	group.Put("/:id", updateCategory)    // <server-url>/api/v1/categories/:id
	group.Delete("/:id", deleteCategory) // <server-url>/api/v1/categories/:id
}

// Get a list of all categories
func getCategories(c *fiber.Ctx) error {
	categories, err := models.GetCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting categories: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.CategoryListSerializer(categories))
}

// Create a new category
func createCategory(c *fiber.Ctx) error {
	var requestBody struct {
		Name        string `json:"name"`
		PhotoURL    string `json:"photo_url"`
		Description string `json:"description"`
		MaxItems    int    `json:"max_items"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.InvalidBody)
	}

	// Validate category name
	if !models.ValidateCategoryName(requestBody.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid category name",
		})
	}

	category := models.Category{
		Name:        requestBody.Name,
		PhotoURL:    requestBody.PhotoURL,
		Description: requestBody.Description,
		MaxItems:    requestBody.MaxItems,
	}

	if err := category.CreateCategory(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error creating category: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(schemas.CategorySerializer(category))
}

// Get a category by id
func getCategory(c *fiber.Ctx) error {
	category, err := models.GetCategoryByID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting category: %s", err.Error()),
		})
	}

	if category == nil {
		return c.Status(fiber.StatusNotFound).JSON(schemas.NotFound)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.CategorySerializer(*category))
}

// Update a category by id
func updateCategory(c *fiber.Ctx) error {
	var requestBody struct {
		Name        string `json:"name"`
		PhotoURL    string `json:"photo_url"`
		Description string `json:"description"`
		MaxItems    int    `json:"max_items"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.InvalidBody)
	}

	category, err := models.GetCategoryByID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting category: %s", err.Error()),
		})
	}

	if category == nil {
		return c.Status(fiber.StatusNotFound).JSON(schemas.NotFound)
	}

	if !models.ValidateCategoryName(requestBody.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid category name",
		})
	}

	category.Name = requestBody.Name
	category.PhotoURL = requestBody.PhotoURL
	category.Description = requestBody.Description
	category.MaxItems = requestBody.MaxItems

	if err := category.UpdateCategory(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error updating category: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(schemas.CategorySerializer(*category))
}

// Delete a category by id
func deleteCategory(c *fiber.Ctx) error {
	category, err := models.GetCategoryByID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting category: %s", err.Error())})
	}

	if category == nil {
		return c.Status(fiber.StatusNotFound).JSON(schemas.NotFound)
	}

	if err := category.DeleteCategory(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error deleting category: %s", err.Error()),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
