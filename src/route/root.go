package route

import (

	"github.com/gofiber/fiber/v2"
)

func RootHandler(c *fiber.Ctx) error {
	return c.Render("index", nil)
}
