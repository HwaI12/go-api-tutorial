package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// db := database.Connect()
	router := mux.NewRouter()
	// api.RegisterRoutes(router, db)
	fmt.Println("Server Start Up........")
	fmt.Println("http://localhost:8080")

	http.ListenAndServe(":8080", router)
}
