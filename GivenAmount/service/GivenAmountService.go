package GivenAmountService

import (
	GivenAmountRepository "REST_API/GivenAmount/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func ShowCurrentGivenAmount(c *fiber.Ctx) error {
	givenAmount := GivenAmountRepository.GetGivenAmountAsTable()
	return c.Status(200).JSON(givenAmount)
}

func SetNewGivenAmountAndDeleteOldOne(c *fiber.Ctx) error {
	GivenAmountRepository.DeleteOldGivenAmount()
	newAmount, _ := strconv.Atoi(c.Params("amount"))
	GivenAmountRepository.SetNewGivenAmount(newAmount)
	return c.SendString("New Given Amount is " + c.Params("amount"))
}
