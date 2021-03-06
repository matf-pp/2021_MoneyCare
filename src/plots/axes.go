package plots

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/src/admin"
	"os"
	"time"

	chart "github.com/wcharczuk/go-chart/v2"
)

func DrawChart(userID primitive.ObjectID, time2 time.Time) {
	/*
	   This is an example of using the `TimeSeries` to automatically coerce time.Time values into a continuous xrange.
	   Note: chart.TimeSeries implements `ValueFormatterProvider` and as a result gives the XAxis the appropriate formatter to use for the ticks.
	*/
	spendingService := admin.SpendingService

	day0 := time2.AddDate(0, 0, -6)
	day1 := time2.AddDate(0, 0, -5)
	day2 := time2.AddDate(0, 0, -4)
	day3 := time2.AddDate(0, 0, -3)
	day4 := time2.AddDate(0, 0, -2)
	day5 := time2.AddDate(0, 0, -1)
	day6 := time2

	amount0 := spendingService.FindUsersSpendingByDate(userID, day0)
	amount1 := spendingService.FindUsersSpendingByDate(userID, day1)
	amount2 := spendingService.FindUsersSpendingByDate(userID, day2)
	amount3 := spendingService.FindUsersSpendingByDate(userID, day3)
	amount4 := spendingService.FindUsersSpendingByDate(userID, day4)
	amount5 := spendingService.FindUsersSpendingByDate(userID, day5)
	amount6 := spendingService.FindUsersSpendingByDate(userID, day6)

	graph := chart.Chart{
		Width: 512,
		Height: 256,
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: []time.Time{
					day0,
					day1,
					day2,
					day3,
					day4,
					day5,
					day6,
				},
				YValues: []float64{amount0, amount1, amount2, amount3, amount4, amount5, amount6},
			},
		},
	}

	f, _ := os.Create("graph.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
