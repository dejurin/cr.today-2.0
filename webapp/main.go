package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("localize/active.uk.toml")

	engine := pug.New("./views", ".pug")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {

		lang := c.Query("lang", "en")
		return c.Render("main/index", fiber.Map{
			"Title": "Test",
			"lang":  lang,
		})
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		lang := c.Query("lang", "en")
		accept := c.AcceptsLanguages()
		localizer := i18n.NewLocalizer(bundle, lang, accept)

		helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "HelloPerson",
				Other: "Hello {{.Name}}",
			},
			TemplateData: map[string]string{
				"Name": "YURII",
			},
		})

		return c.Render("exchange/index", fiber.Map{
			"Title": helloPerson,
		})
	})

	app.Static("/static", "./static")

	log.Fatal(app.Listen(":3000"))
}
