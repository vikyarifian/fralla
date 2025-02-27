package auth

import (
	"fralla/dto"

	"github.com/golang-jwt/jwt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	secretKey = []byte("fr4ll@")
)

func CreateToken(visitor dto.Visitor) (string, error) {

	claims := jwt.MapClaims{
		"id":         visitor.ID,
		"username":   visitor.Username,
		"first_name": cases.Title(language.English, cases.Compact).String(visitor.FirstName),
		"last_name":  cases.Title(language.English, cases.Compact).String(visitor.LastName),
		"email":      visitor.Email,
		"phone":      visitor.Phone,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
