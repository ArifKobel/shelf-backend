package file_handler

import (
	"fmt"

	auth_service "github.com/Arifkobel/shelf/services/auth"
	"github.com/Arifkobel/shelf/services/database"
	"github.com/Arifkobel/shelf/services/database/schemas"
	errors_service "github.com/Arifkobel/shelf/services/errors"
	file_service "github.com/Arifkobel/shelf/services/file"
	"github.com/gofiber/fiber/v3"
)

func StoreFile() fiber.Handler {
	return func(c fiber.Ctx) error {
		jwt := c.Request().Header.Peek("Authorization")
		if len(jwt) == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().Unauthorized,
			})
		}
		email, err := auth_service.GetEmailFromToken(string(jwt))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().Unauthorized,
			})
		}
		fmt.Println(email)
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
				"message": errors_service.GetErrorMessages().DbConnectionFailed,
			})
		}
		userResponse := db.Select("id").Where("email = ?", email).First(&schemas.User{})
		if userResponse.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().UserNotFound,
			})
		}
		user := schemas.User{}
		userResponse.Scan(&user)
		fileModel := schemas.File{
			Name:     file.Filename,
			FileName: file_name,
			UserID:   user.Id,
		}
		result := db.Create(&fileModel)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().FailedToSaveFile,
			})
		}

		err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", file_name))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().FailedToSaveFile,
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "File uploaded successfully",
			"file":    file.Filename,
		})
	}
}

func GetAllFiles() fiber.Handler {
	return func(c fiber.Ctx) error {
		jwt := c.Request().Header.Peek("Authorization")
		if len(jwt) == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().Unauthorized,
			})
		}
		email, err := auth_service.GetEmailFromToken(string(jwt))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().Unauthorized,
			})
		}
		db, err := database.Connect()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().DbConnectionFailed,
			})
		}
		userResponse := db.Select("id").Where("email = ?", email).First(&schemas.User{})
		if userResponse.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().UserNotFound,
			})
		}
		user := schemas.User{}
		userResponse.Scan(&user)
		var files []schemas.File
		db.Where("user_id = ?", user.Id).Find(&files)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"files": files,
		})
	}
}

func DownloadFile() fiber.Handler {
	return func(c fiber.Ctx) error {
		jwt := c.Request().Header.Peek("Authorization")
		if len(jwt) == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().Unauthorized,
			})
		}
		email, err := auth_service.GetEmailFromToken(string(jwt))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().Unauthorized,
			})
		}
		db, err := database.Connect()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().DbConnectionFailed,
			})
		}
		userResponse := db.Select("id").Where("email = ?", email).First(&schemas.User{})
		if userResponse.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().UserNotFound,
			})
		}
		user := schemas.User{}
		userResponse.Scan(&user)
		fileId := c.Params("id")
		file := schemas.File{}
		db.Where("id = ? AND user_id = ?", fileId, user.Id).First(&file)
		if file.Id == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": errors_service.GetErrorMessages().FileNotFound,
			})
		}
		return c.SendFile(fmt.Sprintf("./uploads/%s", file.FileName))
	}
}
