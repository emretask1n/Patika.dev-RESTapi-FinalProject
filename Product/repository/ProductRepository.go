package ProductRepository

import (
	ProductModel "REST_API/Product/model"
	"REST_API/database"
)

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

func GetProductTable() []ProductModel.Product {
	var products []ProductModel.Product
	database.Instance.Find(&products)
	return products
}
