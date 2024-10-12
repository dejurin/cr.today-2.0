package main

import (
	"cr-today-2/webapp/handlers"
	"cr-today-2/webapp/i18n"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug/v2"
)

func main() {
	i18n.InitBundle()

	engine := pug.New("webapp/views", ".pug")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Handlers
	app.Get("/", handlers.Home)
	app.Get("/currencies", handlers.Currencies)

	// Static assets
	app.Static("/static", "./static")

	log.Fatal(app.Listen(":3000"))
}
