package main

import (
	"cr-today-2/webapp/handlers"
	"cr-today-2/webapp/handlers/api"
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	minifier "github.com/beyer-stefan/gofiber-minifier"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
	"golang.org/x/text/language"
)

func main() {

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	engine := jet.New("views", ".jet.html")
	engine.Reload(getEnvAsBool("TEMPLATE_RELOAD", false))
	engine.Debug(getEnvAsBool("TEMPLATE_DEBUG", false))

	AddGlobalVariables(engine)

	app := fiber.New(fiber.Config{
		Views:             engine,
		PassLocalsToViews: true,
		ServerHeader:      "CR.Today",
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

	app.Use(minifier.New(minifier.Config{
		MinifyHTML: true,
	}))

	// Handlers
	app.Get("/", handlers.Home)

	// Api
	app.Get("/api/currency.json", api.Currency)

	// Static assets
	app.Static("/static", "./static")

	// Check
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString(os.Getenv("APP_ENV"))
	})

	// Up
	log.Fatal(app.Listen(":3000"))
}

func getEnvAsBool(key string, defaultVal bool) bool {
	valStr := os.Getenv(key)
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}
