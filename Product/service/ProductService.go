package ProductService

/*
func PriceCalculationForProduct(dto ProductModel.PriceCalculationDTO) (prices, vats []int) {
	if dto.Quantity > 3 && dto.Discount < 10 {
		withoutDiscount := dto.Price * 3
		withDiscount := dto.Price * (dto.Quantity - 3) / 100 * 92
		totalPrice := withoutDiscount + withDiscount
		dto.Prices = append(dto.Prices, totalPrice)
		dto.Vats = append(dto.Vats, totalPrice*dto.Vat/100)
		return dto.Prices, dto.Vats
	} else {
		totalPrice := dto.Price * dto.Quantity * (100 - dto.Discount) / 100
		dto.Prices = append(dto.Prices, totalPrice)
		dto.Vats = append(dto.Vats, totalPrice*dto.Vat/100)
		return dto.Prices, dto.Vats
	}
}
*/
