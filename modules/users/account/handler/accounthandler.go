package handler

import (
	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/account/usecase"
	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	UserUsecase usecase.AccountUsecase
}

func NewAccountHandler(userUsecase usecase.AccountUsecase) *AccountHandler {
	return &AccountHandler{
		UserUsecase: userUsecase,
	}
}

func (h *AccountHandler) CreateAccount(c *fiber.Ctx) error {
	var newAccount entities.RegisterAccount
	if err := c.BodyParser(&newAccount); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err := h.UserUsecase.CreateAccount(newAccount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(newAccount)
}
