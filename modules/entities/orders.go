package entities

import (
	"encoding/json"
	"time"
)

type Order struct {
	OrderID         int                   `db:"order_id"`
	UserID          int                   `db:"user_id"`
	ShippingDetails json.RawMessage       `db:"shipping_details"`
	OrderStatusID   int                   `db:"order_status_id"`
	StartDate       time.Time             `db:"start_date"`
	EndDate         time.Time             `db:"end_date"`
	Products        []OrderProductRequest `json:"products"`
}

type OrderRes struct {
	OrderID         int             `json:"order_id" db:"order_id"`
	UserID          int             `json:"user_id" db:"user_id"`
	FirstName       string          `json:"first_name" db:"first_name"`
	LastName        string          `json:"last_name" db:"last_name"`
	PhoneNumber     string          `json:"phonenumber" db:"phonenumber"`
	OrderStatusID   int             `json:"order_status_id" db:"order_status_id"`
	OrderStatus     string          `json:"status" db:"status"`
	CreatedAT       time.Time       `json:"created_at" db:"created_at"`
	Expired         time.Time       `json:"expired" db:"expired"`
	ShippingDetails json.RawMessage `json:"shipping_details" db:"shipping_details"`
	ProductDetails  json.RawMessage `json:"product_details" db:"product_details"`
}

type CreateOrderRequest struct {
	OrderDate  string `json:"order_date"`
	CustomerID string `json:"customer_id"`
}

// OrderProductRequest represents the expected JSON format for a product in the order
type OrderProductRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Price_Stock struct {
	Price int `db:"price"`
	Stock int `db:"stock"`
}

type PS2 struct {
	Price int `db:"price"`
	Stock int `db:"stock"`
}
