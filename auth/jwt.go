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
	"time"

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

	if strings.Trim(token.Username, " ") == "" {
		token.Username = idhash
	}
	println(func(t time.Time) *time.Time { return &t }(time.Now().In(utils.GetLocationTime())))
	visitor := models.Visitor{
		No:        int(no) + 1,
		ID:        idhash,
		IP:        token.IP,
		Username:  token.Username,
		IsAuth:    false,
		CreatedBy: "System",
		CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now().In(utils.GetLocationTime())),
		UpdatedBy: "System",
		UpdatedAt: func(t time.Time) *time.Time { return &t }(time.Now().In(utils.GetLocationTime())),
	}

	claims := jwt.MapClaims{
		"id":         int(no) + 1,
		"username":   token.Username,
		"first_name": cases.Title(language.English, cases.Compact).String(token.FirstName),
		"last_name":  cases.Title(language.English, cases.Compact).String(token.LastName),
		"email":      token.Email,
		"phone":      token.Phone,
		"level":      token.Level,
		"popup_show": strconv.FormatBool(token.PopUpShow),
	}

	auth_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := auth_token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	if token.Token == "" {
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

	popup_show, _ := strconv.ParseBool(fmt.Sprintf("%s", claims["popup_show"]))
	token = dto.Token{
		IP:        c.IP(),
		Username:  fmt.Sprintf("%s", claims["username"]),
		FirstName: cases.Title(language.English, cases.Compact).String(fmt.Sprintf("%s", claims["first_name"])),
		LastName:  cases.Title(language.English, cases.Compact).String(fmt.Sprintf("%s", claims["last_name"])),
		Level:     fmt.Sprintf("%s", claims["level"]),
		Token:     c.Get("Authorization"),
		IsAuth:    false,
		PopUpShow: popup_show,
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
