package GivenAmountRouter

import (
	GivenAmountService "REST_API/GivenAmount/service"
	"github.com/gofiber/fiber/v2"
)

//GetGivenAmount shows "GivenAmount" for user
func GetGivenAmount(c *fiber.Ctx) error {
	return GivenAmountService.ShowCurrentGivenAmount(c)
}

//SetGivenAmount changes GivenAmount with a new one
func SetGivenAmount(c *fiber.Ctx) error {
	return GivenAmountService.SetNewGivenAmountAndDeleteOldOne(c)
}
