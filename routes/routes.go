package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mowamed/go-admin/controllers"
)

func Setup(app *fiber.App) {

	app.Get("/", controllers.Hello)
}
