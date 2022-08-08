package main

import (
	"github.com/gofiber/fiber/v2"
	"log"

	"REST_API/database"
	"REST_API/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/products", routes.GetAllProducts)
	app.Get("/cart/:user_id", routes.ShowCart)
	app.Get("/givenAmount", routes.GetGivenAmount)
	app.Delete("/delete/:user_id/:product_id", routes.DeleteItemFromCart)
	app.Post("/:product_id/:user_id/:quantity", routes.AddItemToCart)
	app.Post("/order/:id", routes.CompleteOrder)
	app.Post("/amount/:amount", routes.SetGivenAmount)

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
