package v1

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/middleware"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/schemas"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

func submissionsHandler(r fiber.Router) {
	group := r.Group("/submissions")

	// Routes
	group.Use(middleware.JWTAuthMiddleware)
	group.Post("/submit", submitLinks) // <server-url>/api/v1/submissions/submit
}

func submitLinks(c *fiber.Ctx) error {
	team, err := models.GetTeamByName(c.Locals("team").(models.Team).Name)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}

	if team == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"detail": "Team not found",
		})
	}

	if team.Submitted {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Team has already submitted",
		})
	}

	var requestBody struct {
		FigmaURL string `json:"figmaURL"`
		DocURL   string `json:"docURL"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.InvalidBody)
	}

	if requestBody.FigmaURL == "" || requestBody.DocURL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Both figmaURL and docURL are required",
		})
	}

	submission := &models.Submission{
		FigmaURL:         requestBody.FigmaURL,
		DocURL:           requestBody.DocURL,
		Team:             *team,
		ProblemStatement: team.ProblemStatement,
	}

	if err := submission.CreateSubmission(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}

	team.Submitted = true
	err = team.UpdateTeam()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"detail": "Submission successful",
	})
}
