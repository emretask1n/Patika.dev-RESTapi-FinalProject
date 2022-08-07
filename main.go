package main

import (
	"github.com/gofiber/fiber/v2"
	"log"

	"REST_API/database"
	"REST_API/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Get("/products", routes.GetAllProducts)
	app.Get("/cart/:id", routes.ShowCart)
	app.Delete("/delete/:id", routes.Delete)
	app.Post("/:product_id/:user_id/:quantity", routes.AddItemToCart)
	app.Post("/order/:id", routes.CompleteOrder)
	/*
		app.Post("/book", routes.AddBook)
		app.Put("/book", routes.Update)
		app.Delete("/book", routes.Delete)

	*/
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
