package Handlers

import (
	"fmt"
	"mysql-auth/Database"

	"github.com/gofiber/fiber/v2"
)

func User(c *fiber.Ctx) error {
	var reqbody UserInfo
	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	token := reqbody.Token
	user, err := Database.Users(token)
	fmt.Println(token)
    fmt.Println("Handler",user)
	if err != nil {
		c.JSON(fiber.Map{
			"status": "error",
		})
		return err
	}

	c.JSON(fiber.Map{
		"status": "OK",
		"data":   user,
	})
	

	return nil
}