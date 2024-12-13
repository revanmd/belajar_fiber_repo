package main

import (
	"log"

	"belajar-api/config"
	"belajar-api/database"
	"belajar-api/routes"
	"belajar-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load Config
	cfg := config.LoadConfig()

	// Initialize JWT
	utils.InitJWT(cfg.JWTSecret)

	// Connect to MongoDB
	db := database.Connect(cfg)
	database.Migrate(db)

	// Initialize Validator
	utils.InitValidator()

	// Initialize Fiber
	app := fiber.New(fiber.Config{BodyLimit: 35 * 1024 * 1024})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins
	}))
	app.Use(recover.New())

	// Setup Routes with db dependency
	routes.SetupRoutes(app, db, cfg)

	// Start server
	log.Fatal(app.Listen(":" + cfg.Port))
}
