package auth

import (
	"fralla/dto"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func IsAuthenticated(c *fiber.Ctx) (dto.Visitor, bool) {

	token := c.Cookies("session")
	_, err := VerifyToken(token)
	var jwt dto.Visitor
	if err == nil {
		jwt, _ := GetTokenVisitor(c)
		return jwt, jwt.IsAuth
	} else {
		jwt = dto.Visitor{}
		CreateTokenVisitor(jwt)
	}
	return jwt, jwt.IsAuth

}

func AssertAuthenticatedMiddleware(c *fiber.Ctx) error {
	if _, r := IsAuthenticated(c); !r {
		c.Set("HX-Redirect", "/401")
		return c.Redirect("/401")
	}
	return c.Next()
}

func AssertAdminAuthenticatedMiddleware(c *fiber.Ctx) error {
	token, r := IsAuthenticated(c)
	if !r {
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

func SetSession(c *fiber.Ctx) string {
	newSessionId := c.Get("Authorization")
	c.Cookie(&fiber.Cookie{
		Name:     "session",
		Value:    newSessionId,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})
	return newSessionId
}

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
