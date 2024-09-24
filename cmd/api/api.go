package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/donnykd/go-ecom/service/user"
	"github.com/gorilla/mux"
)

type APIserver struct {
	address string
	db      *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIserver {
	return &APIserver{
		address: address,
		db:      db,
	}
}

func (server *APIserver) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStorage := user.NewStorage(server.db)
	userHandler := user.NewHandler(userStorage)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", server.address)
	return http.ListenAndServe(server.address, router)
}
