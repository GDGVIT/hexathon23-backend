package v1

import (
	"fmt"

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
	group.Get("/", getParticipants)                   // <server-url>/api/v1/participants/
	group.Get("/checkedin", getCheckedInParticipants) // <server-url>/api/v1/participants/checkedin")
	group.Get("/:id", getParticipant)                 // <server-url>/api/v1/participants/:id
	group.Post("/", createParticipant)                // <server-url>/api/v1/participants/
	group.Put("/:id", updateParticipant)              // <server-url>/api/v1/participants/:id
	group.Delete("/:id", deleteParticipant)           // <server-url>/api/v1/participants/:id
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting participants: %s", err.Error())})
	}
	return c.Status(fiber.StatusOK).JSON(schemas.ParticipantListSerializer(participants))
}

// Get all checked in participants
func getCheckedInParticipants(c *fiber.Ctx) error {
	var participants []models.Participant
	var err error
	if c.Query("q") != "" {
		participants, err = models.SearchParticipantCheckedIn(c.Query("q"))
	} else {
		participants, err = models.GetParticipantsCheckedIn()
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting participants: %s", err.Error())})
	}
	return c.Status(fiber.StatusOK).JSON(schemas.ParticipantListSerializer(participants))
}

// Get a participant by id
func getParticipant(c *fiber.Ctx) error {
	participant, err := models.GetParticipantByID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting participant: %s", err.Error())})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.ParticipantSerializer(*participant))
}

// Create a new participant
func createParticipant(c *fiber.Ctx) error {
	var requestBody struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		RegNo     string `json:"reg_no"`
		CheckedIn bool   `json:"checked_in"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.InvalidBody)
	}

	// Check if participant exists
	if models.CheckParticipantExists(requestBody.RegNo) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Participant already exists",
		})
	}

	participant := models.Participant{
		Name:      requestBody.Name,
		Email:     requestBody.Email,
		RegNo:     requestBody.RegNo,
		CheckedIn: requestBody.CheckedIn,
	}

	if err := participant.CreateParticipant(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error creating participant: %s", err.Error())})
	}

	return c.Status(fiber.StatusCreated).JSON(schemas.ParticipantSerializer(participant))
}

// Update an existing participant
func updateParticipant(c *fiber.Ctx) error {
	var requestBody struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		RegNo     string `json:"reg_no"`
		CheckedIn bool   `json:"checked_in"`
	}

	participant, err := models.GetParticipantByID(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting participant: %s", err.Error())})
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.InvalidBody)
	}

	if requestBody.Name != "" {
		participant.Name = requestBody.Name
	}
	if requestBody.Email != "" {
		participant.Email = requestBody.Email
	}
	if requestBody.RegNo != "" {
		participant.RegNo = requestBody.RegNo
	}
	if requestBody.CheckedIn {
		participant.CheckedIn = requestBody.CheckedIn
	}

	if err := participant.UpdateParticipant(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error updating participant: %s", err.Error())})
	}

	return c.Status(fiber.StatusAccepted).JSON(schemas.ParticipantSerializer(*participant))
}

// Delete a participant by id
func deleteParticipant(c *fiber.Ctx) error {
	participant, err := models.GetParticipantByID(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting participant: %s", err.Error())})
	}

	if err := participant.DeleteParticipant(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error deleting participant: %s", err.Error())})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
