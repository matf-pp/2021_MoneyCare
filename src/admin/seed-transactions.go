package main

import (
	"main/src/db"
	"main/src/services"
	_ "main/src/services"
	"math/rand"
	"time"
)

var categorynamestonumbers map[int]string
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var users_id [90]string
var rnorm_values [30]int
var connectionString string
var databaseName string
var userCollectionName string
var categoryCollectionName string
var spendingCollectionName string

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
		rnorm_values[i] = int(normalInverse(mi, sigma))
	}
}

func normalInverse(mu float32, sigma float32) float32 {
	return float32(rand.NormFloat64()*float64(sigma) + float64(mu))
}

func main() {
	connection, err := db.NewConnection(connectionString, databaseName)
	if err != nil {
		panic(err)
	}

	defer connection.Disconnect()

	userService := services.NewUserService(connection, userCollectionName)
	categoryService := services.NewCategoryService(connection, categoryCollectionName)
	spendingService := services.NewSpendingService(connection, spendingCollectionName)

	categorynamestonumbers[0] = "food"
	categorynamestonumbers[1] = "chem"
	categorynamestonumbers[2] = "other"
	categorynamestonumbers[3] = "clothes"
	categorynamestonumbers[4] = "bills"

	for j := 0; j < 5; j++ {
		getNormDistro(30, 10000, 2000)
		users_id[j] = RandStringRunes(6)
		userService.InsertOne(users_id[j])
		iduser, _ := userService.FindOne(users_id[j])
		for k := 0; k < 5; k++ {
			idcategory, err := categoryService.FindOne(categorynamestonumbers[k])
			if err != nil {
				panic(err)
			}
			for i := 0; i < 30; i++ {
				spendingService.InsertOne(iduser.ID, idcategory.ID, rnorm_values[i]/30)
			}
		}
	}

	for j := 0; j < 5; j++ {
		getNormDistro(30, 30000, 10000)
		users_id[j] = RandStringRunes(6)
		userService.InsertOne(users_id[j])
		iduser, _ := userService.FindOne(users_id[j])
		for k := 0; k < 5; k++ {
			idcategory, err := categoryService.FindOne(categorynamestonumbers[k])
			if err != nil {
				panic(err)
			}
			for i := 0; i < 30; i++ {
				spendingService.InsertOne(iduser.ID, idcategory.ID, rnorm_values[i]/30)
			}
		}
	}
	for j := 0; j < 5; j++ {
		getNormDistro(30, 50000, 10000)
		users_id[j] = RandStringRunes(6)
		userService.InsertOne(users_id[j])
		iduser, _ := userService.FindOne(users_id[j])
		for k := 0; k < 5; k++ {
			idcategory, err := categoryService.FindOne(categorynamestonumbers[k])
			if err != nil {
				panic(err)
			}
			for i := 0; i < 30; i++ {
				spendingService.InsertOne(iduser.ID, idcategory.ID, rnorm_values[i]/30)
			}
		}
	}
}
