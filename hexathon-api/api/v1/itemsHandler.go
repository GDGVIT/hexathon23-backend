package v1

import (
	"fmt"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/middleware"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/schemas"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ItemsHandler handles all the routes related to items
func itemsHandler(r fiber.Router) {
	group := r.Group("/items")

	// Routes
	group.Use(middleware.JWTAuthMiddleware)
	group.Get("/", getItems)   // <server-url>/api/v1/items/
	group.Get("/:id", getItem) // <server-url>/api/v1/items/:id

	group.Use(middleware.IsAdminMiddleware)
	group.Post("/", createItem)      // <server-url>/api/v1/items/
	group.Put("/:id", updateItem)    // <server-url>/api/v1/items/:id
	group.Delete("/:id", deleteItem) // <server-url>/api/v1/items/:id
}

// Create a new item
func createItem(c *fiber.Ctx) error {
	var requestBody struct {
		Name        string `json:"name"`
		PhotoURL    string `json:"photo_url"`
		Description string `json:"description"`
		Price       int    `json:"price"`
		CategoryID  string `json:"category_id"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.InvalidBody)
	}

	// Validate item name
	if !models.ValidateItemName(requestBody.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid item name",
		})
	}

	// Check if category exists
	if !models.CheckCategoryExists(requestBody.CategoryID) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Category does not exist",
		})
	}

	// String to uuid
	categoryID, err := uuid.Parse(requestBody.CategoryID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": fmt.Sprintf("Invalid category id: %s", err.Error()),
		})
	}

	item := &models.Item{
		Name:        requestBody.Name,
		PhotoURL:    requestBody.PhotoURL,
		Description: requestBody.Description,
		Price:       requestBody.Price,
		CategoryID:  categoryID,
	}

	if err := item.CreateItem(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error creating item: %s", err.Error()),
		})
	}

	item, err = models.GetItemByID(item.ID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting item: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(schemas.ItemSerializer(*item))
}

// Get a list of all items
func getItems(c *fiber.Ctx) error {
	var items []models.Item
	var err error
	// Check if query params exist
	if c.Query("category") != "" {
		items, err = models.GetItemsByCategoryID(c.Query("category"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"detail": fmt.Sprintf("Error getting items: %s", err.Error()),
			})
		}
	} else {
		items, err = models.GetItems()
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting items: %s", err.Error()),
		})
	}

	team := c.Locals("team").(models.Team)
	cart, err := team.GetCart()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting cart: %s", err.Error())})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.ItemCartListSerializer(items, *cart))
}

// Get an item by id
func getItem(c *fiber.Ctx) error {
	item, err := models.GetItemByID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting item: %s", err.Error())})
	}

	if item == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"detail": "Item not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.ItemSerializer(*item))
}

// Update an item by id
func updateItem(c *fiber.Ctx) error {
	var requestBody struct {
		Name        string `json:"name"`
		PhotoURL    string `json:"photo_url"`
		Description string `json:"description"`
		Price       int    `json:"price"`
		CategoryID  string `json:"category_id"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.InvalidBody)
	}

	// Validate item name
	if !models.ValidateItemName(requestBody.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid item name",
		})
	}

	// Check if category exists
	if !models.CheckCategoryExists(requestBody.CategoryID) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Category does not exist",
		})
	}

	item, err := models.GetItemByID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting item: %s", err.Error())})
	}

	if item == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"detail": "Item not found",
		})
	}

	// String to uuid
	categoryID, err := uuid.Parse(requestBody.CategoryID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid category id",
		})
	}

	item.Name = requestBody.Name
	item.PhotoURL = requestBody.PhotoURL
	item.Description = requestBody.Description
	item.Price = requestBody.Price
	item.CategoryID = categoryID

	if err := item.UpdateItem(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error updating item: %s", err.Error())})
	}

	item, err = models.GetItemByID(item.ID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting item: %s", err.Error())})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.ItemSerializer(*item))
}

// Delete an item by id
func deleteItem(c *fiber.Ctx) error {
	item, err := models.GetItemByID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting item: %s", err.Error())})
	}

	if item == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"detail": "Item not found",
		})
	}

	if err := item.DeleteItem(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error deleting item: %s", err.Error())})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
