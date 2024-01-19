package database

import (
	"fmt"

	"github.com/AxAxAxx/go-test-api/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	*sqlx.DB
}

func ConnPgSQL(cfg *config.Configs) (*Database, error) {
	connpostgre := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgreSQL.Host, cfg.PostgreSQL.Port, cfg.PostgreSQL.Username, cfg.PostgreSQL.Password, cfg.PostgreSQL.Database, cfg.PostgreSQL.SSLMode)

	fmt.Println(connpostgre)
	db, err := sqlx.Connect("postgres", connpostgre)
	if err != nil {
		return nil, err
	}
	fmt.Println("-Connect-")
	return &Database{db}, nil
}
