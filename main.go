package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getAllTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get all todos")
}

func getTodoByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get todo by id")
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create todo")
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update todo")
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete todo")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/todo", getAllTodo).Methods("GET")
	r.HandleFunc("/api/todo/{id}", getTodoByID).Methods("GET")
	r.HandleFunc("/api/todo", createTodo).Methods("POST")
	r.HandleFunc("/api/todo/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/api/todo/{id}", deleteTodo)

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}
}
