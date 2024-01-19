package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Routing(route fiber.Router, handler ProductHandler) {
	route.Get("/", handler.GetProduct)
	route.Post("/", handler.CreateProduct)
	route.Patch("/:id", handler.UpdateProduct)
	route.Delete("/:id", handler.DeleteProduct)
}
