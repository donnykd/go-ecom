package main

import (
	"log"

	"github.com/donnykd/go-ecom/cmd/api"
	"github.com/donnykd/go-ecom/config"
	"github.com/donnykd/go-ecom/db"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cfg := config.InitConfig()

	db, err := db.NewPostgreSQL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully Connected!")

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
