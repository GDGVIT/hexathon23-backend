package v1

import "github.com/gofiber/fiber/v2"

// AuthHandler handles all the routes related to authentication
func authHandler(r fiber.Router) {
	group := r.Group("/auth")
	group.Post("/register", register)
	group.Post("/login", login)
}

// Create a team with the given name and password
func register(c *fiber.Ctx) error {
	return nil
}

// Login a team with the using name and password
func login(c *fiber.Ctx) error {
	return nil
}
