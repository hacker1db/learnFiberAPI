package handlers

import "github.com/gofiber/fiber/v2"

func IdErrorHandler(err error, c *fiber.Ctx) error {
	if err != nil {
		return c.Status(400).JSON("please ensure that :id is an integer")
	}
	return err
}
