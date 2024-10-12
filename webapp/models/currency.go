package models

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Модель валюты
type Currency struct {
	Code   string
	Symbol string
}

// Список валют
var Currencies = []Currency{
	{
		Code:   "USD",
		Symbol: "$",
	},
	{
		Code:   "EUR",
		Symbol: "€",
	},
}

// Статические строки для локализации
var Messages = []*i18n.Message{
	{
		ID:    "Currency.USD",
		One:   "US Dollar",             // Единственное число
		Other: "{{.count}} US Dollars", // Множественное число
	},
	{
		ID:    "Currency.EUR",
		One:   "Euro",             // Единственное число
		Other: "{{.count}} Euros", // Множественное число
	},
}

// Метод для получения локализованного названия валюты
func (c *Currency) LocalizeName(localizer *i18n.Localizer, count int) string {
	// Локализация через go-i18n с заранее определенными строками
	localizedName := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:   "Currency." + c.Code,
		PluralCount: count,
		TemplateData: map[string]interface{}{
			"count": count,
		},
	})
	return localizedName
}
