package route

import (
	"github.com/JesusKian/URL-Shortner/src/config"
	"github.com/gofiber/fiber/v2"
)

var (
	key string = ""
)

func GoHandlerGet(c *fiber.Ctx) error {
	key = c.Params("id")

	result, err := config.Database.Query("SELECT URL from data where ID=?", key)
	if err != nil {
		config.SetLog("E", "route.GoHandlerGet() -> Couldn't Get URL")
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&url)

		if err != nil {
			config.SetLog("E", "route.GoHandlerGet() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}
	}

	if url == "" {
		return c.Status(fiber.StatusNotFound).Render("error", fiber.Map{
			"Message": "Entered URL isn't Valid !",
		})
	}

	return c.Redirect(url)
}
