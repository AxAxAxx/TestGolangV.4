package handler

import (
	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/order/controllers"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	OrderController controllers.OrderController
}

func NewOrderHandler(orderController controllers.OrderController) *OrderHandler {
	return &OrderHandler{
		OrderController: orderController,
	}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var newOrder entities.Order

	if err := c.BodyParser(&newOrder); err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}

	err := h.OrderController.CreateOrder(newOrder)
	if err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}
	return c.JSON(entities.OrderSuccessResponse(newOrder))
}

func (o *OrderHandler) GetOrders(c *fiber.Ctx) error {
	orderId := c.Query("order_id")
	status := c.Query("status")
	fname := c.Query("fname")
	lname := c.Query("lname")
	phonenumber := c.Query("phonenumber")
	startdate := c.Query("startdate")
	enddate := c.Query("enddate")
	limit := c.Query("limit")
	orders, err := o.OrderController.GetOrders(orderId, fname, lname, phonenumber, status, startdate, enddate, limit)
	if err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}
	return c.JSON(entities.OrdersSuccessResponse(&orders))
}
