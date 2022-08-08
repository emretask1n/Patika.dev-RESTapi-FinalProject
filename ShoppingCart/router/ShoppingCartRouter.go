package ShoppingCartRouter

import (
	ShoppingCartService "REST_API/ShoppingCart/service"
	"github.com/gofiber/fiber/v2"
)

func AddItemToCart(c *fiber.Ctx) error {
	return ShoppingCartService.AddItemToCartByProductIdUserIdAndQuantity(c)
}

func ShowCart(c *fiber.Ctx) error {
	return ShoppingCartService.ShowAllCart(c)
}

func CompleteOrder(c *fiber.Ctx) error {
	return ShoppingCartService.CompleteOrderByUserId(c)
}

func DeleteItemFromCart(c *fiber.Ctx) error {
	return ShoppingCartService.DeleteItemFromCartByProductIdAndUserId(c)
}
