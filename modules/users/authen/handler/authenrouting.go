package handler

import (
	"github.com/gofiber/fiber/v2"
)

func RoutingGenToken(router fiber.Router, handler AuthenHandler) {
	router.Post("/login", handler.Login)
	router.Post("/refreshToken", handler.RefreshToken)
	router.Post("/logout", handler.Logout)
}
