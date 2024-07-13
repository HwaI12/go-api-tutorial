package api

import (
	"database/sql"

	"github.com/HwaI12/go-api-tutorial/internal/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	bookController := controllers.NewBookController(db)

	router.HandleFunc("/books", bookController.CreateBook).Methods("POST")
}
