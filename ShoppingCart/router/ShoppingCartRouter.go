package ShoppingCartRouter

import (
	PlacedOrderModel "REST_API/PlacedOrders/model"
	"REST_API/ShoppingCart/model"
	ShoppingCartRepository "REST_API/ShoppingCart/repository"
	ShoppingCartService "REST_API/ShoppingCart/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

//TotalAmount is to Calculate sum of prices
var TotalAmount int

//result is to show the names of products, quantities and total prices and vats of Cart
var result []string

//AddItemToCart adds users products to the basket and the total of the basket
//changes accordingly.
func AddItemToCart(c *fiber.Ctx) error {
	productId, _ := strconv.Atoi(c.Params("product_id"))
	userId, _ := strconv.Atoi(c.Params("user_id"))
	quantity, _ := strconv.Atoi(c.Params("quantity"))

	shoppingCart := model.ShoppingCart{ProductID: productId, UserID: userId, Quantity: quantity}

	ShoppingCartRepository.InsertIntoCart(shoppingCart)

	return c.Status(200).JSON(shoppingCart)

}

//ShowCart lists the products users have added to their cart and total price and
//VAT of the cart
func ShowCart(c *fiber.Ctx) error {

	return ShoppingCartService.ShowAllCart(c)
	/*
		TotalAmount = 0
		userId, _ := strconv.Atoi(c.Params("user_id"))
		result = nil

		Discount := map[int]int{
			1:  0,
			8:  0,
			18: 0,
		}

		VATTypes := ProductRepository.GetVATTypes()
		MonthlySpending := PlacedOrderRepository.GetMonthlySpendingByUserId(userId)
		GivenAmount := GivenAmountRepository.GetGivenAmount()
		OrderCountForDiscount := PlacedOrderRepository.GetOrderCountAndSpendingForDiscount(userId, GivenAmount)
		productIds := ShoppingCartRepository.GetIdsOfProductsFromCartByUserId(userId)

		var prices []int
		var vats []int

		Discount = ShoppingCartService.DiscountCalculator(Discount, MonthlySpending, VATTypes, OrderCountForDiscount, GivenAmount)

		for i := 0; i < len(productIds); i++ {

			quantity := ShoppingCartRepository.GetQuantityByProductId(productIds[i])
			price := ProductRepository.GetPriceByProductId(productIds[i])
			vat := ProductRepository.GetVATByProductId(productIds[i])
			name := ProductRepository.GetProductNameByProductId(productIds[i])

			result = append(result, "Product Name: "+name+" Quantity: "+strconv.Itoa(quantity))

			priceCalculationDTO := ProductModel.PriceCalculationDTO{
				Quantity: quantity,
				Discount: Discount[i],
				Price:    price,
				Vat:      vat,
				Prices:   prices,
				Vats:     vats,
			}

			prices, vats = ShoppingCartService.PriceCalculationForProduct(priceCalculationDTO)

			// Business rules B and D are handled here
		}

		//Calculation for TotalAmount and VATs, the reason why TotalAmount was
		//not created here is CompleteOrder needs its reference.
		TotalAmount = ShoppingCartService.SumOfIntSlice(prices)
		totalVATS := ShoppingCartService.SumOfIntSlice(vats)

		return CheckCartStatus(c, totalVATS)


	*/
}

/*
func CheckCartStatus(c *fiber.Ctx, totalVATS int) error {
	if len(result) != 0 {
		result = append(result, "Total Price: "+strconv.Itoa(TotalAmount))
		result = append(result, "Total VAT: "+strconv.Itoa(totalVATS))
		return c.Status(200).JSON(result)
	} else {
		return c.SendString("Cart is Empty")
	}
}

*/

//CompleteOrder creates an order with the products users add to their cart
func CompleteOrder(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Params("user_id"))

	CompletedOrders := PlacedOrderModel.PlacedOrders{UserID: userId, TotalPrice: TotalAmount, CreatedAt: time.Now()}

	ShoppingCartRepository.CompleteOrders(CompletedOrders)

	ShoppingCartRepository.DeleteCartById(userId)
	return c.SendString("Order Completed!")
}

//DeleteItemFromCart removes products of users
func DeleteItemFromCart(c *fiber.Ctx) error {
	IdToBeDeleted, _ := strconv.Atoi(c.Params("product_id"))
	userId, _ := strconv.Atoi(c.Params("user_id"))

	ShoppingCartRepository.DeleteProductFromCartByUProductIdAndUserId(IdToBeDeleted, userId)

	return c.Status(200).JSON("deleted")
}
