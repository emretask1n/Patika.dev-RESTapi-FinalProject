package GivenAmountRouter

import (
	GivenAmountModel "REST_API/GivenAmount/model"
	"REST_API/database"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

//GetGivenAmount shows "GivenAmount" for user
func GetGivenAmount(c *fiber.Ctx) error {
	var givenAmount []GivenAmountModel.GivenAmounts
	database.Instance.Find(&givenAmount)
	return c.Status(200).JSON(givenAmount)
}

//SetGivenAmount changes GivenAmount with a new one
func SetGivenAmount(c *fiber.Ctx) error {
	database.Instance.Exec("DELETE from given_amounts")

	newAmount, _ := strconv.Atoi(c.Params("amount"))
	var givenAmount = GivenAmountModel.GivenAmounts{
		GivenAmount: newAmount,
	}

	database.Instance.Select("given_amount").Create(&givenAmount)
	return c.SendString("New Given Amount is " + c.Params("amount"))
}
