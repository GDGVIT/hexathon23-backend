package v1

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func teamsHandler(r fiber.Router) {
	group := r.Group("/teams")
	group.Use(middleware.JWTAuthMiddleware)

	// Routes
	group.Post("/", createTeam)
	group.Get("/", getTeams)
	group.Get("/:id", getTeam)
	group.Put("/:id", updateTeam)
	group.Delete("/:id", deleteTeam)
}

// Create a new team(Different from register as done from admin side)
func createTeam(c *fiber.Ctx) error {
	return nil
}

// Get a list of all teams
func getTeams(c *fiber.Ctx) error {
	return nil
}

// Get a team by id
func getTeam(c *fiber.Ctx) error {
	return nil
}

// Update a team by id
func updateTeam(c *fiber.Ctx) error {
	return nil
}

// Delete a team by id
func deleteTeam(c *fiber.Ctx) error {
	return nil
}
