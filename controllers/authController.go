package controllers

import (
	"aims/db"
	"aims/models"
	"aims/utilities"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/sujit-baniya/flash"
	"golang.org/x/crypto/bcrypt"
)

func RegisterPostController(c *fiber.Ctx) error {
	formUsername := c.FormValue("username")
	formEmail := c.FormValue("email")
	formPassword := c.FormValue("password")

	password, err := bcrypt.GenerateFromPassword([]byte(formPassword), 14)
	if err != nil {
		return flash.WithError(c, fiber.Map{
			"content": err,
		}).Redirect("/student/register")
	}

	user := models.User{
		Username: formUsername,
		Email:    formEmail,
		Password: string(password),
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return flash.WithError(c, fiber.Map{"content": err}).Redirect("/student/register")
	}

	return c.Redirect("/student/login")
}

func VerifyPostController(c *fiber.Ctx) error {

	formUsername := c.FormValue("username")
	formPassword := c.FormValue("password")

	var user models.User

	req := db.DB.Where("username=?", formUsername).First(&user)
	if err := req.Error; err != nil {
		return err
	}

	if user.Id == 0 {
		c.Status(fiber.StatusUnauthorized)
		return flash.WithError(c, fiber.Map{"content": "User not found"}).Redirect("/student/login")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formPassword)); err != nil {

		c.Status(fiber.StatusUnauthorized)
		return flash.WithError(c, fiber.Map{"content": "Invalid Credentials"}).Redirect("/student/login")
	}

	token, exp, err := utilities.CreateJWTToken(user)
	if err != nil {
		return err
	}

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return flash.WithError(c, fiber.Map{"content": "Server Error"}).Redirect("/student/login")
	}

	cookie := fiber.Cookie{
		Name:     "auth",
		Value:    token,
		Expires:  time.Unix(exp, 0),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Redirect("/applicant")
}

func LogoutController(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "auth",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.Redirect("/student/login")
}
