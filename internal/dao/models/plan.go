package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Plan struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Date        string             `bson:"date" json:"date"`
	Completed   bool               `bson:"completed" json:"completed"`
	TimeSpent   int                `bson:"timeSpent,omitempty" json:"timeSpent,omitempty"`
	ImageURL    string             `bson:"imageUrl,omitempty" json:"imageUrl,omitempty"`
}
