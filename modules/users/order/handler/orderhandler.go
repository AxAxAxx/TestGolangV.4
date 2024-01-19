package handler

import (
	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/order/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type OrderHandler struct {
	OrderHandler usecase.OrderUsecase
}

func NewOrderHandler(orderHandler usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		OrderHandler: orderHandler,
	}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	claims, _ := c.Locals("username").(jwt.MapClaims)
	roleValue, exists := claims["user_id"]
	if !exists || roleValue == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Role not found or nil"})
	}
	user_id := roleValue.(float64)
	var newOrder entities.Order
	if err := c.BodyParser(&newOrder); err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}
	err := h.OrderHandler.CreateOrder(user_id, newOrder)
	if err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}
	return c.JSON(entities.OrderSuccessResponse(newOrder))
}

func (h *OrderHandler) GetOrders(c *fiber.Ctx) error {
	claims, _ := c.Locals("username").(jwt.MapClaims)
	roleValue, exists := claims["user_id"]
	if !exists || roleValue == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Role not found or nil"})
	}
	user_id := roleValue.(float64)
	orders, err := h.OrderHandler.GetOrders(user_id)
	if err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}
	return c.JSON(entities.OrdersSuccessResponse(&orders))
}

func (h *OrderHandler) GetOrdersByAdmid(c *fiber.Ctx) error {
	filter := []string{}
	orderId := c.Query("order_id")
	status := c.Query("status")
	fname := c.Query("fname")
	lname := c.Query("lname")
	phonenumber := c.Query("phonenumber")
	created_at := c.Query("created_at")
	expired := c.Query("expired")
	limit := c.Query("limit")
	sorting := c.Query("sorting")
	offset := c.Query("offset")
	//fmt.Println(orderId, status, fname, lname, phonenumber)
	filter = append(filter, orderId, fname, lname, status, phonenumber, created_at, expired)
	orders, err := h.OrderHandler.GetOrdersByAdmid(filter, limit, sorting, offset)
	if err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}
	return c.JSON(entities.OrdersSuccessResponse(&orders))
}
