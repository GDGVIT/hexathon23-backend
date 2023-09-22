package v1

import (
	"fmt"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/middleware"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/schemas"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

// TeamsHandler handles all the routes related to teams
func teamsHandler(r fiber.Router) {
	group := r.Group("/teams")

	// Routes
	group.Use(middleware.JWTAuthMiddleware)
	group.Get("/me", getMyTeam) // <server-url>/api/v1/teams/me

	group.Use(middleware.IsAdminMiddleware)
	group.Get("/", getTeams)                                           // <server-url>/api/v1/teams/
	group.Post("/:name/regeneratePassword", regenerateTeamPassword)    // <server-url>/api/v1/teams/:name/regeneratePassword
	group.Get("/:name", getTeam)                                       // <server-url>/api/v1/teams/:id
	group.Post("/admin", createAdminTeam)                              // <server-url>/api/v1/teams/admin
	group.Post("/", createTeam)                                        // <server-url>/api/v1/teams/
	group.Put("/:name", updateTeam)                                    // <server-url>/api/v1/teams/:id
	group.Delete("/:name", deleteTeam)                                 // <server-url>/api/v1/teams/:id
	group.Post("/checkout", checkoutTeams)                             // <server-url>/api/v1/teams/checkout
	group.Post("/confirmProblemStatement", confirmAllProblemStatement) // <server-url>/api/v1/teams/confirmProblemStatement
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
		Name      string   `json:"name"`
		MemberIDs []string `json:"member_ids"`
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

	err = team.CreateTeam()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}
	team.SetMembers(requestBody.MemberIDs)

	return c.Status(fiber.StatusCreated).JSON(schemas.TeamCredentialsSerializer(team, pwd))
}

// Get a list of all teams
func getTeams(c *fiber.Ctx) error {
	teams, err := models.GetTeams()

	var resTeams []models.Team
	if c.Query("checked_out") == "true" {

		for teamIndex, team := range teams {
			cart, err := team.GetCart()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"detail": fmt.Sprintf("Error getting cart for team %s: %s", team.Name, err.Error()),
				})
			}
			if cart.CheckedOut {
				resTeams = append(resTeams, teams[teamIndex])
			}
		}
	} else if c.Query("checked_out") == "false" {
		for teamIndex, team := range teams {
			cart, err := team.GetCart()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"detail": fmt.Sprintf("Error getting cart for team %s: %s", team.Name, err.Error()),
				})
			}
			if !cart.CheckedOut {
				resTeams = append(resTeams, teams[teamIndex])
			}
		}
	} else {
		resTeams = teams
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(schemas.TeamListSerializer(resTeams))
}

// Get a team by id
func getTeam(c *fiber.Ctx) error {
	team, err := models.GetTeamByName(c.Params("name"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"detail": "Team not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(schemas.TeamSerializer(*team))
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
		Name                        string   `json:"name"`
		Password                    string   `json:"password"`
		MemberIds                   []string `json:"member_ids"`
		ProblemStatementGenerations int      `json:"ps_generations"`
		StatementConfirmed          bool     `json:"ps_confirmed"`
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
	if requestBody.MemberIds != nil {
		team.SetMembers(requestBody.MemberIds)
	}

	if requestBody.ProblemStatementGenerations != 0 {
		team.StatementGenerations = requestBody.ProblemStatementGenerations
	}

	team.StatementConfirmed = requestBody.StatementConfirmed

	err = team.UpdateTeam()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(schemas.TeamSerializer(*team))
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

// Get my team
func getMyTeam(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(schemas.TeamCheckoutSerializer(c.Locals("team").(models.Team)))
}

// Regenerate password for a team
func regenerateTeamPassword(c *fiber.Ctx) error {
	team, err := models.GetTeamByName(c.Params("name"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"detail": "Team not found",
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

	team.Password = string(passwordHash)

	err = team.UpdateTeam()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(schemas.TeamCredentialsSerializer(*team, pwd))
}

// Checkout all teams who have not checked out
func checkoutTeams(c *fiber.Ctx) error {
	teams, err := models.GetTeams()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Internal Server Error",
		})
	}

	statusCheckout := make(map[string]string)

	for _, team := range teams {
		cart, err := team.GetCart()
		if err != nil {
			statusCheckout[team.Name] = "Error getting cart"
		}
		if !cart.CheckedOut {
			err = cart.CheckoutCart()
			if err != nil {
				statusCheckout[team.Name] = "Error checking out cart"
			} else {
				statusCheckout[team.Name] = "Checked out successfully"
			}
		}
	}

	return c.Status(fiber.StatusOK).JSON(statusCheckout)
}

// Confirm the problem statement for all teams
func confirmAllProblemStatement(c *fiber.Ctx) error {
	teams, err := models.GetTeams()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": fmt.Sprintf("Error getting teams: %s", err.Error())})
	}

	statusCheckout := make(map[string]string)

	for _, team := range teams {
		team.StatementConfirmed = true
		team.StatementGenerations = 0
		err = team.UpdateTeam()
		if err != nil {
			statusCheckout[team.Name] = "Error updating team"
		} else {
			statusCheckout[team.Name] = "Confirmed problem statement"
		}
	}
	return c.Status(fiber.StatusOK).JSON(statusCheckout)
}
