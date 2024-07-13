package api

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	// bookController := controllers.NewBookController(db)

	// router.HandleFunc("/books", bookController.GetBooks).Methods("GET")
	// router.HandleFunc("/books", bookController.CreateBook).Methods("POST")
}
