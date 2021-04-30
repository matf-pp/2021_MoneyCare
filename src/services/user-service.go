package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main/src/db"
	"main/src/models"
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

func (userService *UserService) InsertOne(userName string) (*mongo.InsertOneResult, error) {
	user := models.User{Username: userName}
	res, err := userService.userCollection.InsertOne(*userService.ctx, user)
	return res, err
}

func (userService *UserService) FindOne(Cname string) (models.User, error) {
	var result models.User
	err := userService.userCollection.FindOne(*userService.ctx, bson.M{"username": Cname}).Decode(&result)
	return result, err
}
