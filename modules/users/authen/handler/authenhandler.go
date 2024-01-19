package handler

import (
	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/authen/usecase"
	"github.com/AxAxAxx/go-test-api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

type AuthenHandler struct {
	AuthenHandler usecase.AuthUsecase
}

func NewAuthenHandler(authenHandler usecase.AuthUsecase) *AuthenHandler {
	return &AuthenHandler{
		AuthenHandler: authenHandler,
	}
}

func (h *AuthenHandler) Login(c *fiber.Ctx) error {
	var req entities.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}
	account, isValidUser, err := h.AuthenHandler.AuthenticateUser(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if isValidUser {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Credentials"})
	}
	err = middleware.VerifyPassword(account.Password, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Credentials"})
	}
	accessToken, refreshToken := middleware.GenerateTokens(account.Username, account.UserID, account.Role_id)
	err = h.AuthenHandler.Token(account.AccountID, accessToken, refreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "don't have token"})
	}
	return c.JSON(fiber.Map{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (h *AuthenHandler) RefreshToken(c *fiber.Ctx) error {
	var req entities.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}
	token, err := h.AuthenHandler.RefreshToken(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if token.Role_id == 0 {
		return c.Status(401).SendString("Invalid Role")
	}
	middleware.TokenBlacklist[token.Token] = true
	newAccessToken, refresh_token := middleware.GenerateTokens(token.Username, token.UserID, token.Role_id)
	err = h.AuthenHandler.UpdateRefreshToken(newAccessToken, refresh_token, token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"newAccessToken": newAccessToken,
		"refreshToken":   refresh_token,
	})
}

func (h *AuthenHandler) Logout(c *fiber.Ctx) error {
	var req entities.DeleteToken
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}
	if req.RefreshToken == "" && req.Token == "" {
		return c.JSON(fiber.Map{
			"ERROR": "Can't logout !!!",
		})
	}
	h.AuthenHandler.DeleteToken(req.Token, req.RefreshToken)
	middleware.TokenBlacklist[req.Token] = true
	return c.JSON(fiber.Map{
		"Status": "Logout Success",
	})
}
