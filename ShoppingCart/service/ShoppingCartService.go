package ShoppingCartService

import (
	GivenAmountRepository "REST_API/GivenAmount/repository"
	PlacedOrderRepository "REST_API/PlacedOrders/repository"
	ProductRepository "REST_API/Product/repository"
	ShoppingCartModel "REST_API/ShoppingCart/model"
	ShoppingCartRepository "REST_API/ShoppingCart/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

//TotalAmount is to Calculate sum of prices
var TotalAmount int

//result is to show the names of products, quantities and total prices and vats of Cart
var result []string

//SumOfIntSlice function is to get the sum of int slices.
func SumOfIntSlice(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

//DiscountCalculator calculates discounts based on VAT values, Business Rules A and C are handled here
func DiscountCalculator(dto ShoppingCartModel.DiscountCalculatorDTO) map[int]int {
	Vats := []int{1, 8, 18}
	if dto.MonthlySpending > dto.GivenAmount {
		for l := 0; l < len(Vats); l++ {
			dto.Discount[Vats[l]] = 10
		}
	} else {
	}
	if dto.OrderCountForDiscount%4 == 3 {
		dto.Discount[8] = 10
		dto.Discount[18] = 15
	}
	return dto.Discount
}

func PriceCalculationForProduct(dto ShoppingCartModel.PriceCalculationDTO) (prices, vats []int) {
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

func ShowAllCart(c *fiber.Ctx) error {
	TotalAmount = 0
	userId, _ := strconv.Atoi(c.Params("user_id"))
	result = nil

	Discount := map[int]int{
		1:  0,
		8:  0,
		18: 0,
	}

	MonthlySpending := PlacedOrderRepository.GetMonthlySpendingByUserId(userId)
	GivenAmount := GivenAmountRepository.GetGivenAmount()
	OrderCountForDiscount := PlacedOrderRepository.GetOrderCountAndSpendingForDiscount(userId, GivenAmount)
	productIds := ShoppingCartRepository.GetIdsOfProductsFromCartByUserId(userId)

	var prices []int
	var vats []int

	discountCalculatorDTO := ShoppingCartModel.DiscountCalculatorDTO{
		Discount:              Discount,
		MonthlySpending:       MonthlySpending,
		OrderCountForDiscount: OrderCountForDiscount,
		GivenAmount:           GivenAmount,
	}

	Discount = DiscountCalculator(discountCalculatorDTO)

	for i := 0; i < len(productIds); i++ {
		VatTypes := ShoppingCartRepository.GetVATTypes()
		quantity := ShoppingCartRepository.GetQuantityByProductId(productIds[i])
		price := ProductRepository.GetPriceByProductId(productIds[i])
		vat := ProductRepository.GetVATByProductId(productIds[i])
		name := ProductRepository.GetProductNameByProductId(productIds[i])

		result = append(result, "Product Name: "+name+" Quantity: "+strconv.Itoa(quantity))

		priceCalculationDTO := ShoppingCartModel.PriceCalculationDTO{
			Quantity: quantity,
			Discount: Discount[VatTypes[i]],
			Price:    price,
			Vat:      vat,
			Prices:   prices,
			Vats:     vats,
		}
		prices, vats = PriceCalculationForProduct(priceCalculationDTO)
	}

	TotalAmount = SumOfIntSlice(prices)
	totalVATS := SumOfIntSlice(vats)

	if len(result) != 0 {
		result = append(result, "Total Price: "+strconv.Itoa(TotalAmount))
		result = append(result, "Total VAT: "+strconv.Itoa(totalVATS))
		return c.Status(200).JSON(result)
	} else {
		return c.SendString("Cart is Empty")
	}
}
