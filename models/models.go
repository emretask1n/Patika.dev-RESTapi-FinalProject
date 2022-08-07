package models

import (
	"time"
)

type Product struct {
	ID    int    `json:"Product ID"`
	Name  string `json:"Product Name"`
	Price int    `json:"Product Price"`
	Vat   string `json:"VAT"`
}

type ShoppingCart struct {
	ProductID int `json:"ProductID"`
	UserID    int `json:"UserID"`
	Quantity  int `json:"Quantity"`
}

type User struct {
	UserID int
	Name   string
}

type PlacedOrders struct {
	UserID     int
	TotalPrice int
	CreatedAt  time.Time
}

type Basket struct {
	id       int
	Name     string `json:"ProductName"`
	Quantity int    `json:"Quantity"`
}

type TotalPrice struct {
	TotalPrice int `json:"TotalPrice"`
	TotalVAT   int `json:"TotalVAT"`
}
