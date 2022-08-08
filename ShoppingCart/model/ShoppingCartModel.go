package ShoppingCartModel

type ShoppingCart struct {
	ProductID int `json:"ProductID"`
	UserID    int `json:"UserID"`
	Quantity  int `json:"Quantity"`
}
