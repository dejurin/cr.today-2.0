package main

import (
	"log"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	engine := pug.New("./webapp/views", ".pug")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(
		fiberi18n.New(&fiberi18n.Config{
			RootPath:         "./webapp/localize",
			FormatBundleFile: "toml",
			AcceptLanguages:  []language.Tag{language.Ukrainian, language.English},
			DefaultLanguage:  language.English,
		}),
	)

	app.Static("/static", "./webapp/static")

	app.Get("/", func(c *fiber.Ctx) error {
		lang := c.Query("lang", "en")

		msg, err := fiberi18n.Localize(c, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "Hello",
				One:   "Hello_, {{.Name}}!",
				Other: "Hello_, {{.Name}}s!",
			},
			TemplateData: map[string]interface{}{
				"Name": "Test",
			},
			PluralCount: 1,
		})

		if err != nil {
			return err
		}

		return c.Render("index", fiber.Map{
			"Title": "Test",
			"msg":   msg,
			"lang":  lang,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
