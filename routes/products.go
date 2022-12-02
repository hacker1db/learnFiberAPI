package routes

import (
	"errors"
	"fmt"
	"learnFiberAPI/database"
	"learnFiberAPI/handlers"
	"learnFiberAPI/models"
	"learnFiberAPI/validators"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateResponseProduct(productModel models.Product) models.RoutesProduct {
	return models.RoutesProduct{
		ID:           productModel.ID,
		SerialNumber: productModel.SerialNumber,
		Name:         productModel.Name,
	}
}

func FindProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("product does exist")
	}
	return nil

}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	product.SerialNumber = uuid.New()
	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)

}
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.Database.Db.Find(&products)
	var responseProducts []models.RoutesProduct
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}

func GetProduct(c *fiber.Ctx) error {
	id, product, err := validators.IdProductValidator(c)

	err = handlers.IdErrorHandler(err, c)
	if err != nil {
		return err
	}

	if err := FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	responseUser := CreateResponseProduct(product)

	return c.Status(200).JSON(responseUser)

}

func UpdateProduct(c *fiber.Ctx) error {

	id, product, err := validators.IdProductValidator(c)

	err = handlers.IdErrorHandler(err, c)
	if err != nil {
		return err
	}

	if err := FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())

	}

	var updateData models.Product
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	product.Name = updateData.Name

	database.Database.Db.Save(&product)
	responseUser := CreateResponseProduct(product)

	return c.Status(200).JSON(responseUser)

}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("please ensure that :id is an integer")
	}

	if err := FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())

	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseString := fmt.Sprintf("Successfully Deleted Product. ProductID: %v ProductName: %v", product.ID, product.Name)
	return c.Status(200).SendString(responseString)
}
