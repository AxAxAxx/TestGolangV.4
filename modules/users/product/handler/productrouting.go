package handler

import "github.com/gofiber/fiber/v2"

func Routing(route fiber.Router, handler ProductHandler) {
	route.Get("/", handler.GetProduct)
	route.Get("/id/:productID", handler.GetProduct)
	route.Get("/size/:size", handler.GetProduct)
	route.Get("/style/:style", handler.GetProduct)
	route.Get("/gender/:gender", handler.GetProduct)
	//GET Product 2 Filter
	route.Get("/style&size/:style/:size", handler.GetProduct)
	route.Get("/style&gender/:style/:gender", handler.GetProduct)
	route.Get("/size&gender/:size/:gender", handler.GetProduct)
	//GET Product 3 Filter
	route.Get("/style&size&gender/:style/:size/:gender", handler.GetProduct)
	//haslimit
	route.Get("/size&limit/:size/:limit", handler.GetProduct)
	route.Get("/style&limit/:style/:limit", handler.GetProduct)
	route.Get("/gender&limit/:gender/:limit", handler.GetProduct)
	//GET Product 2 Filter
	route.Get("/style&size&limit/:style/:size/:limit", handler.GetProduct)
	route.Get("/style&gender&limit/:style/:gender/:limit", handler.GetProduct)
	route.Get("/size&gender&limit/:size/:gender/:limit", handler.GetProduct)
	//GET Product 3 Filter
	route.Get("/style&size&gender&limit/:style/:size/:gender/:limit", handler.GetProduct)

	route.Post("/create", handler.CreateProduct)

	route.Put("/update/:ID", handler.UpdateProduct)

	route.Delete("/delete/:id", handler.DeleteProduct)
}
