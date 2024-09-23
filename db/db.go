package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/donnykd/go-ecom/config"
	_ "github.com/lib/pq"
)

func NewPostgreSQL(cfg config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.USER, cfg.PASSWORD, cfg.HOST, cfg.PORT, cfg.NAME)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, nil
}
