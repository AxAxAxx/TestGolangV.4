package entities

import "github.com/gofiber/fiber/v2"

func ProductSuccessResponse(data Product) *fiber.Map {
	product := Product{
		ProductID:    data.ProductID,
		StyleGroup:   data.StyleGroup,
		StyleProduct: data.StyleProduct,
		Size:         data.Size,
		Gender:       data.Gender,
		Price:        data.Price,
	}
	return &fiber.Map{
		"status": "success",
		"data":   product,
		"error":  nil,
	}
}

func ProductsSuccessResponse(data *[]Product) *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"data":   data,
		"error":  nil,
	}
}

func OrderSuccessResponse(data Order) *fiber.Map {
	order := Order{
		OrderID:         data.OrderID,
		UserID:          data.UserID,
		ProductID:       data.ProductID,
		ProductDetails:  data.ProductDetails,
		ShippingDetails: data.ShippingDetails,
		CreatedAt:       data.CreatedAt,
		FirstName:       data.FirstName,
		LastName:        data.LastName,
		PhoneNumber:     data.PhoneNumber,
		Quantity:        data.Quantity,
		Total_Price:     data.Total_Price,
		OrderStatus:     data.OrderStatus,
		StartDate:       data.StartDate,
		EndDate:         data.EndDate,
	}
	return &fiber.Map{
		"status": "success",
		"data":   order,
		"error":  nil,
	}
}

func OrdersSuccessResponse(data *[]Order) *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"data":   data,
		"error":  nil,
	}
}

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func DeleteResponse() *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"data":   "delete successfully",
		"err":    nil}
}

func UpdateResponse() *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"data":   "update successfully",
		"err":    nil}
}
