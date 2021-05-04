package services

import (
	"context"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main/src/db"
	"main/src/models"
	"strconv"
)

type UserService struct {
	ctx            *context.Context
	userCollection *mongo.Collection
}

func NewUserService(connection *db.Connection, collectionName string) *UserService {
	collection := connection.NewCollection(collectionName)
	userService := UserService{ctx: &connection.Ctx, userCollection: collection}
	return &userService
}

func (userService *UserService) InsertOne(userName string, g float64, in float64, out float64) (*mongo.InsertOneResult, error) {
	user := models.User{Username: userName, Goal: g, Income: in, Outgoings: out}
	res, err := userService.userCollection.InsertOne(*userService.ctx, user)
	return res, err
}

func (userService *UserService) InsertFromEntry(entryUN *gtk.Entry, entryG *gtk.Entry, entryIn *gtk.Entry, entryOut *gtk.Entry) {

	un, err := entryUN.GetText()
	if err != nil {
		panic(err)
	}
	result := userService.userCollection.FindOne(*userService.ctx, bson.M{"username": un})
	fmt.Println(result)
	//if result == nil {
	//	entryUN.SetText("Username already taken")
	//	fmt.Println("Username already taken")
	//} else {

	g1, err := entryG.GetText()
	if err != nil {
		panic(err)
	}
	g, err := strconv.ParseFloat(g1, 64)
	if err != nil {
		panic(err)
	}
	in1, err := entryIn.GetText()
	if err != nil {
		panic(err)
	}
	in, err := strconv.ParseFloat(in1, 64)
	if err != nil {
		panic(err)
	}
	out1, err := entryOut.GetText()
	if err != nil {
		panic(err)
	}
	out, err := strconv.ParseFloat(out1, 64)
	if err != nil {
		panic(err)
	}

	_, err = userService.InsertOne(un, g, in, out)
	if err != nil {
		panic(err)
	}
	//}
}

func (userService *UserService) FindOne(Cname string) (models.User, error) {
	var result models.User
	err := userService.userCollection.FindOne(*userService.ctx, bson.M{"username": Cname}).Decode(&result)
	return result, err
}

func (userService *UserService) FindOneFromEntry(entryIn *gtk.Entry) {
	uname, err := entryIn.GetText()
	if err != nil {
		panic(err)
	}
	result := userService.userCollection.FindOne(*userService.ctx, bson.M{"username": uname})
	if result == nil {
		fmt.Println("Username does not exist")
	} else {
	}
}
