package admin

import (
	"main/src/db"
	"main/src/services"
	_ "main/src/services"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var usersID [90]string
var rnormValues [30]float64
var connectionString string
var databaseName string
var userCollectionName string
var categoryCollectionName string
var spendingCollectionName string

var UserService *services.UserService
var SpendingService *services.SpendingService
var CategoryService *services.CategoryService

func init() {
	connectionString = "mongodb://localhost:27017"
	databaseName = "MoneyCare"
	userCollectionName = "users"
	categoryCollectionName = "categories"
	spendingCollectionName = "spending"
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getNormDistro(x int, mi float32, sigma float32) {
	for i := 0; i < x; i++ {
		rnormValues[i] = float64(normalInverse(mi, sigma))
	}
}

func normalInverse(mu float32, sigma float32) float32 {
	return float32(rand.NormFloat64()*float64(sigma) + float64(mu))
}

func SetupSeed() {
	connection, err := db.NewConnection(connectionString, databaseName)
	if err != nil {
		panic(err)
	}

	//defer connection.Disconnect()

	UserService = services.NewUserService(connection, userCollectionName)
	CategoryService = services.NewCategoryService(connection, categoryCollectionName)
	SpendingService = services.NewSpendingService(connection, spendingCollectionName)
	categorynamestonumbers := make(map[int]string)
	categorynamestonumbers[0] = "food"
	categorynamestonumbers[1] = "chem"
	categorynamestonumbers[2] = "other"
	categorynamestonumbers[3] = "clothes"
	categorynamestonumbers[4] = "bills"

	//for j := 0; j < 5; j++ {
	//	getNormDistro(30, 10000, 3000)
	//	usersID[j] = RandStringRunes(6)
	//	UserService.InsertOne(usersID[j], 0, 50000, 0)
	//	iduser, _ := UserService.FindOne(usersID[j])
	//	for k := 0; k < 5; k++ {
	//		idcategory, err := CategoryService.FindOne(categorynamestonumbers[k])
	//		if err != nil {
	//			panic(err)
	//		}
	//		for i := 0; i < 30; i++ {
	//			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
	//		}
	//	}
	//}
	//
	//for j := 0; j < 5; j++ {
	//	getNormDistro(30, 30000, 2000)
	//	usersID[j] = RandStringRunes(6)
	//	UserService.InsertOne(usersID[j], 0, 150000, 0)
	//	iduser, _ := UserService.FindOne(usersID[j])
	//	for k := 0; k < 5; k++ {
	//		idcategory, err := CategoryService.FindOne(categorynamestonumbers[k])
	//		if err != nil {
	//			panic(err)
	//		}
	//		for i := 0; i < 30; i++ {
	//			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
	//		}
	//	}
	//}
	//for j := 0; j < 5; j++ {
	//	getNormDistro(30, 50000, 10000)
	//	usersID[j] = RandStringRunes(6)
	//	UserService.InsertOne(usersID[j], 0, 250000, 0)
	//	iduser, _ := UserService.FindOne(usersID[j])
	//	for k := 0; k < 5; k++ {
	//		idcategory, err := CategoryService.FindOne(categorynamestonumbers[k])
	//		if err != nil {
	//			panic(err)
	//		}
	//		for i := 0; i < 30; i++ {
	//			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
	//		}
	//	}
	//}
}
