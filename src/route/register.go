package route

import (
	"fmt"

	"github.com/JesusKian/URL-Shortner/src/config"
	"github.com/JesusKian/URL-Shortner/src/structure"
	"github.com/gofiber/fiber/v2"
)

var (
	err      error           = nil
	Channels                 = make(chan string, 3)
	Data     *structure.Data = &structure.Data{
		Title: "Default-Title",
		Url:   "Default-Title",
	}
)

func RegisterHandlerPost(c *fiber.Ctx) error {
	if err = c.BodyParser(Data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": string(err.Error()),
		})
	}

	_, err = config.Database.Query(`
	INSERT INTO data
	(Title, URL, ID)
	VALUES
	(?, ?, ?)`, Data.Title, Data.Url, config.Generator())

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
