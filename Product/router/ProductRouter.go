package ProductRouter

import (
	ProductService "REST_API/Product/service"
	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	return ProductService.GetProducts(c)
}
