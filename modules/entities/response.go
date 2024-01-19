package entities

import (
	"github.com/gofiber/fiber/v2"
)

func ProductSuccessResponse(data Product) *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"data":   data,
		"error":  nil,
	}
}

func ProductsSuccessResponse(data *[]ProductRes) *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"data":   data,
		"error":  nil,
	}
}

func OrderSuccessResponse(data Order) *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"data":   data,
		"error":  nil,
	}
}

func OrdersSuccessResponse(data *[]OrderRes) *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"data":   data,
		"error":  nil,
	}
}

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}

func DeleteResponse() *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"data":   "delete successfully",
		"err":    nil,
	}
}

func UpdateResponse() *fiber.Map {
	return &fiber.Map{
		"status": "success",
		"data":   "update successfully",
		"err":    nil,
	}
}
