package auth

import (
	"fmt"
	"fralla/dto"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	secretKey = []byte(os.Getenv("JWT_KEY"))
)

func CreateTokenVisitor(visitor dto.Visitor) (string, error) {

	claims := jwt.MapClaims{
		"id":         visitor.ID,
		"username":   visitor.Username,
		"first_name": cases.Title(language.English, cases.Compact).String(visitor.FirstName),
		"last_name":  cases.Title(language.English, cases.Compact).String(visitor.LastName),
		"email":      visitor.Email,
		"phone":      visitor.Phone,
		"level":      visitor.Level,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetTokenVisitor(c *fiber.Ctx) (dto.Visitor, error) {

	var token dto.Visitor
	claims, err := VerifyToken(c.Cookies("session"))
	if err != nil {
		return token, err
	}

	token = dto.Visitor{
		Username:  fmt.Sprintf("%s", claims["username"]),
		FirstName: cases.Title(language.English, cases.Compact).String(fmt.Sprintf("%s", claims["first_name"])),
		LastName:  cases.Title(language.English, cases.Compact).String(fmt.Sprintf("%s", claims["last_name"])),
		Level:     fmt.Sprintf("%s", claims["level"]),
		Token:     c.Get("Authorization"),
		IsAuth:    false,
	}

	if strings.Trim(token.Level, " ") != "visitor" {
		token.IsAuth = true
	}

	return token, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fiber.ErrUnauthorized
	}

}
