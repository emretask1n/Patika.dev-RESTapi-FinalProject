package routes

import (
	"REST_API/database"
	"REST_API/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

var GivenAmount int
var TotalAmount int
var result []string

func Hello(c *fiber.Ctx) error {
	return c.SendString("fiber")
}

func AddItemToCart(c *fiber.Ctx) error {
	productId, _ := strconv.Atoi(c.Params("product_id"))
	userId, _ := strconv.Atoi(c.Params("user_id"))
	quantity, _ := strconv.Atoi(c.Params("quantity"))

	shoppingCart := models.ShoppingCart{ProductID: productId, UserID: userId, Quantity: quantity}

	database.Instance.Select("ProductID", "UserID", "Quantity").Create(&shoppingCart)

	return c.Status(200).JSON(shoppingCart)

}

func ShowCart(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Params("id"))

	var VATTypes int
	database.Instance.Raw("SELECT COUNT( DISTINCT vat ) FROM products").Scan(&VATTypes)

	Discount := map[int]int{
		1:  1,
		8:  1,
		18: 1}

	var MonthlySpending int
	database.Instance.Raw("select sum(total_price) from placed_orders where created_at > current_date - interval 30 day and user_id = ?", userId).Scan(&MonthlySpending)
	if MonthlySpending < GivenAmount {
	} else {
		for l := 0; l < VATTypes; l++ {
			Discount[l] = 10
		}
	}

	var OrderCountForDiscount int
	database.Instance.Raw("SELECT COUNT( user_id ) FROM placed_orders where user_id = ? and total_price > ?", userId, GivenAmount).Scan(&OrderCountForDiscount)

	if OrderCountForDiscount%4 == 3 {
		Discount[8] = 10
		Discount[18] = 15
	}

	// GivenAmountInAMonth assigned
	GivenAmount = 100000

	var ids []int
	var prices []int
	var vats []int

	database.Instance.Raw("Select product_id FROM shopping_carts where user_id = ?", userId).Scan(&ids)

	for i := 0; i < len(ids); i++ {

		var quantity int
		database.Instance.Raw("Select product_id FROM shopping_carts where product_id = ?", ids[i]).Scan(&quantity)

		var price int
		database.Instance.Raw("Select price FROM products where id = ?", ids[i]).Scan(&price)

		var vat int

		var name string
		database.Instance.Raw("Select name FROM products where id = ?", ids[i]).Scan(&name)
		result = append(result, "Product Name: "+name+" Quantity: "+strconv.Itoa(quantity))

		if quantity > 3 && Discount[vat] < 10 {
			totalPrice := (price * 3) + (price * (quantity - 3) * 92 % 100)
			prices = append(prices, totalPrice)
			database.Instance.Raw("Select vat FROM products where id = ?", ids[i]).Scan(&vat)
			vats = append(vats, totalPrice*vat/100)
		} else {
			totalPrice := price * quantity * (100 - Discount[vat]) * 100
			prices = append(prices, totalPrice)
			database.Instance.Raw("Select vat FROM products where id = ?", ids[i]).Scan(&vat)
			vats = append(vats, totalPrice*vat/100)
		}
	}

	TotalAmount = sum(prices)
	totalVATS := sum(vats)

	if len(result) == 0 {
		return c.SendString("Cart is Empty")
	} else {
		result = append(result, "Total Price: "+strconv.Itoa(TotalAmount))
		result = append(result, "Total VAT: "+strconv.Itoa(totalVATS))
		return c.Status(200).JSON(result)
	}
}

func CompleteOrder(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Params("id"))

	CompletedOrders := models.PlacedOrders{UserID: userId, TotalPrice: TotalAmount, CreatedAt: time.Now()}

	database.Instance.Select("UserID", "TotalPrice", "CreatedAt").Create(&CompletedOrders)

	database.Instance.Exec("DELETE FROM shopping_carts where user_id = ?", userId)
	result = nil
	return c.SendString("Order Completed!")
}

func GetAllProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	database.Instance.Find(&products)

	return c.Status(200).JSON(products)
}

func Delete(c *fiber.Ctx) error {
	ShoppingCart := []models.ShoppingCart{}
	IdToBeDeleted, _ := strconv.Atoi(c.Params("id"))

	database.Instance.Delete(&ShoppingCart, IdToBeDeleted)

	return c.Status(200).JSON("deleted")
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
