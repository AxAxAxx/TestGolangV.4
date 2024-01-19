package usecase

import (
	"time"

	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/order/repositories"
)

type OrderUsecase struct {
	OrderRepository repositories.OrderRepositoty
}

func NewOrderUsecase(orderRepository repositories.OrderRepositoty) *OrderUsecase {
	return &OrderUsecase{
		OrderRepository: orderRepository,
	}
}

func (u *OrderUsecase) CreateOrder(user_id float64, newOrder entities.Order) error {
	created_at := time.Now()
	expired := created_at.Add(24 * time.Hour)
	var t_price int
	var price_stock entities.Price_Stock
	var total_price, total_stock []int
	pc_sc, err := u.OrderRepository.GetPriceAndStock(price_stock, newOrder)
	if err != nil {
		return err
	}
	for index, value := range pc_sc {
		for indexProduct, productReq := range newOrder.Products {
			if index == indexProduct {
				value.Price = productReq.Quantity * value.Price
				if value.Stock < productReq.Quantity {
					return err
				}
				value.Stock = value.Stock - productReq.Quantity
				total_stock = append(total_stock, value.Stock)
				total_price = append(total_price, value.Price)
			}
		}
	}
	shippingdeteils, err := u.OrderRepository.GetShipDetails(user_id, newOrder)
	if err != nil {
		return err
	}
	id, err := u.OrderRepository.CreateOrders(user_id, created_at, expired, shippingdeteils, newOrder)
	if err != nil {
		return err
	}
	err = u.OrderRepository.AddProduct(id, t_price, total_price, created_at, newOrder)
	if err != nil {
		return err
	}
	err = u.OrderRepository.UpdatedStock(total_stock, newOrder)
	if err != nil {
		return err
	}
	return nil
}

func (u *OrderUsecase) GetOrders(filter float64) ([]entities.OrderRes, error) {
	orders, err := u.OrderRepository.GetOrders(filter, []entities.OrderRes{})
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (u *OrderUsecase) GetOrdersByAdmid(filter []string, limit, sorting, offset string) ([]entities.OrderRes, error) {
	orders, err := u.OrderRepository.GetOrdersByAdmid(filter, limit, sorting, offset, []entities.OrderRes{})
	if err != nil {
		return nil, err
	}
	return orders, nil
}
