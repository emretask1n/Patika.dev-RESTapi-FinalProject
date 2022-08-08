package ProductModel

type Product struct {
	ID    int    `json:"Product ID"`
	Name  string `json:"Product Name"`
	Price int    `json:"Product Price"`
	Vat   string `json:"VAT"`
}
