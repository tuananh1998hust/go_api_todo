package main

import (
	"log"
	"net/http"

	"github.com/tuananh1998hust/go_api_todo/config"
	"github.com/tuananh1998hust/go_api_todo/routes"
)

func main() {
	config.CheckConnection()

	r := routes.SetUpRoutes()

	log.Fatal(http.ListenAndServe(":8080", r))
}
