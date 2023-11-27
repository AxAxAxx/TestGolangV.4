package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Routing(router fiber.Router, handler OrderHandler) {
	//Get
	router.Get("/", handler.GetOrders)
	router.Get("/id/:order_id", handler.GetOrders)
	router.Get("/firstname/:fname", handler.GetOrders)
	router.Get("/lastname/:lname", handler.GetOrders)
	router.Get("/phonenumber/:phonenumber", handler.GetOrders)
	router.Get("/status/:status", handler.GetOrders)
	router.Get("/paiddate/:startdate/:enddate", handler.GetOrders)
	//haslimit
	router.Get("/orderbyidlimit/:order_id/:limit", handler.GetOrders)
	router.Get("/orderbyfirstnamelimit/:fname/:limit", handler.GetOrders)
	router.Get("/orderbylastnamelimit/:lname/:limit", handler.GetOrders)
	router.Get("/orderbyphonenumberlimit/:phonenumber/:limit", handler.GetOrders)
	router.Get("/orderbystatuslimit/:status/:limit", handler.GetOrders)
	router.Get("/orderbypaiddatelimit/:startdate/:enddate/:limit", handler.GetOrders)
	//Post
	router.Post("/create", handler.CreateOrder)

}
