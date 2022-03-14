package controllers

import (
	"net/http"

	"github.com/Amir-Hossein-shamsi/fiber-resapi/src/db"
	"github.com/Amir-Hossein-shamsi/fiber-resapi/src/db/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handller struct {
	Handler db.Repository
}

func (r *Handller) CreateUser(context *fiber.Ctx) error {
	user := models.User{}
	err := context.BodyParser(&user)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	validator := validator.New()
	err = validator.Struct(models.User{})
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": err},
		)
		return err
	}

	err = r.Handler.DB.Create(&user).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book has been successfully added",
	})
	return nil
}

func (r *Handller) UpdateUser(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	userModel := &models.User{}
	user := models.User{}

	err := context.BodyParser(&user)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.Handler.DB.Model(userModel).Where("id = ?", id).Updates(user).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not update book",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book has been successfully updated",
	})
	return nil

}

func (r *Handller) GetUsers(context *fiber.Ctx) error {
	userModel := &[]models.User{}

	err := r.Handler.DB.Find(userModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books gotten successfully",
		"data":    userModel,
	})
	return nil

}
