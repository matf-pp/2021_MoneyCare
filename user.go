package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID     			primitive.ObjectID `bson:"_id,omitempty"` //
	Username		string				`bson:"username,omitempty"`
	Month  			time.Month         `bson:"month,omitempty"`
	Income 			float64            	`bson:"income,omitempty"` //ukupni prihodi za ceo mesec (navodi se na pocetku meseca, moguce dodati jos tokom meseca)
	Outgoings	    float64 		   `bson:"outgoings,omitempty"` //ukupni rashodi za ceo mesec (inicijalno jednaki fiksnim rashodima); mozda da vidimo da imamo ono sto se puni pa kad predje odredjenu granicu da se zacrveni
	ConstOutgoings	float64			   `bson:"const_outgoings,omitempty"`
	Goal 			float64			   `bson:"goal,omitempty"` //koliko zeli da ustedi tokom ovog meseca
	Food			float64			   `bson:"food,omitempty"` //koliko je zapravo ustedeo (prihodi-rashodi); moze da se racuna sve vreme tokom meseca, ali mozda to nema smisla
	Chem			float64			   `bson:"chem,omitempty"`
	Clothes			float64			   `bson:"clothes,omitempty"`
	Bills			float64			   `bson:"bills,omitempty"`
	Other			float64			   `bson:"other,omitempty"`
}

//dodavanje korisnika nakon registracije
func add_user (uname string, month time.Month, income float64, outgoings float64, goal float64) {
	//poziva se na signUpOK
	//da li je unet username
	//provera da li vec postoji taj username
	user := User{
		Username: uname,
		Month: month,
		Income: income,
		Outgoings: 0.0,
		ConstOutgoings: outgoings,
		Goal: goal,
		Food: 0.0,
		Chem: 0.0,
		Clothes: 0.0,
		Bills: 0.0,
		Other: 0.0,
	}
	transactions.InsertOne(ctx, user)

}



