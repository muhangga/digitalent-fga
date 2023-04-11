package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhangga/docs"
	"github.com/muhangga/internal/delivery"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {

	docs.SwaggerInfo.Title = "Book API Documentation"
	docs.SwaggerInfo.Description = "This is a sample server Book server."
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	server := fiber.New()

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	server.Get("/swagger/*", fiberSwagger.WrapHandler)

	api := server.Group("/api")

	api.Get("/books", delivery.GetAllBook)
	api.Get("/book/:id", delivery.GetBookByID)
	api.Post("/book", delivery.SaveBook)
	api.Put("/book/:id", delivery.UpdateBook)
	api.Delete("/book/:id", delivery.DeleteBook)

	server.Listen(":8000")
}
