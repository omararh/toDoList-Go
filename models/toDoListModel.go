package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	taskTitle   string             `json:"task,omitempty"`
	taskState   string             `json:"task,omitempty"`
	Color   string             `json:"task,omitempty"`
}