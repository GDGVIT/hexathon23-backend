package api

import (
	v1 "github.com/GDGVIT/hexathon23-backend/hexathon-api/api/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewWebApi() *fiber.App {
	app := fiber.New()

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     "*",
			AllowHeaders:     "Origin, Content-Type, Accept",
			AllowCredentials: true,
			AllowMethods:     "GET,POST,DELETE,PATCH,PUT,OPTIONS",
		},
	))

	// Root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Welcome to Hexathon API!🎉")
	})

	// Ping endpoint
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"detail": "pong",
			})
	})

	// Register version routers
	api := app.Group("/api")
	v1.V1handler(api)

	return app
}
