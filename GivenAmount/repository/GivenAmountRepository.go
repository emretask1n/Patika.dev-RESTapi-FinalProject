package GivenAmountRepository

import "REST_API/database"

func GetGivenAmount() int {
	var GivenAmount int
	database.Instance.Raw("Select given_amount FROM given_amounts").Scan(&GivenAmount)
	return GivenAmount
}
