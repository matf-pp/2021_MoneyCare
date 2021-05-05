package plots

import (
	"github.com/wcharczuk/go-chart/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/src/admin"
	"os"
	"time"
)

func PieChart(userID primitive.ObjectID, month time.Month) {
	categoryService := admin.CategoryService
	spendingService := admin.SpendingService

	foodID, err := categoryService.FindOne("food")
	if err != nil {
		panic(err)
	}
	foodAmount := spendingService.FindUsersSpendingByCategoryByMonth(userID, foodID.ID, month)
	billsID, err := categoryService.FindOne("bills")
	if err != nil {
		panic(err)
	}
	billsAmount := spendingService.FindUsersSpendingByCategoryByMonth(userID, billsID.ID, month)
	clothesID, err := categoryService.FindOne("clothes")
	if err != nil {
		panic(err)
	}
	clothesAmount := spendingService.FindUsersSpendingByCategoryByMonth(userID, clothesID.ID, month)
	otherID, err := categoryService.FindOne("other")
	if err != nil {
		panic(err)
	}
	otherAmount := spendingService.FindUsersSpendingByCategoryByMonth(userID, otherID.ID, month)
	chemID, err := categoryService.FindOne("chem")
	if err != nil {
		panic(err)
	}
	chemAmount := spendingService.FindUsersSpendingByCategoryByMonth(userID, chemID.ID, month)

	pie := chart.PieChart{
		Width:  256,
		Height: 256,
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
