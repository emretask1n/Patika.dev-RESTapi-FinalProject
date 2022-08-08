package common

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func DiscountCalculator(Discount map[int]int, MonthlySpending int, VATTypes int, OrderCountForDiscount int, GivenAmount int) map[int]int {
	Vats := []int{1, 8, 18}
	if MonthlySpending < GivenAmount {
	} else if MonthlySpending == 0 && GivenAmount == 0 {
	} else {
		for l := 0; l < VATTypes; l++ {
			Discount[Vats[l]] = 10
		}
	}

	if OrderCountForDiscount%4 == 3 {
		Discount[8] = 10
		Discount[18] = 15
	}

	return Discount
}
