package v1

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/middleware"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/schemas"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func teamsHandler(r fiber.Router) {
	group := r.Group("/teams")
	group.Use(middleware.JWTAuthMiddleware)
	group.Use(middleware.IsAdminMiddleware)

	// Routes
	group.Post("/admin", createAdminTeam) // <server-url>/api/v1/teams/admin
	group.Post("/", createTeam)           // <server-url>/api/v1/teams/
	group.Get("/", getTeams)              // <server-url>/api/v1/teams/
	group.Get("/:name", getTeam)          // <server-url>/api/v1/teams/:id
	group.Put("/:name", updateTeam)       // <server-url>/api/v1/teams/:id
	group.Delete("/:name", deleteTeam)    // <server-url>/api/v1/teams/:id
}

// Create an admin team
func createAdminTeam(c *fiber.Ctx) error {
	var requestBody struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid request body",
		})
	}

	if requestBody.Name == "" || requestBody.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid request body",
		})
	}

	// Validate team name
	if !models.ValidateTeamName(requestBody.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid team name",
		})
	}

	// Check if team name exists
	if models.CheckTeamNameExists(requestBody.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Team name already exists",
		})
	}

	// Encrypt the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}

	teamPassword := string(passwordHash)

	team := models.Team{
		Name:     requestBody.Name,
		Password: teamPassword,
		Role:     "admin",
	}

	err = team.CreateTeam()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(schemas.TeamCredentialsSerializer(team, requestBody.Password))
}

// Create a new team(Different from register as done from admin side)
func createTeam(c *fiber.Ctx) error {
	var requestBody struct {
		Name    string   `json:"name"`
		Members []string `json:"members"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid request body",
		})
	}

	if requestBody.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid request body",
		})
	}

	// Validate team name
	if !models.ValidateTeamName(requestBody.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid team name",
		})
	}

	// Check if team name exists
	if models.CheckTeamNameExists(requestBody.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Team name already exists",
		})
	}

	// Generate a random password with 12 characters of length, 3 digits and 3 symbols
	pwd, err := password.Generate(12, 3, 3, false, false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}

	// Encrypt the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}

	teamPassword := string(passwordHash)

	team := models.Team{
		Name:     requestBody.Name,
		Password: teamPassword,
	}
	team.SetMembers(requestBody.Members)

	err = team.CreateTeam()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(schemas.TeamCredentialsSerializer(team, pwd))
}

// Get a list of all teams
func getTeams(c *fiber.Ctx) error {
	teams, err := models.GetTeams()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}
	return c.JSON(schemas.TeamListSerializer(teams))
}

// Get a team by id
func getTeam(c *fiber.Ctx) error {
	team, err := models.GetTeamByName(c.Params("name"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"detail": "Team not found",
		})
	}
	return c.JSON(schemas.TeamSerializer(*team))
}

// Update a team by id
func updateTeam(c *fiber.Ctx) error {
	team, err := models.GetTeamByName(c.Params("name"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"detail": "Team not found",
		})
	}

	var requestBody struct {
		Name     string   `json:"name"`
		Password string   `json:"password"`
		Members  []string `json:"members"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid request body",
		})
	}

	if requestBody.Name != "" {
		// Validate team name
		if !models.ValidateTeamName(requestBody.Name) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"detail": "Invalid team name",
			})
		}

		// Check if team name exists
		if models.CheckTeamNameExists(requestBody.Name) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"detail": "Team name already exists",
			})
		}
		team.Name = requestBody.Name
	}
	if requestBody.Password != "" {
		// Validate team password
		if !models.ValidateTeamPassword(requestBody.Password) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"detail": "Invalid team password",
			})
		}

		// Encrypt the password
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"detail": "Internal Server Error",
			})
		}

		team.Password = string(passwordHash)
	}
	if requestBody.Members != nil {
		team.SetMembers(requestBody.Members)
	}

	err = team.UpdateTeam()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}
	return c.JSON(schemas.TeamSerializer(*team))
}

// Delete a team by id
func deleteTeam(c *fiber.Ctx) error {
	team, err := models.GetTeamByName(c.Params("name"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"detail": "Team not found",
		})
	}

	err = team.DeleteTeam()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
