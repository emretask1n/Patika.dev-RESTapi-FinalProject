package common

//SumOfIntSlice function is to get the sum of int slices.
func SumOfIntSlice(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

//DiscountCalculator calculates discounts based on VAT values, Business Rules A and C are handled here
func DiscountCalculator(Discount map[int]int, MonthlySpending int, VATTypes int, OrderCountForDiscount int, GivenAmount int) map[int]int {
	Vats := []int{1, 8, 18}
	if MonthlySpending > GivenAmount {
		for l := 0; l < VATTypes; l++ {
			Discount[Vats[l]] = 10
		}
	} else {
	}

	if OrderCountForDiscount%4 == 3 {
		Discount[8] = 10
		Discount[18] = 15
	}

	return Discount
}
