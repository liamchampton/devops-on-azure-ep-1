package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func renderIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":        "DevOps on Azure",
		"Description":  "A series all about DevOps",
		"WelcomeTitle": "Welcome to this 3 part series!",
	}, "layouts/main")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", renderIndex)
}

func main() {
	engine := handlebars.New("./views", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	setupRoutes(app)

	app.Listen(":3000")
}
