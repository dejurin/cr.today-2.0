package handlers

import (
	"cr-today-2/webapp/models"

	"encoding/json"

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

	currencyListJSON, err := json.Marshal(currencyList)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Render("main/index", fiber.Map{
		"Title":            mainTitle,
		"CurrencyList":     currencyList,
		"CurrencyListJSON": currencyListJSON,
	})
}
