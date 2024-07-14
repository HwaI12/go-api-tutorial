package api

import (
	"database/sql"

	controller "github.com/HwaI12/go-api-tutorial/internal/controller"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	bookController := controller.NewBookController(db)

	router.HandleFunc("/books", bookController.CreateBook).Methods("POST")
	router.HandleFunc("/books", bookController.GetBooks).Methods("GET")
}
