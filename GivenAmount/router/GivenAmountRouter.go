package GivenAmountRouter

import (
	GivenAmountService "REST_API/GivenAmount/service"
	"github.com/gofiber/fiber/v2"
)

func GetGivenAmount(c *fiber.Ctx) error {
	return GivenAmountService.ShowCurrentGivenAmount(c)
}

func SetGivenAmount(c *fiber.Ctx) error {
	return GivenAmountService.SetNewGivenAmountAndDeleteOldOne(c)
}
