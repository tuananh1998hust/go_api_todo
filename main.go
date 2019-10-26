package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Todo :
type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Mock Data
var todosCollection []Todo

func getAllTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todosCollection)
}

func getTodoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range todosCollection {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Todo{})
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = strconv.Itoa(rand.Intn(1000000))
	todo.Completed = false
	todosCollection = append(todosCollection, todo)

	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index := range todosCollection {
		if todosCollection[index].ID == params["id"] {
			todosCollection[index].Completed = true
			json.NewEncoder(w).Encode(todosCollection[index])
			return
		}
	}

	json.NewEncoder(w).Encode(&Todo{})
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var todo Todo

	for index := range todosCollection {
		if todosCollection[index].ID == params["id"] {
			todosCollection[index] = todosCollection[len(todosCollection)-1]
			todosCollection[len(todosCollection)-1] = todo
			todosCollection = todosCollection[:len(todosCollection)-1]
			return
		}
	}

	fmt.Fprintf(w, "Delete Success")
}

func main() {
	r := mux.NewRouter()

	// Mock Data
	todosCollection = append(todosCollection, Todo{ID: "1", Title: "Study Golang", Completed: true})
	todosCollection = append(todosCollection, Todo{ID: "2", Title: "Golang RestAPI", Completed: false})

	r.HandleFunc("/api/todo", getAllTodo).Methods("GET")
	r.HandleFunc("/api/todo/{id}", getTodoByID).Methods("GET")
	r.HandleFunc("/api/todo", createTodo).Methods("POST")
	r.HandleFunc("/api/todo/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/api/todo/{id}", deleteTodo).Methods("DELETE")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}
}
