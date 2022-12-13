package main

import (
	"github.com/gofiber/fiber/v2"
	"learnFiberAPI/database"
	"learnFiberAPI/routes"
	"log"
	"os"
)

func HealthCheck(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		err := c.Status(200).JSON(" Everything appears to be Healthy")
		if err != nil {
			return c.Status(400).JSON(err.Error())
		}
		return err
	})
}

func SetUpUserRoutes(app *fiber.App) {

	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/Users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
}

func SetUpProductRoutes(app *fiber.App) {

	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)
}

func SetUpOrdersRoutes(app *fiber.App) {
	app.Post("/api/order", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/order/:id", routes.GetOrder)

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	HealthCheck(app)
	SetUpUserRoutes(app)
	SetUpProductRoutes(app)

	if os.Getenv("Env") == "production" {
		log.Fatal(app.Listen(":3000"))

	}
	log.Fatal(app.Listen("localhost:3000"))
}
