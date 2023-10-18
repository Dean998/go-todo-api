package main

import (
	"log"
	"net/http"
	"todo-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	// Placeholder for future
	// r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	// r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
