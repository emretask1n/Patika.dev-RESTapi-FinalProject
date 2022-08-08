package PlacedOrderModel

import "time"

type PlacedOrders struct {
	UserID     int
	TotalPrice int
	CreatedAt  time.Time
}
