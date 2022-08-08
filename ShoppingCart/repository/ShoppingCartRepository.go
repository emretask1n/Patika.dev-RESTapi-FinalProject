package ShoppingCartRepository

import (
	PlacedOrderModel "REST_API/PlacedOrders/model"
	ShoppingCartModel "REST_API/ShoppingCart/model"
	"REST_API/database"
	"gorm.io/gorm"
)

func GetIdsOfProductsFromCartByUserId(userId int) []int {
	var ids []int
	database.Instance.Raw("Select product_id FROM shopping_carts where user_id = ?", userId).Scan(&ids)
	return ids
}

func GetQuantityByProductId(productId int) int {
	var quantity int
	database.Instance.Raw("Select quantity FROM shopping_carts where product_id = ?", productId).Scan(&quantity)
	return quantity
}

func InsertIntoCart(shoppingCart ShoppingCartModel.ShoppingCart) *gorm.DB {
	return database.Instance.Select("ProductID", "UserID", "Quantity").Create(&shoppingCart)
}

func CompleteOrders(CompletedOrders PlacedOrderModel.PlacedOrders) *gorm.DB {
	return database.Instance.Select("UserID", "TotalPrice", "CreatedAt").Create(&CompletedOrders)
}

func DeleteCartById(userId int) *gorm.DB {
	return database.Instance.Exec("DELETE FROM shopping_carts where user_id = ?", userId)
}

func DeleteProductFromCartByUProductIdAndUserId(IdToBeDeleted int, userId int) *gorm.DB {
	return database.Instance.Exec("delete from shopping_carts where product_id = ? and user_id = ?", IdToBeDeleted, userId)
}
