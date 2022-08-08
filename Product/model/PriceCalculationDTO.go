package ProductModel

type PriceCalculationDTO struct {
	Quantity int
	Discount int
	Price    int
	Vat      int
	Prices   []int
	Vats     []int
}
