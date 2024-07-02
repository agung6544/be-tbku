package main

import (
	"log"
	"os"

	"github.com/agung6544/be-tbku/config"
	"github.com/agung6544/be-tbku/url"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/agung6544/be-tbku/docs"

)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/agung6544
// @contact.email 714220039@std.ulbi.ac.id

// @host be-tbku1-179bbc23801d.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(config.Cors))

	// Routes
	url.Web(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("Sorry, page not found!")
	})

	// Get port from environment variable or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3200"
	}

	// Start server
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
