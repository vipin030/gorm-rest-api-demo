package main

import (
	"book-management/db"
	"book-management/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	handler := handlers.New(DB)

	router := mux.NewRouter()

	router.HandleFunc("/books", handler.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", handler.AddBook).Methods(http.MethodPost)

	log.Println("Server has started")
	http.ListenAndServe(":8000", router)
}
