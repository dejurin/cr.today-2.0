package main

import (
	"cr-today-2/webapp/handlers"
	"cr-today-2/webapp/i18n"
	"log"

	"github.com/gofiber/template/jet/v2"

	"github.com/gofiber/fiber/v2"
)

func main() {
	i18n.InitBundle()

	engine := jet.New("views", ".jet.html")
	engine.Reload(true)
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views:             engine,
		PassLocalsToViews: true,
		ViewsLayout:       "layouts/main",
	})

	// Handlers
	app.Get("/", handlers.Home)

	// Static assets
	app.Static("/static", "./static")

	log.Fatal(app.Listen(":3000"))
}
