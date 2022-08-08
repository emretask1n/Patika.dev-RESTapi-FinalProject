package ShoppingCartModel

type PriceCalculationDTO struct {
	Quantity int
	Discount int
	Price    int
	Vat      int
	Prices   []int
	Vats     []int
}

type DiscountCalculatorDTO struct {
	Discount              map[int]int
	MonthlySpending       int
	OrderCountForDiscount int
	GivenAmount           int
}
