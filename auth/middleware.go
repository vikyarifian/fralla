package auth

import (
	"fralla/dto"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func IsAuthenticated(c *fiber.Ctx) dto.Token {

	auth_token := c.Cookies("session")
	_, err := VerifyToken(auth_token)
	var token dto.Token
	if err == nil {
		jwt, _ := GetToken(c)
		return jwt
	} else {
		token = dto.Token{
			Token:     "",
			Level:     "VISITOR",
			PopUpShow: true,
			IP:        c.IP(),
		}
		auth_token, _ := CreateToken(token)
		c.Cookie(&fiber.Cookie{
			Name:     "session",
			Value:    auth_token,
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})
	}
	return token

}

func AssertAuthenticatedMiddleware(c *fiber.Ctx) error {
	if r := IsAuthenticated(c); !r.IsAuth {
		c.Set("HX-Redirect", "/401")
		return c.Redirect("/401")
	}
	return c.Next()
}

func AssertAdminAuthenticatedMiddleware(c *fiber.Ctx) error {
	token := IsAuthenticated(c)
	if !token.IsAuth {
		c.Set("HX-Redirect", "/401")
		return c.Redirect("/401")
	} else {
		if token.Level != "ADMIN" {
			c.Set("HX-Redirect", "/401")
			return c.Redirect("/401")
		}
	}
	return c.Next()
}

func GetUserSessionId(c *fiber.Ctx) string {
	return c.Cookies("session")
}

// func SetSession(c *fiber.Ctx) string {
// 	newSessionId := c.Get("Authorization")
// 	c.Cookie(&fiber.Cookie{
// 		Name:     "session",
// 		Value:    newSessionId,
// 		HTTPOnly: true,
// 		Secure:   true,
// 		SameSite: "Strict",
// 	})
// 	return newSessionId
// }

func ClearSession(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    "",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
