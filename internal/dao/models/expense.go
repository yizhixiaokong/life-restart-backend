package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Expense struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type     string             `bson:"type" json:"type"`
	Category string             `bson:"category" json:"category"`
	Amount   float64            `bson:"amount" json:"amount"`
	Date     string             `bson:"date" json:"date"`
}
