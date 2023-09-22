package v1

import (
	"strings"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api/schemas"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/auth"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// AuthHandler handles all the routes related to authentication
func authHandler(r fiber.Router) {
	group := r.Group("/auth")

	// Routes
	group.Post("/register", register) // <server-url>/api/v1/auth/register
	group.Post("/login", login)       // <server-url>/api/v1/auth/login
}

// Create a team with the given name and password
func register(c *fiber.Ctx) error {
	var requestBody struct {
		Name      string   `json:"name"`
		Password  string   `json:"password"`
		MemberIDs []string `json:"member_ids"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(schemas.InvalidBody)
	}

	// Check if the team already exists
	if models.CheckTeamNameExists(requestBody.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Team with the given name already exists",
		})
	}

	// Validate team name and password
	if !models.ValidateTeamName(requestBody.Name) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid team name",
		})
	}

	requestBody.Name = strings.ToLower(requestBody.Name)
	if !models.ValidateTeamPassword(requestBody.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid password. Password must be atleast 8 characters long, contain atleast one uppercase letter, one lowercase letter, one digit and one special character",
		})
	}

	// Encrypt the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Error while encrypting password",
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
			"detail": "Error while creating team",
		})
	}
	team.SetMembers(requestBody.MemberIDs)

	token, err := auth.CreateJWTToken(team.Name, team.Role, auth.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Error while creating JWT token",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(schemas.TeamLoginSerializer(team, token))
}

// Login a team with the using name and password
func login(c *fiber.Ctx) error {
	var requestBody struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid request body",
		})
	}

	// Check if the team exists
	team, err := models.GetTeamByName(requestBody.Name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Team with the given name does not exist",
		})
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(team.Password), []byte(requestBody.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Incorrect password",
		})
	}
	// if team.Password != requestBody.Password {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"detail": "Incorrect password",
	// 	})
	// }

	token, err := auth.CreateJWTToken(team.Name, team.Role, auth.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"detail": "Error while creating JWT token",
		})
	}
	return c.Status(fiber.StatusOK).JSON(schemas.TeamLoginSerializer(*team, token))
}
