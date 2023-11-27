package entities

import "time"

type Order struct {
	OrderID         int             `db:"order_id"`
	UserID          int             `db:"user_id"`
	ProductID       int             `db:"product_id"`
	OrderStatusID   int             `db:"orderstatus_id"`
	ProductDetails  ProductDetails  `db:"product_details"`
	ShippingDetails ShippingDetails `db:"shipping_details"`
	CreatedAt       time.Time       `db:"created_at"`
	FirstName       string          `db:"first_name"`
	LastName        string          `db:"last_name"`
	PhoneNumber     string          `db:"phonenumber"`
	Quantity        int             `db:"quantity"`
	Total_Price     int             `db:"total_price"`
	OrderStatus     string          `db:"status"`
	StartDate       time.Time       `db:"start_date"`
	EndDate         time.Time       `db:"end_date"`
}

type ShippingDetails struct {
	FirstName      string `db:"first_name"`
	LastName       string `db:"last_name"`
	PhoneNumber    string `db:"phonenumber"`
	AddressDetails string `db:"address_details"`
	PostalCode     string `db:"postal_code"`
	Province       string `db:"province"`
	Country        string `db:"country"`
}

type ProductDetails struct {
	StyleProduct string `db:"style_name"`
	Size         string `db:"size"`
	Gender       string `db:"gender_name"`
	Price        int    `db:"price"`
}
