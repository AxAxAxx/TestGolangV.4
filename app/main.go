package main

import (
	"log"

	"github.com/AxAxAxx/go-test-api/modules/servers"
	"github.com/AxAxAxx/go-test-api/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db, err := database.ConnPgSQL()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the T-shirts Shop!"))
	})
	servers.Handlers(app, db.DB)

	app.Listen(":3000")
}
