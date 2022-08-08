package GivenAmountRepository

import (
	GivenAmountModel "REST_API/GivenAmount/model"
	"REST_API/database"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetGivenAmount() int {
	var GivenAmount int
	database.Instance.Raw("Select given_amount FROM given_amounts").Scan(&GivenAmount)
	return GivenAmount
}

func SetNewGivenAmountAndDeleteOldOne(c *fiber.Ctx) {
	database.Instance.Exec("DELETE from given_amounts")

	newAmount, _ := strconv.Atoi(c.Params("amount"))
	var givenAmount = GivenAmountModel.GivenAmounts{
		GivenAmount: newAmount,
	}

	database.Instance.Select("given_amount").Create(&givenAmount)
}
