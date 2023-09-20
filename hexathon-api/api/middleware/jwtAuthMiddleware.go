package middleware

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/auth"
	"github.com/gofiber/fiber/v2"
)

// JWTAuthMiddleware is a go-fiber middleware to authenticate the user and add the user to the context if authenticated
func JWTAuthMiddleware(c *fiber.Ctx) error {
	// Get the JWT token from the Authorization header
	authorizationString := c.Get("Authorization")

	// If the Authorization header is not present, return an error
	if authorizationString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": "Authorization header not present",
		})
	}

	// Get team from JWT token
	team, parseErr := auth.GetTeamFromJWTToken(authorizationString, auth.JWTSecret)
	if parseErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": parseErr.Error(),
		})
	}

	// Add the team to the context
	c.Locals("team", team)
	c.Locals("role", team.Role)
	return c.Next()
}
