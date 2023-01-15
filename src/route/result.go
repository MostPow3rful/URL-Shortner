package route

import (
	"fmt"

	"github.com/JesusKian/URL-Shortner/src/config"
	"github.com/gofiber/fiber/v2"
)

var (
	title    string = ""
	url      string = ""
	uniqueID string = ""
)

func getDataFromDB() {
	result, err := config.Database.Query("SELECT * FROM data")
	if err != nil {
		config.SetLog("E", "route.getDataFromDB() -> Couldn't Get data from Database")
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&title, &url, &uniqueID)

		if err != nil {
			config.SetLog("E", "route.getDataFromDB() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}
	}
}

func ResultHandlerGet(c *fiber.Ctx) error {
	getDataFromDB()

	return c.Render("result", fiber.Map{
		"TITLE": title,
		"URL":   url,
		"LINK":  fmt.Sprintf("127.0.0.1:8569/go/%s", uniqueID),
	})
}
