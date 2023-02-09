package route

import (
	"fmt"

	"github.com/JesusKian/URL-Shortner/src/config"
	"github.com/JesusKian/URL-Shortner/src/sql"
	"github.com/JesusKian/URL-Shortner/src/structure"
	"github.com/JesusKian/URL-Shortner/src/timer"
	"github.com/gofiber/fiber/v2"
)

var (
	err      error           = nil
	Channels                 = make(chan string, 3)
	Data     *structure.Data = &structure.Data{
		Title:    "Default-Title",
		Url:      "Default-URL",
		Expire:   "0",
		UniqueID: "Default-ID",
	}
)

func ShortnerHandlerPost(c *fiber.Ctx) error {
	if err = c.BodyParser(Data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": string(err.Error()),
		})
	}

	Data.UniqueID = config.Generator()
	Data.Expire = timer.SetExpire(Data.Expire)
	_, err = sql.Database.Query(`
	INSERT INTO data
	(Title, URL, ID, Expire)
	VALUES
	(?, ?, ?, ?)`,
		Data.Title, Data.Url, Data.UniqueID, Data.Expire,
	)

	if err != nil {
		config.SetLog("E", "route.ShortnerHandlerMw() -> Couldn't Add Data In Database")
		config.SetLog("D", err.Error())
	}

	config.SetLog("I", fmt.Sprintf("New Data : Title=%s , URL=%s , ID=%s, Expire=%s", Data.Title, Data.Url, Data.UniqueID, Data.Expire))
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"title": Data.Title,
		"url":   Data.Url,
	})
}
