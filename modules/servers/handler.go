package servers

import (
	"github.com/AxAxAxx/go-test-api/modules/users/controllers"
	"github.com/AxAxAxx/go-test-api/modules/users/repositories"
	"github.com/AxAxAxx/go-test-api/modules/users/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Handlers(app *fiber.App, db *sqlx.DB) {
	//Product
	productRepository := repositories.NewProductRepository(db)
	productUsecase := usecases.NewProductUsecase(*productRepository)
	productHandler := controllers.NewProductHandler(*productUsecase)
	app.Get("/products", productHandler.GetProduct)
	app.Get("/productbyid/:productID", productHandler.GetProduct)
	app.Get("/productbysize/:size", productHandler.GetProduct)
	app.Get("/productbystyle/:style", productHandler.GetProduct)
	app.Get("/productbygender/:gender", productHandler.GetProduct)
	//GET Product 2 Filter
	app.Get("/products/style&size/:style/:size", productHandler.GetProduct)
	app.Get("/products/style&gender/:style/:gender", productHandler.GetProduct)
	app.Get("/products/size&gender/:size/:gender", productHandler.GetProduct)
	//GET Product 3 Filter
	app.Get("/products/style&size&gender/:style/:size/:gender", productHandler.GetProduct)

	app.Get("/productbysizelimit/:size/:limit", productHandler.GetProduct)
	app.Get("/productbystylelimit/:style/:limit", productHandler.GetProduct)
	app.Get("/productbygenderlimit/:gender/:limit", productHandler.GetProduct)
	//GET Product 2 Filter
	app.Get("/products/style&sizelimit/:style/:size/:limit", productHandler.GetProduct)
	app.Get("/products/style&genderlimit/:style/:gender/:limit", productHandler.GetProduct)
	app.Get("/products/size&genderlimit/:size/:gender/:limit", productHandler.GetProduct)
	//GET Product 3 Filter
	app.Get("/products/style&size&genderlimit/:style/:size/:gender/:limit", productHandler.GetProduct)
	//Order
	orderRepository := repositories.NewOrderRepository(db)
	orderUsecase := usecases.NewOrderUsecase(*orderRepository)
	orderHandler := controllers.NewOrderHandler(*orderUsecase)
	app.Get("orders", orderHandler.GetOrders)
	//Get
	app.Get("orderbyid/:order_id", orderHandler.GetOrders)
	app.Get("orderbyfirstname/:fname", orderHandler.GetOrders)
	app.Get("orderbylastname/:lname", orderHandler.GetOrders)
	app.Get("orderbyphonenumber/:phonenumber", orderHandler.GetOrders)
	app.Get("orderbystatus/:status", orderHandler.GetOrders)
	app.Get("orderbypaiddate/:startdate/:enddate", orderHandler.GetOrders)

	app.Get("orderbyidlimit/:order_id/:limit", orderHandler.GetOrders)
	app.Get("orderbyfirstnamelimit/:fname/:limit", orderHandler.GetOrders)
	app.Get("orderbylastnamelimit/:lname/:limit", orderHandler.GetOrders)
	app.Get("orderbyphonenumberlimit/:phonenumber/:limit", orderHandler.GetOrders)
	app.Get("orderbystatuslimit/:status/:limit", orderHandler.GetOrders)
	app.Get("orderbypaiddatelimit/:startdate/:enddate/:limit", orderHandler.GetOrders)

	//Post
	app.Post("/orders/create", orderHandler.CreateOrder)

	userRepository := repositories.NewUserRepository(db)
	userUsercase := usecases.NewUserUsecase(*userRepository)
	userHandler := controllers.NewUserHandler(*userUsercase)
	app.Post("account/user", userHandler.CreateUser)
	app.Post("account/address", userHandler.AddAddress)
	app.Post("account/create", userHandler.CreateAccount)

}
