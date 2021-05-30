package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mowamed/go-admin/controllers"
	"github.com/mowamed/go-admin/middlewares"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user", controllers.User)
}
