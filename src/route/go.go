package route

import (
	"github.com/JesusKian/URL-Shortner/src/config"
	"github.com/JesusKian/URL-Shortner/src/sql"
	"github.com/JesusKian/URL-Shortner/src/timer"
	"github.com/gofiber/fiber/v2"
)

var (
	key       string = ""
	nowExpire string = ""
)

func GoHandlerGet(c *fiber.Ctx) error {
	key = c.Params("id")

	result, err := sql.Database.Query("SELECT URL,Expire from data where ID=?", key)
	if err != nil {
		config.SetLog("E", "route.GoHandlerGet() -> Couldn't Get URL")
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&url, &nowExpire)

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

	if timer.CheckExpire(nowExpire) == true {
		return c.Redirect(url)
	}

	return c.Status(fiber.StatusNotFound).Render("error", fiber.Map{
		"Message": "Entered URL Have Been Expired !",
	})
}
