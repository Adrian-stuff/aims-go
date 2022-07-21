package main

import (
	"aims/db"
	"aims/routes"
	"aims/utilities"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	jwtware "github.com/gofiber/jwt/v3"

	"github.com/gofiber/template/html"
)

func main() {
	db.Connect()
	engine := html.New("./templates", ".html")

	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false

	// Debug will print each template that is parsed, good for debugging
	engine.Debug(true) // Optional. Default: false

	// After you created your engine, you can pass it to Fiber's Views Engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	// app.Use(csrf.New())
	app.Use(logger.New())
	public := app.Group("/")
	routes.PublicRoutes(public)
	private := app.Group("/")
	private.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(utilities.SecretKey),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Redirect("/")
		}, TokenLookup: "cookie:auth",
	}))
	routes.PrivateRoutes(private)

	log.Fatal(app.Listen(":3000"))
}
