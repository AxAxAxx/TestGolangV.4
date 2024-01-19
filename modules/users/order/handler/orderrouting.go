package handler

import (
	"github.com/AxAxAxx/go-test-api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func Routing(router fiber.Router, handler OrderHandler) {
	router.Get("/admin", middleware.AdminOnly(), handler.GetOrdersByAdmid)
	router.Get("/", handler.GetOrders)
	router.Post("/", handler.CreateOrder)
}
