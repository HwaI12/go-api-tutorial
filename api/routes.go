package api

import (
	"database/sql"
	"go-api-tutorial/internal/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	bookController := controllers.NewBookController(db)

	router.HandleFunc("/books", bookController.GetBooks).Methods("GET")
	router.HandleFunc("/books", bookController.CreateBook).Methods("POST")
	// router.HandleFunc("/books/{id}", bookController.GetBook).Methods("GET")
	// router.HandleFunc("/books/{id}", bookController.UpdateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", bookController.DeleteBook).Methods("DELETE")
}
