package v1

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/middleware"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/schemas"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

// ProblemStatementHandler handles all the routes related to problemStatements
func problemStatementHandler(r fiber.Router) {
	group := r.Group("/problemStatements")

	// Routes
	group.Use(middleware.JWTAuthMiddleware)
	group.Get("/team", getProblemStatementForTeam)       // <server-url>/api/v1/problemStatements/team
	group.Post("/team", generateProblemStatementForTeam) // <server-url>/api/v1/problemStatements/team
	group.Post("/confirm", confirmProblemStatement)      // <server-url>/api/v1/problemStatements/confirm

	group.Use(middleware.IsAdminMiddleware)
	group.Get("/", getProblemStatements)         // <server-url>/api/v1/problemStatements/
	group.Get("/:id", getProblemStatement)       // <server-url>/api/v1/problemStatements/:id
	group.Post("/", createProblemStatement)      // <server-url>/api/v1/problemStatements/
	group.Put("/:id", updateProblemStatement)    // <server-url>/api/v1/problemStatements/:id
	group.Delete("/:id", deleteProblemStatement) // <server-url>/api/v1/problemStatements/:id
}

// Get a list of all problemStatements
func getProblemStatements(c *fiber.Ctx) error {
	problemStatements, err := models.GetProblemStatements()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.ProblemStatementListSerializer(problemStatements))
}

// Get a problemStatement by id
func getProblemStatement(c *fiber.Ctx) error {
	problemStatement, err := models.GetProblemStatement(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(schemas.ProblemStatementSerializer(problemStatement))
}

// Create a new problemStatement
func createProblemStatement(c *fiber.Ctx) error {
	var requestBody struct {
		Name        string `json:"name"`
		OneLiner    string `json:"one_liner"`
		Description string `json:"description"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.InvalidBody)
	}

	// Validate problemStatement name
	if !models.ValidateProblemStatementName(requestBody.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid problemStatement name",
		})
	}

	problemStatement := models.ProblemStatement{
		Name:        requestBody.Name,
		Description: requestBody.Description,
		OneLiner:    requestBody.OneLiner,
	}

	if err := problemStatement.CreateProblemStatement(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(schemas.ProblemStatementSerializer(problemStatement))
}

// Update an existing problemStatement
func updateProblemStatement(c *fiber.Ctx) error {
	problemStatement, err := models.GetProblemStatement(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}

	var requestBody struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		OneLiner    string `json:"one_liner"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.InvalidBody)
	}

	if requestBody.Name != "" {
		// Validate problemStatement name
		if !models.ValidateProblemStatementName(requestBody.Name) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"detail": "Invalid problemStatement name",
			})
		}
		problemStatement.Name = requestBody.Name
	}

	if requestBody.Description != "" {
		problemStatement.Description = requestBody.Description
	}

	if requestBody.OneLiner != "" {
		problemStatement.OneLiner = requestBody.OneLiner
	}

	if err := problemStatement.UpdateProblemStatement(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}

	return c.Status(fiber.StatusAccepted).JSON(schemas.ProblemStatementSerializer(problemStatement))
}

// Delete an existing problemStatement
func deleteProblemStatement(c *fiber.Ctx) error {
	problemStatement, err := models.GetProblemStatement(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}

	if err := problemStatement.DeleteProblemStatement(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Get a problemStatement for a team
func getProblemStatementForTeam(c *fiber.Ctx) error {
	team, err := models.GetTeamByName(c.Locals("team").(models.Team).Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}

	if team == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"detail": "Team not found",
		})
	}

	// Check if query param type is present
	if c.Query("type") == "one_liner" {
		return c.Status(fiber.StatusOK).JSON(schemas.ProblemStatementOneLinerSerializer(team.ProblemStatement))
	}

	return c.Status(fiber.StatusOK).JSON(schemas.ProblemStatementGenerationSerializer(team.ProblemStatement, team.StatementGenerations))
}

// Generate a problemStatement for a team
func generateProblemStatementForTeam(c *fiber.Ctx) error {
	team, err := models.GetTeamByName(c.Locals("team").(models.Team).Name)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}

	if team == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"detail": "Team not found",
		})
	}

	problemStatement, err := models.GenerateProblemStatementForTeam(team)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}
	if problemStatement == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"detail": "All 3 problem generations exhausted",
		})
	}

	return c.Status(fiber.StatusOK).JSON(schemas.ProblemStatementGenerationSerializer(*problemStatement, team.StatementGenerations))
}

// Confirm the problem statement selected
func confirmProblemStatement(c *fiber.Ctx) error {
	team, err := models.GetTeamByName(c.Locals("team").(models.Team).Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(schemas.InternalServerError)
	}

	if team == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"detail": "Team not found",
		})
	}

	team.StatementGenerations = 0
	err = team.UpdateTeam()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(schemas.TeamSerializer(*team))
}
