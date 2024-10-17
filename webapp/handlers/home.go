package handlers

import (
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Home(c *fiber.Ctx) error {

	mainTitle := fiberi18n.MustLocalize(c, &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "mainTitle",
			Other: "CR.Today - Currency Converter",
		},
	})

	return c.Render("main/index", fiber.Map{
		"Title": mainTitle,
	})
}
