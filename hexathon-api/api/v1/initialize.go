package v1

import "github.com/gofiber/fiber/v2"

// V1handler handles all the routes related to api version v1
func V1handler(r fiber.Router) {
	group := r.Group("/v1")

	// Register all the handlers
	categoriesHandler(group)
	itemsHandler(group)
	teamsHandler(group)
	problemStatementHandler(group)
	authHandler(group)
	submissionsHandler(group)
	cartHandler(group)
	participantsHandler(group)
}
