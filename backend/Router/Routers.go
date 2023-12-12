package Routers

import (
	Handler "mysql-auth/Handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Initalize(app *fiber.App){
	v1 := app.Group("/api/v1")
	
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders:  "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	app.Use(logger.New(logger.Config{
		Format: "${ip}:${port} ${status} - ${method} ${path}\n",
	}))

	v1.Post("/user", Handler.User) 
    v1.Post("/login", Handler.Login)
    v1.Post("/register", Handler.Register)
	app.Get("/", Handler.Home)
	
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}