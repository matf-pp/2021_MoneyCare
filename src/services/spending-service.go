package services

import (
	"context"
	"fmt"
	bson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"main/src/db"
	"main/src/models"
)

type SpendingService struct {
	ctx                *context.Context
	spendingCollection *mongo.Collection
}

func NewSpendingService(connection *db.Connection, collectionName string) *SpendingService {
	collection := connection.NewCollection(collectionName)
	spendingService := SpendingService{ctx: &connection.Ctx, spendingCollection: collection}
	return &spendingService
}

func (spendingService *SpendingService) InsertOne(userId primitive.ObjectID, categoryId primitive.ObjectID, amount int) (*mongo.InsertOneResult, error) {
	spending := models.Spending{UserId: userId, CategoryId: categoryId, Amount: amount}
	res, err := spendingService.spendingCollection.InsertOne(*spendingService.ctx, spending)
	return res, err
}

func (spendingService *SpendingService) FindUsersSpending(userIdp primitive.ObjectID) {
	var spendingsFiltered []models.Spending
	filterCursor, err := spendingService.spendingCollection.Find(*spendingService.ctx, bson.M{"userId": userIdp})
	if err != nil {
		panic(err)
	}

	if err = filterCursor.All(*spendingService.ctx, &spendingsFiltered); err != nil {
		log.Fatal(err)
	}

	n := len(spendingsFiltered)
	spent := 0
	for i := 0; i < n; i++ {
		spent += spendingsFiltered[i].Amount
	}

	fmt.Println(spent)

	//matchStage := bson.D{{"$match", bson.D{{"userId", userIdp}}}}
	//groupStage := bson.D{{"$group", bson.D{{"_id", "$userId"}, {"total", bson.D{{"$sum", "$amount"}}}}}}
	//showInfoCursor, err := spendingService.spendingCollection.Aggregate(*spendingService.ctx, mongo.Pipeline{matchStage, groupStage})
	//if err != nil {
	//	panic(err)
	//}
	//
	//var showsWithInfo []bson.M
	//if err = showInfoCursor.All(*spendingService.ctx, &showsWithInfo); err != nil {
	//	panic(err)
	//}
	//fmt.Println(showsWithInfo) //return showInfoCursor, err
}
func (spendingService *SpendingService) Find(userIdp primitive.ObjectID) (models.Spending, error) {
	var result models.Spending
	err := spendingService.spendingCollection.FindOne(*spendingService.ctx, bson.M{"username": userIdp}).Decode(&result)
	return result, err
}
