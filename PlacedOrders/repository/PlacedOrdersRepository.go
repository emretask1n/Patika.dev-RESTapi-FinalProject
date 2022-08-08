package PlacedOrderRepository

import "REST_API/database"

func GetMonthlySpendingByUserId(userId int) int {
	var MonthlySpending int
	database.Instance.Raw("select sum(total_price) from placed_orders where created_at > current_date - interval 30 day and user_id = ?", userId).Scan(&MonthlySpending)
	return MonthlySpending
}

func GetOrderCountAndSpendingForDiscount(userId int, GivenAmount int) int {
	var OrderCountForDiscount int
	database.Instance.Raw("SELECT COUNT( user_id ) FROM placed_orders where user_id = ? and total_price > ?", userId, GivenAmount).Scan(&OrderCountForDiscount)
	return OrderCountForDiscount
}
