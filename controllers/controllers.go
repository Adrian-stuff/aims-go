package controllers

import (
	"aims/db"
	"aims/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sujit-baniya/flash"
)

func RegisterGetController(c *fiber.Ctx) error {
	return c.Render("views/register/index", fiber.Map{
		"title":   "Register",
		"message": flash.Get(c),
	}, "layouts/login/base")
}
func LoginController(c *fiber.Ctx) error {
	return c.Render("views/login/index", fiber.Map{
		"title":   "login",
		"message": flash.Get(c),
	}, "layouts/login/base")
}

func ErrorController(c *fiber.Ctx) error {
	return flash.WithError(c, fiber.Map{"content": "Invalid Creds"}).Redirect("/student/login")
}

func ApplicantController(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var userModel models.User

	db.DB.Where("id=?", claims["user_id"]).First(&userModel)

	return c.Render("views/applicant/index", fiber.Map{"user": userModel}, "layouts/applicant/base")
}

func IndexController(c *fiber.Ctx) error {
	return c.Redirect("/student/login")
}
