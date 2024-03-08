package main

import (
	"log"

	auth_handler "github.com/Arifkobel/shelf/handler/auth"
	file_handler "github.com/Arifkobel/shelf/handler/file"
	"github.com/Arifkobel/shelf/services/database"
	defaults_tables "github.com/Arifkobel/shelf/services/database/defaults"
	"github.com/Arifkobel/shelf/services/database/schemas"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 1024,
	})
	db, err := database.Connect()
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(schemas.User{}, schemas.File{}, schemas.Provider{})
	defaults_tables.SetProviders(db)
	fileRoute := app.Group("/file")
	fileRoute.Post("/upload", file_handler.StoreFile())
	fileRoute.Get("/list", file_handler.GetAllFiles())
	fileRoute.Get("/download/:id", file_handler.DownloadFile())
	authRoute := app.Group("/auth")
	authRoute.Post("/register", auth_handler.Register())
	authRoute.Post("/login", auth_handler.Login())
	app.Listen(":3000")
}
