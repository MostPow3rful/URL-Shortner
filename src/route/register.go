package route

import (
	"fmt"

	"github.com/JesusKian/URL-Shortner/src/config"
	"github.com/JesusKian/URL-Shortner/src/sql"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandlerPost(c *fiber.Ctx) error {
	if err = c.BodyParser(Data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": string(err.Error()),
		})
	}

	uniqueID := config.Generator()
	_, err = sql.Database.Query(`
	INSERT INTO data
	(Title, URL, ID)
	VALUES
	(?, ?, ?)`, Data.Title, Data.Url, uniqueID)

	if err != nil {
		config.SetLog("E", "route.RegisterHandlerMw() -> Couldn't Add Data In Database")
		config.SetLog("D", err.Error())
	}

	config.SetLog("I", fmt.Sprintf("New Data : Title=%s , URL=%s , ID=%s", Data.Title, Data.Url, uniqueID))
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"title": Data.Title,
		"url":   Data.Url,
	})
}
