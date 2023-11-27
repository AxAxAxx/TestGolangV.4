package controllers

import (
	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/order/repositories"
)

type OrderController struct {
	OrderRepository repositories.OrderRepositoty
}

func NewOrderController(orderRepository repositories.OrderRepositoty) *OrderController {
	return &OrderController{
		OrderRepository: orderRepository,
	}
}

func (u *OrderController) CreateOrder(newOrder entities.Order) error {
	err := u.OrderRepository.CreateOrder(newOrder)
	if err != nil {
		return err
	}
	return nil
}

func (u *OrderController) GetOrders(id, fname, lname, phonenumber, status, startdate, enddate, limit string) ([]entities.Order, error) {
	orders, err := u.OrderRepository.GetOrders(id, fname, lname, phonenumber, status, startdate, enddate, limit, []entities.Order{})
	if err != nil {
		return nil, err
	}
	return orders, nil
}
