package main

import (
	"log"

	"github.com/JesusKian/URL-Shortner/src/config"
	"github.com/JesusKian/URL-Shortner/src/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config.ConnectToSqlDatabase()
	config.DatabaseStatus()
}

func main() {
	var (
		engine *html.Engine = html.New("./static/html", ".html")
		app    *fiber.App   = fiber.New(fiber.Config{
			Views: engine,
		})
	)

	app.Use(logger.New())
	app.Static("/static", "./static")

	app.Use(route.RoothandlerMW)
	app.Get("/", route.RootHandler)
	app.Post("/register", route.RegisterHandlerPost)
	app.Get("/result", route.ResultHandlerGet)
	app.Get("/go/:id", route.GoHandlerGet)

	log.Fatal(app.Listen(":8569"))
}
