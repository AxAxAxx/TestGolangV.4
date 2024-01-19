package servers

import (
	"github.com/AxAxAxx/go-test-api/modules/users/authen/handler"
	"github.com/AxAxAxx/go-test-api/modules/users/authen/repositories"
	"github.com/AxAxAxx/go-test-api/modules/users/authen/usecase"
	orderhandler "github.com/AxAxAxx/go-test-api/modules/users/order/handler"
	orderrepository "github.com/AxAxAxx/go-test-api/modules/users/order/repositories"
	orderusecase "github.com/AxAxAxx/go-test-api/modules/users/order/usecase"
	"github.com/AxAxAxx/go-test-api/pkg/middleware"

	accounthandler "github.com/AxAxAxx/go-test-api/modules/users/account/handler"
	accountrepository "github.com/AxAxAxx/go-test-api/modules/users/account/repositories"
	accountusecase "github.com/AxAxAxx/go-test-api/modules/users/account/usecase"
	producthandler "github.com/AxAxAxx/go-test-api/modules/users/product/handler"
	productrepository "github.com/AxAxAxx/go-test-api/modules/users/product/repositories"
	productusecase "github.com/AxAxAxx/go-test-api/modules/users/product/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Server(app *fiber.App, db *sqlx.DB) {
	v1 := app.Group("/v1")
	//Account
	accountRepo := accountrepository.NewAccountRepository(db)
	accountUsecase := accountusecase.NewAccountUsecase(*accountRepo)
	accountHandler := accounthandler.NewAccountHandler(*accountUsecase)
	accountGroup := v1.Group("/register")
	accounthandler.Routing(accountGroup, *accountHandler)

	//Authen
	authRepo := repositories.NewAccountRepository(db)
	authUsecaes := usecase.NewAuthUsecase(*authRepo)
	authHandler := handler.NewAuthenHandler(*authUsecaes)
	authenGroup := v1.Group("/")
	handler.RoutingGenToken(authenGroup, *authHandler)

	//Order
	orderRepository := orderrepository.NewOrderRepository(db)
	orderController := orderusecase.NewOrderUsecase(*orderRepository)
	orderHandler := orderhandler.NewOrderHandler(*orderController)
	orderGroup := v1.Group("/order", middleware.AuthMiddleware())
	orderhandler.Routing(orderGroup, *orderHandler)

	//Product
	productRepository := productrepository.NewProductRepository(db)
	productUsecase := productusecase.NewProductUsecase(*productRepository)
	productHandler := producthandler.NewProductHandler(*productUsecase)
	productGroup := v1.Group("/product")
	producthandler.Routing(productGroup, *productHandler)
}
