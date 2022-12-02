package routes

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	database "learnFiberAPI/database"
	"learnFiberAPI/handlers"
	"learnFiberAPI/models"
	"learnFiberAPI/validators"
)

func CreateResponseUser(userModel models.User) models.UserSerializer {

	return models.UserSerializer{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}

}

func CreateUser(c *fiber.Ctx) error {
	userModels := validators.UserModelCreate()

	if err := c.BodyParser(&userModels); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&userModels)
	responseUser := CreateResponseUser(userModels)
	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.Database.Db.Find(&users)
	var responseUsers []models.UserSerializer
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)

}

func FindUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does exist")
	}
	return nil

}

func GetUser(c *fiber.Ctx) error {
	id, user, err := validators.IdUserValidator(c)

	err = handlers.IdErrorHandler(err, c)
	if err != nil {
		return err
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {

	id, user, err := validators.IdUserValidator(c)

	err = handlers.IdErrorHandler(err, c)
	if err != nil {
		return err
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())

	}

	var updateData models.UpdateUser
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName
	database.Database.Db.Save(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)

}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("please ensure that :id is an integer")
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())

	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseString := fmt.Sprintf("Successfully Deleted user. UserID: %v UserName: %v %v", user.ID, user.FirstName, user.LastName)
	return c.Status(200).SendString(responseString)
}
