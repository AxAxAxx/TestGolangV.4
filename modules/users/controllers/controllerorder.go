package controllers

import (
	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/usecases"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	OrderUsecase usecases.OrderUsecase
}

func NewOrderHandler(oroductUsecase usecases.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		OrderUsecase: oroductUsecase,
	}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var newOrder entities.Order

	if err := c.BodyParser(&newOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err := h.OrderUsecase.CreateOrder(newOrder)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return nil
}

func (o *OrderHandler) GetOrders(c *fiber.Ctx) error {
	orderId := c.Params("order_id")
	status := c.Params("status")
	fname := c.Params("fname")
	lname := c.Params("lname")
	phonenumber := c.Params("phonenumber")
	startdate := c.Params("startdate")
	enddate := c.Params("enddate")
	limit := c.Params("limit")
	orders, err := o.OrderUsecase.GetOrders(orderId, fname, lname, phonenumber, status, startdate, enddate, limit)
	if err != nil {
		return err
	}
	return c.JSON(orders)
}
