package v1

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func teamsHandler(r fiber.Router) {
	group := r.Group("/teams")
	group.Use(middleware.JWTAuthMiddleware)
	group.Use(middleware.IsAdminMiddleware)

	// Routes
	group.Post("/", createTeam)      // <server-url>/api/v1/teams/
	group.Get("/", getTeams)         // <server-url>/api/v1/teams/
	group.Get("/:id", getTeam)       // <server-url>/api/v1/teams/:id
	group.Put("/:id", updateTeam)    // <server-url>/api/v1/teams/:id
	group.Delete("/:id", deleteTeam) // <server-url>/api/v1/teams/:id
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
