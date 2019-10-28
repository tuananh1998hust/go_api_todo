package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tuananh1998hust/go_api_todo/config"
	"github.com/tuananh1998hust/go_api_todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client = config.MongoClient()

// TodoCollection :
var TodoCollection *mongo.Collection = mongoClient.Database("GoAPI").Collection("todo")

// FindAll :
func FindAll(w http.ResponseWriter, r *http.Request) {
	var todos []*models.Todo

	findOptions := options.Find()
	findOptions.SetLimit(20)

	cur, err := TodoCollection.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for cur.Next(context.TODO()) {
		var item models.Todo

		err = cur.Decode(&item)

		todos = append(todos, &item)
	}

	if err := cur.Err(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	cur.Close(context.TODO())

	RespondWithJSON(w, http.StatusOK, todos)
}

// FindByID :
func FindByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		return
	}

	var todo models.Todo
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	err = TodoCollection.FindOne(context.TODO(), filter).Decode(&todo)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	result, err := TodoCollection.InsertOne(context.TODO(), bson.M{
		"title":      todo.Title,
		"completed":  false,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	})

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, result)
}

// UpdateTodo :
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.M{"$set": bson.D{bson.E{Key: "completed", Value: true}}}
	result, err := TodoCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, result)
}

// DeleteTodo :
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	filter := bson.D{{"_id", id}}
	result, err := TodoCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, result)
}
