package main

import (
	Routers "mysql-auth/Router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	Routers.Initalize(app)

	app.Listen(":3000")
}