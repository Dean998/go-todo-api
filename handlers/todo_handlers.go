package handlers

import (
	"encoding/json"
	// "fmt"
	"net/http"
	// "strconv"
	"todo-api/models"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assigning an ID to the todo
	todo.ID = len(models.Todos) + 1
	models.Todos = append(models.Todos, todo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// ... Implementation to come later
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// ... Implementation to come later
}
