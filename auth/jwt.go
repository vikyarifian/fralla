package auth

import (
	"fmt"
	"fralla/db"
	"fralla/dto"
	"fralla/models"
	"fralla/utils"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

var (
	secretKey = []byte(os.Getenv("JWT_KEY"))
)

func CreateToken(token dto.Token) (string, error) {

	var no int64
	db.PgSql.Session(&gorm.Session{PrepareStmt: true})

	db.PgSql.Table("public.visitors").Select("max(no)").Row().Scan(&no)

	idhash := utils.GetMD5Hash(strconv.Itoa(int(no) + 1))
	visitor := models.Visitor{
		No:        int(no) + 1,
		ID:        idhash,
		IP:        token.IP,
		IsAuth:    false,
		CreatedBy: "System",
		UpdatedBy: "System",
	}

	claims := jwt.MapClaims{
		"id":         int(no) + 1,
		"username":   idhash,
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
		visitor.Token = tokenString
		db.PgSql.Create(&visitor)
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
