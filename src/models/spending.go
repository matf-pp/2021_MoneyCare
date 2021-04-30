package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Spending struct {
	UserId     primitive.ObjectID `bson:"userId,omitempty"`
	Username   string             `bson:"username,omitempty"`
	CategoryId primitive.ObjectID `bson:"categoryId,omitempty"`
	Amount     int                `bson:"amount,omitempty"`
}
