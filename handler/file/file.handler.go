package file_handler

import (
	"fmt"

	"github.com/Arifkobel/shelf/services/database"
	file_service "github.com/Arifkobel/shelf/services/file"
	"github.com/gofiber/fiber/v3"
)

func StoreFile() fiber.Handler {
	return func(c fiber.Ctx) error {
		file, err := c.FormFile("document")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "File not found",
			})
		}
		file_name := file_service.GenerateFileName(file.Filename)
		db, err := database.Connect()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Database connection failed",
			})
		}
		defer db.Close()
		fmt.Println(file_name)
		_, err = db.Exec("INSERT INTO files (name, file_name) VALUES ($1, $2)", file.Filename, file_name)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to store file",
			})
		}

		err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", file_name))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to save file",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "File uploaded successfully",
			"file":    file.Filename,
		})
	}
}
