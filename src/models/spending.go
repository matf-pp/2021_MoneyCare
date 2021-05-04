package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Spending struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserId     primitive.ObjectID `bson:"userId,omitempty"`
	Username   string             `bson:"username,omitempty"`
	CategoryId primitive.ObjectID `bson:"categoryId,omitempty"`
	Amount     float64            `bson:"amount,omitempty"`
	Time       time.Time          `bson:"time,omitempty"`
}
