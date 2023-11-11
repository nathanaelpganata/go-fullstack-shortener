package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nathanaelpganata/go-fullstack-shortener/internal/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.IndexPage)
	app.Get("/shorten", controllers.ShortenPage)
	app.Post("/shorten", controllers.Shorten)
	app.Get("/:shortUrl", controllers.Redirect)
}
