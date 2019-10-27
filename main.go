package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
)

// Todo :
type Todo struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Completed bool               `bson:"completed"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

var collection *mongo.Collection

func getAllTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todos []*Todo

	findOptions := options.Find()
	findOptions.SetLimit(20)

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var item Todo

		err := cur.Decode(&item)

		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, &item)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	json.NewEncoder(w).Encode(todos)
}

func getTodoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	var todo Todo
	filter := bson.D{{"_id", id}}
	err = collection.FindOne(context.TODO(), filter).Decode(&todo)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(todo)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	result, err := collection.InsertOne(context.TODO(), bson.M{
		"title":      todo.Title,
		"completed":  false,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	})

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", id}}
	update := bson.M{"$set": bson.D{{"completed", true}}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", id}}
	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check Connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB is connected...")

	collection = client.Database("GoRestAPI").Collection("todo")

	r := mux.NewRouter()

	r.HandleFunc("/api/todo", getAllTodo).Methods("GET")
	r.HandleFunc("/api/todo/{id}", getTodoByID).Methods("GET")
	r.HandleFunc("/api/todo", createTodo).Methods("POST")
	r.HandleFunc("/api/todo/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/api/todo/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
