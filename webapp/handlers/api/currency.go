package api

import (
	"cr-today-2/webapp/models"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func Currency(c *fiber.Ctx) error {
	count := 1
	var currencyList []fiber.Map

	for index, currency := range models.Currencies {
		name := fiberi18n.MustLocalize(c, &i18n.LocalizeConfig{
			DefaultMessage: models.Messages[index],
			PluralCount:    count,
			TemplateData: map[string]interface{}{
				"count": count,
			},
		})

		currencyList = append(currencyList, fiber.Map{
			"Name":   name,
			"Code":   currency.Code,
			"Symbol": currency.Symbol,
		})
	}

	return c.JSON(currencyList)
}
