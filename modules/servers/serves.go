package servers

import (
	ordercontroller "github.com/AxAxAxx/go-test-api/modules/users/order/controllers"
	orderhandler "github.com/AxAxAxx/go-test-api/modules/users/order/handler"
	orderrepository "github.com/AxAxAxx/go-test-api/modules/users/order/repositories"
	productcontroller "github.com/AxAxAxx/go-test-api/modules/users/product/controllers"
	producthandler "github.com/AxAxAxx/go-test-api/modules/users/product/handler"
	productrepository "github.com/AxAxAxx/go-test-api/modules/users/product/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Server(app *fiber.App, db *sqlx.DB) {
	//Order
	orderRepository := orderrepository.NewOrderRepository(db)
	orderController := ordercontroller.NewOrderController(*orderRepository)
	orderHandler := orderhandler.NewOrderHandler(*orderController)
	orderGroup := app.Group("/order")
	orderhandler.Routing(orderGroup, *orderHandler)

	//Product
	productRepository := productrepository.NewProductRepository(db)
	productController := productcontroller.NewProductController(*productRepository)
	productHandler := producthandler.NewProductHandler(*productController)

	productGroup := app.Group("/product")

	producthandler.Routing(productGroup, *productHandler)
}
