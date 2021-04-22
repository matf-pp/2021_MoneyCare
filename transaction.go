package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Transaction struct {
	ID   		 primitive.ObjectID `bson:"_id,omitempty"`  //ovo je 'ID' transakcije
	Username	 string				`bson:"korisnik,omitempty"`
	Category   	Kategorija             `bson:"kategorija,omitempty"`
	Description	string             `bson:"opis,omitempty"`
	Amount	     float64            `bson:"iznos,omitempty"`
	Date		 time.Time			`bson:"datum,omitempty"`
}