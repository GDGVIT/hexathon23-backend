package v1

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/middleware"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/schemas"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

// ParticipantsHandler handles all the routes related to participants
func participantsHandler(r fiber.Router) {
	group := r.Group("/participants")

	// Routes
	group.Use(middleware.JWTAuthMiddleware)
	group.Use(middleware.IsAdminMiddleware)
	group.Get("/", getParticipants) // <server-url>/api/v1/participants/
}

// Get all participants
func getParticipants(c *fiber.Ctx) error {
	var participants []models.Participant
	var err error
	if c.Query("q") != "" {
		participants, err = models.SearchParticipantNotCheckedIn(c.Query("q"))
	} else {
		participants, err = models.GetParticipantsNotCheckedIn()
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(schemas.ParticipantListSerializer(participants))
}
