package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Spending struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     primitive.ObjectID `bson:"userId,omitempty"`
	CategoryID primitive.ObjectID `bson:"categoryId,omitempty"`
	Amount     float64            `bson:"amount,omitempty"`
	Time       time.Time          `bson:"time,omitempty"`
}
