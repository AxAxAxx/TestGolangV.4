package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AxAxAxx/go-test-api/modules/servers"
	"github.com/AxAxAxx/go-test-api/pkg/config"
	"github.com/AxAxAxx/go-test-api/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func newConfig() *config.Configs {
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}

	return &config.Configs{
		Server: config.Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		PostgreSQL: config.PostgreSQL{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
		},
	}
}

func main() {
	cfg := newConfig()

	db, err := database.ConnPgSQL(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the T-shirts Shop!"))
	})
	servers.Server(app, db.DB)

	connectionURL := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	app.Listen(connectionURL)
}
