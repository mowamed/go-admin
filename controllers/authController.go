package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mowamed/go-admin/models"
)

func Register(c *fiber.Ctx) error {

	user := models.User{
		FirstName: "momo",
		LastName:  "bakus",
	}
	return c.JSON(user)
}
