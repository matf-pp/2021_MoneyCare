package admin

import (
	"main/src/db"
	"main/src/services"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var UsersID0 [30]string
var UsersID1 [30]string
var UsersID2 [30]string
var currentuser int
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
	UsersID0[29] = string(1)
}

func randStringRunes(n int) string {
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

	incomeclasses := make(map[[30]string]int)

	categorynamestonumbers := make(map[int]string)
	categorynamestonumbers[0] = "food"
	categorynamestonumbers[1] = "chem"
	categorynamestonumbers[2] = "other"
	categorynamestonumbers[3] = "clothes"
	categorynamestonumbers[4] = "bills"

	for j := 0; j < 5; j++ {
		UsersID0[j] = randStringRunes(6)
		UserService.InsertOne(UsersID0[j], 0, 50000, 0)
	}
	incomeclasses[UsersID0] = 0

	k := 0
	//food
	idcategory, err := CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 10000, 3000)
		iduser, _ := UserService.FindOne(UsersID0[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}
	k++
	// chem
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 2000, 1000)
		iduser, _ := UserService.FindOne(UsersID0[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}
	k++
	// other
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 10000, 7000)
		iduser, _ := UserService.FindOne(UsersID0[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}
	k++
	//clot
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 7000, 3000)
		iduser, _ := UserService.FindOne(UsersID0[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}
	k++
	//bills
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 12000, 3000)
		iduser, _ := UserService.FindOne(UsersID0[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}

	for j := 0; j < 5; j++ {
		UsersID1[j] = randStringRunes(6)
		UserService.InsertOne(UsersID1[j], 0, 70000, 0)
	}
	incomeclasses[UsersID1] = 1

	//food
	k = 0
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 20000, 5000)
		iduser, _ := UserService.FindOne(UsersID1[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}
	k++
	// chem
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 5000, 1000)

		iduser, _ := UserService.FindOne(UsersID1[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}
	k++
	// other
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 10000, 7000)
		iduser, _ := UserService.FindOne(UsersID1[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}
	k++
	//clot
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 15000, 5000)
		iduser, _ := UserService.FindOne(UsersID1[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}
	k++
	//bills
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 30000, 5000)
		iduser, _ := UserService.FindOne(UsersID1[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}

	for j := 0; j < 5; j++ {
		UsersID2[j] = randStringRunes(6)
		UserService.InsertOne(UsersID2[j], 0, 150000, 0)
	}
	incomeclasses[UsersID2] = 2

	//food
	k = 0
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 30000, 5000)
		iduser, _ := UserService.FindOne(UsersID2[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}
	k++
	// chem
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 10000, 1000)
		iduser, _ := UserService.FindOne(UsersID2[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}
	k++
	// other
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 40000, 20000)
		iduser, _ := UserService.FindOne(UsersID2[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}
	k++
	//clot
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 20000, 5000)
		iduser, _ := UserService.FindOne(UsersID2[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}
	k++
	//bills
	idcategory, err = CategoryService.FindOne(categorynamestonumbers[k])
	for j := 0; j < 5; j++ {
		getNormDistro(30, 40000, 10000)
		iduser, _ := UserService.FindOne(UsersID2[j])
		if err != nil {
			panic(err)
		}
		for i := 0; i < 30; i++ {
			SpendingService.InsertOne(iduser.ID, idcategory.ID, rnormValues[i]/30)
		}
	}

}
