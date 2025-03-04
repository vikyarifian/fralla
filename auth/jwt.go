package auth

import (
	"fmt"
	"fralla/db"
	"fralla/dto"
	"fralla/models"
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

func CreateToken(token dto.Token) (string, error) {

	claims := jwt.MapClaims{
		"id":         token.ID,
		"username":   token.Username,
		"first_name": cases.Title(language.English, cases.Compact).String(token.FirstName),
		"last_name":  cases.Title(language.English, cases.Compact).String(token.LastName),
		"email":      token.Email,
		"phone":      token.Phone,
		"level":      token.Level,
	}

	auth_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := auth_token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	if strings.Trim(token.Level, " ") == "VISITOR" {
		visitor := models.Visitor{
			IP:        token.IP,
			Token:     tokenString,
			IsAuth:    false,
			CreatedBy: "System",
			UpdatedBy: "System",
		}

		// maxId := db.PgSql.Table("visitors").Select("max(no)").Row()
		// _ = maxId.Scan(&visitor.No)
		// idhash := utils.GetMD5Hash(strconv.Itoa(visitor.No + 1))
		// visitor.No = visitor.No + 1
		// visitor.ID = idhash
		db.PgSql.Save(&visitor)
	}

	return tokenString, nil
}

func GetToken(c *fiber.Ctx) (dto.Token, error) {

	var token dto.Token
	claims, err := VerifyToken(c.Cookies("session"))
	if err != nil {
		return token, err
	}

	token = dto.Token{
		IP:        c.IP(),
		Username:  fmt.Sprintf("%s", claims["username"]),
		FirstName: cases.Title(language.English, cases.Compact).String(fmt.Sprintf("%s", claims["first_name"])),
		LastName:  cases.Title(language.English, cases.Compact).String(fmt.Sprintf("%s", claims["last_name"])),
		Level:     fmt.Sprintf("%s", claims["level"]),
		Token:     c.Get("Authorization"),
		IsAuth:    false,
	}

	if strings.Trim(token.Level, " ") == "USER" || strings.Trim(token.Level, " ") == "ADMIN" {
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
