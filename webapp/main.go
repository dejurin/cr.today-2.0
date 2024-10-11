package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug/v2"
	"golang.org/x/text/language"
)

func localizeOrDefault(c *fiber.Ctx, key, defaultValue string) string {
	localized, err := fiberi18n.Localize(c, key)
	if err != nil {
		return defaultValue
	}
	return localized
}

func main() {
	engine := pug.New("./webapp/views", ".pug")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	i18nMiddleware := fiberi18n.New(&fiberi18n.Config{
		RootPath:         "./webapp/localize",
		AcceptLanguages:  []language.Tag{language.Ukrainian, language.English},
		DefaultLanguage:  language.English,
		FormatBundleFile: "toml",
		UnmarshalFunc:    toml.Unmarshal,
	})

	app.Use(i18nMiddleware)
	app.Use(func(c *fiber.Ctx) error {
		t_darkModeSwitch := localizeOrDefault(c, "darkModeSwitch", "Dark Mode")
		log.Println(t_darkModeSwitch)
		c.Locals("t_darkMode", t_darkModeSwitch)
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		lang := c.Query("lang", "en")

		return c.Render("index", fiber.Map{
			"Title": "Test",
			"lang":  lang,
		})
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		localize, err := fiberi18n.Localize(c, "hello")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString(localize)
	})

	app.Static("/static", "./webapp/static")

	log.Fatal(app.Listen(":3000"))
}
