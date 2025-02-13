package v1

import (
	"fmt"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/middleware"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/schemas"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

// CartHandler handles all the routes related to carts
func cartHandler(r fiber.Router) {
	group := r.Group("/carts")

	// Routes
	group.Use(middleware.JWTAuthMiddleware)
	group.Get("/", getMyCart)                // <server-url>/api/v1/carts/
	group.Post("/checkout", checkoutCart)    // <server-url>/api/v1/carts/checkout
	group.Post("/:itemId", addToCart)        // <server-url>/api/v1/carts/:itemId
	group.Delete("/:itemId", deleteFromCart) // <server-url>/api/v1/carts/:itemId
}

// Get my cart
func getMyCart(c *fiber.Ctx) error {
	team := c.Locals("team").(models.Team)
	cart, err := team.GetCart()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting cart: %s", err.Error()),
		})
	}
	return c.Status(fiber.StatusOK).JSON(schemas.CartSerializer(*cart))
}

// Add an item to my cart
func addToCart(c *fiber.Ctx) error {
	itemID := c.Params("itemId")
	item, err := models.GetItemByID(itemID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting item: %s", err.Error()),
		})
	}

	cTeam := c.Locals("team").(models.Team)
	team, err := models.GetTeamByName(cTeam.Name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting team: %s", err.Error()),
		})
	}

	cart, err := team.GetCart()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting cart: %s", err.Error()),
		})
	}

	if cart.CheckedOut {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Cart already checked out",
		})
	}

	err = cart.AddToCart(*item)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error adding item to cart: %s", err.Error()),
		})
	}
	return c.Status(fiber.StatusOK).JSON(schemas.CartSerializer(*cart))
}

// Delete an item from my cart
func deleteFromCart(c *fiber.Ctx) error {
	itemID := c.Params("itemId")
	item, err := models.GetItemByID(itemID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting item: %s", err.Error()),
		})
	}

	team := c.Locals("team").(models.Team)
	cart, err := team.GetCart()
	if cart.CheckedOut {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Cart already checked out",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting cart: %s", err.Error()),
		})
	}
	err = cart.DeleteFromCart(*item)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error deleting item from cart: %s", err.Error()),
		})
	}
	return c.Status(fiber.StatusOK).JSON(schemas.CartSerializer(*cart))
}

// Checkout my cart
func checkoutCart(c *fiber.Ctx) error {
	team := c.Locals("team").(models.Team)
	cart, err := team.GetCart()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting cart: %s", err.Error()),
		})
	}
	if cart.CheckedOut {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Cart already checked out",
		})
	}
	err = cart.CheckoutCart()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error checking out cart: %s", err.Error()),
		})
	}
	retTeam, err := models.GetTeamByName(team.Name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting team: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.TeamCheckoutSerializer(*retTeam))
}
