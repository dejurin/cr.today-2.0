package main

import (
	"cr-today-2/webapp/handlers"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
	"golang.org/x/text/language"
)

func main() {

	engine := jet.New("views", ".jet.html")
	engine.Reload(true)
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views:             engine,
		PassLocalsToViews: true,
		ViewsLayout:       "layouts/main",
	})

	app.Use(
		fiberi18n.New(&fiberi18n.Config{
			RootPath:         "./localize",
			FormatBundleFile: "toml",
			UnmarshalFunc:    toml.Unmarshal,
			AcceptLanguages:  []language.Tag{language.Ukrainian, language.English},
			DefaultLanguage:  language.English,
		}),
	)

	// Handlers
	app.Get("/", handlers.Home)

	// Static assets
	app.Static("/static", "./static")

	log.Fatal(app.Listen(":3000"))
}
