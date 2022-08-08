package GivenAmountRepository

import (
	GivenAmountModel "REST_API/GivenAmount/model"
	"REST_API/database"
	"gorm.io/gorm"
)

func GetGivenAmount() int {
	var GivenAmount int
	database.Instance.Raw("Select given_amount FROM given_amounts").Scan(&GivenAmount)
	return GivenAmount
}

func DeleteOldGivenAmount() *gorm.DB {
	return database.Instance.Exec("DELETE from given_amounts")
}

func GetGivenAmountAsTable() []GivenAmountModel.GivenAmounts {
	var givenAmount []GivenAmountModel.GivenAmounts
	database.Instance.Find(&givenAmount)
	return givenAmount
}

func SetNewGivenAmount(newAmount int) {
	var givenAmount = GivenAmountModel.GivenAmounts{
		GivenAmount: newAmount,
	}
	database.Instance.Select("given_amount").Create(&givenAmount)
}
