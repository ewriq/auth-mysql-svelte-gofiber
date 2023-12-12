package Handlers

import (
	"github.com/gofiber/fiber/v2"

)

func Home(c *fiber.Ctx) error {
	c.JSON(fiber.Map{
		"hello":"world",
	})
	return nil
}