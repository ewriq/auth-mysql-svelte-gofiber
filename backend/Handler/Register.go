package Handlers

import (
	"fmt"

	"mysql-auth/Database"
	"mysql-auth/Utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var reqbody UserBody

	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	email := reqbody.Email
	password := reqbody.Password
	username := reqbody.Username

	fmt.Print(email)
	fmt.Print(password)
	fmt.Print(username)
	
	if isValidEmail(email) {
		if isPasswordValid(password) {
            token := Utils.Token(10)
			fmt.Print(token)
			err := Database.Register(username, password, email, token)
			if err != true {
				c.JSON(fiber.Map{
					"status": 502,
					"error":  err,
				})
			} else {
				c.JSON(fiber.Map{
					"status": "OK",
					"message": "User registered successfully",
					"token": token,
				})
			}
		} else {
			c.JSON(fiber.Map{
				"status": "ERROR",
				"error":  "Geçersiz parola.",
			})
		}
	} else {
		c.JSON(fiber.Map{
			"status": "ERROR",
			"error":  "Geçersiz email.",
		})
	}

	return nil
}