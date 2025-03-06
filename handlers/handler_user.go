package handlers

import (
	"fralla/auth"
	"fralla/db"
	"fralla/dto"
	"fralla/models"
	"fralla/templ/components"
	"fralla/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func HandleRegister(c *fiber.Ctx) error {
	var users []models.User
	var newUser models.User
	var count int64
	username := c.FormValue("username")
	email := c.FormValue("email")
	phone := c.FormValue("phone")
	password := c.FormValue("password")
	err := db.PgSql.Where("username=? or email=? or phone=?", username, email, phone).Find(&users).Count(&count).Error
	if err != nil {
		return utils.Render(c, components.ErrorAlert(err.Error(), "register"), templ.WithStatus(http.StatusBadRequest))
	}
	if count > 0 {
		msg := ""
		for _, user := range users {
			if strings.Trim(user.Username, " ") == strings.Trim(username, " ") {
				if msg != "" {
					msg = msg + ", "
				}
				msg = "Username"
			}
			if strings.Trim(user.Email, " ") == strings.Trim(email, " ") {
				if msg != "" {
					msg = msg + ", "
				}
				msg = msg + "Email"
			}
			if strings.Trim(user.Phone, " ") == strings.Trim(phone, " ") {
				if msg != "" {
					msg = msg + ", "
				}
				msg = msg + "Phone"
			}
		}

		return utils.Render(c, components.ErrorAlert(msg+" already exist!", "register"), templ.WithStatus(http.StatusBadRequest))
	}

	if len(strings.Trim(password, " ")) < 6 {
		return utils.Render(c, components.ErrorAlert("Password must be at least 6 characters!", "register"), templ.WithStatus(http.StatusBadRequest))
	}

	no := 0
	maxId := db.PgSql.Table("users").Select("max(no)").Row()
	_ = maxId.Scan(&no)
	idhash := utils.GetMD5Hash(strconv.Itoa(no + 1))

	t := time.Now()
	hash, _ := auth.HashPassword(password)
	newUser.Username = username
	newUser.Password = string(hash)
	newUser.FirstName = username
	newUser.Email = email
	newUser.Phone = phone
	newUser.No = no + 1
	newUser.ID = idhash
	newUser.Level = "user"
	newUser.CreatedAt = &t
	newUser.CreatedBy = newUser.Username
	newUser.UpdatedAt = &t
	newUser.UpdatedBy = newUser.Username

	err = db.PgSql.Create(&newUser).Error
	if err != nil {
		return utils.Render(c, components.ErrorAlert(err.Error(), "register"), templ.WithStatus(http.StatusBadRequest))
	}

	return utils.Render(c, components.SuccessAlert("Register success! Go to login form.", "register"), templ.WithStatus(http.StatusOK))
}

func HandleLogin(c *fiber.Ctx) error {
	username := c.FormValue("username")
	var user models.User
	err := db.PgSql.Where("username=? or email=?", username, username).First(&user).Error
	if err != nil {
		log.Println(err.Error())
		return utils.Render(c, components.ErrorAlert("Bad Credentials", "login"), templ.WithStatus(http.StatusBadRequest))
	}

	password := c.FormValue("password")
	if bcrypt.CompareHashAndPassword([]byte(strings.Trim(user.Password, " ")), []byte(password)) == nil {
		token := dto.Token{
			No:        user.No,
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			Username:  user.Username,
			Level:     user.Level,
			IP:        c.IP(),
			IsAuth:    true,
			PopUpShow: true,
		}
		tokenString, err := auth.CreateToken(token)
		if err != nil {
			log.Println(err.Error())
			return utils.Render(c, components.ErrorAlert("Bad Credentials", "login"), templ.WithStatus(http.StatusBadRequest))
		}

		c.Cookie(&fiber.Cookie{
			Name:     "session",
			Value:    tokenString,
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})
		c.Response().Header.Set("HX-Redirect", "/")
		return c.SendStatus(fiber.StatusOK)

	} else {
		return utils.Render(c, components.ErrorAlert("Invalid password", "login"), templ.WithStatus(http.StatusBadRequest))
	}
}
