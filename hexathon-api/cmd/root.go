package cmd

import (
	"log"
	"os"
	"strconv"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/api"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/hexCli"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/auth"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/database"
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
)

// HexathonApp is the main application struct
type HexathonApp struct {
	env    Env
	cliApp *cli.App
	webApp *fiber.App
}

// Env is the environment variables struct
type Env struct {
	// Fiber Variables
	fiberPort string
	debug     string

	// Database Variables
	postgresUrl string

	// Auth Variables
	jwtSecret string

	// Other Variables
	defaultAmount int
}

// Create a new HexathonApp
func NewHexathonApp() *HexathonApp {
	var app HexathonApp
	// Initialize the HexathonApp
	app.init()
	return &app
}

// Set the environment variables
func (app *HexathonApp) setEnv() {
	app.env.debug = os.Getenv("DEBUG")
	app.env.fiberPort = os.Getenv("FIBER_PORT")
	app.env.postgresUrl = os.Getenv("POSTGRES_URL")
	app.env.jwtSecret = os.Getenv("JWT_SECRET")

	var err error
	app.env.defaultAmount, err = strconv.Atoi(os.Getenv("DEFAULT_AMOUNT"))
	if err != nil {
		app.env.defaultAmount = 1000
	}
}

// Initialize the CLI app
func (app *HexathonApp) initCliApp() {
	app.cliApp = hexCli.NewCliApp()

	app.cliApp.Commands = append(app.cliApp.Commands, &cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "Run the server",
		Action: func(c *cli.Context) error {
			app.webApp.Listen(app.env.fiberPort)
			return nil
		},
	})
}

// Initialize the Web app
func (app *HexathonApp) initWebApp() {
	app.webApp = api.NewWebApi()
}

// Initialize the HexathonApp
func (app *HexathonApp) init() {
	// Set the environment variables
	app.setEnv()

	// Initialize the database, models and auth
	database.Connect(app.env.debug, app.env.postgresUrl)
	models.InitializeModels(app.env.defaultAmount)
	auth.InitializeAuth(app.env.jwtSecret)

	// Initialize the Web app
	app.initWebApp()

	// Initialize the CLI app
	app.initCliApp()
}

// Run the HexathonApp
func (app *HexathonApp) Run() {
	// Run the CLI app
	err := app.cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
