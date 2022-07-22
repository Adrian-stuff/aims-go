package routes

import (
	"aims/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(g fiber.Router) {
	g.Get("/", controllers.IndexController)
	g.Get("/error", controllers.ErrorController)
	g.Get("/student/login", controllers.LoginController)
	g.Get("/student/register", controllers.RegisterGetController)
	g.Post("/student/register", controllers.RegisterPostController)

	g.Post("/student/verify", controllers.VerifyPostController)
}

func PrivateRoutes(g fiber.Router) {
	g.Get("/applicant", controllers.ApplicantController)
	g.Get("/student/logout", controllers.LogoutController)
}
