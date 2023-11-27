package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Routing(router fiber.Router, handler OrderHandler) {
	router.Get("/", handler.GetOrders)
	router.Post("/create", handler.CreateOrder)

}
