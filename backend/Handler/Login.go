package Handlers

import (
	"fmt"
	"mysql-auth/Database"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var reqbody UserBody

	if err := c.BodyParser(&reqbody); err != nil {
		return err
	}

	email := reqbody.Email
	password := reqbody.Password
	fmt.Print(email)
	fmt.Print(password)

	if isValidEmail(email) {
		if isPasswordValid(password) {
			err, token  := Database.Login(email, password)
			fmt.Println(token,err,"----------------------------------------------------")
			if err == "" {
				c.JSON(fiber.Map{
					"status": "ERROR",
					"msg":    "Kullanıcı bulunamadı veya e-posta, şifreniz hatalı ",
				})
			} else if err == "." {
				c.JSON(fiber.Map{
					"status": "ERROR",
					"msg":    "Kullanıcı bulunamadı lütfen kayıt olunuz",
				})
			} else {
				c.JSON(fiber.Map{
					"status": "OK",
					"msg":    "Kullanıcı başarıyla giriş yaptı",
					"token":  err,
				})
			}
		} else {
			c.JSON(fiber.Map{
				"status": "ERROR",
				"msg":    "Geçersiz parola.",
			})
		}
	} else {
		c.JSON(fiber.Map{
			"status": "ERROR",
			"msg":    "Geçersiz email.",
		})
	}

	return nil
}
func isValidEmail(email string) bool {
	return strings.Contains(email, "@")
}

func isPasswordValid(password string) bool {
	return len(password) >= 8
}