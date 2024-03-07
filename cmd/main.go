package main

import (
	file_handler "github.com/Arifkobel/shelf/handler/file"
	"github.com/Arifkobel/shelf/services/database"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 1024,
	})
	database.Migrate()
	fileRoute := app.Group("/file")
	fileRoute.Post("/upload", file_handler.StoreFile())
	app.Listen(":3000")
}
