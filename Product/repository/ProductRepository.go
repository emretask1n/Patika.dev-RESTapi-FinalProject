package ProductRepository

import "REST_API/database"

func GetPriceByProductId(productId int) int {
	var price int
	database.Instance.Raw("Select price FROM products where id = ?", productId).Scan(&price)
	return price
}

func GetVATByProductId(productId int) int {
	var vat int
	database.Instance.Raw("Select vat FROM products where id = ?", productId).Scan(&vat)
	return vat
}

func GetProductNameByProductId(productId int) string {
	var name string
	database.Instance.Raw("Select name FROM products where id = ?", productId).Scan(&name)
	return name
}

func GetVATTypes() int {
	var vatTypes int
	database.Instance.Raw("SELECT COUNT( DISTINCT vat ) FROM products").Scan(&vatTypes)
	return vatTypes
}
