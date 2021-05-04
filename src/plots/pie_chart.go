package plots

import (
	"github.com/wcharczuk/go-chart/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/src/admin"
	"os"
)

func PieChart(userID primitive.ObjectID) {
	categoryService := admin.CategoryService
	spendingService := admin.SpendingService

	foodID, err := categoryService.FindOne("food")
	if err != nil {
		panic(err)
	}
	foodAmount := spendingService.FindUsersSpendingByCategory(userID, foodID.ID)
	billsID, err := categoryService.FindOne("bills")
	if err != nil {
		panic(err)
	}
	billsAmount := spendingService.FindUsersSpendingByCategory(userID, billsID.ID)
	clothesID, err := categoryService.FindOne("clothes")
	if err != nil {
		panic(err)
	}
	clothesAmount := spendingService.FindUsersSpendingByCategory(userID, clothesID.ID)
	otherID, err := categoryService.FindOne("other")
	if err != nil {
		panic(err)
	}
	otherAmount := spendingService.FindUsersSpendingByCategory(userID, otherID.ID)
	chemID, err := categoryService.FindOne("chem")
	if err != nil {
		panic(err)
	}
	chemAmount := spendingService.FindUsersSpendingByCategory(userID, chemID.ID)

	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: foodAmount, Label: "Food"},
			{Value: billsAmount, Label: "Bills"},
			{Value: clothesAmount, Label: "Clothes"},
			{Value: otherAmount, Label: "Other"},
			{Value: chemAmount, Label: "Chem"},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}
