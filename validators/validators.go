package validators

import (
	"github.com/gofiber/fiber/v2"
	"learnFiberAPI/models"
)

func IdUserValidator(c *fiber.Ctx) (int, models.User, error) {
	id, err := c.ParamsInt("id")
	user := UserModelCreate()

	return id, user, err
}
func IdProductValidator(c *fiber.Ctx) (int, models.Product, error) {
	id, err := c.ParamsInt("id")
	user := models.Product{}

	return id, user, err
}
func UserModelCreate() models.User {

	var userModels models.User
	return userModels
}
