package models

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Currency struct {
	Code   string
	Symbol string
}

var Currencies = []Currency{
	{
		Code:   "USD",
		Symbol: "$",
	},
	{
		Code:   "EUR",
		Symbol: "€",
	},
	{
		Code:   "UAH",
		Symbol: "₴",
	},
}

var Messages = []*i18n.Message{
	{
		ID:    "Currency.USD",
		One:   "US Dollar",
		Other: "{{.count}} US Dollars",
	},
	{
		ID:    "Currency.EUR",
		One:   "Euro",
		Other: "{{.count}} Euros",
	},
	{
		ID:    "Currency.UAH",
		One:   "Ukrainian Hryvnia",
		Other: "{{.count}} Ukrainian Hryvnias",
	},
}
