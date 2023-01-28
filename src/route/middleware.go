package route

import (
	"github.com/gofiber/fiber/v2"

	"fmt"
)

var (
	whiteList map[string][]string = map[string][]string{
		"GET": {
			"/",
			"/result",
		},
		"POST": {
			"/shortner",
		},
	}
)

func PathHandler(c *fiber.Ctx) error {
	var path string = string(c.Request().URI().Path())

	if (len(path) >= 4) && (path[0:4] == "/go/") {
		return c.Next()
	}

	for allowedMethod, whitePaths := range whiteList {
		for _, whitePath := range whitePaths {
			if path == whitePath {
				if string(c.Request().Header.Method()) == string((allowedMethod)) {
					return c.Next()
				}
				return fiber.NewError(fiber.StatusMethodNotAllowed, fmt.Sprintf("[ERR] : You Must Use [%s] Method", allowedMethod))
			}
		}
	}

	return fiber.NewError(404, fmt.Sprintf("[ERR] : Invalid Path [%s]", path))
}
