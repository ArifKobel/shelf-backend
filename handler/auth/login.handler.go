package auth_handler

import (
	auth_service "github.com/Arifkobel/shelf/services/auth"
	"github.com/Arifkobel/shelf/services/database"
	"github.com/Arifkobel/shelf/services/database/schemas"
	"github.com/gofiber/fiber/v3"
)

func Login() fiber.Handler {
	return func(c fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		db, err := database.Connect()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to connect to database",
			})
		}
		res := db.Where("email = ? AND password = ?", email, password).First(&schemas.User{})
		if res.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Email or password is incorrect",
			})
		}
		jwt, err := auth_service.GenerateJwt(email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to generate token",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Login success",
			"token":   jwt,
		})
	}
}
