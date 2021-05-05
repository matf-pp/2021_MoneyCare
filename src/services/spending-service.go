package services

import (
	"context"
	"github.com/gotk3/gotk3/gtk"
	bson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"main/src/db"
	"main/src/models"
	"strconv"
	"time"
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

func (spendingService *SpendingService) InsertOne(userID primitive.ObjectID, categoryId primitive.ObjectID, amount float64) (*mongo.InsertOneResult, error) {
	spending := models.Spending{UserID: userID, CategoryID: categoryId, Amount: amount, Time: time.Now()}
	res, err := spendingService.spendingCollection.InsertOne(*spendingService.ctx, spending)
	return res, err
}

func (spendingService *SpendingService) InsertFromEntry(userID primitive.ObjectID, categoryID primitive.ObjectID, entryAmount *gtk.Entry) (*mongo.InsertOneResult, error) {

	amount1, err := entryAmount.GetText()
	if err != nil {
		panic(err)
	}
	amount, err := strconv.ParseFloat(amount1, 64)
	if err != nil {
		panic(err)
	}

	res, err := spendingService.InsertOne(userID, categoryID, amount)
	return res, err
}

func (spendingService *SpendingService) FindUsersSpending(userIdp primitive.ObjectID) float64 {
	var spendingsFiltered []models.Spending
	filterCursor, err := spendingService.spendingCollection.Find(*spendingService.ctx, bson.M{"userId": userIdp})
	if err != nil {
		panic(err)
	}

	if err = filterCursor.All(*spendingService.ctx, &spendingsFiltered); err != nil {
		log.Fatal(err)
	}

	n := len(spendingsFiltered)
	spent := 0.0
	for i := 0; i < n; i++ {
		spent += spendingsFiltered[i].Amount
	}

	//fmt.Println(spent)
	return spent
}

func (spendingService *SpendingService) FindUsersSpendingByMonth(userIdp primitive.ObjectID, month time.Month) float64 {
	var spendingsFiltered []models.Spending
	filterCursor, err := spendingService.spendingCollection.Find(*spendingService.ctx, bson.M{"userId": userIdp})
	if err != nil {
		panic(err)
	}

	if err = filterCursor.All(*spendingService.ctx, &spendingsFiltered); err != nil {
		log.Fatal(err)
	}

	n := len(spendingsFiltered)
	spent := 0.0
	for i := 0; i < n; i++ {
		if spendingsFiltered[i].Time.Month() == month {
			spent += spendingsFiltered[i].Amount
		}
	}

	//fmt.Println(spent)
	return spent
}

func (spendingService *SpendingService) FindUsersSpendingByDate(userIdp primitive.ObjectID, ourTime time.Time) float64 {
	var spendingsFiltered []models.Spending
	filterCursor, err := spendingService.spendingCollection.Find(*spendingService.ctx, bson.M{"userId": userIdp})
	if err != nil {
		panic(err)
	}

	if err = filterCursor.All(*spendingService.ctx, &spendingsFiltered); err != nil {
		log.Fatal(err)
	}

	n := len(spendingsFiltered)
	spent := 0.0
	for i := 0; i < n; i++ {
		y1, m1, d1 := spendingsFiltered[i].Time.Date()
		y2, m2, d2 := ourTime.Date()
		if y1 == y2 && m1 == m2 && d1 == d2 {
			spent += spendingsFiltered[i].Amount
		}
	}

	//fmt.Println(spent)
	return spent
}

func (spendingService *SpendingService) FindUsersSpendingByCategory(userIdp primitive.ObjectID, categoryIDp primitive.ObjectID) float64 {
	var spendingsFiltered []models.Spending
	filterCursor, err := spendingService.spendingCollection.Find(*spendingService.ctx, bson.M{"userId": userIdp, "categoryId": categoryIDp})
	if err != nil {
		panic(err)
	}

	if err = filterCursor.All(*spendingService.ctx, &spendingsFiltered); err != nil {
		log.Fatal(err)
	}

	n := len(spendingsFiltered)
	spent := 0.0
	for i := 0; i < n; i++ {
		spent += spendingsFiltered[i].Amount
	}

	//fmt.Println(spent)
	return spent
}

func (spendingService *SpendingService) FindUsersSpendingByCategoryByMonth(userIdp primitive.ObjectID, categoryIDp primitive.ObjectID, month time.Month) float64 {
	var spendingsFiltered []models.Spending
	filterCursor, err := spendingService.spendingCollection.Find(*spendingService.ctx, bson.M{"userId": userIdp, "categoryId": categoryIDp})
	if err != nil {
		panic(err)
	}

	if err = filterCursor.All(*spendingService.ctx, &spendingsFiltered); err != nil {
		log.Fatal(err)
	}

	n := len(spendingsFiltered)
	spent := 0.0
	for i := 0; i < n; i++ {
		if spendingsFiltered[i].Time.Month() == month {
			spent += spendingsFiltered[i].Amount
		}
	}

	return spent
}

func (spendingService *SpendingService) Find(userIdp primitive.ObjectID) (models.Spending, error) {
	var result models.Spending
	err := spendingService.spendingCollection.FindOne(*spendingService.ctx, bson.M{"username": userIdp}).Decode(&result)
	return result, err
}
