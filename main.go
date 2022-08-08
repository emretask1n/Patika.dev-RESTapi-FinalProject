package main

import (
	GivenAmountRouter "REST_API/GivenAmount/router"
	ProductRouter "REST_API/Product/router"
	ShoppingCartRouter "REST_API/ShoppingCart/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"

	"REST_API/database"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/products", ProductRouter.GetAllProducts)
	app.Get("/cart/:user_id", ShoppingCartRouter.ShowCart)
	app.Get("/givenAmount", GivenAmountRouter.GetGivenAmount)
	app.Delete("/delete/:user_id/:product_id", ShoppingCartRouter.DeleteItemFromCart)
	app.Post("/:product_id/:user_id/:quantity", ShoppingCartRouter.AddItemToCart)
	app.Post("/order/:user_id", ShoppingCartRouter.CompleteOrder)
	app.Post("/amount/:amount", GivenAmountRouter.SetGivenAmount)

}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
