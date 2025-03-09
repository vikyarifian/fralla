package handlers

import (
	"fralla/auth"

	"github.com/gofiber/fiber/v2"
)

func HidePopUp(c *fiber.Ctx) error {
	token, _ := auth.GetToken(c)
	print(c.Query("popup"))
	if c.Query("popup") == "on" {
		token.PopUpShow = false
	} else {
		token.PopUpShow = true
	}

	auth_token, _ := auth.CreateToken(token)
	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    auth_token,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
	return c.JSON(fiber.Map{"status": "success"})
}
