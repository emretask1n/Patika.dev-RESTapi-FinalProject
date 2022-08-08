package GivenAmountRouter

import (
	GivenAmountModel "REST_API/GivenAmount/model"
	GivenAmountRepository "REST_API/GivenAmount/repository"
	"REST_API/database"
	"github.com/gofiber/fiber/v2"
)

//GetGivenAmount shows "GivenAmount" for user
func GetGivenAmount(c *fiber.Ctx) error {
	var givenAmount []GivenAmountModel.GivenAmounts
	database.Instance.Find(&givenAmount)
	return c.Status(200).JSON(givenAmount)
}

//SetGivenAmount changes GivenAmount with a new one
func SetGivenAmount(c *fiber.Ctx) error {
	GivenAmountRepository.SetNewGivenAmountAndDeleteOldOne(c)
	return c.SendString("New Given Amount is " + c.Params("amount"))
}
