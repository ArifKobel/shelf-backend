package auth_handler

import (
	auth_service "github.com/Arifkobel/shelf/services/auth"
	"github.com/Arifkobel/shelf/services/database"
	"github.com/Arifkobel/shelf/services/database/schemas"
	errors_service "github.com/Arifkobel/shelf/services/errors"
	"github.com/gofiber/fiber/v3"
)

func Register() fiber.Handler {
	return func(c fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		db, err := database.Connect()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().DbConnectionFailed,
			})
		}
		result := db.Create(&schemas.User{
			Email:    email,
			Password: password,
			Providers: []schemas.Provider{
				{
					Name: "email",
				},
			},
			Files: nil,
		})
		if result.Error.Error() == `ERROR: duplicate key value violates unique constraint "uni_users_email" (SQLSTATE 23505)` {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().EmailAlreadyExists,
			})
		}
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().FailedToSaveUser,
			})
		}
		jwt, err := auth_service.GenerateJwt(email)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Register success",
			"token":   jwt,
		})
	}
}
