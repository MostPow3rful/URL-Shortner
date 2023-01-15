package route

import (
	"github.com/gofiber/fiber/v2"
)

var (
	whiteList []string = []string{
		"/",
		"/register",
		"/reslt",
	}
)

func RootHandler(c *fiber.Ctx) error {
	return c.Render("index", nil)
}

func RoothandlerMW(c *fiber.Ctx) error {
	var path string = string(c.Request().URI().Path())[0:4]

	if (len(path) >= 4) && (path[0:4] == "/go/") {
		c.Next()
	}

	for _, whitePath := range whiteList {
		if path == whitePath {
			c.Next()
		}

	}

	return fiber.NewError(404, "Invalid Path !")
}