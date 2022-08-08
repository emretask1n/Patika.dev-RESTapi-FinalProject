package routes

import (
	"REST_API/common"
	"REST_API/database"
	"REST_API/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

var TotalAmount int
var result []string

func AddItemToCart(c *fiber.Ctx) error {
	productId, _ := strconv.Atoi(c.Params("product_id"))
	userId, _ := strconv.Atoi(c.Params("user_id"))
	quantity, _ := strconv.Atoi(c.Params("quantity"))

	shoppingCart := models.ShoppingCart{ProductID: productId, UserID: userId, Quantity: quantity}

	database.Instance.Select("ProductID", "UserID", "Quantity").Create(&shoppingCart)

	return c.Status(200).JSON(shoppingCart)

}

func ShowCart(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Params("user_id"))
	result = nil

	Discount := map[int]int{
		1:  1,
		8:  1,
		18: 1,
	}

	var VATTypes int
	database.Instance.Raw("SELECT COUNT( DISTINCT vat ) FROM products").Scan(&VATTypes)

	var MonthlySpending int
	database.Instance.Raw("select sum(total_price) from placed_orders where created_at > current_date - interval 30 day and user_id = ?", userId).Scan(&MonthlySpending)
	fmt.Println(MonthlySpending)

	var GivenAmount int
	database.Instance.Raw("Select given_amount FROM given_amounts").Scan(&GivenAmount)

	var OrderCountForDiscount int
	database.Instance.Raw("SELECT COUNT( user_id ) FROM placed_orders where user_id = ? and total_price > ?", userId, GivenAmount).Scan(&OrderCountForDiscount)

	var ids []int
	database.Instance.Raw("Select product_id FROM shopping_carts where user_id = ?", userId).Scan(&ids)

	var prices []int
	var vats []int

	Discount = common.DiscountCalculator(Discount, MonthlySpending, VATTypes, OrderCountForDiscount, GivenAmount)

	for i := 0; i < len(ids); i++ {

		var quantity int
		database.Instance.Raw("Select quantity FROM shopping_carts where product_id = ?", ids[i]).Scan(&quantity)

		var price int
		database.Instance.Raw("Select price FROM products where id = ?", ids[i]).Scan(&price)

		var vat int
		database.Instance.Raw("Select vat FROM products where id = ?", ids[i]).Scan(&vat)

		var name string
		database.Instance.Raw("Select name FROM products where id = ?", ids[i]).Scan(&name)
		result = append(result, "Product Name: "+name+" Quantity: "+strconv.Itoa(quantity))

		fmt.Println(Discount[vat])
		if quantity > 3 && Discount[vat] < 10 {
			withoutDiscount := price * 3
			withDiscount := price * (quantity - 3) / 100 * 92
			totalPrice := withoutDiscount + withDiscount
			prices = append(prices, totalPrice)
			database.Instance.Raw("Select vat FROM products where id = ?", ids[i]).Scan(&vat)
			vats = append(vats, totalPrice*vat/100)
		} else {
			totalPrice := price * quantity * (100 - Discount[vat]) / 100
			prices = append(prices, totalPrice)
			database.Instance.Raw("Select vat FROM products where id = ?", ids[i]).Scan(&vat)
			vats = append(vats, totalPrice*vat/100)
		}
	}

	TotalAmount = common.Sum(prices)
	totalVATS := common.Sum(vats)

	if len(result) == 0 {
		return c.SendString("Cart is Empty")
	} else {
		result = append(result, "Total Price: "+strconv.Itoa(TotalAmount))
		result = append(result, "Total VAT: "+strconv.Itoa(totalVATS))
		prices = nil
		vats = nil
		TotalAmount = 0
		totalVATS = 0
		return c.Status(200).JSON(result)
	}
}

func CompleteOrder(c *fiber.Ctx) error {
	userId, _ := strconv.Atoi(c.Params("user_id"))

	CompletedOrders := models.PlacedOrders{UserID: userId, TotalPrice: TotalAmount, CreatedAt: time.Now()}

	database.Instance.Select("UserID", "TotalPrice", "CreatedAt").Create(&CompletedOrders)

	database.Instance.Exec("DELETE FROM shopping_carts where user_id = ?", userId)
	return c.SendString("Order Completed!")
}

func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product

	database.Instance.Find(&products)

	return c.Status(200).JSON(products)
}

func DeleteItemFromCart(c *fiber.Ctx) error {
	IdToBeDeleted, _ := strconv.Atoi(c.Params("product_id"))
	userId, _ := strconv.Atoi(c.Params("user_id"))

	database.Instance.Exec("delete from shopping_carts where product_id = ? and user_id = ?", IdToBeDeleted, userId)

	return c.Status(200).JSON("deleted")
}

func GetGivenAmount(c *fiber.Ctx) error {
	var givenAmount []models.GivenAmounts
	database.Instance.Find(&givenAmount)
	return c.Status(200).JSON(givenAmount)
}

func SetGivenAmount(c *fiber.Ctx) error {
	database.Instance.Exec("DELETE from given_amounts")

	newAmount, _ := strconv.Atoi(c.Params("amount"))
	var givenAmount = models.GivenAmounts{
		GivenAmount: newAmount,
	}

	database.Instance.Select("given_amount").Create(&givenAmount)
	return c.SendString("New Given Amount is " + c.Params("amount"))
}
