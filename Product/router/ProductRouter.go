package ProductRouter

import (
	"REST_API/Product/model"
	"REST_API/database"
	"github.com/gofiber/fiber/v2"
)

//GetAllProducts lists all products
func GetAllProducts(c *fiber.Ctx) error {
	var products []ProductModel.Product

	database.Instance.Find(&products)

	return c.Status(200).JSON(products)
}
