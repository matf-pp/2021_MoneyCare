package main

import (
	"fmt"
	"main/src/db"
	"main/src/services"
	"math/rand"
	"time"
)

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

	getNormDistro(30, 25000, 1000)
	for i := 1; i < 30; i++ {
		users_id[i] = RandStringRunes(5)
		_, err := userService.InsertOne(users_id[i])
		if err != nil {
			panic(err)
		}
	}

	for i := 1; i < 30; i++ {
		userId, err := userService.FindOne(users_id[i])
		if err != nil {
			panic(err)
		}
		categoryId, err := categoryService.FindOne("food")
		if err != nil {
			panic(err)
		}

		_, err = spendingService.InsertOne(userId.ID, categoryId.ID, rnorm_values[i]/30)
	}

	getNormDistro(30, 50000, 1000)
	for i := 30; i < 60; i++ {
		users_id[i] = RandStringRunes(5)
		_, err := userService.InsertOne(users_id[i])
		if err != nil {
			panic(err)
		}
	}
	j := 0
	for i := 30; i < 60; i++ {
		userId, err := userService.FindOne(users_id[i])
		if err != nil {
			panic(err)
		}
		categoryId, err := categoryService.FindOne("food")
		if err != nil {
			panic(err)
		}

		spendingService.InsertOne(userId.ID, categoryId.ID, rnorm_values[j]/30)
		j++
	}

	getNormDistro(30, 10000, 1000)
	for i := 60; i < 90; i++ {
		users_id[i] = RandStringRunes(5)
		_, err := userService.InsertOne(users_id[i])
		if err != nil {
			panic(err)
		}
	}
	j = 0
	for i := 60; i < 90; i++ {
		userId, err := userService.FindOne(users_id[i])
		if err != nil {
			panic(err)
		}
		categoryId, err := categoryService.FindOne("food")
		if err != nil {
			panic(err)
		}

		spendingService.InsertOne(userId.ID, categoryId.ID, rnorm_values[j]/30)
		j++
	}

	id, err := userService.FindOne("tanxw")
	if err != nil {
		panic(err)
	}
	fmt.Println(id.ID)
	spendingService.FindUsersSpending(id.ID)

	for i := 0; i < 2; i++ {

	}

}

type ProductModel struct {
}

//func (this ProductModel) SumQuantities() (float64, error) {
//
//		pipeline := []bson.M{
//			{
//				"$group": bson.M{
//					"_id":   "",
//					"total": bson.M{"$sum": "$quantity"},
//				},
//			},
//		}
//		result := []bson.M{}
//		err = db.Connection("product").Pipe(pipeline).All(&result)
//		return result[0]["total"].(float64), nil
//
//}
