package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Routing(router fiber.Router, handler AccountHandler) {
	router.Post("/", handler.CreateAccount)
}
