package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Todo :
type Todo struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Title     string             `bson:"title" json:"title"`
	Completed bool               `bson:"completed" json:"completed"`
	CreatedAt time.Time          `bson:"created_at" json:"createAt"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updateAt"`
}
