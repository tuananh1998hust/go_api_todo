package routes

import (
	"github.com/gorilla/mux"
	"github.com/tuananh1998hust/go_api_todo/controllers"
)

// SetUpRoutes :
func SetUpRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/todo", controllers.FindAll).Methods("GET")
	r.HandleFunc("/api/todo/{id}", controllers.FindByID).Methods("GET")
	r.HandleFunc("/api/todo", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/api/todo/{id}", controllers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/api/todo/{id}", controllers.DeleteTodo).Methods("DELETE")

	return r
}
