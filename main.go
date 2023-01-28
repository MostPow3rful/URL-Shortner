package main

import (
	"github.com/JesusKian/URL-Shortner/src/route"
	"github.com/JesusKian/URL-Shortner/src/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"

	"log"
)

func init() {
	sql.ConnectToSqlDatabase()
	sql.DatabaseStatus()
}

func main() {
	const VERSION string = "2.0"

	var (
		engine *html.Engine = html.New("./static/html", ".html")
		app    *fiber.App   = fiber.New(fiber.Config{
			Views: engine,
		})
	)

	app.Static("/static", "./static")

	// Registered Middlewares
	app.Use(logger.New())
	app.Use(etag.New())
	app.Use(route.PathHandler)

	// Paths With GET Method
	app.Get("/", route.RootHandler)
	app.Get("/result", route.ResultHandlerGet)
	app.Get("/go/:id<str>", route.GoHandlerGet)

	// Paths With POST Method
	app.Post("/shortner", route.ShortnerHandlerPost)

	log.Fatal(app.Listen(":8569"))
}
