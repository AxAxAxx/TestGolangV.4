package handler

import "github.com/gofiber/fiber/v2"

func Routing(route fiber.Router, handler ProductHandler) {
	route.Get("/", handler.GetProduct)
	route.Post("/create", handler.CreateProduct)
	route.Put("/update/:ID", handler.UpdateProduct)
	route.Delete("/delete/:id", handler.DeleteProduct)
}
