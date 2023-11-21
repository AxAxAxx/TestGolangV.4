package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

type Database struct {
	*sqlx.DB
}

func ConnPgSQL() (*Database, error) {
	var config Config
	config.host = "localhost"
	config.port = 5432
	config.user = "postgres"
	config.password = "1234"
	config.dbname = "testGolang"
	connpostgre := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.host, config.port, config.user, config.password, config.dbname)
	db, err := sqlx.Connect("postgres", connpostgre)
	if err != nil {
		return nil, err
	}
	fmt.Println("-Connect Successful-")
	return &Database{db}, nil
}
