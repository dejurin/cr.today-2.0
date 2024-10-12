package handlers

import (
	"cr-today-2/webapp/i18n"
	"cr-today-2/webapp/models"

	"github.com/gofiber/fiber/v2"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

func Currencies(c *fiber.Ctx) error {
	lang := c.Query("lang", "en")
	localizer := goi18n.NewLocalizer(i18n.Bundle, lang)

	var localizedCurrencies []fiber.Map

	for _, currency := range models.Currencies {
		name := currency.LocalizeName(localizer, 1)

		localizedCurrencies = append(localizedCurrencies, fiber.Map{
			"Code":   currency.Code,
			"Name":   name,
			"Symbol": currency.Symbol,
		})
	}

	return c.Render("main/index", fiber.Map{
		"Currencies": localizedCurrencies,
	})
}
