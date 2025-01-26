package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todolist-app/models"

	"github.com/gorilla/mux"
)

// Get Todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Todos)
}

// Create Todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo
	_ = json.NewDecoder(r.Body).Decode(newTodo)
	newTodo.ID = fmt.Sprintf("%d", len(models.Todos)+1)
	models.Todos = append(models.Todos, newTodo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}

// Update Todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, todo := range models.Todos {
		if todo.ID == params["id"] {
			models.Todos[index].Checked = !models.Todos[index].Checked
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.Todos[index])
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)

}

// Delete todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, todo := range models.Todos {
		if todo.ID == params["id"] {
			models.Todos = append(models.Todos[:index], models.Todos[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.Todos)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)

}
