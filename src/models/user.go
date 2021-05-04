package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//mora da bude jedinstveno po ID-ju, ali vise istih username-ova, zbog razlicitih meseci??
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username,omitempty"`
	Goal      float64            `bson:"goal,omitempty"`
	Income    float64            `bson:"income,omitempty"`
	Outgoings float64            `bson:"outgoings,omitempty"`
}
