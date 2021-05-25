package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mowamed/go-admin/database"
	"github.com/mowamed/go-admin/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
