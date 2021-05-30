package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mowamed/go-admin/util"
)

func IsAuthenticated(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})

	}

	return c.Next()

}
