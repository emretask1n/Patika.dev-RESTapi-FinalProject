package ProductService

import (
	ProductRepository "REST_API/Product/repository"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	products := ProductRepository.GetProductTable()

	return c.Status(200).JSON(products)
}
