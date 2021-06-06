package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mowamed/go-admin/database"
	"github.com/mowamed/go-admin/middlewares"
	"github.com/mowamed/go-admin/models"
	"strconv"
)

func AllProducts(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Product{}, page))
}

func CreateProduct(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&product)

	return nil
}
