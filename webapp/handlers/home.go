package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Render("main/index", fiber.Map{
		"Title": "Welcome to the Currency App",
	})
}
