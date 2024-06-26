package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReadingEntry struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Author      string             `bson:"author" json:"author"`
	Type        string             `bson:"type" json:"type"`
	Total       int                `bson:"total" json:"total"`
	Progress    int                `bson:"progress" json:"progress"`
	TimeSpent   int                `bson:"timeSpent" json:"timeSpent"`
	LastUpdated string             `bson:"lastUpdated" json:"lastUpdated"`
}
